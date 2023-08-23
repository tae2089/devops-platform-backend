package domain

import (
	"bytes"
	"text/template"

	"github.com/tae2089/bob-logging/logger"
)

type JobTemplateWritable interface {
	Write() string
}

type JenkinsFrontFile struct {
	GitURL       string
	Branch       string
	SlackChannel string
	BuckName     string
	CloudFrontID string
	AwsProfile   string
}

func (j *JenkinsFrontFile) Write() string {
	tmpl, err := template.New("Create front deploy").Parse(frontJobFileTemplate)
	if err != nil {
		logger.Error(err)
		return ""
	}
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, j); err != nil {
		logger.Error(err)
		return ""
	}
	return tpl.String()
}

type JenkinsJob struct {
	BranchName   string
	TeamName     string
	WebhookToken string
	GitURL       string
	FileName     string
}

func (j *JenkinsJob) Write() string {
	tmpl, err := template.New("Create Job").Parse(jobTemplate)
	if err != nil {
		logger.Error(err)
		return ""
	}
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, j); err != nil {
		logger.Error(err)
		return ""
	}
	return tpl.String()
}

type ResultMessageJenkinsJob struct {
	Project      string
	WebhookToken string
	GitURL       string
	Branch       string
	FileName     string
}

func (r *ResultMessageJenkinsJob) Write() string {
	tmpl, err := template.New("Create Job").Parse(registJobResultTemplate)
	if err != nil {
		logger.Error(err)
		return ""
	}
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, r); err != nil {
		logger.Error(err)
		return ""
	}
	return tpl.String()
}

const registJobResultTemplate = `
	project: {{.Project}}
	branch: {{.Branch}}
    webhook_token: {{.WebhookToken}}
	file_name: {{.FileName}}
	git_repo_url: {{.GitURL}}
`

const frontJobFileTemplate = `

//파이프라인 시작
pipeline {
//필요한 사전 선언
    environment {
//공통변수 시작
        GITURL = "{{.GitURL}}"
        BRANCH = "{{.Branch}}"
        SLACKCHANNEL = "#{{.SlackChannel}}"
        S3 = "{{.BuckName}}"
        CFID = "{{.CloudFrontID}}"
        AWSPROFILE = "{{.AwsProfile}}"
 //추가변수 설정
        }
    tools {
        nodejs "nodejs"
}

  agent any
  stages {
    stage('Cloning Git') {
      steps {
        slackSend (channel: SLACKCHANNEL, color: '#00FF00', message: "${env.JOB_NAME}앱의 CI 과정이 시작되었습니다 \n Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})")
        git([url: GITURL, branch: BRANCH, credentialsId: 'sshkey'])
          }
            } 
    
    stage('Yarn Build') {
      steps{
            sh '''
            yarn install
            yarn build
            '''
    }
        }

    stage('S3 Deploy') {
    steps{
                sh 'aws s3 sync ./build s3://$S3 --delete --profile $AWSPROFILE'

                sh '''
                INVALIDATION_ID=$(aws cloudfront create-invalidation --profile $AWSPROFILE --distribution-id $CFID --paths "/*" | jq -r '.Invalidation.Id')
                aws cloudfront wait invalidation-completed --profile $AWSPROFILE --distribution-id $CFID --id $INVALIDATION_ID
                '''
                }
            }
        }
             
            
    post {
        always {
            echo 'One way or another, I have finished'
        }
        success {
            slackSend (channel: SLACKCHANNEL, color: '#00FF00', message: "빌드 완료 \n ${env.JOB_NAME}앱의 CI 과정과 무효화가 성공적으로 끝났습니다 \n 변경된 사항은 바로 적용 됩니다 <@U02NG39P1GW>\n Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})")
        }
        unstable {
            echo 'I am unstable :/!'
        }
        failure {
            slackSend (channel: SLACKCHANNEL, color: '#00FF00', message: "빌드가 실패하였습니다 \n ${env.JOB_NAME}앱의 젠킨스 콘솔을 확인해주세요 <@U02NG39P1GW> \n Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})")
            slackSend (channel: '#devops-alarm', color: '#00FF00', message: "빌드가 실패하였습니다 \n ${env.JOB_NAME}앱의 젠킨스 콘솔을 확인해주세요 \n Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})")
        }
        changed {
            echo 'Things were different before..........'
        }
    }
}
`

const backendJobFileTemplate = ``

const jobTemplate = `
<flow-definition plugin="workflow-job@1254.v3f64639b_11dd">
  <actions>
    <org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction plugin="pipeline-model-definition@2.2118.v31fd5b_9944b_5"/>
    <org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction plugin="pipeline-model-definition@2.2118.v31fd5b_9944b_5">
      <jobProperties/>
      <triggers/>
      <parameters/>
      <options/>
    </org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction>
  </actions>
  <description></description>
  <keepDependencies>false</keepDependencies>
  <properties>
    <hudson.security.AuthorizationMatrixProperty>
      <inheritanceStrategy class="org.jenkinsci.plugins.matrixauth.inheritance.NonInheritingStrategy"/>
      <permission>GROUP:hudson.model.Item.Build:{{.TeamName}}</permission>
      <permission>GROUP:hudson.model.Item.Cancel:{{.TeamName}}</permission>
      <permission>USER:hudson.model.Item.Configure:devops</permission>
      <permission>USER:hudson.model.Item.Create:devops</permission>
      <permission>USER:hudson.model.Item.Delete:devops</permission>
      <permission>GROUP:hudson.model.Item.Read:{{.TeamName}}</permission>
      <permission>USER:hudson.model.Item.Read:devops</permission>
      <permission>GROUP:hudson.model.Item.Workspace:{{.TeamName}}</permission>
      <permission>USER:hudson.model.Item.Workspace:devops</permission>
      <permission>GROUP:hudson.model.Run.Delete:{{.TeamName}}</permission>
      <permission>GROUP:hudson.model.Run.Replay:{{.TeamName}}</permission>
      <permission>GROUP:hudson.model.Run.Update:{{.TeamName}}</permission>
      <permission>GROUP:hudson.model.View.Read:{{.TeamName}}</permission>
      <permission>USER:hudson.model.View.Read:devops</permission>
      <permission>GROUP:hudson.scm.SCM.Tag:{{.TeamName}}</permission>
      <permission>USER:hudson.scm.SCM.Tag:devops</permission>
    </hudson.security.AuthorizationMatrixProperty>
    <jenkins.model.BuildDiscarderProperty>
      <strategy class="hudson.tasks.LogRotator">
        <daysToKeep>-1</daysToKeep>
        <numToKeep>10</numToKeep>
        <artifactDaysToKeep>-1</artifactDaysToKeep>
        <artifactNumToKeep>-1</artifactNumToKeep>
      </strategy>
    </jenkins.model.BuildDiscarderProperty>
    <org.jenkinsci.plugins.workflow.job.properties.PipelineTriggersJobProperty>
      <triggers>
        <org.jenkinsci.plugins.gwt.GenericTrigger plugin="generic-webhook-trigger@1.85.2">
          <spec></spec>
          <genericVariables>
            <org.jenkinsci.plugins.gwt.GenericVariable>
              <expressionType>JSONPath</expressionType>
              <key>ref</key>
              <value>$.ref</value>
              <regexpFilter></regexpFilter>
              <defaultValue></defaultValue>
            </org.jenkinsci.plugins.gwt.GenericVariable>
          </genericVariables>
          <regexpFilterText>$ref</regexpFilterText>
          <regexpFilterExpression>^(refs/heads/{{.BranchName}})</regexpFilterExpression>
          <printPostContent>false</printPostContent>
          <printContributedVariables>false</printContributedVariables>
          <causeString>Generic Cause</causeString>
          <token>{{.WebhookToken}}</token>
          <tokenCredentialId></tokenCredentialId>
          <silentResponse>false</silentResponse>
          <overrideQuietPeriod>false</overrideQuietPeriod>
          <shouldNotFlattern>false</shouldNotFlattern>
        </org.jenkinsci.plugins.gwt.GenericTrigger>
      </triggers>
    </org.jenkinsci.plugins.workflow.job.properties.PipelineTriggersJobProperty>
  </properties>
  <definition class="org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition" plugin="workflow-cps@3565.v4b_d9b_8c29a_b_3">
    <scm class="hudson.plugins.git.GitSCM" plugin="git@4.14.3">
      <configVersion>2</configVersion>
      <userRemoteConfigs>
        <hudson.plugins.git.UserRemoteConfig>
          <refspec>+refs/heads/{{.BranchName}}:refs/remotes/origin/{{.BranchName}}</refspec>
          <url>{{.GitURL}}</url>
          <credentialsId>sshkey</credentialsId>
        </hudson.plugins.git.UserRemoteConfig>
      </userRemoteConfigs>
      <branches>
        <hudson.plugins.git.BranchSpec>
          <name>*/{{.BranchName}}</name>
        </hudson.plugins.git.BranchSpec>
      </branches>
      <doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations>
      <submoduleCfg class="empty-list"/>
      <extensions/>
    </scm>
    <scriptPath>{{.FileName}}</scriptPath>
    <lightweight>true</lightweight>
  </definition>
  <triggers/>
  <disabled>false</disabled>
</flow-definition>
`

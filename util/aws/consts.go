package aws

type ManagedCache string

const (
	Amplify                                ManagedCache = "2e54312d-136d-493c-8eb9-b001f22f67d2" //Policy for Amplify Origin
	CachingDisabled                        ManagedCache = "4135ea2d-6df8-44a3-9df3-4b5a84be39ad" //Policy with caching disabled
	CachingOptimized                       ManagedCache = "658327ea-f89d-4fab-a63d-7e88639e58f6" //Policy with caching enabled. Supports Gzip and Brotli compression.
	CachingOptimizedForUncompressedObjects ManagedCache = "b2884449-e4de-46a7-ac36-70bc7f1ddd6d" //Default policy when compression is disabled
	ElementalMediaPackage                  ManagedCache = "08627262-05a9-4f76-9ded-b50ca2e3a84f" //Policy for Elemental MediaPackage Origin
)

func (c ManagedCache) String() string {
	switch c {
	case Amplify:
		return "2e54312d-136d-493c-8eb9-b001f22f67d2"
	case CachingDisabled:
		return "4135ea2d-6df8-44a3-9df3-4b5a84be39ad"
	case CachingOptimized:
		return "658327ea-f89d-4fab-a63d-7e88639e58f6"
	case CachingOptimizedForUncompressedObjects:
		return "b2884449-e4de-46a7-ac36-70bc7f1ddd6d"
	case ElementalMediaPackage:
		return "08627262-05a9-4f76-9ded-b50ca2e3a84f"
	}
	return "unknown"
}

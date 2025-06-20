// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailscale

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-tailscale/sdk/go/tailscale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "tailscale:index/acl:Acl":
		r = &Acl{}
	case "tailscale:index/awsExternalId:AwsExternalId":
		r = &AwsExternalId{}
	case "tailscale:index/contacts:Contacts":
		r = &Contacts{}
	case "tailscale:index/deviceAuthorization:DeviceAuthorization":
		r = &DeviceAuthorization{}
	case "tailscale:index/deviceKey:DeviceKey":
		r = &DeviceKey{}
	case "tailscale:index/deviceSubnetRoutes:DeviceSubnetRoutes":
		r = &DeviceSubnetRoutes{}
	case "tailscale:index/deviceTags:DeviceTags":
		r = &DeviceTags{}
	case "tailscale:index/dnsNameservers:DnsNameservers":
		r = &DnsNameservers{}
	case "tailscale:index/dnsPreferences:DnsPreferences":
		r = &DnsPreferences{}
	case "tailscale:index/dnsSearchPaths:DnsSearchPaths":
		r = &DnsSearchPaths{}
	case "tailscale:index/dnsSplitNameservers:DnsSplitNameservers":
		r = &DnsSplitNameservers{}
	case "tailscale:index/logstreamConfiguration:LogstreamConfiguration":
		r = &LogstreamConfiguration{}
	case "tailscale:index/oauthClient:OauthClient":
		r = &OauthClient{}
	case "tailscale:index/postureIntegration:PostureIntegration":
		r = &PostureIntegration{}
	case "tailscale:index/tailnetKey:TailnetKey":
		r = &TailnetKey{}
	case "tailscale:index/tailnetSettings:TailnetSettings":
		r = &TailnetSettings{}
	case "tailscale:index/webhook:Webhook":
		r = &Webhook{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:tailscale" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/acl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/awsExternalId",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/contacts",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/deviceAuthorization",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/deviceKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/deviceSubnetRoutes",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/deviceTags",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/dnsNameservers",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/dnsPreferences",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/dnsSearchPaths",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/dnsSplitNameservers",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/logstreamConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/oauthClient",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/postureIntegration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/tailnetKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/tailnetSettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/webhook",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"tailscale",
		&pkg{version},
	)
}

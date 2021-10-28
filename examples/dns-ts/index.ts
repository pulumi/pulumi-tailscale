import * as pulumi from "@pulumi/pulumi";
import * as tailscale from "@pulumi/tailscale";

const nameservers = new tailscale.DnsNameservers("ts-demo", {
    nameservers: ["1.1.1.1"]
});

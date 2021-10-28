"""A Python Pulumi program"""

import json
import pulumi_tailscale as tailscale

acl = tailscale.Acl("demo-py",
                    acl=json.dumps({
                        "ACLs": [{
                            "Action": "accept",
                            "User": ["*"],
                            "Ports": ["*:*"],
                        }]
                    }))


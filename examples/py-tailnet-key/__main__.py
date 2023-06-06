"""A Python Pulumi program"""

import pulumi_tailscale as tailscale

tailscale.TailnetKey("demo-py", tags=["tag:server"])

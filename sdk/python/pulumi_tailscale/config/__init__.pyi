# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

apiKey: Optional[str]
"""
The API key to use for authenticating requests to the API. Can be set via the TAILSCALE_API_KEY environment variable.
Conflicts with 'oauth_client_id' and 'oauth_client_secret'.
"""

baseUrl: Optional[str]
"""
The base URL of the Tailscale API. Defaults to https://api.tailscale.com. Can be set via the TAILSCALE_BASE_URL
environment variable.
"""

oauthClientId: Optional[str]
"""
The OAuth application's ID when using OAuth client credentials. Can be set via the TAILSCALE_OAUTH_CLIENT_ID environment
variable. Both 'oauth_client_id' and 'oauth_client_secret' must be set. Conflicts with 'api_key'.
"""

oauthClientSecret: Optional[str]
"""
The OAuth application's secret when using OAuth client credentials. Can be set via the TAILSCALE_OAUTH_CLIENT_SECRET
environment variable. Both 'oauth_client_id' and 'oauth_client_secret' must be set. Conflicts with 'api_key'.
"""

scopes: Optional[str]
"""
The OAuth 2.0 scopes to request when for the access token generated using the supplied OAuth client credentials. See
https://tailscale.com/kb/1215/oauth-clients/#scopes for available scopes. Only valid when both 'oauth_client_id' and
'oauth_client_secret' are set.
"""

tailnet: Optional[str]
"""
The organization name of the Tailnet in which to perform actions. Can be set via the TAILSCALE_TAILNET environment
variable. Default is the tailnet that owns API credentials passed to the provider.
"""

userAgent: Optional[str]
"""
User-Agent header for API requests.
"""


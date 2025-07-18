# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs

__all__ = [
    'GetUsersResult',
    'AwaitableGetUsersResult',
    'get_users',
    'get_users_output',
]

@pulumi.output_type
class GetUsersResult:
    """
    A collection of values returned by getUsers.
    """
    def __init__(__self__, id=None, role=None, type=None, users=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def role(self) -> Optional[builtins.str]:
        """
        Filters the users list to elements whose role is the provided value.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def type(self) -> Optional[builtins.str]:
        """
        Filters the users list to elements whose type is the provided value.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def users(self) -> Sequence['outputs.GetUsersUserResult']:
        """
        The list of users in the tailnet
        """
        return pulumi.get(self, "users")


class AwaitableGetUsersResult(GetUsersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUsersResult(
            id=self.id,
            role=self.role,
            type=self.type,
            users=self.users)


def get_users(role: Optional[builtins.str] = None,
              type: Optional[builtins.str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUsersResult:
    """
    The users data source describes a list of users in a tailnet

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tailscale as tailscale

    all_users = tailscale.get_users()
    ```


    :param builtins.str role: Filters the users list to elements whose role is the provided value.
    :param builtins.str type: Filters the users list to elements whose type is the provided value.
    """
    __args__ = dict()
    __args__['role'] = role
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tailscale:index/getUsers:getUsers', __args__, opts=opts, typ=GetUsersResult).value

    return AwaitableGetUsersResult(
        id=pulumi.get(__ret__, 'id'),
        role=pulumi.get(__ret__, 'role'),
        type=pulumi.get(__ret__, 'type'),
        users=pulumi.get(__ret__, 'users'))
def get_users_output(role: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                     type: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetUsersResult]:
    """
    The users data source describes a list of users in a tailnet

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tailscale as tailscale

    all_users = tailscale.get_users()
    ```


    :param builtins.str role: Filters the users list to elements whose role is the provided value.
    :param builtins.str type: Filters the users list to elements whose type is the provided value.
    """
    __args__ = dict()
    __args__['role'] = role
    __args__['type'] = type
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('tailscale:index/getUsers:getUsers', __args__, opts=opts, typ=GetUsersResult)
    return __ret__.apply(lambda __response__: GetUsersResult(
        id=pulumi.get(__response__, 'id'),
        role=pulumi.get(__response__, 'role'),
        type=pulumi.get(__response__, 'type'),
        users=pulumi.get(__response__, 'users')))

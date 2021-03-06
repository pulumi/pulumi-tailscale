# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['TailnetKeyArgs', 'TailnetKey']

@pulumi.input_type
class TailnetKeyArgs:
    def __init__(__self__, *,
                 ephemeral: Optional[pulumi.Input[bool]] = None,
                 preauthorized: Optional[pulumi.Input[bool]] = None,
                 reusable: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a TailnetKey resource.
        :param pulumi.Input[bool] ephemeral: Indicates if the key is ephemeral.
        :param pulumi.Input[bool] preauthorized: Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default.
        :param pulumi.Input[bool] reusable: Indicates if the key is reusable or single-use.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: List of tags to apply to the machines authenticated by the key.
        """
        if ephemeral is not None:
            pulumi.set(__self__, "ephemeral", ephemeral)
        if preauthorized is not None:
            pulumi.set(__self__, "preauthorized", preauthorized)
        if reusable is not None:
            pulumi.set(__self__, "reusable", reusable)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def ephemeral(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if the key is ephemeral.
        """
        return pulumi.get(self, "ephemeral")

    @ephemeral.setter
    def ephemeral(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ephemeral", value)

    @property
    @pulumi.getter
    def preauthorized(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default.
        """
        return pulumi.get(self, "preauthorized")

    @preauthorized.setter
    def preauthorized(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "preauthorized", value)

    @property
    @pulumi.getter
    def reusable(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if the key is reusable or single-use.
        """
        return pulumi.get(self, "reusable")

    @reusable.setter
    def reusable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reusable", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of tags to apply to the machines authenticated by the key.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _TailnetKeyState:
    def __init__(__self__, *,
                 ephemeral: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 preauthorized: Optional[pulumi.Input[bool]] = None,
                 reusable: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering TailnetKey resources.
        :param pulumi.Input[bool] ephemeral: Indicates if the key is ephemeral.
        :param pulumi.Input[str] key: The authentication key
        :param pulumi.Input[bool] preauthorized: Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default.
        :param pulumi.Input[bool] reusable: Indicates if the key is reusable or single-use.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: List of tags to apply to the machines authenticated by the key.
        """
        if ephemeral is not None:
            pulumi.set(__self__, "ephemeral", ephemeral)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if preauthorized is not None:
            pulumi.set(__self__, "preauthorized", preauthorized)
        if reusable is not None:
            pulumi.set(__self__, "reusable", reusable)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def ephemeral(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if the key is ephemeral.
        """
        return pulumi.get(self, "ephemeral")

    @ephemeral.setter
    def ephemeral(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ephemeral", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The authentication key
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def preauthorized(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default.
        """
        return pulumi.get(self, "preauthorized")

    @preauthorized.setter
    def preauthorized(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "preauthorized", value)

    @property
    @pulumi.getter
    def reusable(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if the key is reusable or single-use.
        """
        return pulumi.get(self, "reusable")

    @reusable.setter
    def reusable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reusable", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of tags to apply to the machines authenticated by the key.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class TailnetKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ephemeral: Optional[pulumi.Input[bool]] = None,
                 preauthorized: Optional[pulumi.Input[bool]] = None,
                 reusable: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        The tailnet_key resource allows you to create pre-authentication keys that can register new nodes without needing to sign in via a web browser. See https://tailscale.com/kb/1085/auth-keys for more information

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tailscale as tailscale

        sample_key = tailscale.TailnetKey("sampleKey",
            ephemeral=False,
            preauthorized=True,
            reusable=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] ephemeral: Indicates if the key is ephemeral.
        :param pulumi.Input[bool] preauthorized: Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default.
        :param pulumi.Input[bool] reusable: Indicates if the key is reusable or single-use.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: List of tags to apply to the machines authenticated by the key.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TailnetKeyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The tailnet_key resource allows you to create pre-authentication keys that can register new nodes without needing to sign in via a web browser. See https://tailscale.com/kb/1085/auth-keys for more information

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tailscale as tailscale

        sample_key = tailscale.TailnetKey("sampleKey",
            ephemeral=False,
            preauthorized=True,
            reusable=True)
        ```

        :param str resource_name: The name of the resource.
        :param TailnetKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TailnetKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ephemeral: Optional[pulumi.Input[bool]] = None,
                 preauthorized: Optional[pulumi.Input[bool]] = None,
                 reusable: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TailnetKeyArgs.__new__(TailnetKeyArgs)

            __props__.__dict__["ephemeral"] = ephemeral
            __props__.__dict__["preauthorized"] = preauthorized
            __props__.__dict__["reusable"] = reusable
            __props__.__dict__["tags"] = tags
            __props__.__dict__["key"] = None
        super(TailnetKey, __self__).__init__(
            'tailscale:index/tailnetKey:TailnetKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ephemeral: Optional[pulumi.Input[bool]] = None,
            key: Optional[pulumi.Input[str]] = None,
            preauthorized: Optional[pulumi.Input[bool]] = None,
            reusable: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'TailnetKey':
        """
        Get an existing TailnetKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] ephemeral: Indicates if the key is ephemeral.
        :param pulumi.Input[str] key: The authentication key
        :param pulumi.Input[bool] preauthorized: Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default.
        :param pulumi.Input[bool] reusable: Indicates if the key is reusable or single-use.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: List of tags to apply to the machines authenticated by the key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TailnetKeyState.__new__(_TailnetKeyState)

        __props__.__dict__["ephemeral"] = ephemeral
        __props__.__dict__["key"] = key
        __props__.__dict__["preauthorized"] = preauthorized
        __props__.__dict__["reusable"] = reusable
        __props__.__dict__["tags"] = tags
        return TailnetKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def ephemeral(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if the key is ephemeral.
        """
        return pulumi.get(self, "ephemeral")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        The authentication key
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def preauthorized(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default.
        """
        return pulumi.get(self, "preauthorized")

    @property
    @pulumi.getter
    def reusable(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if the key is reusable or single-use.
        """
        return pulumi.get(self, "reusable")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of tags to apply to the machines authenticated by the key.
        """
        return pulumi.get(self, "tags")


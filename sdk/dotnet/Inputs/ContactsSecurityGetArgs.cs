// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tailscale.Inputs
{

    public sealed class ContactsSecurityGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Email address to send communications to
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        public ContactsSecurityGetArgs()
        {
        }
        public static new ContactsSecurityGetArgs Empty => new ContactsSecurityGetArgs();
    }
}

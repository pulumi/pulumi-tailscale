// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tailscale.Outputs
{

    [OutputType]
    public sealed class GetUsersUserResult
    {
        /// <summary>
        /// The time the user joined their tailnet.
        /// </summary>
        public readonly string Created;
        /// <summary>
        /// true when the user has a node currently connected to the control server.
        /// </summary>
        public readonly bool CurrentlyConnected;
        /// <summary>
        /// Number of devices the user owns.
        /// </summary>
        public readonly int DeviceCount;
        /// <summary>
        /// The name of the user.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The unique identifier for the user.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The later of either: a) The last time any of the user's nodes were connected to the network or b) The last time the user authenticated to any tailscale service, including the admin panel.
        /// </summary>
        public readonly string LastSeen;
        /// <summary>
        /// The emailish login name of the user.
        /// </summary>
        public readonly string LoginName;
        /// <summary>
        /// The profile pic URL for the user.
        /// </summary>
        public readonly string ProfilePicUrl;
        /// <summary>
        /// The role of the user.
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// The status of the user.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The tailnet that owns the user.
        /// </summary>
        public readonly string TailnetId;
        /// <summary>
        /// The type of relation this user has to the tailnet associated with the request.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetUsersUserResult(
            string created,

            bool currentlyConnected,

            int deviceCount,

            string displayName,

            string id,

            string lastSeen,

            string loginName,

            string profilePicUrl,

            string role,

            string status,

            string tailnetId,

            string type)
        {
            Created = created;
            CurrentlyConnected = currentlyConnected;
            DeviceCount = deviceCount;
            DisplayName = displayName;
            Id = id;
            LastSeen = lastSeen;
            LoginName = loginName;
            ProfilePicUrl = profilePicUrl;
            Role = role;
            Status = status;
            TailnetId = tailnetId;
            Type = type;
        }
    }
}
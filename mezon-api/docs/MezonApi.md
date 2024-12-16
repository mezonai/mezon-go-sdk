# \MezonApi

All URIs are relative to *http://127.0.0.1:7350*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MezonAddApp**](MezonApi.md#MezonAddApp) | **Post** /v2/apps/add | Add a new apps.
[**MezonAddAppToClan**](MezonApi.md#MezonAddAppToClan) | **Post** /v2/apps/app/{appId}/clan/{clanId} | Add an app to clan.
[**MezonAddChannelFavorite**](MezonApi.md#MezonAddChannelFavorite) | **Post** /v2/channel/favorite | 
[**MezonAddChannelUsers**](MezonApi.md#MezonAddChannelUsers) | **Post** /v2/channel/{channelId}/add | Add users to a channel.
[**MezonAddClanSticker**](MezonApi.md#MezonAddClanSticker) | **Post** /v2/sticker | Add a new sticker
[**MezonAddFriends**](MezonApi.md#MezonAddFriends) | **Post** /v2/friend | Add friends by ID or username to a user&#39;s account.
[**MezonAddRolesChannelDesc**](MezonApi.md#MezonAddRolesChannelDesc) | **Post** /v2/rolechannel/addrole | 
[**MezonAuthenticate**](MezonApi.md#MezonAuthenticate) | **Post** /v2/apps/authenticate/token | Authenticate a app with a token against the server.
[**MezonAuthenticateApple**](MezonApi.md#MezonAuthenticateApple) | **Post** /v2/account/authenticate/apple | Authenticate a user with an Apple ID against the server.
[**MezonAuthenticateCustom**](MezonApi.md#MezonAuthenticateCustom) | **Post** /v2/account/authenticate/custom | Authenticate a user with a custom id against the server.
[**MezonAuthenticateDevice**](MezonApi.md#MezonAuthenticateDevice) | **Post** /v2/account/authenticate/device | Authenticate a user with a device id against the server.
[**MezonAuthenticateEmail**](MezonApi.md#MezonAuthenticateEmail) | **Post** /v2/account/authenticate/email | Authenticate a user with an email+password against the server.
[**MezonAuthenticateFacebook**](MezonApi.md#MezonAuthenticateFacebook) | **Post** /v2/account/authenticate/facebook | Authenticate a user with a Facebook OAuth token against the server.
[**MezonAuthenticateFacebookInstantGame**](MezonApi.md#MezonAuthenticateFacebookInstantGame) | **Post** /v2/account/authenticate/facebookinstantgame | Authenticate a user with a Facebook Instant Game token against the server.
[**MezonAuthenticateGameCenter**](MezonApi.md#MezonAuthenticateGameCenter) | **Post** /v2/account/authenticate/gamecenter | Authenticate a user with Apple&#39;s GameCenter against the server.
[**MezonAuthenticateGoogle**](MezonApi.md#MezonAuthenticateGoogle) | **Post** /v2/account/authenticate/google | Authenticate a user with Google against the server.
[**MezonAuthenticateGoogleRedirect**](MezonApi.md#MezonAuthenticateGoogleRedirect) | **Get** /v2/account/authenticate/google | Authenticate a user with Google against the server.
[**MezonAuthenticateSteam**](MezonApi.md#MezonAuthenticateSteam) | **Post** /v2/account/authenticate/steam | Authenticate a user with Steam against the server.
[**MezonBanApp**](MezonApi.md#MezonBanApp) | **Post** /v2/apps/app/{id}/ban | Ban a app.
[**MezonBlockFriends**](MezonApi.md#MezonBlockFriends) | **Post** /v2/friend/block | Block one or more users by ID or username.
[**MezonChangeChannelCategory**](MezonApi.md#MezonChangeChannelCategory) | **Patch** /v2/rolechannel/category/{newCategoryId} | update the category of a channel
[**MezonCheckDuplicateClanName**](MezonApi.md#MezonCheckDuplicateClanName) | **Get** /v2/clandesc/{clanName} | check duplicate clan name
[**MezonCheckLoginRequest**](MezonApi.md#MezonCheckLoginRequest) | **Post** /v2/account/authenticate/checklogin | 
[**MezonCloseDirectMess**](MezonApi.md#MezonCloseDirectMess) | **Put** /v2/direct/close | close direct message.
[**MezonConfirmLogin**](MezonApi.md#MezonConfirmLogin) | **Post** /v2/account/authenticate/confirmlogin | 
[**MezonCreateActiviy**](MezonApi.md#MezonCreateActiviy) | **Post** /v2/activity | Create user activity
[**MezonCreateCategoryDesc**](MezonApi.md#MezonCreateCategoryDesc) | **Post** /v2/createcategory | 
[**MezonCreateChannelDesc**](MezonApi.md#MezonCreateChannelDesc) | **Post** /v2/channeldesc | Create a new channel with the current user as the owner.
[**MezonCreateClanDesc**](MezonApi.md#MezonCreateClanDesc) | **Post** /v2/clandesc | Create a clan
[**MezonCreateClanEmoji**](MezonApi.md#MezonCreateClanEmoji) | **Post** /v2/emoji/create | Post clan Emoji  /v2/emoji/create
[**MezonCreateEvent**](MezonApi.md#MezonCreateEvent) | **Post** /v2/eventmanagement/create | Create a new event for clan.
[**MezonCreateLinkInviteUser**](MezonApi.md#MezonCreateLinkInviteUser) | **Post** /v2/invite | Add users to a channel.
[**MezonCreateOnboarding**](MezonApi.md#MezonCreateOnboarding) | **Post** /v2/onboarding | create onboarding.
[**MezonCreatePinMessage**](MezonApi.md#MezonCreatePinMessage) | **Post** /v2/pinmessage/set | set notification user channel.
[**MezonCreateQRLogin**](MezonApi.md#MezonCreateQRLogin) | **Post** /v2/account/authenticate/createqrlogin | 
[**MezonCreateRole**](MezonApi.md#MezonCreateRole) | **Post** /v2/roles | Create a new role for clan.
[**MezonCreateSystemMessage**](MezonApi.md#MezonCreateSystemMessage) | **Post** /v2/systemmessages | Create a system messages.
[**MezonDeleteAccount**](MezonApi.md#MezonDeleteAccount) | **Delete** /v2/account | Delete the current user&#39;s account.
[**MezonDeleteApp**](MezonApi.md#MezonDeleteApp) | **Delete** /v2/apps/app/{id} | Delete all information stored for an app.
[**MezonDeleteByIdClanEmoji**](MezonApi.md#MezonDeleteByIdClanEmoji) | **Delete** /v2/emoji/{id} | Delete a emoji by ID.
[**MezonDeleteCategoryDesc**](MezonApi.md#MezonDeleteCategoryDesc) | **Delete** /v2/deletecategory/category_id/{categoryId}/clan_id/{clanId} | 
[**MezonDeleteCategoryOrder**](MezonApi.md#MezonDeleteCategoryOrder) | **Delete** /v2/deletecategoryorder/clan_id/{clanId} | 
[**MezonDeleteChannelCanvas**](MezonApi.md#MezonDeleteChannelCanvas) | **Delete** /v2/canvases/{canvasId} | 
[**MezonDeleteChannelDesc**](MezonApi.md#MezonDeleteChannelDesc) | **Delete** /v2/channeldesc/{channelId} | Delete a channel by ID.
[**MezonDeleteClanDesc**](MezonApi.md#MezonDeleteClanDesc) | **Delete** /v2/clandesc/{clanDescId} | Delete a clan desc by ID.
[**MezonDeleteClanStickerById**](MezonApi.md#MezonDeleteClanStickerById) | **Delete** /v2/sticker/{id} | Delete a sticker by ID
[**MezonDeleteClanWebhookById**](MezonApi.md#MezonDeleteClanWebhookById) | **Delete** /v2/clanwebhooks/{id} | Disabled clan webhook.
[**MezonDeleteEvent**](MezonApi.md#MezonDeleteEvent) | **Delete** /v2/eventmanagement/{eventId} | Delete a event by ID.
[**MezonDeleteFriends**](MezonApi.md#MezonDeleteFriends) | **Delete** /v2/friend | Delete one or more users by ID or username.
[**MezonDeleteNotiReactMessage**](MezonApi.md#MezonDeleteNotiReactMessage) | **Delete** /v2/notifireactmessage/delete | 
[**MezonDeleteNotificationCategorySetting**](MezonApi.md#MezonDeleteNotificationCategorySetting) | **Delete** /v2/notificationusercategory/delete | 
[**MezonDeleteNotificationChannel**](MezonApi.md#MezonDeleteNotificationChannel) | **Delete** /v2/notificationuserchannel/delete | 
[**MezonDeleteNotifications**](MezonApi.md#MezonDeleteNotifications) | **Delete** /v2/notification | Delete one or more notifications for the current user.
[**MezonDeleteOnboarding**](MezonApi.md#MezonDeleteOnboarding) | **Delete** /v2/onboarding/{id} | delete onboarding.
[**MezonDeletePinMessage**](MezonApi.md#MezonDeletePinMessage) | **Delete** /v2/pinmessage/delete | 
[**MezonDeleteRole**](MezonApi.md#MezonDeleteRole) | **Delete** /v2/roles/{roleId} | Delete a role by ID.
[**MezonDeleteRoleChannelDesc**](MezonApi.md#MezonDeleteRoleChannelDesc) | **Put** /v2/rolechannel/delete | Update a role when Delete a role by ID.
[**MezonDeleteSystemMessage**](MezonApi.md#MezonDeleteSystemMessage) | **Delete** /v2/systemmessages/{clanId} | Delete a specific system messages.
[**MezonDeleteWebhookById**](MezonApi.md#MezonDeleteWebhookById) | **Patch** /v2/webhooks/{id} | disabled webhook
[**MezonEditChannelCanvases**](MezonApi.md#MezonEditChannelCanvases) | **Post** /v2/canvases/editor | Channel canvas editor
[**MezonEvent**](MezonApi.md#MezonEvent) | **Post** /v2/event | Submit an event for processing in the server&#39;s registered runtime custom events handler.
[**MezonGenerateClanWebhook**](MezonApi.md#MezonGenerateClanWebhook) | **Post** /v2/clanwebhooks | Generate clan webhook.
[**MezonGenerateWebhook**](MezonApi.md#MezonGenerateWebhook) | **Post** /v2/webhooks/generate | create webhook
[**MezonGetAccount**](MezonApi.md#MezonGetAccount) | **Get** /v2/account | Fetch the current user&#39;s account.
[**MezonGetApp**](MezonApi.md#MezonGetApp) | **Get** /v2/apps/app/{id} | Get detailed app information.
[**MezonGetChanEncryptionMethod**](MezonApi.md#MezonGetChanEncryptionMethod) | **Get** /v2/channel/{channelId}/encrypt_method | get channel encryption method
[**MezonGetChannelCanvasDetail**](MezonApi.md#MezonGetChannelCanvasDetail) | **Get** /v2/canvases/{id} | 
[**MezonGetChannelCanvasList**](MezonApi.md#MezonGetChannelCanvasList) | **Get** /v2/channel-canvases/{channelId} | 
[**MezonGetChannelCategoryNotiSettingsList**](MezonApi.md#MezonGetChannelCategoryNotiSettingsList) | **Get** /v2/getChannelCategoryNotiSettingsList | List GetChannelCategoryNotiSettingsList
[**MezonGetClanDescProfile**](MezonApi.md#MezonGetClanDescProfile) | **Get** /v2/clandescprofile/{clanId} | Get a clan desc profile
[**MezonGetKeyServer**](MezonApi.md#MezonGetKeyServer) | **Get** /v2/e2ee/key_server | get key server
[**MezonGetLinkInvite**](MezonApi.md#MezonGetLinkInvite) | **Get** /v2/invite/{inviteId} | Add users to a channel.
[**MezonGetListEmojisByUserId**](MezonApi.md#MezonGetListEmojisByUserId) | **Get** /v2/emojis | get list emoji by user id
[**MezonGetListFavoriteChannel**](MezonApi.md#MezonGetListFavoriteChannel) | **Get** /v2/channel/favorite/{clanId} | 
[**MezonGetListPermission**](MezonApi.md#MezonGetListPermission) | **Get** /v2/permissions | Get permission list
[**MezonGetListStickersByUserId**](MezonApi.md#MezonGetListStickersByUserId) | **Get** /v2/stickers | get list sticker by user id
[**MezonGetNotificationCategory**](MezonApi.md#MezonGetNotificationCategory) | **Get** /v2/getnotificationcategory | List GetNotificationChannel
[**MezonGetNotificationChannel**](MezonApi.md#MezonGetNotificationChannel) | **Get** /v2/getnotificationchannel | List GetNotificationChannel
[**MezonGetNotificationClan**](MezonApi.md#MezonGetNotificationClan) | **Get** /v2/getnotificationclan | List GetNotificationClan
[**MezonGetNotificationReactMessage**](MezonApi.md#MezonGetNotificationReactMessage) | **Get** /v2/getnotificationreactmessage | List GetNotificationReactMessage
[**MezonGetOnboardingDetail**](MezonApi.md#MezonGetOnboardingDetail) | **Get** /v2/onboarding/{id} | get detailed onboarding information.
[**MezonGetPermissionByRoleIdChannelId**](MezonApi.md#MezonGetPermissionByRoleIdChannelId) | **Get** /v2/permissions/roles/channels/users | GetPermissionByRoleIdChannelId
[**MezonGetPinMessagesList**](MezonApi.md#MezonGetPinMessagesList) | **Get** /v2/pinmessage/get | 
[**MezonGetPubKeys**](MezonApi.md#MezonGetPubKeys) | **Get** /v2/pubkey | get pubkey
[**MezonGetRoleOfUserInTheClan**](MezonApi.md#MezonGetRoleOfUserInTheClan) | **Get** /v2/roleuserinclan/{clanId} | 
[**MezonGetSystemMessageByClanId**](MezonApi.md#MezonGetSystemMessageByClanId) | **Get** /v2/systemmessages/{clanId} | Get details of a specific system messages.
[**MezonGetSystemMessagesList**](MezonApi.md#MezonGetSystemMessagesList) | **Get** /v2/systemmessages | Get the list of system messages.
[**MezonGetUserProfileOnClan**](MezonApi.md#MezonGetUserProfileOnClan) | **Get** /v2/getclanprofile/{clanId} | 
[**MezonGetUserStatus**](MezonApi.md#MezonGetUserStatus) | **Get** /v2/userstatus | Get user status
[**MezonGetUsers**](MezonApi.md#MezonGetUsers) | **Get** /v2/user | Fetch zero or more users by ID and/or username.
[**MezonGiveMeACoffee**](MezonApi.md#MezonGiveMeACoffee) | **Post** /v2/givecoffee | Give a coffee
[**MezonHandlerClanWebhook**](MezonApi.md#MezonHandlerClanWebhook) | **Post** /v2/clanwebhooks/{token}/{username} | Handler clan webhook.
[**MezonHandlerWebhook**](MezonApi.md#MezonHandlerWebhook) | **Post** /v2/webhooks/{channelId}/{token} | webhook
[**MezonHashtagDMList**](MezonApi.md#MezonHashtagDMList) | **Get** /v2/hashtagdmlist | List HashtagDMList
[**MezonHealthcheck**](MezonApi.md#MezonHealthcheck) | **Get** /healthcheck | A healthcheck which load balancers can use to check the service.
[**MezonImportFacebookFriends**](MezonApi.md#MezonImportFacebookFriends) | **Post** /v2/friend/facebook | Import Facebook friends and add them to a user&#39;s account.
[**MezonImportSteamFriends**](MezonApi.md#MezonImportSteamFriends) | **Post** /v2/friend/steam | Import Steam friends and add them to a user&#39;s account.
[**MezonInviteUser**](MezonApi.md#MezonInviteUser) | **Post** /v2/invite/{inviteId} | Add users to a channel.
[**MezonLeaveThread**](MezonApi.md#MezonLeaveThread) | **Post** /v2/channel/{channelId}/leave | Leave a channel the user is a member of.
[**MezonLinkApple**](MezonApi.md#MezonLinkApple) | **Post** /v2/account/link/apple | Add an Apple ID to the social profiles on the current user&#39;s account.
[**MezonLinkCustom**](MezonApi.md#MezonLinkCustom) | **Post** /v2/account/link/custom | Add a custom ID to the social profiles on the current user&#39;s account.
[**MezonLinkDevice**](MezonApi.md#MezonLinkDevice) | **Post** /v2/account/link/device | Add a device ID to the social profiles on the current user&#39;s account.
[**MezonLinkEmail**](MezonApi.md#MezonLinkEmail) | **Post** /v2/account/link/email | Add an email+password to the social profiles on the current user&#39;s account.
[**MezonLinkFacebook**](MezonApi.md#MezonLinkFacebook) | **Post** /v2/account/link/facebook | Add Facebook to the social profiles on the current user&#39;s account.
[**MezonLinkFacebookInstantGame**](MezonApi.md#MezonLinkFacebookInstantGame) | **Post** /v2/account/link/facebookinstantgame | Add Facebook Instant Game to the social profiles on the current user&#39;s account.
[**MezonLinkGameCenter**](MezonApi.md#MezonLinkGameCenter) | **Post** /v2/account/link/gamecenter | Add Apple&#39;s GameCenter to the social profiles on the current user&#39;s account.
[**MezonLinkGoogle**](MezonApi.md#MezonLinkGoogle) | **Post** /v2/account/link/google | Add Google to the social profiles on the current user&#39;s account.
[**MezonLinkSteam**](MezonApi.md#MezonLinkSteam) | **Post** /v2/account/link/steam | Add Steam to the social profiles on the current user&#39;s account.
[**MezonListActivity**](MezonApi.md#MezonListActivity) | **Get** /v2/activity | List activity
[**MezonListApps**](MezonApi.md#MezonListApps) | **Get** /v2/apps/app | List (and optionally filter) accounts.
[**MezonListAuditLog**](MezonApi.md#MezonListAuditLog) | **Get** /v2/audit_log | 
[**MezonListCategoryDescs**](MezonApi.md#MezonListCategoryDescs) | **Get** /v2/categorydesc/{clanId} | 
[**MezonListChannelApps**](MezonApi.md#MezonListChannelApps) | **Get** /v2/channel-apps | List channel apps.
[**MezonListChannelAttachment**](MezonApi.md#MezonListChannelAttachment) | **Get** /v2/channel/{channelId}/attachment | List all attachment that are part of a channel.
[**MezonListChannelByUserId**](MezonApi.md#MezonListChannelByUserId) | **Get** /v2/listchannelbyuserid | List HashtagDMList
[**MezonListChannelDescs**](MezonApi.md#MezonListChannelDescs) | **Get** /v2/channeldesc | List user channels
[**MezonListChannelMessages**](MezonApi.md#MezonListChannelMessages) | **Get** /v2/channel/{channelId} | List a channel&#39;s message history.
[**MezonListChannelSetting**](MezonApi.md#MezonListChannelSetting) | **Get** /v2/channelsetting/{clanId} | List channel setting
[**MezonListChannelUsers**](MezonApi.md#MezonListChannelUsers) | **Get** /v2/channel/{channelId}/user | List all users that are part of a channel.
[**MezonListChannelVoiceUsers**](MezonApi.md#MezonListChannelVoiceUsers) | **Get** /v2/channelvoice | List all users that are part of a channel.
[**MezonListClanDescs**](MezonApi.md#MezonListClanDescs) | **Get** /v2/clandesc | List clans
[**MezonListClanUsers**](MezonApi.md#MezonListClanUsers) | **Get** /v2/clandesc/{clanId}/user | List all users that are part of a clan.
[**MezonListClanWebhook**](MezonApi.md#MezonListClanWebhook) | **Get** /v2/clanwebhooks/{clanId} | List clan webhook.
[**MezonListEvents**](MezonApi.md#MezonListEvents) | **Get** /v2/eventmanagement | List user events
[**MezonListFriends**](MezonApi.md#MezonListFriends) | **Get** /v2/friend | List all friends for the current user.
[**MezonListNotifications**](MezonApi.md#MezonListNotifications) | **Get** /v2/notification | Fetch list of notifications.
[**MezonListOnboarding**](MezonApi.md#MezonListOnboarding) | **Get** /v2/onboarding | list onboarding.
[**MezonListOnboardingStep**](MezonApi.md#MezonListOnboardingStep) | **Get** /v2/onboardingsteps | List onboarding step.
[**MezonListPTTChannelUsers**](MezonApi.md#MezonListPTTChannelUsers) | **Get** /v2/ptt_channels/users | List all users in ptt channel.
[**MezonListRolePermissions**](MezonApi.md#MezonListRolePermissions) | **Get** /v2/roles/{roleId}/permissions | List role permissions
[**MezonListRoleUsers**](MezonApi.md#MezonListRoleUsers) | **Get** /v2/roles/{roleId}/users | List role permissions
[**MezonListRoles**](MezonApi.md#MezonListRoles) | **Get** /v2/roles | ListRoles
[**MezonListStreamingChannelUsers**](MezonApi.md#MezonListStreamingChannelUsers) | **Get** /v2/streaming-channels/users | List all users that are part of a channel.
[**MezonListStreamingChannels**](MezonApi.md#MezonListStreamingChannels) | **Get** /v2/streaming-channels | List streaming channels.
[**MezonListThreadDescs**](MezonApi.md#MezonListThreadDescs) | **Get** /v2/thread/{channelId} | List user channels
[**MezonListUserClansByUserId**](MezonApi.md#MezonListUserClansByUserId) | **Get** /v2/users/clans | ListUserClansByUserId
[**MezonListUserPermissionInChannel**](MezonApi.md#MezonListUserPermissionInChannel) | **Get** /v2/users/clans/channels | ListUserPermissionInChannel
[**MezonListUsersAddChannelByChannelId**](MezonApi.md#MezonListUsersAddChannelByChannelId) | **Get** /v2/channeldesc/users/add | list user add channel by channel ids
[**MezonListWalletLedger**](MezonApi.md#MezonListWalletLedger) | **Get** /v2/walletledger | Get user status
[**MezonListWebhookByChannelId**](MezonApi.md#MezonListWebhookByChannelId) | **Get** /v2/webhooks/{channelId} | list webhook belong to the channel
[**MezonMarkAsRead**](MezonApi.md#MezonMarkAsRead) | **Post** /v2/markasread | Mark as read
[**MezonOpenDirectMess**](MezonApi.md#MezonOpenDirectMess) | **Put** /v2/direct/open | open direct message.
[**MezonPushPubKey**](MezonApi.md#MezonPushPubKey) | **Post** /v2/pubkey/push | store pubkey for e2ee
[**MezonRegistFCMDeviceToken**](MezonApi.md#MezonRegistFCMDeviceToken) | **Post** /v2/devicetoken | regist fcm device token
[**MezonRegisterStreamingChannel**](MezonApi.md#MezonRegisterStreamingChannel) | **Post** /v2/streaming-channels | Register streaming in channel ( for bot - get streaming key)
[**MezonRegistrationEmail**](MezonApi.md#MezonRegistrationEmail) | **Post** /v2/account/registry | Authenticate a user with an email+password against the server.
[**MezonRemoveChannelFavorite**](MezonApi.md#MezonRemoveChannelFavorite) | **Delete** /v2/channel/favorite/{channelId} | 
[**MezonRemoveChannelUsers**](MezonApi.md#MezonRemoveChannelUsers) | **Post** /v2/channel/{channelId}/remove | Kick a set of users from a channel.
[**MezonRemoveClanUsers**](MezonApi.md#MezonRemoveClanUsers) | **Post** /v2/clandesc/{clanId}/kick | Kick a set of users from a clan.
[**MezonRpcFunc**](MezonApi.md#MezonRpcFunc) | **Post** /v2/rpc/{id} | Execute a Lua function on the server.
[**MezonRpcFunc2**](MezonApi.md#MezonRpcFunc2) | **Get** /v2/rpc/{id} | Execute a Lua function on the server.
[**MezonSearchMessage**](MezonApi.md#MezonSearchMessage) | **Post** /v2/es/search | Search message from elasticsearch service.
[**MezonSendToken**](MezonApi.md#MezonSendToken) | **Post** /v2/sendtoken | UpdateWallets
[**MezonSessionLogout**](MezonApi.md#MezonSessionLogout) | **Post** /v2/session/logout | Log out a session, invalidate a refresh token, or log out all sessions/refresh tokens for a user.
[**MezonSessionRefresh**](MezonApi.md#MezonSessionRefresh) | **Post** /v2/account/session/refresh | Refresh a user&#39;s session using a refresh token retrieved from a previous authentication request.
[**MezonSetChanEncryptionMethod**](MezonApi.md#MezonSetChanEncryptionMethod) | **Post** /v2/channel/{channelId}/encrypt_method | store channel encryption method
[**MezonSetMuteNotificationCategory**](MezonApi.md#MezonSetMuteNotificationCategory) | **Post** /v2/mutenotificationcategory/set | set mute notification user channel.
[**MezonSetMuteNotificationChannel**](MezonApi.md#MezonSetMuteNotificationChannel) | **Post** /v2/mutenotificationchannel/set | set mute notification user channel.
[**MezonSetNotificationCategorySetting**](MezonApi.md#MezonSetNotificationCategorySetting) | **Post** /v2/notificationucategory/set | set notification user channel.
[**MezonSetNotificationChannelSetting**](MezonApi.md#MezonSetNotificationChannelSetting) | **Post** /v2/notificationchannel/set | set notification user channel.
[**MezonSetNotificationClanSetting**](MezonApi.md#MezonSetNotificationClanSetting) | **Post** /v2/notificationclan/set | set notification user channel.
[**MezonSetNotificationReactMessage**](MezonApi.md#MezonSetNotificationReactMessage) | **Post** /v2/notifireactmessage/set | 
[**MezonSetRoleChannelPermission**](MezonApi.md#MezonSetRoleChannelPermission) | **Post** /v2/permissionrolechannel/set | set permission role channel.
[**MezonStreamingServerCallback**](MezonApi.md#MezonStreamingServerCallback) | **Post** /v2/ossrs/callback | Ossrs http callback.
[**MezonUnbanApp**](MezonApi.md#MezonUnbanApp) | **Post** /v2/apps/app/{id}/unban | Unban an app.
[**MezonUnlinkApple**](MezonApi.md#MezonUnlinkApple) | **Post** /v2/account/unlink/apple | Remove the Apple ID from the social profiles on the current user&#39;s account.
[**MezonUnlinkCustom**](MezonApi.md#MezonUnlinkCustom) | **Post** /v2/account/unlink/custom | Remove the custom ID from the social profiles on the current user&#39;s account.
[**MezonUnlinkDevice**](MezonApi.md#MezonUnlinkDevice) | **Post** /v2/account/unlink/device | Remove the device ID from the social profiles on the current user&#39;s account.
[**MezonUnlinkEmail**](MezonApi.md#MezonUnlinkEmail) | **Post** /v2/account/unlink/email | Remove the email+password from the social profiles on the current user&#39;s account.
[**MezonUnlinkFacebook**](MezonApi.md#MezonUnlinkFacebook) | **Post** /v2/account/unlink/facebook | Remove Facebook from the social profiles on the current user&#39;s account.
[**MezonUnlinkFacebookInstantGame**](MezonApi.md#MezonUnlinkFacebookInstantGame) | **Post** /v2/account/unlink/facebookinstantgame | Remove Facebook Instant Game profile from the social profiles on the current user&#39;s account.
[**MezonUnlinkGameCenter**](MezonApi.md#MezonUnlinkGameCenter) | **Post** /v2/account/unlink/gamecenter | Remove Apple&#39;s GameCenter from the social profiles on the current user&#39;s account.
[**MezonUnlinkGoogle**](MezonApi.md#MezonUnlinkGoogle) | **Post** /v2/account/unlink/google | Remove Google from the social profiles on the current user&#39;s account.
[**MezonUnlinkSteam**](MezonApi.md#MezonUnlinkSteam) | **Post** /v2/account/unlink/steam | Remove Steam from the social profiles on the current user&#39;s account.
[**MezonUpdateAccount**](MezonApi.md#MezonUpdateAccount) | **Put** /v2/account | Update fields in the current user&#39;s account.
[**MezonUpdateApp**](MezonApi.md#MezonUpdateApp) | **Put** /v2/apps/app/{id} | Update one or more fields on a app.
[**MezonUpdateCategory**](MezonApi.md#MezonUpdateCategory) | **Put** /v2/updatecategory/{clanId} | Update fields in a given category.
[**MezonUpdateCategoryOrder**](MezonApi.md#MezonUpdateCategoryOrder) | **Put** /v2/category/orders | 
[**MezonUpdateChannelDesc**](MezonApi.md#MezonUpdateChannelDesc) | **Put** /v2/channeldesc/{channelId} | Update fields in a given channel.
[**MezonUpdateChannelPrivate**](MezonApi.md#MezonUpdateChannelPrivate) | **Put** /v2/updatechannelprivate | Update channel private.
[**MezonUpdateClanDesc**](MezonApi.md#MezonUpdateClanDesc) | **Put** /v2/clandesc/{clanId} | Update fields in a given clan.
[**MezonUpdateClanDescProfile**](MezonApi.md#MezonUpdateClanDescProfile) | **Put** /v2/clandescprofile/{clanId} | Update fields in a given clan profile.
[**MezonUpdateClanEmojiById**](MezonApi.md#MezonUpdateClanEmojiById) | **Patch** /v2/emoji/{id} | Update ClanEmoj By id
[**MezonUpdateClanStickerById**](MezonApi.md#MezonUpdateClanStickerById) | **Patch** /v2/sticker/{id} | Update a sticker by ID
[**MezonUpdateClanWebhookById**](MezonApi.md#MezonUpdateClanWebhookById) | **Put** /v2/clanwebhooks/{id} | Update clan webhook by id.
[**MezonUpdateEvent**](MezonApi.md#MezonUpdateEvent) | **Put** /v2/eventmanagement/{eventId} | Update fields in a given event.
[**MezonUpdateEventUser**](MezonApi.md#MezonUpdateEventUser) | **Put** /v2/eventmanagement/user | Update fields in a given event.
[**MezonUpdateOnboarding**](MezonApi.md#MezonUpdateOnboarding) | **Put** /v2/onboarding/{id} | update onboarding.
[**MezonUpdateOnboardingStepByClanId**](MezonApi.md#MezonUpdateOnboardingStepByClanId) | **Put** /v2/onboardingsteps/{clanId} | Update onboarding step.
[**MezonUpdateRole**](MezonApi.md#MezonUpdateRole) | **Put** /v2/roles/{roleId} | Update fields in a given role.
[**MezonUpdateSystemMessage**](MezonApi.md#MezonUpdateSystemMessage) | **Put** /v2/systemmessages/{clanId} | Update a system messages.
[**MezonUpdateUser**](MezonApi.md#MezonUpdateUser) | **Put** /v2/user/update | 
[**MezonUpdateUserProfileByClan**](MezonApi.md#MezonUpdateUserProfileByClan) | **Put** /v2/updateclanprofile/{clanId} | 
[**MezonUpdateUserStatus**](MezonApi.md#MezonUpdateUserStatus) | **Put** /v2/userstatus | Update user status
[**MezonUpdateWebhookById**](MezonApi.md#MezonUpdateWebhookById) | **Patch** /v2/webhooks/update/{id} | update webhook name by id
[**MezonUploadAttachmentFile**](MezonApi.md#MezonUploadAttachmentFile) | **Post** /v2/uploadattachmentfile | Upload attachment
[**MezonWithdrawToken**](MezonApi.md#MezonWithdrawToken) | **Post** /v2/withdrawtoken | WithdrawToken


# **MezonAddApp**
> interface{} MezonAddApp(ctx, body)
Add a new apps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAddAppRequest**](ApiAddAppRequest.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonAddAppToClan**
> interface{} MezonAddAppToClan(ctx, appId, clanId)
Add an app to clan.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**| The unique identifier of the app. | 
  **clanId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonAddChannelFavorite**
> ApiAddFavoriteChannelResponse MezonAddChannelFavorite(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAddFavoriteChannelRequest**](ApiAddFavoriteChannelRequest.md)|  | 

### Return type

[**ApiAddFavoriteChannelResponse**](apiAddFavoriteChannelResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonAddChannelUsers**
> interface{} MezonAddChannelUsers(ctx, channelId, optional)
Add users to a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| The channel to add users to. | 
 **optional** | ***MezonAddChannelUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonAddChannelUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userIds** | [**optional.Interface of []string**](string.md)| The users to add. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonAddClanSticker**
> interface{} MezonAddClanSticker(ctx, body)
Add a new sticker

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiClanStickerAddRequest**](ApiClanStickerAddRequest.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonAddFriends**
> interface{} MezonAddFriends(ctx, optional)
Add friends by ID or username to a user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonAddFriendsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonAddFriendsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []string**](string.md)| The account id of a user. | 
 **usernames** | [**optional.Interface of []string**](string.md)| The account username of a user. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonAddRolesChannelDesc**
> interface{} MezonAddRolesChannelDesc(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAddRoleChannelDescRequest**](ApiAddRoleChannelDescRequest.md)| Add a role for channel. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonAuthenticate**
> ApiSession MezonAuthenticate(ctx, body)
Authenticate a app with a token against the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAuthenticateRequest**](ApiAuthenticateRequest.md)| Authenticate against the server with a device ID. | 

### Return type

[**ApiSession**](apiSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonAuthenticateApple**
> ApiSession MezonAuthenticateApple(ctx, account, optional)
Authenticate a user with an Apple ID against the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **account** | [**ApiAccountApple**](ApiAccountApple.md)| The Apple account details. | 
 **optional** | ***MezonAuthenticateAppleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonAuthenticateAppleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **create** | **optional.Bool**| Register the account if the user does not already exist. | 
 **username** | **optional.String**| Set the username on the account at register. Must be unique. | 

### Return type

[**ApiSession**](apiSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonAuthenticateCustom**
> ApiSession MezonAuthenticateCustom(ctx, account, optional)
Authenticate a user with a custom id against the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **account** | [**ApiAccountCustom**](ApiAccountCustom.md)| The custom account details. | 
 **optional** | ***MezonAuthenticateCustomOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonAuthenticateCustomOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **create** | **optional.Bool**| Register the account if the user does not already exist. | 
 **username** | **optional.String**| Set the username on the account at register. Must be unique. | 

### Return type

[**ApiSession**](apiSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonAuthenticateDevice**
> ApiSession MezonAuthenticateDevice(ctx, account, optional)
Authenticate a user with a device id against the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **account** | [**ApiAccountDevice**](ApiAccountDevice.md)| The device account details. | 
 **optional** | ***MezonAuthenticateDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonAuthenticateDeviceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **create** | **optional.Bool**| Register the account if the user does not already exist. | 
 **username** | **optional.String**| Set the username on the account at register. Must be unique. | 

### Return type

[**ApiSession**](apiSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonAuthenticateEmail**
> ApiSession MezonAuthenticateEmail(ctx, account, optional)
Authenticate a user with an email+password against the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **account** | [**ApiAccountEmail**](ApiAccountEmail.md)| The email account details. | 
 **optional** | ***MezonAuthenticateEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonAuthenticateEmailOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **create** | **optional.Bool**| Register the account if the user does not already exist. | 
 **username** | **optional.String**| Set the username on the account at register. Must be unique. | 

### Return type

[**ApiSession**](apiSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonAuthenticateFacebook**
> ApiSession MezonAuthenticateFacebook(ctx, account, optional)
Authenticate a user with a Facebook OAuth token against the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **account** | [**ApiAccountFacebook**](ApiAccountFacebook.md)| The Facebook account details. | 
 **optional** | ***MezonAuthenticateFacebookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonAuthenticateFacebookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **create** | **optional.Bool**| Register the account if the user does not already exist. | 
 **username** | **optional.String**| Set the username on the account at register. Must be unique. | 
 **sync** | **optional.Bool**| Import Facebook friends for the user. | 

### Return type

[**ApiSession**](apiSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonAuthenticateFacebookInstantGame**
> ApiSession MezonAuthenticateFacebookInstantGame(ctx, account, optional)
Authenticate a user with a Facebook Instant Game token against the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **account** | [**ApiAccountFacebookInstantGame**](ApiAccountFacebookInstantGame.md)| The Facebook Instant Game account details. | 
 **optional** | ***MezonAuthenticateFacebookInstantGameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonAuthenticateFacebookInstantGameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **create** | **optional.Bool**| Register the account if the user does not already exist. | 
 **username** | **optional.String**| Set the username on the account at register. Must be unique. | 

### Return type

[**ApiSession**](apiSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonAuthenticateGameCenter**
> ApiSession MezonAuthenticateGameCenter(ctx, account, optional)
Authenticate a user with Apple's GameCenter against the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **account** | [**ApiAccountGameCenter**](ApiAccountGameCenter.md)| The Game Center account details. | 
 **optional** | ***MezonAuthenticateGameCenterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonAuthenticateGameCenterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **create** | **optional.Bool**| Register the account if the user does not already exist. | 
 **username** | **optional.String**| Set the username on the account at register. Must be unique. | 

### Return type

[**ApiSession**](apiSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonAuthenticateGoogle**
> ApiSession MezonAuthenticateGoogle(ctx, account, optional)
Authenticate a user with Google against the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **account** | [**ApiAccountGoogle**](ApiAccountGoogle.md)| The Google account details. | 
 **optional** | ***MezonAuthenticateGoogleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonAuthenticateGoogleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **create** | **optional.Bool**| Register the account if the user does not already exist. | 
 **username** | **optional.String**| Set the username on the account at register. Must be unique. | 

### Return type

[**ApiSession**](apiSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonAuthenticateGoogleRedirect**
> ApiSession MezonAuthenticateGoogleRedirect(ctx, optional)
Authenticate a user with Google against the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonAuthenticateGoogleRedirectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonAuthenticateGoogleRedirectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **optional.String**| The code | 

### Return type

[**ApiSession**](apiSession.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonAuthenticateSteam**
> ApiSession MezonAuthenticateSteam(ctx, account, optional)
Authenticate a user with Steam against the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **account** | [**ApiAccountSteam**](ApiAccountSteam.md)| The Steam account details. | 
 **optional** | ***MezonAuthenticateSteamOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonAuthenticateSteamOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **create** | **optional.Bool**| Register the account if the user does not already exist. | 
 **username** | **optional.String**| Set the username on the account at register. Must be unique. | 
 **sync** | **optional.Bool**| Import Steam friends for the user. | 

### Return type

[**ApiSession**](apiSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonBanApp**
> interface{} MezonBanApp(ctx, id)
Ban a app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The unique identifier of the app. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonBlockFriends**
> interface{} MezonBlockFriends(ctx, optional)
Block one or more users by ID or username.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonBlockFriendsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonBlockFriendsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []string**](string.md)| The account id of a user. | 
 **usernames** | [**optional.Interface of []string**](string.md)| The account username of a user. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonChangeChannelCategory**
> interface{} MezonChangeChannelCategory(ctx, newCategoryId, body)
update the category of a channel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newCategoryId** | **string**|  | 
  **body** | [**MezonChangeChannelCategoryBody**](MezonChangeChannelCategoryBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonCheckDuplicateClanName**
> ApiCheckDuplicateClanNameResponse MezonCheckDuplicateClanName(ctx, clanName)
check duplicate clan name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanName** | **string**|  | 

### Return type

[**ApiCheckDuplicateClanNameResponse**](apiCheckDuplicateClanNameResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonCheckLoginRequest**
> ApiSession MezonCheckLoginRequest(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiConfirmLoginRequest**](ApiConfirmLoginRequest.md)|  | 

### Return type

[**ApiSession**](apiSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonCloseDirectMess**
> interface{} MezonCloseDirectMess(ctx, body)
close direct message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiDeleteChannelDescRequest**](ApiDeleteChannelDescRequest.md)| Delete a channel the user has access to. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonConfirmLogin**
> interface{} MezonConfirmLogin(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiConfirmLoginRequest**](ApiConfirmLoginRequest.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonCreateActiviy**
> ApiUserActivity MezonCreateActiviy(ctx, body)
Create user activity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiCreateActivityRequest**](ApiCreateActivityRequest.md)|  | 

### Return type

[**ApiUserActivity**](apiUserActivity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonCreateCategoryDesc**
> ApiCategoryDesc MezonCreateCategoryDesc(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiCreateCategoryDescRequest**](ApiCreateCategoryDescRequest.md)|  | 

### Return type

[**ApiCategoryDesc**](apiCategoryDesc.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonCreateChannelDesc**
> ApiChannelDescription MezonCreateChannelDesc(ctx, body)
Create a new channel with the current user as the owner.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiCreateChannelDescRequest**](ApiCreateChannelDescRequest.md)| Create a channel within clan. | 

### Return type

[**ApiChannelDescription**](apiChannelDescription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonCreateClanDesc**
> ApiClanDesc MezonCreateClanDesc(ctx, body)
Create a clan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiCreateClanDescRequest**](ApiCreateClanDescRequest.md)|  | 

### Return type

[**ApiClanDesc**](apiClanDesc.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonCreateClanEmoji**
> interface{} MezonCreateClanEmoji(ctx, body)
Post clan Emoji  /v2/emoji/create

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiClanEmojiCreateRequest**](ApiClanEmojiCreateRequest.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonCreateEvent**
> ApiEventManagement MezonCreateEvent(ctx, body)
Create a new event for clan.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiCreateEventRequest**](ApiCreateEventRequest.md)| Create a event within clan. | 

### Return type

[**ApiEventManagement**](apiEventManagement.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonCreateLinkInviteUser**
> ApiLinkInviteUser MezonCreateLinkInviteUser(ctx, body)
Add users to a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiLinkInviteUserRequest**](ApiLinkInviteUserRequest.md)| Add link invite users to. | 

### Return type

[**ApiLinkInviteUser**](apiLinkInviteUser.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonCreateOnboarding**
> interface{} MezonCreateOnboarding(ctx, body)
create onboarding.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiCreateOnboardingRequest**](ApiCreateOnboardingRequest.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonCreatePinMessage**
> ApiChannelMessageHeader MezonCreatePinMessage(ctx, body)
set notification user channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiPinMessageRequest**](ApiPinMessageRequest.md)|  | 

### Return type

[**ApiChannelMessageHeader**](apiChannelMessageHeader.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonCreateQRLogin**
> ApiLoginIdResponse MezonCreateQRLogin(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiLoginRequest**](ApiLoginRequest.md)|  | 

### Return type

[**ApiLoginIdResponse**](apiLoginIDResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonCreateRole**
> ApiRole MezonCreateRole(ctx, body)
Create a new role for clan.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiCreateRoleRequest**](ApiCreateRoleRequest.md)| Create a role within clan. | 

### Return type

[**ApiRole**](apiRole.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonCreateSystemMessage**
> interface{} MezonCreateSystemMessage(ctx, body)
Create a system messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiSystemMessageRequest**](ApiSystemMessageRequest.md)| Request to get system message by clan and channel IDs. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteAccount**
> interface{} MezonDeleteAccount(ctx, )
Delete the current user's account.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteApp**
> interface{} MezonDeleteApp(ctx, id, optional)
Delete all information stored for an app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The unique identifier of the app. | 
 **optional** | ***MezonDeleteAppOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonDeleteAppOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recordDeletion** | **optional.Bool**| Record the app deletion - used for GDPR compliance. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteByIdClanEmoji**
> interface{} MezonDeleteByIdClanEmoji(ctx, id, optional)
Delete a emoji by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***MezonDeleteByIdClanEmojiOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonDeleteByIdClanEmojiOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clanId** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteCategoryDesc**
> interface{} MezonDeleteCategoryDesc(ctx, categoryId, clanId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryId** | **string**|  | 
  **clanId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteCategoryOrder**
> interface{} MezonDeleteCategoryOrder(ctx, clanId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteChannelCanvas**
> interface{} MezonDeleteChannelCanvas(ctx, canvasId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **canvasId** | **string**| canvas id | 
 **optional** | ***MezonDeleteChannelCanvasOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonDeleteChannelCanvasOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clanId** | **optional.String**| clan id | 
 **channelId** | **optional.String**| channel id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteChannelDesc**
> interface{} MezonDeleteChannelDesc(ctx, channelId)
Delete a channel by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| The id of a channel. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteClanDesc**
> interface{} MezonDeleteClanDesc(ctx, clanDescId)
Delete a clan desc by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanDescId** | **string**| The id of a group. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteClanStickerById**
> interface{} MezonDeleteClanStickerById(ctx, id, optional)
Delete a sticker by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***MezonDeleteClanStickerByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonDeleteClanStickerByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clanId** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteClanWebhookById**
> interface{} MezonDeleteClanWebhookById(ctx, id, optional)
Disabled clan webhook.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| id. | 
 **optional** | ***MezonDeleteClanWebhookByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonDeleteClanWebhookByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clanId** | **optional.String**| clan id. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteEvent**
> interface{} MezonDeleteEvent(ctx, eventId, optional)
Delete a event by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventId** | **string**| The id of a event. | 
 **optional** | ***MezonDeleteEventOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonDeleteEventOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clanId** | **optional.String**| clan id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteFriends**
> interface{} MezonDeleteFriends(ctx, optional)
Delete one or more users by ID or username.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonDeleteFriendsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonDeleteFriendsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []string**](string.md)| The account id of a user. | 
 **usernames** | [**optional.Interface of []string**](string.md)| The account username of a user. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteNotiReactMessage**
> interface{} MezonDeleteNotiReactMessage(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonDeleteNotiReactMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonDeleteNotiReactMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteNotificationCategorySetting**
> interface{} MezonDeleteNotificationCategorySetting(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonDeleteNotificationCategorySettingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonDeleteNotificationCategorySettingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryId** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteNotificationChannel**
> interface{} MezonDeleteNotificationChannel(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonDeleteNotificationChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonDeleteNotificationChannelOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteNotifications**
> interface{} MezonDeleteNotifications(ctx, optional)
Delete one or more notifications for the current user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonDeleteNotificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonDeleteNotificationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []string**](string.md)| The id of notifications. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteOnboarding**
> interface{} MezonDeleteOnboarding(ctx, id, optional)
delete onboarding.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| id | 
 **optional** | ***MezonDeleteOnboardingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonDeleteOnboardingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clanId** | **optional.String**| clan id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeletePinMessage**
> interface{} MezonDeletePinMessage(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonDeletePinMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonDeletePinMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageId** | **optional.String**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteRole**
> interface{} MezonDeleteRole(ctx, roleId, optional)
Delete a role by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **string**| The id of a role. | 
 **optional** | ***MezonDeleteRoleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonDeleteRoleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channelId** | **optional.String**| The id of a channel | 
 **clanId** | **optional.String**| clan_id | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteRoleChannelDesc**
> interface{} MezonDeleteRoleChannelDesc(ctx, body)
Update a role when Delete a role by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiDeleteRoleRequest**](ApiDeleteRoleRequest.md)| Delete a role the user has access to. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteSystemMessage**
> interface{} MezonDeleteSystemMessage(ctx, clanId)
Delete a specific system messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanId** | **string**| Clan ID | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonDeleteWebhookById**
> interface{} MezonDeleteWebhookById(ctx, id, body)
disabled webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **body** | [**MezonDeleteWebhookByIdBody**](MezonDeleteWebhookByIdBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonEditChannelCanvases**
> ApiEditChannelCanvasResponse MezonEditChannelCanvases(ctx, body)
Channel canvas editor

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiEditChannelCanvasRequest**](ApiEditChannelCanvasRequest.md)|  | 

### Return type

[**ApiEditChannelCanvasResponse**](apiEditChannelCanvasResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonEvent**
> interface{} MezonEvent(ctx, body)
Submit an event for processing in the server's registered runtime custom events handler.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MezonapiEvent**](MezonapiEvent.md)| Represents an event to be passed through the server to registered event handlers. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGenerateClanWebhook**
> ApiGenerateClanWebhookResponse MezonGenerateClanWebhook(ctx, body)
Generate clan webhook.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiGenerateClanWebhookRequest**](ApiGenerateClanWebhookRequest.md)|  | 

### Return type

[**ApiGenerateClanWebhookResponse**](apiGenerateClanWebhookResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGenerateWebhook**
> ApiWebhookGenerateResponse MezonGenerateWebhook(ctx, body)
create webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiWebhookCreateRequest**](ApiWebhookCreateRequest.md)|  | 

### Return type

[**ApiWebhookGenerateResponse**](apiWebhookGenerateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetAccount**
> ApiAccount MezonGetAccount(ctx, )
Fetch the current user's account.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiAccount**](apiAccount.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetApp**
> ApiApp MezonGetApp(ctx, id)
Get detailed app information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The unique identifier of the app. | 

### Return type

[**ApiApp**](apiApp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetChanEncryptionMethod**
> ApiChanEncryptionMethod MezonGetChanEncryptionMethod(ctx, channelId, optional)
get channel encryption method

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**|  | 
 **optional** | ***MezonGetChanEncryptionMethodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonGetChanEncryptionMethodOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **method** | **optional.String**|  | 

### Return type

[**ApiChanEncryptionMethod**](apiChanEncryptionMethod.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetChannelCanvasDetail**
> ApiChannelCanvasDetailResponse MezonGetChannelCanvasDetail(ctx, id, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| id | 
 **optional** | ***MezonGetChannelCanvasDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonGetChannelCanvasDetailOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clanId** | **optional.String**| clan id | 
 **channelId** | **optional.String**| channel id | 

### Return type

[**ApiChannelCanvasDetailResponse**](apiChannelCanvasDetailResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetChannelCanvasList**
> ApiChannelCanvasListResponse MezonGetChannelCanvasList(ctx, channelId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| channel id | 
 **optional** | ***MezonGetChannelCanvasListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonGetChannelCanvasListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clanId** | **optional.String**| clan id | 
 **limit** | **optional.Int32**| limit | 
 **page** | **optional.Int32**| page | 

### Return type

[**ApiChannelCanvasListResponse**](apiChannelCanvasListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetChannelCategoryNotiSettingsList**
> ApiNotificationChannelCategorySettingList MezonGetChannelCategoryNotiSettingsList(ctx, optional)
List GetChannelCategoryNotiSettingsList

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonGetChannelCategoryNotiSettingsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonGetChannelCategoryNotiSettingsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clanId** | **optional.String**|  | 

### Return type

[**ApiNotificationChannelCategorySettingList**](apiNotificationChannelCategorySettingList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetClanDescProfile**
> ApiClanDescProfile MezonGetClanDescProfile(ctx, clanId)
Get a clan desc profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanId** | **string**| Clan id | 

### Return type

[**ApiClanDescProfile**](apiClanDescProfile.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetKeyServer**
> ApiGetKeyServerResp MezonGetKeyServer(ctx, )
get key server

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiGetKeyServerResp**](apiGetKeyServerResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetLinkInvite**
> ApiInviteUserRes MezonGetLinkInvite(ctx, inviteId)
Add users to a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inviteId** | **string**| id clan to add link to . | 

### Return type

[**ApiInviteUserRes**](apiInviteUserRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetListEmojisByUserId**
> ApiEmojiListedResponse MezonGetListEmojisByUserId(ctx, )
get list emoji by user id

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiEmojiListedResponse**](apiEmojiListedResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetListFavoriteChannel**
> ApiListFavoriteChannelResponse MezonGetListFavoriteChannel(ctx, clanId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanId** | **string**|  | 

### Return type

[**ApiListFavoriteChannelResponse**](apiListFavoriteChannelResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetListPermission**
> ApiPermissionList MezonGetListPermission(ctx, )
Get permission list

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiPermissionList**](apiPermissionList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetListStickersByUserId**
> ApiStickerListedResponse MezonGetListStickersByUserId(ctx, )
get list sticker by user id

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiStickerListedResponse**](apiStickerListedResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetNotificationCategory**
> ApiNotificationUserChannel MezonGetNotificationCategory(ctx, optional)
List GetNotificationChannel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonGetNotificationCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonGetNotificationCategoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryId** | **optional.String**|  | 

### Return type

[**ApiNotificationUserChannel**](apiNotificationUserChannel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetNotificationChannel**
> ApiNotificationUserChannel MezonGetNotificationChannel(ctx, optional)
List GetNotificationChannel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonGetNotificationChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonGetNotificationChannelOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **optional.String**|  | 

### Return type

[**ApiNotificationUserChannel**](apiNotificationUserChannel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetNotificationClan**
> ApiNotificationSetting MezonGetNotificationClan(ctx, optional)
List GetNotificationClan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonGetNotificationClanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonGetNotificationClanOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clanId** | **optional.String**|  | 

### Return type

[**ApiNotificationSetting**](apiNotificationSetting.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetNotificationReactMessage**
> ApiNotifiReactMessage MezonGetNotificationReactMessage(ctx, optional)
List GetNotificationReactMessage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonGetNotificationReactMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonGetNotificationReactMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **optional.String**|  | 

### Return type

[**ApiNotifiReactMessage**](apiNotifiReactMessage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetOnboardingDetail**
> ApiOnboardingItem MezonGetOnboardingDetail(ctx, id, optional)
get detailed onboarding information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| id | 
 **optional** | ***MezonGetOnboardingDetailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonGetOnboardingDetailOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clanId** | **optional.String**| clan id | 

### Return type

[**ApiOnboardingItem**](apiOnboardingItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetPermissionByRoleIdChannelId**
> ApiPermissionRoleChannelListEventResponse MezonGetPermissionByRoleIdChannelId(ctx, optional)
GetPermissionByRoleIdChannelId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonGetPermissionByRoleIdChannelIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonGetPermissionByRoleIdChannelIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleId** | **optional.String**| role id | 
 **channelId** | **optional.String**| channel id | 
 **userId** | **optional.String**| user id | 

### Return type

[**ApiPermissionRoleChannelListEventResponse**](apiPermissionRoleChannelListEventResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetPinMessagesList**
> ApiPinMessagesList MezonGetPinMessagesList(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonGetPinMessagesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonGetPinMessagesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageId** | **optional.String**|  | 
 **channelId** | **optional.String**|  | 
 **clanId** | **optional.String**|  | 

### Return type

[**ApiPinMessagesList**](apiPinMessagesList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetPubKeys**
> ApiGetPubKeysResponse MezonGetPubKeys(ctx, optional)
get pubkey

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonGetPubKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonGetPubKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIds** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**ApiGetPubKeysResponse**](apiGetPubKeysResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetRoleOfUserInTheClan**
> ApiRoleList MezonGetRoleOfUserInTheClan(ctx, clanId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanId** | **string**| clan_id. | 
 **optional** | ***MezonGetRoleOfUserInTheClanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonGetRoleOfUserInTheClanOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **channelId** | **optional.String**| channel_id | 

### Return type

[**ApiRoleList**](apiRoleList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetSystemMessageByClanId**
> ApiSystemMessage MezonGetSystemMessageByClanId(ctx, clanId)
Get details of a specific system messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanId** | **string**| Clan ID | 

### Return type

[**ApiSystemMessage**](apiSystemMessage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetSystemMessagesList**
> ApiSystemMessagesList MezonGetSystemMessagesList(ctx, )
Get the list of system messages.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiSystemMessagesList**](apiSystemMessagesList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetUserProfileOnClan**
> ApiClanProfile MezonGetUserProfileOnClan(ctx, clanId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanId** | **string**| id clanc | 

### Return type

[**ApiClanProfile**](apiClanProfile.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetUserStatus**
> ApiUserStatus MezonGetUserStatus(ctx, )
Get user status

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiUserStatus**](apiUserStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGetUsers**
> ApiUsers MezonGetUsers(ctx, optional)
Fetch zero or more users by ID and/or username.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonGetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonGetUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []string**](string.md)| The account id of a user. | 
 **usernames** | [**optional.Interface of []string**](string.md)| The account username of a user. | 
 **facebookIds** | [**optional.Interface of []string**](string.md)| The Facebook ID of a user. | 

### Return type

[**ApiUsers**](apiUsers.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonGiveMeACoffee**
> interface{} MezonGiveMeACoffee(ctx, body)
Give a coffee

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiGiveCoffeeEvent**](ApiGiveCoffeeEvent.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonHandlerClanWebhook**
> interface{} MezonHandlerClanWebhook(ctx, token, username, body)
Handler clan webhook.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| token. | 
  **username** | **string**| username. | 
  **body** | [**ApiClanWebhookHandlerBody**](ApiClanWebhookHandlerBody.md)| body. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonHandlerWebhook**
> interface{} MezonHandlerWebhook(ctx, channelId, token, body)
webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**|  | 
  **token** | **string**|  | 
  **body** | [**interface{}**](interface{}.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonHashtagDMList**
> ApiHashtagDmList MezonHashtagDMList(ctx, optional)
List HashtagDMList

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonHashtagDMListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonHashtagDMListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | [**optional.Interface of []string**](string.md)| user Id | 
 **limit** | **optional.Int32**| Max number of records to return. Between 1 and 100. | 

### Return type

[**ApiHashtagDmList**](apiHashtagDmList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonHealthcheck**
> interface{} MezonHealthcheck(ctx, )
A healthcheck which load balancers can use to check the service.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonImportFacebookFriends**
> interface{} MezonImportFacebookFriends(ctx, account, optional)
Import Facebook friends and add them to a user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **account** | [**ApiAccountFacebook**](ApiAccountFacebook.md)| The Facebook account details. | 
 **optional** | ***MezonImportFacebookFriendsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonImportFacebookFriendsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reset** | **optional.Bool**| Reset the current user&#39;s friends list. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonImportSteamFriends**
> interface{} MezonImportSteamFriends(ctx, account, optional)
Import Steam friends and add them to a user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **account** | [**ApiAccountSteam**](ApiAccountSteam.md)| The Facebook account details. | 
 **optional** | ***MezonImportSteamFriendsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonImportSteamFriendsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reset** | **optional.Bool**| Reset the current user&#39;s friends list. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonInviteUser**
> ApiInviteUserRes MezonInviteUser(ctx, inviteId)
Add users to a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inviteId** | **string**| id clan to add link to . | 

### Return type

[**ApiInviteUserRes**](apiInviteUserRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonLeaveThread**
> interface{} MezonLeaveThread(ctx, channelId)
Leave a channel the user is a member of.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| The channel ID to leave. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonLinkApple**
> interface{} MezonLinkApple(ctx, body)
Add an Apple ID to the social profiles on the current user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAccountApple**](ApiAccountApple.md)| Send a Apple Sign In token to the server. Used with authenticate/link/unlink. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonLinkCustom**
> interface{} MezonLinkCustom(ctx, body)
Add a custom ID to the social profiles on the current user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAccountCustom**](ApiAccountCustom.md)| Send a custom ID to the server. Used with authenticate/link/unlink. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonLinkDevice**
> interface{} MezonLinkDevice(ctx, body)
Add a device ID to the social profiles on the current user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAccountDevice**](ApiAccountDevice.md)| Send a device to the server. Used with authenticate/link/unlink and user. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonLinkEmail**
> interface{} MezonLinkEmail(ctx, body)
Add an email+password to the social profiles on the current user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAccountEmail**](ApiAccountEmail.md)| Send an email with password to the server. Used with authenticate/link/unlink. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonLinkFacebook**
> interface{} MezonLinkFacebook(ctx, account, optional)
Add Facebook to the social profiles on the current user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **account** | [**ApiAccountFacebook**](ApiAccountFacebook.md)| The Facebook account details. | 
 **optional** | ***MezonLinkFacebookOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonLinkFacebookOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sync** | **optional.Bool**| Import Facebook friends for the user. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonLinkFacebookInstantGame**
> interface{} MezonLinkFacebookInstantGame(ctx, body)
Add Facebook Instant Game to the social profiles on the current user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAccountFacebookInstantGame**](ApiAccountFacebookInstantGame.md)| Send a Facebook Instant Game token to the server. Used with authenticate/link/unlink. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonLinkGameCenter**
> interface{} MezonLinkGameCenter(ctx, body)
Add Apple's GameCenter to the social profiles on the current user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAccountGameCenter**](ApiAccountGameCenter.md)| Send Apple&#39;s Game Center account credentials to the server. Used with authenticate/link/unlink.  https://developer.apple.com/documentation/gamekit/gklocalplayer/1515407-generateidentityverificationsign | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonLinkGoogle**
> interface{} MezonLinkGoogle(ctx, body)
Add Google to the social profiles on the current user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAccountGoogle**](ApiAccountGoogle.md)| Send a Google token to the server. Used with authenticate/link/unlink. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonLinkSteam**
> interface{} MezonLinkSteam(ctx, body)
Add Steam to the social profiles on the current user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiLinkSteamRequest**](ApiLinkSteamRequest.md)| Link Steam to the current user&#39;s account. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListActivity**
> ApiListUserActivity MezonListActivity(ctx, )
List activity

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiListUserActivity**](apiListUserActivity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListApps**
> ApiAppList MezonListApps(ctx, optional)
List (and optionally filter) accounts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonListAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListAppsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| User ID or username filter. | 
 **tombstones** | **optional.Bool**| Search only recorded deletes. | 
 **cursor** | **optional.String**| Cursor to start from | 

### Return type

[**ApiAppList**](apiAppList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListAuditLog**
> MezonapiListAuditLog MezonListAuditLog(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonListAuditLogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListAuditLogOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **actionLog** | **optional.String**|  | 
 **userId** | **optional.String**|  | 
 **clanId** | **optional.String**|  | 
 **page** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 

### Return type

[**MezonapiListAuditLog**](mezonapiListAuditLog.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListCategoryDescs**
> ApiCategoryDescList MezonListCategoryDescs(ctx, clanId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanId** | **string**| the Clan that category belong to | 
 **optional** | ***MezonListCategoryDescsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListCategoryDescsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **creatorId** | **optional.String**| Category creator | 
 **categoryName** | **optional.String**| Category name | 
 **categoryId** | **optional.String**|  | 
 **categoryOrder** | **optional.Int32**|  | 

### Return type

[**ApiCategoryDescList**](apiCategoryDescList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListChannelApps**
> ApiListChannelAppsResponse MezonListChannelApps(ctx, optional)
List channel apps.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonListChannelAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListChannelAppsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clanId** | **optional.String**| clan id | 

### Return type

[**ApiListChannelAppsResponse**](apiListChannelAppsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListChannelAttachment**
> ApiChannelAttachmentList MezonListChannelAttachment(ctx, channelId, optional)
List all attachment that are part of a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| The channel ID to list from. | 
 **optional** | ***MezonListChannelAttachmentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListChannelAttachmentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clanId** | **optional.String**| The clan id | 
 **fileType** | **optional.String**| The channel type | 
 **limit** | **optional.Int32**| Max number of records to return. Between 1 and 100. | 
 **state** | **optional.Int32**| The group user state to list. | 
 **before** | **optional.String**| An optional previous id for page. | 
 **after** | **optional.String**| An optional next id for page. | 
 **around** | **optional.String**| An optional around id for page. | 

### Return type

[**ApiChannelAttachmentList**](apiChannelAttachmentList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListChannelByUserId**
> ApiChannelDescList MezonListChannelByUserId(ctx, )
List HashtagDMList

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiChannelDescList**](apiChannelDescList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListChannelDescs**
> ApiChannelDescList MezonListChannelDescs(ctx, optional)
List user channels

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonListChannelDescsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListChannelDescsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Max number of records to return. Between 1 and 100. | 
 **state** | **optional.Int32**| The channel state to list. | 
 **cursor** | **optional.String**| Cursor to start from | 
 **clanId** | **optional.String**| The clan of this channel | 
 **channelType** | **optional.Int32**| channel type | 

### Return type

[**ApiChannelDescList**](apiChannelDescList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListChannelMessages**
> ApiChannelMessageList MezonListChannelMessages(ctx, channelId, optional)
List a channel's message history.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| The channel ID to list from. | 
 **optional** | ***MezonListChannelMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListChannelMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clanId** | **optional.String**| the clan id | 
 **messageId** | **optional.String**| The current message ID. | 
 **limit** | **optional.Int32**| Max number of records to return. Between 1 and 100. | 
 **direction** | **optional.Int32**| True if listing should be older messages to newer, false if reverse. | 

### Return type

[**ApiChannelMessageList**](apiChannelMessageList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListChannelSetting**
> ApiChannelSettingListResponse MezonListChannelSetting(ctx, clanId, optional)
List channel setting

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanId** | **string**| clan id | 
 **optional** | ***MezonListChannelSettingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListChannelSettingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentId** | **optional.String**| parent id of thread | 
 **categoryId** | **optional.String**| category id | 
 **privateChannel** | **optional.Int32**| is private channel | 
 **active** | **optional.Int32**| is active | 
 **status** | **optional.Int32**| status | 
 **type_** | **optional.Int32**| type | 
 **limit** | **optional.Int32**| limit | 
 **page** | **optional.Int32**| page | 
 **channelLabel** | **optional.String**| channel label | 

### Return type

[**ApiChannelSettingListResponse**](apiChannelSettingListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListChannelUsers**
> ApiChannelUserList MezonListChannelUsers(ctx, channelId, optional)
List all users that are part of a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| The channel ID to list from. | 
 **optional** | ***MezonListChannelUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListChannelUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clanId** | **optional.String**| The clan id | 
 **channelType** | **optional.Int32**| The channel type | 
 **limit** | **optional.Int32**| Max number of records to return. Between 1 and 100. | 
 **state** | **optional.Int32**| The group user state to list. | 
 **cursor** | **optional.String**| An optional next page cursor. | 

### Return type

[**ApiChannelUserList**](apiChannelUserList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListChannelVoiceUsers**
> ApiVoiceChannelUserList MezonListChannelVoiceUsers(ctx, optional)
List all users that are part of a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonListChannelVoiceUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListChannelVoiceUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clanId** | **optional.String**| The clan id | 
 **channelId** | **optional.String**| The channel ID to list from. | 
 **channelType** | **optional.Int32**| The channel type | 
 **limit** | **optional.Int32**| Max number of records to return. Between 1 and 100. | 
 **state** | **optional.Int32**| The group user state to list. | 
 **cursor** | **optional.String**| An optional next page cursor. | 

### Return type

[**ApiVoiceChannelUserList**](apiVoiceChannelUserList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListClanDescs**
> ApiClanDescList MezonListClanDescs(ctx, optional)
List clans

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonListClanDescsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListClanDescsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Max number of records to return. Between 1 and 100. | 
 **state** | **optional.Int32**| The friend state to list. | 
 **cursor** | **optional.String**| Cursor to start from | 

### Return type

[**ApiClanDescList**](apiClanDescList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListClanUsers**
> ApiClanUserList MezonListClanUsers(ctx, clanId)
List all users that are part of a clan.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanId** | **string**| The clan ID to list from. | 

### Return type

[**ApiClanUserList**](apiClanUserList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListClanWebhook**
> ApiListClanWebhookResponse MezonListClanWebhook(ctx, clanId)
List clan webhook.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanId** | **string**| clan_id. | 

### Return type

[**ApiListClanWebhookResponse**](apiListClanWebhookResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListEvents**
> ApiEventList MezonListEvents(ctx, optional)
List user events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonListEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clanId** | **optional.String**| The clan of this event | 

### Return type

[**ApiEventList**](apiEventList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListFriends**
> ApiFriendList MezonListFriends(ctx, optional)
List all friends for the current user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonListFriendsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListFriendsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Max number of records to return. Between 1 and 100. | 
 **state** | **optional.Int32**| The friend state to list. | 
 **cursor** | **optional.String**| An optional next page cursor. | 

### Return type

[**ApiFriendList**](apiFriendList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListNotifications**
> ApiNotificationList MezonListNotifications(ctx, optional)
Fetch list of notifications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonListNotificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListNotificationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of notifications to get. Between 1 and 100. | 
 **clanId** | **optional.String**| The clan id | 
 **notificationId** | **optional.String**| The current notification Id. | 
 **code** | **optional.Int32**| The code. | 
 **direction** | **optional.Int32**| True if listing should be older notifications to newer, false if reverse. | 

### Return type

[**ApiNotificationList**](apiNotificationList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListOnboarding**
> ApiListOnboardingResponse MezonListOnboarding(ctx, optional)
list onboarding.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonListOnboardingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListOnboardingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clanId** | **optional.String**| clan id | 
 **guideType** | **optional.Int32**| guide_type: 0 &#x3D; greeting, 1 &#x3D; rule, 2 &#x3D; task, 3 &#x3D; question | 
 **limit** | **optional.Int32**| limit | 
 **page** | **optional.Int32**| page | 

### Return type

[**ApiListOnboardingResponse**](apiListOnboardingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListOnboardingStep**
> ApiListOnboardingStepResponse MezonListOnboardingStep(ctx, optional)
List onboarding step.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonListOnboardingStepOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListOnboardingStepOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clanId** | **optional.String**| clan id. | 
 **limit** | **optional.Int32**| limit | 
 **page** | **optional.Int32**| page | 

### Return type

[**ApiListOnboardingStepResponse**](apiListOnboardingStepResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListPTTChannelUsers**
> ApiPttChannelUserList MezonListPTTChannelUsers(ctx, optional)
List all users in ptt channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonListPTTChannelUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListPTTChannelUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clanId** | **optional.String**| The clan id | 
 **channelId** | **optional.String**| The channel ID to list from. | 
 **channelType** | **optional.Int32**| The channel type | 
 **limit** | **optional.Int32**| Max number of records to return. Between 1 and 100. | 
 **state** | **optional.Int32**| The group user state to list. | 
 **cursor** | **optional.String**| An optional next page cursor. | 

### Return type

[**ApiPttChannelUserList**](apiPTTChannelUserList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListRolePermissions**
> ApiPermissionList MezonListRolePermissions(ctx, roleId)
List role permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **string**|  | 

### Return type

[**ApiPermissionList**](apiPermissionList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListRoleUsers**
> ApiRoleUserList MezonListRoleUsers(ctx, roleId, optional)
List role permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **string**|  | 
 **optional** | ***MezonListRoleUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListRoleUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Max number of records to return. Between 1 and 100. | 
 **cursor** | **optional.String**| An optional next page cursor. | 

### Return type

[**ApiRoleUserList**](apiRoleUserList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListRoles**
> ApiRoleListEventResponse MezonListRoles(ctx, optional)
ListRoles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonListRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListRolesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clanId** | **optional.String**| clan id | 
 **limit** | **optional.String**| limit | 
 **state** | **optional.String**| state | 
 **cursor** | **optional.String**| cursor | 

### Return type

[**ApiRoleListEventResponse**](apiRoleListEventResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListStreamingChannelUsers**
> ApiStreamingChannelUserList MezonListStreamingChannelUsers(ctx, optional)
List all users that are part of a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonListStreamingChannelUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListStreamingChannelUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clanId** | **optional.String**| The clan id | 
 **channelId** | **optional.String**| The channel ID to list from. | 
 **channelType** | **optional.Int32**| The channel type | 
 **limit** | **optional.Int32**| Max number of records to return. Between 1 and 100. | 
 **state** | **optional.Int32**| The group user state to list. | 
 **cursor** | **optional.String**| An optional next page cursor. | 

### Return type

[**ApiStreamingChannelUserList**](apiStreamingChannelUserList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListStreamingChannels**
> ApiListStreamingChannelsResponse MezonListStreamingChannels(ctx, optional)
List streaming channels.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonListStreamingChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListStreamingChannelsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clanId** | **optional.String**| clan id | 

### Return type

[**ApiListStreamingChannelsResponse**](apiListStreamingChannelsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListThreadDescs**
> ApiChannelDescList MezonListThreadDescs(ctx, channelId, optional)
List user channels

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| channel id | 
 **optional** | ***MezonListThreadDescsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListThreadDescsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Max number of records to return. Between 1 and 100. | 
 **state** | **optional.Int32**| The channel state to list. | 
 **clanId** | **optional.String**| The clan of this channel | 
 **threadId** | **optional.String**| thread id | 

### Return type

[**ApiChannelDescList**](apiChannelDescList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListUserClansByUserId**
> ApiAllUserClans MezonListUserClansByUserId(ctx, )
ListUserClansByUserId

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiAllUserClans**](apiAllUserClans.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListUserPermissionInChannel**
> ApiUserPermissionInChannelListResponse MezonListUserPermissionInChannel(ctx, optional)
ListUserPermissionInChannel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonListUserPermissionInChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListUserPermissionInChannelOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clanId** | **optional.String**| clan id | 
 **channelId** | **optional.String**| channel id | 

### Return type

[**ApiUserPermissionInChannelListResponse**](apiUserPermissionInChannelListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListUsersAddChannelByChannelId**
> ApiAllUsersAddChannelResponse MezonListUsersAddChannelByChannelId(ctx, optional)
list user add channel by channel ids

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonListUsersAddChannelByChannelIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListUsersAddChannelByChannelIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **optional.String**|  | 
 **limit** | **optional.Int32**|  | 

### Return type

[**ApiAllUsersAddChannelResponse**](apiAllUsersAddChannelResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListWalletLedger**
> ApiWalletLedgerList MezonListWalletLedger(ctx, optional)
Get user status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonListWalletLedgerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListWalletLedgerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | 
 **cursor** | **optional.String**|  | 
 **transactionId** | **optional.String**|  | 

### Return type

[**ApiWalletLedgerList**](apiWalletLedgerList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonListWebhookByChannelId**
> ApiWebhookListResponse MezonListWebhookByChannelId(ctx, channelId, optional)
list webhook belong to the channel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**|  | 
 **optional** | ***MezonListWebhookByChannelIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonListWebhookByChannelIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clanId** | **optional.String**|  | 

### Return type

[**ApiWebhookListResponse**](apiWebhookListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonMarkAsRead**
> interface{} MezonMarkAsRead(ctx, body)
Mark as read

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiMarkAsReadRequest**](ApiMarkAsReadRequest.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonOpenDirectMess**
> interface{} MezonOpenDirectMess(ctx, body)
open direct message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiDeleteChannelDescRequest**](ApiDeleteChannelDescRequest.md)| Delete a channel the user has access to. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonPushPubKey**
> interface{} MezonPushPubKey(ctx, body)
store pubkey for e2ee

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiPushPubKeyRequest**](ApiPushPubKeyRequest.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonRegistFCMDeviceToken**
> ApiRegistFcmDeviceTokenResponse MezonRegistFCMDeviceToken(ctx, optional)
regist fcm device token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MezonRegistFCMDeviceTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonRegistFCMDeviceTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| The token | 
 **deviceId** | **optional.String**|  | 
 **platform** | **optional.String**|  | 

### Return type

[**ApiRegistFcmDeviceTokenResponse**](apiRegistFcmDeviceTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonRegisterStreamingChannel**
> ApiRegisterStreamingChannelResponse MezonRegisterStreamingChannel(ctx, body)
Register streaming in channel ( for bot - get streaming key)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiRegisterStreamingChannelRequest**](ApiRegisterStreamingChannelRequest.md)|  | 

### Return type

[**ApiRegisterStreamingChannelResponse**](apiRegisterStreamingChannelResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonRegistrationEmail**
> ApiSession MezonRegistrationEmail(ctx, body)
Authenticate a user with an email+password against the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiRegistrationEmailRequest**](ApiRegistrationEmailRequest.md)|  | 

### Return type

[**ApiSession**](apiSession.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonRemoveChannelFavorite**
> interface{} MezonRemoveChannelFavorite(ctx, channelId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonRemoveChannelUsers**
> interface{} MezonRemoveChannelUsers(ctx, channelId, optional)
Kick a set of users from a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| The channel ID to kick from. | 
 **optional** | ***MezonRemoveChannelUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonRemoveChannelUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userIds** | [**optional.Interface of []string**](string.md)| The users to kick. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonRemoveClanUsers**
> interface{} MezonRemoveClanUsers(ctx, clanId, optional)
Kick a set of users from a clan.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanId** | **string**| The clan ID to kick from. | 
 **optional** | ***MezonRemoveClanUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonRemoveClanUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userIds** | [**optional.Interface of []string**](string.md)| The users to kick. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonRpcFunc**
> ApiRpc MezonRpcFunc(ctx, id, payload, optional)
Execute a Lua function on the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The identifier of the function. | 
  **payload** | **string**| The payload of the function which must be a JSON object. | 
 **optional** | ***MezonRpcFuncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonRpcFuncOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **httpKey** | **optional.String**| The authentication key used when executed as a non-client HTTP request. | 

### Return type

[**ApiRpc**](apiRpc.md)

### Authorization

[HttpKeyAuth](../README.md#HttpKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonRpcFunc2**
> ApiRpc MezonRpcFunc2(ctx, id, optional)
Execute a Lua function on the server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The identifier of the function. | 
 **optional** | ***MezonRpcFunc2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MezonRpcFunc2Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **payload** | **optional.String**| The payload of the function which must be a JSON object. | 
 **httpKey** | **optional.String**| The authentication key used when executed as a non-client HTTP request. | 

### Return type

[**ApiRpc**](apiRpc.md)

### Authorization

[HttpKeyAuth](../README.md#HttpKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonSearchMessage**
> ApiSearchMessageResponse MezonSearchMessage(ctx, body)
Search message from elasticsearch service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiSearchMessageRequest**](ApiSearchMessageRequest.md)|  | 

### Return type

[**ApiSearchMessageResponse**](apiSearchMessageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonSendToken**
> interface{} MezonSendToken(ctx, body)
UpdateWallets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiTokenSentEvent**](ApiTokenSentEvent.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonSessionLogout**
> interface{} MezonSessionLogout(ctx, body)
Log out a session, invalidate a refresh token, or log out all sessions/refresh tokens for a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiSessionLogoutRequest**](ApiSessionLogoutRequest.md)| Log out a session, invalidate a refresh token, or log out all sessions/refresh tokens for a user. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonSessionRefresh**
> ApiSession MezonSessionRefresh(ctx, body)
Refresh a user's session using a refresh token retrieved from a previous authentication request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiSessionRefreshRequest**](ApiSessionRefreshRequest.md)| Authenticate against the server with a refresh token. | 

### Return type

[**ApiSession**](apiSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonSetChanEncryptionMethod**
> interface{} MezonSetChanEncryptionMethod(ctx, channelId, body)
store channel encryption method

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**|  | 
  **body** | [**MezonSetChanEncryptionMethodBody**](MezonSetChanEncryptionMethodBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonSetMuteNotificationCategory**
> interface{} MezonSetMuteNotificationCategory(ctx, body)
set mute notification user channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiSetMuteNotificationRequest**](ApiSetMuteNotificationRequest.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonSetMuteNotificationChannel**
> interface{} MezonSetMuteNotificationChannel(ctx, body)
set mute notification user channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiSetMuteNotificationRequest**](ApiSetMuteNotificationRequest.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonSetNotificationCategorySetting**
> interface{} MezonSetNotificationCategorySetting(ctx, body)
set notification user channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiSetNotificationRequest**](ApiSetNotificationRequest.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonSetNotificationChannelSetting**
> interface{} MezonSetNotificationChannelSetting(ctx, body)
set notification user channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiSetNotificationRequest**](ApiSetNotificationRequest.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonSetNotificationClanSetting**
> interface{} MezonSetNotificationClanSetting(ctx, body)
set notification user channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiSetDefaultNotificationRequest**](ApiSetDefaultNotificationRequest.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonSetNotificationReactMessage**
> interface{} MezonSetNotificationReactMessage(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiNotificationChannel**](ApiNotificationChannel.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonSetRoleChannelPermission**
> interface{} MezonSetRoleChannelPermission(ctx, body)
set permission role channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiUpdateRoleChannelRequest**](ApiUpdateRoleChannelRequest.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonStreamingServerCallback**
> ApiOssrsHttpCallbackResponse MezonStreamingServerCallback(ctx, body)
Ossrs http callback.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiOssrsHttpCallbackRequest**](ApiOssrsHttpCallbackRequest.md)|  | 

### Return type

[**ApiOssrsHttpCallbackResponse**](apiOssrsHttpCallbackResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUnbanApp**
> interface{} MezonUnbanApp(ctx, id)
Unban an app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The unique identifier of the app. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUnlinkApple**
> interface{} MezonUnlinkApple(ctx, body)
Remove the Apple ID from the social profiles on the current user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAccountApple**](ApiAccountApple.md)| Send a Apple Sign In token to the server. Used with authenticate/link/unlink. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUnlinkCustom**
> interface{} MezonUnlinkCustom(ctx, body)
Remove the custom ID from the social profiles on the current user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAccountCustom**](ApiAccountCustom.md)| Send a custom ID to the server. Used with authenticate/link/unlink. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUnlinkDevice**
> interface{} MezonUnlinkDevice(ctx, body)
Remove the device ID from the social profiles on the current user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAccountDevice**](ApiAccountDevice.md)| Send a device to the server. Used with authenticate/link/unlink and user. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUnlinkEmail**
> interface{} MezonUnlinkEmail(ctx, body)
Remove the email+password from the social profiles on the current user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAccountEmail**](ApiAccountEmail.md)| Send an email with password to the server. Used with authenticate/link/unlink. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUnlinkFacebook**
> interface{} MezonUnlinkFacebook(ctx, body)
Remove Facebook from the social profiles on the current user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAccountFacebook**](ApiAccountFacebook.md)| Send a Facebook token to the server. Used with authenticate/link/unlink. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUnlinkFacebookInstantGame**
> interface{} MezonUnlinkFacebookInstantGame(ctx, body)
Remove Facebook Instant Game profile from the social profiles on the current user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAccountFacebookInstantGame**](ApiAccountFacebookInstantGame.md)| Send a Facebook Instant Game token to the server. Used with authenticate/link/unlink. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUnlinkGameCenter**
> interface{} MezonUnlinkGameCenter(ctx, body)
Remove Apple's GameCenter from the social profiles on the current user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAccountGameCenter**](ApiAccountGameCenter.md)| Send Apple&#39;s Game Center account credentials to the server. Used with authenticate/link/unlink.  https://developer.apple.com/documentation/gamekit/gklocalplayer/1515407-generateidentityverificationsign | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUnlinkGoogle**
> interface{} MezonUnlinkGoogle(ctx, body)
Remove Google from the social profiles on the current user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAccountGoogle**](ApiAccountGoogle.md)| Send a Google token to the server. Used with authenticate/link/unlink. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUnlinkSteam**
> interface{} MezonUnlinkSteam(ctx, body)
Remove Steam from the social profiles on the current user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiAccountSteam**](ApiAccountSteam.md)| Send a Steam token to the server. Used with authenticate/link/unlink. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateAccount**
> interface{} MezonUpdateAccount(ctx, body)
Update fields in the current user's account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiUpdateAccountRequest**](ApiUpdateAccountRequest.md)| Update a user&#39;s account details. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateApp**
> interface{} MezonUpdateApp(ctx, id, body)
Update one or more fields on a app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| User ID to update. | 
  **body** | [**MezonUpdateAppBody**](MezonUpdateAppBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateCategory**
> interface{} MezonUpdateCategory(ctx, clanId, body)
Update fields in a given category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanId** | **string**|  | 
  **body** | [**MezonUpdateCategoryBody**](MezonUpdateCategoryBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateCategoryOrder**
> interface{} MezonUpdateCategoryOrder(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiUpdateCategoryOrderRequest**](ApiUpdateCategoryOrderRequest.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateChannelDesc**
> interface{} MezonUpdateChannelDesc(ctx, channelId, body)
Update fields in a given channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| The ID of the channel to update. | 
  **body** | [**MezonUpdateChannelDescBody**](MezonUpdateChannelDescBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateChannelPrivate**
> interface{} MezonUpdateChannelPrivate(ctx, body)
Update channel private.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiChangeChannelPrivateRequest**](ApiChangeChannelPrivateRequest.md)| Update fields in a given channel. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateClanDesc**
> interface{} MezonUpdateClanDesc(ctx, clanId, body)
Update fields in a given clan.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanId** | **string**|  | 
  **body** | [**MezonUpdateClanDescBody**](MezonUpdateClanDescBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateClanDescProfile**
> interface{} MezonUpdateClanDescProfile(ctx, clanId, body)
Update fields in a given clan profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanId** | **string**| Clan id | 
  **body** | [**MezonUpdateClanDescProfileBody**](MezonUpdateClanDescProfileBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateClanEmojiById**
> interface{} MezonUpdateClanEmojiById(ctx, id, body)
Update ClanEmoj By id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **body** | [**MezonUpdateClanEmojiByIdBody**](MezonUpdateClanEmojiByIdBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateClanStickerById**
> interface{} MezonUpdateClanStickerById(ctx, id, body)
Update a sticker by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **body** | [**MezonUpdateClanStickerByIdBody**](MezonUpdateClanStickerByIdBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateClanWebhookById**
> interface{} MezonUpdateClanWebhookById(ctx, id, body)
Update clan webhook by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| webhook id. | 
  **body** | [**MezonUpdateClanWebhookByIdBody**](MezonUpdateClanWebhookByIdBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateEvent**
> interface{} MezonUpdateEvent(ctx, eventId, body)
Update fields in a given event.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventId** | **string**|  | 
  **body** | [**MezonUpdateEventBody**](MezonUpdateEventBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateEventUser**
> interface{} MezonUpdateEventUser(ctx, body)
Update fields in a given event.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiDeleteEventRequest**](ApiDeleteEventRequest.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateOnboarding**
> interface{} MezonUpdateOnboarding(ctx, id, body)
update onboarding.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| id | 
  **body** | [**MezonUpdateOnboardingBody**](MezonUpdateOnboardingBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateOnboardingStepByClanId**
> interface{} MezonUpdateOnboardingStepByClanId(ctx, clanId, body)
Update onboarding step.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanId** | **string**| clan id. | 
  **body** | [**MezonUpdateOnboardingStepByClanIdBody**](MezonUpdateOnboardingStepByClanIdBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateRole**
> interface{} MezonUpdateRole(ctx, roleId, body)
Update fields in a given role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **string**| The ID of the role to update. | 
  **body** | [**MezonUpdateRoleBody**](MezonUpdateRoleBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateSystemMessage**
> interface{} MezonUpdateSystemMessage(ctx, clanId, body)
Update a system messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanId** | **string**| Clan ID | 
  **body** | [**MezonUpdateSystemMessageBody**](MezonUpdateSystemMessageBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateUser**
> interface{} MezonUpdateUser(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiUpdateUsersRequest**](ApiUpdateUsersRequest.md)| Fetch a batch of zero or more users from the server. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateUserProfileByClan**
> interface{} MezonUpdateUserProfileByClan(ctx, clanId, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clanId** | **string**| id clanc | 
  **body** | [**MezonUpdateUserProfileByClanBody**](MezonUpdateUserProfileByClanBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateUserStatus**
> interface{} MezonUpdateUserStatus(ctx, body)
Update user status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiUserStatusUpdate**](ApiUserStatusUpdate.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUpdateWebhookById**
> interface{} MezonUpdateWebhookById(ctx, id, body)
update webhook name by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| webhook Id | 
  **body** | [**MezonUpdateWebhookByIdBody**](MezonUpdateWebhookByIdBody.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonUploadAttachmentFile**
> ApiUploadAttachment MezonUploadAttachmentFile(ctx, body)
Upload attachment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiUploadAttachmentRequest**](ApiUploadAttachmentRequest.md)|  | 

### Return type

[**ApiUploadAttachment**](apiUploadAttachment.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MezonWithdrawToken**
> interface{} MezonWithdrawToken(ctx, body)
WithdrawToken

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiWithdrawTokenRequest**](ApiWithdrawTokenRequest.md)|  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


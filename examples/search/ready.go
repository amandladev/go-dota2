// AckPartyReadyCheck checks for/from a ack party ready.
// Request ID: k_EMsgPartyReadyCheckAcknowledge
// Request type: CMsgPartyReadyCheckAcknowledge
func (d *Dota2) AckPartyReadyCheck(
	readyStatus protocol.EReadyCheckStatus,
) {
	req := &protocol.CMsgPartyReadyCheckAcknowledge{
		ReadyStatus: &readyStatus,
	}
	d.write(uint32(protocol.EDOTAGCMsg_k_EMsgPartyReadyCheckAcknowledge), req)
}

// SendPartyReadyCheck sends a party ready check.
// Request ID: k_EMsgPartyReadyCheckRequest
// Response ID: k_EMsgPartyReadyCheckResponse
// Request type: CMsgPartyReadyCheckRequest
// Response type: CMsgPartyReadyCheckResponse
func (d *Dota2) SendPartyReadyCheck(
	ctx context.Context,
) (*protocol.CMsgPartyReadyCheckResponse, error) {
	req := &protocol.CMsgPartyReadyCheckRequest{}
	resp := &protocol.CMsgPartyReadyCheckResponse{}

	return resp, d.MakeRequest(
		ctx,
		uint32(protocol.EDOTAGCMsg_k_EMsgPartyReadyCheckRequest),
		req,
		uint32(protocol.EDOTAGCMsg_k_EMsgPartyReadyCheckResponse),
		resp,
	)
}

local:msg_overrides.go
// msgMethodNameOverrides overrides the generated client method names.
	dm.EDOTAGCMsg_k_EMsgPartyReadyCheckRequest:           "SendPartyReadyCheck",
	dm.EDOTAGCMsg_k_EMsgPartyReadyCheckAcknowledge:       "AckPartyReadyCheck",

CMsgPartyReadyCheckRequest
Request: k_EMsgPartyReadyCheckRequest matched to type: CMsgPartyReadyCheckRequest
Request: k_EMsgPartyReadyCheckRequest matched to response: k_EMsgPartyReadyCheckResponse
CMsgPartyReadyCheckResponse
Request: k_EMsgPartyReadyCheckResponse matched to type: CMsgPartyReadyCheckResponse
CMsgPartyReadyCheckAcknowledge
Request: k_EMsgPartyReadyCheckAcknowledge matched to type: CMsgPartyReadyCheckAcknowledge

8262	CLIENT	k_EMsgPartyReadyCheckRequest	SendPartyReadyCheck
8263	GC    	k_EMsgPartyReadyCheckResponse
8264	CLIENT	k_EMsgPartyReadyCheckAcknowledge	AckPartyReadyCheck

Computed client methods: [JoinChatChannel StartFindingMatch AbandonLobby StopFindingMatch LeaveLobby LaunchLobby ListLobbies JoinLobby SetLobbyDetails SendInitialQuestionnaireResponse ListChatChannel SendReadyUp SpectateFriendGame RequestReportsRemaining SubmitPlayerReport KickLobbyMember SubmitPlayerReportV2 SubmitPlayerReportResponseV2 SendWatchGame RequestMatchDetails CancelWatchGame ListFriendLobby CreateTeam InvitePlayerToTeam RespondToTeamInvite SendTeamInvite_GCResponseToInvitee KickTeamMember LeaveTeam ApplyTeamToLobby TransferTeamAdmin JoinLobbyBroadcastChannel EditTeamDetails SendBalancedShuffleLobby RequestMatchmakingStats CreateBotGame SetMatchHistoryAccess UpgradeLeagueItem RejoinAllChatChannels LeaveChatChannel SendChatMessage GetHeroStandings RequestItemEditorReservations ReserveEditorItemItemDef ReleaseEditorItemReservation FlipLobbyTeams SetProfilePrivacy SetMemberPartyCoach SetLobbyCoach GetEventPoints SetCompendiumSelection RequestCompendiumData GetPlayerMatchHistory RequestNotifications RequestNotificationsMarkRead SubmitInfoPlayer GetWeekendTourneySchedule RequestJoinableCustomGameModes RequestJoinableCustomLobbies JoinQuickCustomLobby QueryHasItem RequestEmoticonData ToggleLobbyBroadcastChannelCameramanStatus RedeemItem GetAllHeroProgress ListTrophies GetProfileCard GetBattleReport SetProfileCardSlots GetBattleReportAggregateStats GetBattleReportInfo CreateHeroStatue ReportBattleAcknowledge GetBattleReportMatchHistory SendLobbyEventPoints RerollPlayerChallenge SetPartyLeader CancelPartyInvites ApplyGemCombiner GetAllHeroOrder PurchasePlayerCardSpecific RequestLeagueAvailableLobbyNodes SendLeagueAvailableLobbyNodes GetFilteredPlayers SendRemoveFilteredPlayer SendUpdatePartyBeacon RequestActiveBeaconParties SendManageFavorites JoinPartyFromBeacon GetFavoritePlayers SendVerifyFavoritePlayers SendMMInfo PurchaseLabyrinthBlessings PurchaseFilteredPlayerSlot SendUpdateFilteredPlayerNote ClaimSwag RequestPlayerStats FindTopSourceTVGames RequestSocialFeedPostComment RequestCustomGamesFriendsPlayed RequestFriendsPlayedCustomGame ListCustomGamesTop SetPartyOpen SendMergePartyInvite RequestTopLeagueMatches RequestTopFriendMatches KickLobbyMemberFromTeam GetChatMemberCount RequestSocialFeedPostMessage SendCustomGameListenServerStartedLoading SendCustomGameClientFinishedLoading CloseLobbyBroadcastChannel RequestMatchesMinimal SendPingData GetProfileTickets SendH264Unsupported GetQuestProgress GetHeroStatsHistory InvitePrivateChatMember KickPrivateChatMember PromotePrivateChatMember DemotePrivateChatMember RequestLatestConductScorecard SendLatestConductScorecard RequestWagering RequestEventGoals SendHasPlayerVotedForMVP VoteForMVP LeaveTourneyWeekend RequestTeammateStats GetGiftPermissions VoteForArcana RequestArcanaVotesRemaining RequestMyTeamInfo PublishUserStat SubmitLobbyMVPVote SetSpectatorLobbyDetails CreateSpectatorLobby ListLobbySpectator SendSpectatorLobbyGameDetails OpenPlayerCardPack SelectCompendiumInGamePrediction GetTourneyWeekendPlayerStats RecyclePlayerCard CreatePlayerCardPack RequestGetPlayerCardRoster RequestSetPlayerCardRoster SendLobbyBattleCupVictory GetPlayerCardItemInfo RequestSteamDatagramTicket RequestTransferSeasonalMMR ReportChatPublicSpam SetPartyBuilderOptions JoinPlaytest SendLobbyPlaytestDetails SetFavoriteTeam ClaimEventAction GetPeriodicResource SendPeriodicResourceUpdated SubmitTriviaQuestionAnswer StartTriviaSession RequestAnchorPhoneNumber RequestUnanchorPhoneNumber RequestQuickStats RequestSelectionPriorityChoice AutographReward DestroyLobby PurchaseItemWithEventPoints PurchaseHeroRandomRelic ClaimEventActionUsingItem SendPartyReadyCheck AckPartyReadyCheck RequestGetRecentPlayTimeFriends RequestProfile SendProfileUpdate RequestHeroGlobalData RequestPlusWeeklyChallengeResult RequestPrivateMetadataKey ClaimCrawlCavernRoom SendCavernCrawlUseItemOnRoom SendCavernCrawlUseItemOnPath RequestCrawlCavernMapState RequestEventPointLogV2 RequestEventPointLogResponseV2 RequestEventTipsSummary RequestSocialFeed RequestSocialFeedComments GetCrawlCavernClaimedRoomCount RecordContestVote SendLobbyEventGameDetails GrantDevEventPoints GrantDevEventAction SendDevResetEventState GrantEventSupportConsumeItem RequestPlayerRecentAccomplishments RequestPlayerHeroRecentAccomplishments RequestPlayerCoachMatches SubmitCoachTeammateRating RequestPlayerCoachMatch RequestContestVotes VoteMVPTimeout SendDetailedGameStats SendMatchMatchmakingStats SubmitPlayerMatchSurvey SendDevDeleteEventActions RequestSubmitPlayerAvoid SendUnderDraftBuy RerollDraftUnder SendNeutralItemStats CreateGuild SetGuildInfo SendAddGuildRole SendModifyGuildRole SendRemoveGuildRole JoinGuild LeaveGuild SendInviteToGuild SendDeclineInviteToGuild CancelInviteToGuild KickGuildMember SetGuildMemberRole RequestGuildData RequestGuildMembership SendAcceptInviteToGuild SetGuildRoleOrder RequestGuildFeed RequestAccountGuildEventData RequestActiveGuildContracts SelectGuildContract SendAddPlayerToGuildChat SendUnderDraftSell RequestUnderDraft RedeemDraftUnderReward RequestActiveGuildChallenge RequestReporterUpdates SendAcknowledgeReporterUpdates RequestGuildEventMembers ReportGuildContent RequestAccountGuildPersonaInfo RequestAccountGuildPersonaInfoBatch SendLobbyFeaturedGamemodeProgress SubmitDraftTriviaMatchAnswer SendUnderDraftRollBackBench SendLobbyEventGameData GetOWMatchDetails SubmitOWConviction ClaimLeaderboardRewards SendRecalibrateMMR RequestChinaSSAURL RequestChinaSSAAccepted StartWatchingOverwatch StopWatchingOverwatch GetDPCFavorites SetDPCFavoriteState SendOverwatchReplayError SendCoachFriend RequestPrivateCoachingSession SendAcceptPrivateCoachingSession LeavePrivateCoachingSession GetCurrentPrivateCoachingSession SubmitPrivateCoachingSessionRating GetAvailablePrivateCoachingSessions GetAvailablePrivateCoachingSessionsSummary JoinPrivateCoachingSessionLobby RequestRespondToCoachFriend SetEventActiveSeasonID CreateTeamPlayerCardPack RequestBatchGetPlayerCardRoster RequestGetStickerbook RequestCreateStickerbookPage RequestDeleteStickerbookPage RequestPlaceStickers RequestPlaceCollectionStickers RequestOrderStickerbookTeamPage GetShopCandyUserData PurchaseShopCandyReward SendCandyShopDoExchange SendCandyShopDoVariableExchange RerollShopCandyRewards SetHeroSticker GetHeroStickers SetFavoritePage GrantShopDevCandyCandy SendCandyShopDevClearInventory OpenShopCandyBags GrantShopDevCandyCandyBags SendCandyShopDevShuffleExchange GrantShopDevCandyRerollCharges RequestCollectorsCacheAvailableData SendUploadMatchClip RequestRank RequestMapStats GetShowcaseUserData SetShowcaseUserData GetCraftingFantasyData SendFantasyCraftingPerformOperation SendFantasyCraftingDevModifyTablet GetToTIRoadQuests GetToTIRoadActiveQuest GetBingoUserData ClaimBingoRow RerollDevBingoCard GetBingoStatsData SendRoadToTIUseItem SubmitShowcaseReport GetAdminShowcaseReportsRollupList GetAdminShowcaseReportsRollup GetAdminShowcaseUserDetails SendShowcaseAdminConvict SendShowcaseAdminExonerate SendShowcaseAdminReset SendShowcaseAdminLockAccount SelectCraftingFantasyPlayer SendFantasyCraftingGenerateTablets UpgradeToGcFantasyCraftingClientTablets SendRoadToTIDevForceQuest RerollCraftingFantasyOptions SendLobbyRoadToTIMatchQuestData]

local:snapshot-type-list
CMsgPartyReadyCheckAcknowledge
CMsgPartyReadyCheckRequest
CMsgPartyReadyCheckResponse

CMsgReadyCheckStatus
CMsgReadyCheckStatus_ReadyMember

local:dota_gcmessages_common_match_management.proto
enum EReadyCheckStatus {
	k_EReadyCheckStatus_Unknown = 0;
	k_EReadyCheckStatus_NotReady = 1;
	k_EReadyCheckStatus_Ready = 2;
}

enum EReadyCheckRequestResult {
	k_EReadyCheckRequestResult_Success = 0;
	k_EReadyCheckRequestResult_AlreadyInProgress = 1;
	k_EReadyCheckRequestResult_NotInParty = 2;
	k_EReadyCheckRequestResult_SendError = 3;
	k_EReadyCheckRequestResult_UnknownError = 4;
}

message CSODOTAParty {
	enum State {
		UI = 0;
		FINDING_MATCH = 1;
		IN_MATCH = 2;
	}
optional .CMsgReadyCheckStatus ready_check = 62;
}

message CMsgReadyCheckStatus {
	message ReadyMember {
		optional uint32 account_id = 1;
		optional .EReadyCheckStatus ready_status = 2 [default = k_EReadyCheckStatus_Unknown];
	}

	optional uint32 start_timestamp = 1;
	optional uint32 finish_timestamp = 2;
	optional uint32 initiator_account_id = 3;
	repeated .CMsgReadyCheckStatus.ReadyMember ready_members = 4;
}

message CMsgPartyReadyCheckRequest {
}

message CMsgPartyReadyCheckResponse {
	optional .EReadyCheckRequestResult result = 1 [default = k_EReadyCheckRequestResult_Success];
}

message CMsgPartyReadyCheckAcknowledge {
	optional .EReadyCheckStatus ready_status = 1 [default = k_EReadyCheckStatus_Unknown];
}

local: dota_gcmessages_msgid.proto
k_EMsgPartyReadyCheckRequest = 8262;
k_EMsgPartyReadyCheckResponse = 8263;
k_EMsgPartyReadyCheckAcknowledge = 8264;

dota_gcmessages_common_match_management.pb.go
type EReadyCheckStatus int32

const (
	EReadyCheckStatus_k_EReadyCheckStatus_Unknown  EReadyCheckStatus = 0
	EReadyCheckStatus_k_EReadyCheckStatus_NotReady EReadyCheckStatus = 1
	EReadyCheckStatus_k_EReadyCheckStatus_Ready    EReadyCheckStatus = 2
)

var EReadyCheckStatus_name = map[int32]string{
	0: "k_EReadyCheckStatus_Unknown",
	1: "k_EReadyCheckStatus_NotReady",
	2: "k_EReadyCheckStatus_Ready",
}

var EReadyCheckStatus_value = map[string]int32{
	"k_EReadyCheckStatus_Unknown":  0,
	"k_EReadyCheckStatus_NotReady": 1,
	"k_EReadyCheckStatus_Ready":    2,
}

func (x EReadyCheckStatus) Enum() *EReadyCheckStatus {
	p := new(EReadyCheckStatus)
	*p = x
	return p
}

func (x EReadyCheckStatus) String() string {
	return proto.EnumName(EReadyCheckStatus_name, int32(x))
}

func (x *EReadyCheckStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EReadyCheckStatus_value, data, "EReadyCheckStatus")
	if err != nil {
		return err
	}
	*x = EReadyCheckStatus(value)
	return nil
}

func (EReadyCheckStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3efdba6b0593baab, []int{4}
}

type EReadyCheckRequestResult int32

const (
	EReadyCheckRequestResult_k_EReadyCheckRequestResult_Success           EReadyCheckRequestResult = 0
	EReadyCheckRequestResult_k_EReadyCheckRequestResult_AlreadyInProgress EReadyCheckRequestResult = 1
	EReadyCheckRequestResult_k_EReadyCheckRequestResult_NotInParty        EReadyCheckRequestResult = 2
	EReadyCheckRequestResult_k_EReadyCheckRequestResult_SendError         EReadyCheckRequestResult = 3
	EReadyCheckRequestResult_k_EReadyCheckRequestResult_UnknownError      EReadyCheckRequestResult = 4
)

var EReadyCheckRequestResult_name = map[int32]string{
	0: "k_EReadyCheckRequestResult_Success",
	1: "k_EReadyCheckRequestResult_AlreadyInProgress",
	2: "k_EReadyCheckRequestResult_NotInParty",
	3: "k_EReadyCheckRequestResult_SendError",
	4: "k_EReadyCheckRequestResult_UnknownError",
}

var EReadyCheckRequestResult_value = map[string]int32{
	"k_EReadyCheckRequestResult_Success":           0,
	"k_EReadyCheckRequestResult_AlreadyInProgress": 1,
	"k_EReadyCheckRequestResult_NotInParty":        2,
	"k_EReadyCheckRequestResult_SendError":         3,
	"k_EReadyCheckRequestResult_UnknownError":      4,
}

func (x EReadyCheckRequestResult) Enum() *EReadyCheckRequestResult {
	p := new(EReadyCheckRequestResult)
	*p = x
	return p
}

func (x EReadyCheckRequestResult) String() string {
	return proto.EnumName(EReadyCheckRequestResult_name, int32(x))
}

func (x *EReadyCheckRequestResult) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EReadyCheckRequestResult_value, data, "EReadyCheckRequestResult")
	if err != nil {
		return err
	}
	*x = EReadyCheckRequestResult(value)
	return nil
}

func (EReadyCheckRequestResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3efdba6b0593baab, []int{5}
}

type CSODOTAParty struct {
	ReadyCheck                      *CMsgReadyCheckStatus       protobuf:"bytes,62,opt,name=ready_check,json=readyCheck" json:"ready_check,omitempty"
}

func (m *CSODOTAParty) GetReadyCheck() *CMsgReadyCheckStatus {
	if m != nil {
		return m.ReadyCheck
	}
	return nil
}

type CMsgReadyCheckStatus struct {
	StartTimestamp       *uint32                             protobuf:"varint,1,opt,name=start_timestamp,json=startTimestamp" json:"start_timestamp,omitempty"
	FinishTimestamp      *uint32                             protobuf:"varint,2,opt,name=finish_timestamp,json=finishTimestamp" json:"finish_timestamp,omitempty"
	InitiatorAccountId   *uint32                             protobuf:"varint,3,opt,name=initiator_account_id,json=initiatorAccountId" json:"initiator_account_id,omitempty"
	ReadyMembers         []*CMsgReadyCheckStatus_ReadyMember protobuf:"bytes,4,rep,name=ready_members,json=readyMembers" json:"ready_members,omitempty"
	XXX_NoUnkeyedLiteral struct{}                            json:"-"
	XXX_unrecognized     []byte                              json:"-"
	XXX_sizecache        int32                               json:"-"
}

func (m *CMsgReadyCheckStatus) Reset()         { *m = CMsgReadyCheckStatus{} }
func (m *CMsgReadyCheckStatus) String() string { return proto.CompactTextString(m) }
func (*CMsgReadyCheckStatus) ProtoMessage()    {}
func (*CMsgReadyCheckStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_3efdba6b0593baab, []int{4}
}

func (m *CMsgReadyCheckStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgReadyCheckStatus.Unmarshal(m, b)
}
func (m *CMsgReadyCheckStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgReadyCheckStatus.Marshal(b, m, deterministic)
}
func (m *CMsgReadyCheckStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgReadyCheckStatus.Merge(m, src)
}
func (m *CMsgReadyCheckStatus) XXX_Size() int {
	return xxx_messageInfo_CMsgReadyCheckStatus.Size(m)
}
func (m *CMsgReadyCheckStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgReadyCheckStatus.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgReadyCheckStatus proto.InternalMessageInfo

func (m *CMsgReadyCheckStatus) GetStartTimestamp() uint32 {
	if m != nil && m.StartTimestamp != nil {
		return *m.StartTimestamp
	}
	return 0
}

func (m *CMsgReadyCheckStatus) GetFinishTimestamp() uint32 {
	if m != nil && m.FinishTimestamp != nil {
		return *m.FinishTimestamp
	}
	return 0
}

func (m *CMsgReadyCheckStatus) GetInitiatorAccountId() uint32 {
	if m != nil && m.InitiatorAccountId != nil {
		return *m.InitiatorAccountId
	}
	return 0
}

func (m *CMsgReadyCheckStatus) GetReadyMembers() []*CMsgReadyCheckStatus_ReadyMember {
	if m != nil {
		return m.ReadyMembers
	}
	return nil
}

type CMsgReadyCheckStatus_ReadyMember struct {
	AccountId            *uint32            protobuf:"varint,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"
	ReadyStatus          *EReadyCheckStatus protobuf:"varint,2,opt,name=ready_status,json=readyStatus,enum=protocol.EReadyCheckStatus,def=0" json:"ready_status,omitempty"
	XXX_NoUnkeyedLiteral struct{}           json:"-"
	XXX_unrecognized     []byte             json:"-"
	XXX_sizecache        int32              json:"-"
}

func (m *CMsgReadyCheckStatus_ReadyMember) Reset()         { *m = CMsgReadyCheckStatus_ReadyMember{} }
func (m *CMsgReadyCheckStatus_ReadyMember) String() string { return proto.CompactTextString(m) }
func (*CMsgReadyCheckStatus_ReadyMember) ProtoMessage()    {}
func (*CMsgReadyCheckStatus_ReadyMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_3efdba6b0593baab, []int{4, 0}
}

func (m *CMsgReadyCheckStatus_ReadyMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgReadyCheckStatus_ReadyMember.Unmarshal(m, b)
}
func (m *CMsgReadyCheckStatus_ReadyMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgReadyCheckStatus_ReadyMember.Marshal(b, m, deterministic)
}
func (m *CMsgReadyCheckStatus_ReadyMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgReadyCheckStatus_ReadyMember.Merge(m, src)
}
func (m *CMsgReadyCheckStatus_ReadyMember) XXX_Size() int {
	return xxx_messageInfo_CMsgReadyCheckStatus_ReadyMember.Size(m)
}
func (m *CMsgReadyCheckStatus_ReadyMember) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgReadyCheckStatus_ReadyMember.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgReadyCheckStatus_ReadyMember proto.InternalMessageInfo

const Default_CMsgReadyCheckStatus_ReadyMember_ReadyStatus EReadyCheckStatus = EReadyCheckStatus_k_EReadyCheckStatus_Unknown

func (m *CMsgReadyCheckStatus_ReadyMember) GetAccountId() uint32 {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return 0
}

func (m *CMsgReadyCheckStatus_ReadyMember) GetReadyStatus() EReadyCheckStatus {
	if m != nil && m.ReadyStatus != nil {
		return *m.ReadyStatus
	}
	return Default_CMsgReadyCheckStatus_ReadyMember_ReadyStatus
}

type CMsgPartyReadyCheckRequest struct {
	XXX_NoUnkeyedLiteral struct{} json:"-"
	XXX_unrecognized     []byte   json:"-"
	XXX_sizecache        int32    json:"-"
}

func (m *CMsgPartyReadyCheckRequest) Reset()         { *m = CMsgPartyReadyCheckRequest{} }
func (m *CMsgPartyReadyCheckRequest) String() string { return proto.CompactTextString(m) }
func (*CMsgPartyReadyCheckRequest) ProtoMessage()    {}
func (*CMsgPartyReadyCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3efdba6b0593baab, []int{5}
}

func (m *CMsgPartyReadyCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgPartyReadyCheckRequest.Unmarshal(m, b)
}
func (m *CMsgPartyReadyCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgPartyReadyCheckRequest.Marshal(b, m, deterministic)
}
func (m *CMsgPartyReadyCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgPartyReadyCheckRequest.Merge(m, src)
}
func (m *CMsgPartyReadyCheckRequest) XXX_Size() int {
	return xxx_messageInfo_CMsgPartyReadyCheckRequest.Size(m)
}
func (m *CMsgPartyReadyCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgPartyReadyCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgPartyReadyCheckRequest proto.InternalMessageInfo

type CMsgPartyReadyCheckResponse struct {
	Result               *EReadyCheckRequestResult protobuf:"varint,1,opt,name=result,enum=protocol.EReadyCheckRequestResult,def=0" json:"result,omitempty"
	XXX_NoUnkeyedLiteral struct{}                  json:"-"
	XXX_unrecognized     []byte                    json:"-"
	XXX_sizecache        int32                     json:"-"
}

func (m *CMsgPartyReadyCheckResponse) Reset()         { *m = CMsgPartyReadyCheckResponse{} }
func (m *CMsgPartyReadyCheckResponse) String() string { return proto.CompactTextString(m) }
func (*CMsgPartyReadyCheckResponse) ProtoMessage()    {}
func (*CMsgPartyReadyCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3efdba6b0593baab, []int{6}
}

func (m *CMsgPartyReadyCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgPartyReadyCheckResponse.Unmarshal(m, b)
}
func (m *CMsgPartyReadyCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgPartyReadyCheckResponse.Marshal(b, m, deterministic)
}
func (m *CMsgPartyReadyCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgPartyReadyCheckResponse.Merge(m, src)
}
func (m *CMsgPartyReadyCheckResponse) XXX_Size() int {
	return xxx_messageInfo_CMsgPartyReadyCheckResponse.Size(m)
}
func (m *CMsgPartyReadyCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgPartyReadyCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgPartyReadyCheckResponse proto.InternalMessageInfo

const Default_CMsgPartyReadyCheckResponse_Result EReadyCheckRequestResult = EReadyCheckRequestResult_k_EReadyCheckRequestResult_Success

func (m *CMsgPartyReadyCheckResponse) GetResult() EReadyCheckRequestResult {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return Default_CMsgPartyReadyCheckResponse_Result
}

type CMsgPartyReadyCheckAcknowledge struct {
	ReadyStatus          *EReadyCheckStatus protobuf:"varint,1,opt,name=ready_status,json=readyStatus,enum=protocol.EReadyCheckStatus,def=0" json:"ready_status,omitempty"
	XXX_NoUnkeyedLiteral struct{}           json:"-"
	XXX_unrecognized     []byte             json:"-"
	XXX_sizecache        int32              json:"-"
}

func (m *CMsgPartyReadyCheckAcknowledge) Reset()         { *m = CMsgPartyReadyCheckAcknowledge{} }
func (m *CMsgPartyReadyCheckAcknowledge) String() string { return proto.CompactTextString(m) }
func (*CMsgPartyReadyCheckAcknowledge) ProtoMessage()    {}
func (*CMsgPartyReadyCheckAcknowledge) Descriptor() ([]byte, []int) {
	return fileDescriptor_3efdba6b0593baab, []int{7}
}

func (m *CMsgPartyReadyCheckAcknowledge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CMsgPartyReadyCheckAcknowledge.Unmarshal(m, b)
}
func (m *CMsgPartyReadyCheckAcknowledge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CMsgPartyReadyCheckAcknowledge.Marshal(b, m, deterministic)
}
func (m *CMsgPartyReadyCheckAcknowledge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CMsgPartyReadyCheckAcknowledge.Merge(m, src)
}
func (m *CMsgPartyReadyCheckAcknowledge) XXX_Size() int {
	return xxx_messageInfo_CMsgPartyReadyCheckAcknowledge.Size(m)
}
func (m *CMsgPartyReadyCheckAcknowledge) XXX_DiscardUnknown() {
	xxx_messageInfo_CMsgPartyReadyCheckAcknowledge.DiscardUnknown(m)
}

var xxx_messageInfo_CMsgPartyReadyCheckAcknowledge proto.InternalMessageInfo

const Default_CMsgPartyReadyCheckAcknowledge_ReadyStatus EReadyCheckStatus = EReadyCheckStatus_k_EReadyCheckStatus_Unknown

func (m *CMsgPartyReadyCheckAcknowledge) GetReadyStatus() EReadyCheckStatus {
	if m != nil && m.ReadyStatus != nil {
		return *m.ReadyStatus
	}
	return Default_CMsgPartyReadyCheckAcknowledge_ReadyStatus
}

func init() {
	proto.RegisterEnum("protocol.EReadyCheckStatus", EReadyCheckStatus_name, EReadyCheckStatus_value)
	proto.RegisterEnum("protocol.EReadyCheckRequestResult", EReadyCheckRequestResult_name, EReadyCheckRequestResult_value)
	proto.RegisterType((*CMsgReadyCheckStatus)(nil), "protocol.CMsgReadyCheckStatus")
	proto.RegisterType((*CMsgReadyCheckStatus_ReadyMember)(nil), "protocol.CMsgReadyCheckStatus.ReadyMember")
	proto.RegisterType((*CMsgPartyReadyCheckRequest)(nil), "protocol.CMsgPartyReadyCheckRequest")
	proto.RegisterType((*CMsgPartyReadyCheckResponse)(nil), "protocol.CMsgPartyReadyCheckResponse")
	proto.RegisterType((*CMsgPartyReadyCheckAcknowledge)(nil), "protocol.CMsgPartyReadyCheckAcknowledge")
}

local:dota_gcmessages_common_match_management.proto
enum EReadyCheckStatus {
	k_EReadyCheckStatus_Unknown = 0;
	k_EReadyCheckStatus_NotReady = 1;
	k_EReadyCheckStatus_Ready = 2;
}

enum EReadyCheckRequestResult {
	k_EReadyCheckRequestResult_Success = 0;
	k_EReadyCheckRequestResult_AlreadyInProgress = 1;
	k_EReadyCheckRequestResult_NotInParty = 2;
	k_EReadyCheckRequestResult_SendError = 3;
	k_EReadyCheckRequestResult_UnknownError = 4;
}

message CSODOTAParty {
	enum State {
		UI = 0;
		FINDING_MATCH = 1;
		IN_MATCH = 2;
	}
optional CMsgReadyCheckStatus ready_check = 62;
}

message CMsgReadyCheckStatus {
	message ReadyMember {
		optional uint32 account_id = 1;
		optional EReadyCheckStatus ready_status = 2 [default = k_EReadyCheckStatus_Unknown];
	}

	optional uint32 start_timestamp = 1;
	optional uint32 finish_timestamp = 2;
	optional uint32 initiator_account_id = 3;
	repeated CMsgReadyCheckStatus.ReadyMember ready_members = 4;
}

message CMsgPartyReadyCheckRequest {
}

message CMsgPartyReadyCheckResponse {
	optional EReadyCheckRequestResult result = 1 [default = k_EReadyCheckRequestResult_Success];
}

message CMsgPartyReadyCheckAcknowledge {
	optional EReadyCheckStatus ready_status = 1 [default = k_EReadyCheckStatus_Unknown];
}

local:dota_gcmessages_msgid.pb.go
const (	
	EDOTAGCMsg_k_EMsgPartyReadyCheckRequest                                       EDOTAGCMsg = 8262
	EDOTAGCMsg_k_EMsgPartyReadyCheckResponse                                      EDOTAGCMsg = 8263
	EDOTAGCMsg_k_EMsgPartyReadyCheckAcknowledge                                   EDOTAGCMsg = 8264
)

var EDOTAGCMsg_name = map[int32]string{
	8262: "k_EMsgPartyReadyCheckRequest",
	8263: "k_EMsgPartyReadyCheckResponse",
	8264: "k_EMsgPartyReadyCheckAcknowledge",
}

var EDOTAGCMsg_value = map[string]int32{
	"k_EMsgPartyReadyCheckRequest":                                       8262,
	"k_EMsgPartyReadyCheckResponse":                                      8263,
	"k_EMsgPartyReadyCheckAcknowledge":                                   8264,
}

enum EDOTAGCMsg {
	k_EMsgPartyReadyCheckRequest = 8262;
	k_EMsgPartyReadyCheckResponse = 8263;
	k_EMsgPartyReadyCheckAcknowledge = 8264;
}
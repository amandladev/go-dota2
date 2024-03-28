package dota2

import (
	bgcm "github.com/paralin/go-dota2/protocol"
)

// LeaveParty attempts to leave the current party.
func (d *Dota2) LeaveParty() {
	d.write(uint32(bgcm.EGCBaseMsg_k_EMsgGCLeaveParty), &bgcm.CMsgLeaveParty{})
}

// RespondPartyInvite attempts to respond to a party invite.
func (d *Dota2) RespondPartyInvite(partyId uint64, accept bool) {
	d.write(uint32(bgcm.EGCBaseMsg_k_EMsgGCPartyInviteResponse), &bgcm.CMsgPartyInviteResponse{
		PartyId: &partyId,
		Accept:  &accept,
	})
}

// InviteToParty invites a player to your party.
func (d *Dota2) InviteToParty(steamID uint64) {
	d.write(uint32(bgcm.EGCBaseMsg_k_EMsgGCInviteToParty), &bgcm.CMsgInviteToParty{
		SteamId: &steamID,
	})
}

// KickFromParty kicks a player from your party.
func (d *Dota2) KickFromParty(steamID uint64) {
	d.write(uint32(bgcm.EGCBaseMsg_k_EMsgGCKickFromParty), &bgcm.CMsgKickFromParty{
		SteamId: &steamID,
	})
}

// SetPartyCoach announces whether you want to be the coach of the current party.
func (d *Dota2) SetPartyCoach(coach bool) {
	d.write(uint32(bgcm.EDOTAGCMsg_k_EMsgGCPartyMemberSetCoach), &bgcm.CMsgDOTAPartyMemberSetCoach{
		WantsCoach: &coach,
	})
}

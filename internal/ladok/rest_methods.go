package ladok

// CheckMyRights check if the certificate has enuff premissions
//func (s *RestService) isLadokPermissionsSufficent(ctx context.Context) (bool, error) {
//	egna, _, err := s.EgnaBehorigheter(ctx)
//	if err != nil {
//		return false, err
//	}
//
//	if len(egna.Anvandarbehorighet) < 1 {
//		s.logger.Warn("missing Användarbehörigheter al together")
//		return false, nil
//	}
//	uid := egna.Anvandarbehorighet[0].BehorighetsprofilRef.UID
//
//	behorighet, _, err := s.Behorigheter(ctx, uid)
//	if err != nil {
//		return false, err
//	}
//
//	havePermissions := map[int]string{}
//	for _, s := range behorighet.Systemaktiviteter {
//		havePermissions[s.ID] = s.Rattighetsniva
//	}
//
//	for id, wantPermission := range model.LadokPermissions {
//		havePermission, ok := havePermissions[id]
//		if !ok {
//			s.logger.Warn("missing permission:", id)
//			return false, nil
//		}
//		k := wantPermission[havePermission]
//		if !k {
//			s.logger.Warn("missing permission", id, wantPermission)
//			return false, nil
//		}
//	}
//	return true, nil
//}

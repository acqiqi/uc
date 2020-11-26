list,err := models.PacketGetAll()
	if err != nil {
		log.Println(err.Error())
		return
	}

	excel_data := make([]map[string]interface{},len(list))

	for i,v := range list{
		log.Println(i)
		log.Println(v.User.Nickname,v.SUser.Nickname)
		imap := map[string]interface{}{}
		imap["nickname"] = v.User.Nickname
		imap["mobile"] = v.User.Mobile
		imap["price"] = v.Price
		imap["title"] = v.Title
		imap["describe"] = v.Describe
		imap["status"] = v.Status
		imap["type"] = v.Type
		imap["created_at"] = v.CreatedAt
		imap["updated_at"] = v.UpdatedAt
		imap["snickname"] = v.SUser.Nickname
		imap["smobile"] = v.SUser.Mobile
		excel_data[i] = imap
	}

	keys := make([]utils.ExcelKey,11)
	keys[0] = utils.ExcelKey{Name:  "昵称", Key:   "nickname", Type:  "", Width: 100}
	keys[1] = utils.ExcelKey{Name:  "手机号", Key:   "mobile", Type:  "", Width: 100}
	keys[2] = utils.ExcelKey{Name:  "金额", Key:   "price", Type:  "", Width: 100}
	keys[3] = utils.ExcelKey{Name:  "标题", Key:   "title", Type:  "", Width: 100}
	keys[4] = utils.ExcelKey{Name:  "描述", Key:   "describe", Type:  "", Width: 100}
	keys[5] = utils.ExcelKey{Name:  "状态 0未使用 1使用 -1到期", Key:   "status", Type:  "", Width: 100}
	keys[6] = utils.ExcelKey{Name:  "类型", Key:   "type", Type:  "", Width: 100}
	keys[7] = utils.ExcelKey{Name:  "创建时间", Key:   "created_at", Type:  "", Width: 100}
	keys[8] = utils.ExcelKey{Name:  "更新时间", Key:   "updated_at", Type:  "", Width: 100}
	keys[9] = utils.ExcelKey{Name:  "分享人昵称", Key:   "snickname", Type:  "", Width: 100}
	keys[10] = utils.ExcelKey{Name:  "分享人手机号", Key:   "smobile", Type:  "", Width: 100}

	u := utils.ExcelUtils{
		Title:      "分享红包表",
		Header:     utils.Header{
			Title:  "分享红包表",
			Style:  "",
			Width:  0,
			Height: 0,
		},
		Keys:       keys,
		DataBase:   excel_data,
		SaveType:   "",
		FilePath:   "xx",
		FileName:   "account",
		SheetName:  "sheet1",
		F:          nil,
		ExcelStyle: utils.ExcelStyle{},
	}
	u.Export()




list,err := models.UcenterAccountsS()
	if err != nil {
		log.Println(err.Error())
		return
	}

	excel_data := make([]map[string]interface{},len(list))

	for i,v := range list{
		log.Println(i)
		log.Println(v.User.Nickname,v.SUser.Nickname)
		imap := map[string]interface{}{}
		imap["nickname"] = v.User.Nickname
		imap["mobile"] = v.User.Mobile
		imap["price"] = v.Price
		imap["title"] = v.Title
		imap["content"] = v.Content
		imap["is_dz"] = v.IsDz
		imap["type"] = v.Type
		imap["created_at"] = v.CreatedAt
		imap["updated_at"] = v.UpdatedAt
		imap["snickname"] = v.SUser.Nickname
		imap["smobile"] = v.SUser.Mobile
		excel_data[i] = imap
	}

	keys := make([]utils.ExcelKey,11)
	keys[0] = utils.ExcelKey{Name:  "昵称", Key:   "nickname", Type:  "", Width: 50}
	keys[1] = utils.ExcelKey{Name:  "手机号", Key:   "mobile", Type:  "", Width: 50}
	keys[2] = utils.ExcelKey{Name:  "金额", Key:   "price", Type:  "", Width: 50}
	keys[4] = utils.ExcelKey{Name:  "描述", Key:   "content", Type:  "", Width: 50}
	keys[5] = utils.ExcelKey{Name:  "状态 0未到账 1已到账", Key:   "is_dz", Type:  "", Width: 50}
	keys[7] = utils.ExcelKey{Name:  "创建时间", Key:   "created_at", Type:  "", Width: 50}
	keys[8] = utils.ExcelKey{Name:  "更新时间", Key:   "updated_at", Type:  "", Width: 50}
	keys[9] = utils.ExcelKey{Name:  "分享人昵称", Key:   "snickname", Type:  "", Width: 50}
	keys[10] = utils.ExcelKey{Name:  "分享人手机号", Key:   "smobile", Type:  "", Width: 50}

	u := utils.ExcelUtils{
		Title:      "分享红包表",
		Header:     utils.Header{
			Title:  "分享红包表",
			Style:  "",
			Width:  0,
			Height: 0,
		},
		Keys:       keys,
		DataBase:   excel_data,
		SaveType:   "",
		FilePath:   "xx",
		FileName:   "account1",
		SheetName:  "sheet1",
		F:          nil,
		ExcelStyle: utils.ExcelStyle{},
	}
	u.Export()












		users,err := models.GetUserAllList()
    	if err != nil {
    		return
    	}
    	//for i,v := range users {
    	//	if i < 4000 {
    	//		price := models.PacketGetSharePrice(v.Id)
    	//		if price > 0 {
    	//			log.Println(price)
    	//		}
    	//		models.UsersEdit(v.Id,map[string]interface{}{
    	//			"share_accounts": price,
    	//		})
    	//	}
    	//}
    	//for i,v := range users {
    	//	if i < 4000 {
    	//		price := models.UcenterAccountsList(v.Id)
    	//		if price > 0 {
    	//			log.Println(price)
    	//		}
    	//		models.UsersEdit(v.Id,map[string]interface{}{
    	//			"share_accounts": v.ShareAccounts + price,
    	//		})
    	//	}
    	//}









    		list,err := models.QuotationTplLinkAllList()
        	if err != nil {
        		log.Println(err.Error())
        		return
        	}
        	log.Println(list)
        	for _,v := range list {
        		tpl_arr := []Link{}
        		err := utils.JsonDecode(v.TableJson,&tpl_arr)
        		if err != nil {
        			continue
        		}
        		for _,value := range tpl_arr {
        			_,err := models.SmQuotationKeyGetInfo(value.Content,value.Describe)
        			if err != nil {
        				//新增数据
        				m := new(models.SmQuotationKey)
        				m.CatId = v.CatId
        				m.Zl = e.ToFloat64(value.Zl)
        				m.Fl = e.ToFloat64(value.Fl)
        				m.Work = e.ToFloat64(value.Work)
        				m.Content = value.Content
        				m.Describe = value.Describe
        				m.Unit = value.Unit
        				m.Flag = 1
        				models.SmQuotationKeyAdd(m)
        				log.Println(utils.JsonEncode(m))
        			}
        		}
        	}
package mail

type MailServer struct {
	Host string
	Port int
}

var SupportLists map[string]MailServer

func init() {
	SupportLists = make(map[string]MailServer)
	SupportLists["qq.com"] = MailServer{Host: "smtp.qq.com", Port: 587}
	SupportLists["foxmail.com"] = MailServer{Host: "smtp.qq.com", Port: 587}
	SupportLists["163.com"] = MailServer{Host: "smtp.163.com", Port: 587} // or 465 [never verify]
}

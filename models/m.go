package models

import "github.com/jinzhu/gorm"

type Mb struct {
	Id                      int64   `json:"id"`
	CreateDate              string  `json:"create_date"`
	ModifyDate              string  `json:"modify_date"`
	AcceptBindDate          string  `json:"accept_bind_date"`
	Attendance              int64   `json:"attendance"`
	Bind                    int64   `json:"bind"`
	BindOpenid              string  `json:"bind_openid"`
	DeleteFlag              int64   `json:"delete_flag"`
	Duties                  int64   `json:"duties"`
	FaceImageUrl            string  `json:"face_image_url"`
	HeadImage               string  `json:"head_image"`
	IdCardfmurl             string  `json:"id_cardfmurl"`
	IdCardzmurl             string  `json:"id_cardzmurl"`
	IdNo                    string  `json:"id_no"`
	MemberType              int64   `json:"member_type"`
	Mobile                  string  `json:"mobile"`
	Money                   float64 `json:"money"`
	Nickname                string  `json:"nickname"`
	WorkerType              string  `json:"worker_type"`
	WorkingState            int64   `json:"working_state"`
	WorkingStateApproveDate string  `json:"working_state_approve_date"`
	Admin                   int64   `json:"admin"`
	AdminProject            int64   `json:"admin_project"`
	MemberGroup             int64   `json:"member_group"`
	Signature               string  `json:"signature"`
	Unionid                 string  `json:"unionid"`
	Clouduid                string  `json:"clouduid"`
	Feature                 string  `json:"feature"`
}

type Ma struct {
	Id                        int64   `json:"id"`
	CreateDate                string  `json:"create_date"`
	ModifyDate                string  `json:"modify_date"`
	BindOpenid                string  `json:"bind_openid"`
	Birthday                  string  `json:"birthday"`
	DeleteFlag                int64   `json:"delete_flag"`
	HeadImage                 string  `json:"head_image"`
	InvitationCode            string  `json:"invitation_code"`
	LoginDate                 string  `json:"login_date"`
	Mobile                    string  `json:"mobile"`
	Money                     float64 `json:"money"`
	MyinvitationCode          string  `json:"myinvitation_code"`
	Nickname                  string  `json:"nickname"`
	Password                  string  `json:"password"`
	Score                     int64   `json:"score"`
	CompanyName               string  `json:"company_name"`
	ActivateProlocutor        int64   `json:"activate_prolocutor"`
	MemberType                int64   `json:"member_type"`
	Point64                   float64 `json:"point"`
	RequireType               int64   `json:"require_type"`
	CustomPoint64             float64 `json:"custom_point"`
	PopularizeMember          int64   `json:"popularize_member"`
	WxQrCodeUrl               string  `json:"wx_qr_code_url"`
	Address                   string  `json:"address"`
	Telephone                 string  `json:"telephone"`
	TopContacts               string  `json:"top_contacts"`
	Department                string  `json:"department"`
	Admin                     int64   `json:"admin"`
	JobNumber                 string  `json:"job_number"`
	Bind                      int64   `json:"bind"`
	Sources                   string  `json:"sources"`
	Textarea                  string  `json:"textarea"`
	ExtendFields              string  `json:"extend_fields"`
	Attendance                int64   `json:"attendance"`
	IdNo                      string  `json:"id_no"`
	WorkerType                string  `json:"worker_type"`
	FaceImageUrl              string  `json:"face_image_url"`
	WorkerPlatformReceiveFlag int64   `json:"worker_platform_receive_flag"`
	Locations                 string  `json:"locations"`
	IdCardfmurl               string  `json:"id_cardfmurl"`
	IdCardRecognition         string  `json:"id_card_recognition"`
	IdCardzmurl               string  `json:"id_cardzmurl"`
	HoursPerDay               int64   `json:"hours_per_day"`
	Remarks                   string  `json:"remarks"`
	WorkingState              int64   `json:"working_state"`
	WorkingStateApproveDate   string  `json:"working_state_approve_date"`
	Duties                    int64   `json:"duties"`
	MemberGroup               int64   `json:"member_group"`
	MemberGroupId             int64   `json:"member_group_id"`
	FaceInfo                  string  `json:"face_info"`
	AcceptBindDate            string  `json:"accept_bind_date"`
}

// 产品类型列表
func MaGetAllList() ([]*Ma, error) {
	var ma []*Ma
	err := db.Model(&Ma{}).Order("id desc").Find(&ma).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ma, nil
}
func MbGetAllList() ([]*Mb, error) {
	var mb []*Mb
	err := db.Model(&Mb{}).Order("id desc").Find(&mb).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return mb, nil
}

var WorkType = []string{
	"普工",
	"木工",
	"油漆工",
	"泥水工",
	"水电工",
	"焊工",
	"安装工",
	"瓦工",
}

var DescArray = []string{
	"承接各种拆除，打墙、打地板、开槽、拆墙、拆天花搬运上材料等各种拆除，云贵川人干活卖力为人耿直人员充足点包均可以",
	"",
	"专业钻孔。钻空调孔。下水道钻孔。消防钻孔。",
	"",
	"专业砌加气砖放砌体线。 需要老板联系",
	"",
	"专业主体天花拼缝打磨，本团队以质量求生存，以效益求发展，期待各位老板来电。全是90后",
	"",
	"专业放电缆，压电缆头，做桥架，装箱子，送电一条龙，人员多，设备全。欢迎老板骚扰！ 我们的服务宗旨就一句话，服从，服从，再服从，客户就是爹，一切以客户利益为上。百分百满足客户需求！",
	"",
	"有需要工地小工的吗，只限深圳，东莞，地区别的地方匆扰钱好拿的目前有两个人钱一定好拿，人力资源的中介，匆扰，住在地下室，信号不好，最好是加微信，手机微信同号，谢谢！",
	"",
	"专业装修与水电安装、拆卸、家居室内装修，广州等地专业装修水电、拆卸、家居装修、口罩厂三十万级以上、厂房排水电，消防工程等、团队大速度快，价格便宜。全广东任何地方都可以安排开工，欢迎加微信咨询",
	"",
	"深圳找活有需要的老板可以联系我，最好在深圳，装卸工，工地清垃圾，搬厂，搬家，砸墙，上吊篮，打地板，力工、电动基本都会，长期深圳附近。备注也可以保存电话，有需要的时候联系，工钱好结的来，不好结工钱的勿扰谢谢，我不是中介的",
	"",
	"绿化工，种花，种小树苗，种树，挖树，铺草皮，都可以做其他杂，工人力资源，有老板需要的，请联系我谢谢",
	"",
	"本人主打装配式建筑，墙板安装，专业的施工人数50人，主体木、泥、铁各施工人80",
	"",
	"本班组承接，开荒绿化种植，铺草皮，挖树，平地，楼盘保洁，室内外杂工工作，突击长期均可，我们常驻深圳，只在深圳干活，望老板能够长期合作（日结优先）微信同号",
	"",
	"本人从事室内外装修工程（油工）十余年，有专业的施工队伍，价格合理，保质保量，欢迎来电详询",
	"",
	"本人从事铁皮棚，钢结构安装，制作，简易铁棚制作安装，旧铁棚钢结构拆装，各种高空作业项目，石粉厂设备安装，提升机，收尘机，空汽斜槽等十多年，有专业的团队请各位老板来电合作共赢。",
	"",
	"贵州专业木工团队找活，点工、帮工、包工都做。都是吃苦耐劳的云贵川人，多年筛选，技术过硬好管理。以信誉质量为本，以回头合作为方向！欢迎您至电私聊…只限广东省内。",
	"",
	"贵州铁工突击队，团队，队长樊红，为各位老板和农民工弟兄服务！本团队全是云贵川老乡，做事积极认真，听安排，需要的老板请联系！樊红和其他突击队不一样，为老板赚钱就是为自己赚钱…欢迎来电。",
	"",
	"专业从事水电工作15年，有电工证书，有专业团队！专业预埋到二次安装一条龙承包。",
	"",
	"承接大小商铺，厂房，内外墙油漆。",
	"",
	"四川达州钢管架班组（3-10人队伍）在深圳从事多年主体外架，装修架，广告架，满堂红各种钢管架搭建及拆除！有需要的老板可联系（可包工、包料）点工突击！手机号与微信号同步，有时候忙可能没听到老板来电可备注添加微信谈！期待与各位老板精诚合作！联系人：小徐",
	"",
	"本人长期从事施工现场钢筋翻样工作料单易懂准确率高保质保量诚接所有类型钢筋翻样活可住现场欢迎有需要的老板联系一次合作终身朋友电话微信同号",
	"",
	"大量提供建筑各类大工各类小工专业清理地铁隧道清扫楼盘清运建筑垃圾以及家政保洁长期临时突击班组只要你一个电话就可以帮你解决问题欢迎致电任先生",
	"",
	"承接物流，打杂，搬运等等。已承接多个物流外包，有经验，和力气搬运。服务完善，有保险 提供各种厂工。",
	"",
	"本人从事装修行业15年。现有自己的施工团队，有自己的装饰公司，价格亲民，免费报价。免费出图。欢迎广东省内有需要的电联",
	"",
	"本人专业承包建筑木模工，有60人的队伍，包工点工师傅，人员非常实在，都是一些来自五湖四海的木工大师傅，需要的老板可以联系我。",
	"",
	"本人从事木工装修有一团队，承接范围室内外装修，店面，厂房，天花，隔墙，招牌，铝塑板，背景墙，柜，水电安装。可包工",
	"",
	"做下绳子的高空作业，除了批灰不会，做外墙下绳工作，认真，老实。。",
	"",
	"承接大小商铺厂房内外墙油漆涂料防水补漏",
	"",
	"水电施工队伍3人，经验丰富，长期找老板合作！ 承接经营范围，写字楼，商品房，厂房，商铺，展厅，家装， 给排水管，布线安装，桥架制作，。安防监控。。 价钱美丽，手艺精致 找有实力老板合作，不拖欠工资！ ！ 微信同号",
	"",
	"本人曾在碧桂园，万科，中建等。多个公司做过水电，消防弱电，管理（带班）因工地接近尾声，有团队想找活干。有需要的公司或老板，请多关照并联系！以上工程技术本人有一条龙服务。现实中的NL线，少了一条不来电！",
	"",
	"打空气转，安装路灯.穿线.立电杆！挖沟，拉线，种草坪，浇水，修补！打电转，手磨机打磨！打磨地板！修补路边地板砖！拆墙出渣！修地铁站打钢管架二杆子！修地铁站打杂，空气转破桩！钢结构小工！工地上吊材料！",
	"",
	"专业： 建筑内外墙腻子涂料工程。精装房装饰工程。 公寓房内外墙翻新，及设计美化工程。 地下室，厂房一条龙机械服务工程。",
	"",
	"专业木纹漆，铁艺油漆，大门围栏钢架油漆及翻新，本人有10年以上的工作经验，包工，点工都可以，能为您服务是我的荣幸，欢迎您的来电。",
	"",
	"从业防雷七八年有自己的精英团队，承接厂房防雷施工工程，防雷整改，防雷检测，合作共赢！",
	"",
	"本人有五六个人的油漆团队，有需要的老板可以下单，只做包公外墙不做谢谢",
	"",
	"本人找活干有需要的老板请打电话联系我本人能吃苦耐劳做事踏实工资结算最好是日结有需要拆除墙板门框搬运卸货清理垃圾之类的老板请联系我工资低于260一天不要来电打扰我",
	"",
	"本人团队有专业浇灌楼层地面混凝土，主体室内砌砖批荡，承接酒店商场拆除，本班组有专业师傅量少勿扰，主要做广东地区广州最好，本班组所有人员都服从公司安排人员精干人员170多，有需要的公司老总或者老板欢迎来电",
	"",
	"班组团队找活广东范围风管空调保温包保温玻璃棉或者橡塑保温团队人员4－10可以包或点工有工地需要联系或者以后需要的也可以联系",
	"",
	"承接钻孔，拆除，电焊焊接大小工程",
	"",
	"专业做路面硬化，停车场硬化，楼面混凝土，超低价",
	"",
	"本人专业承接各种室外大理石铺装，铺装范围:室外人行道，广场，小区，园林景观等，土石方，混泥土，有专业的施工队伍，技术一流，质量保证",
	"",
	"本人从事泥水工二十多年，做砖批墙，贴瓷砖，修补，打突击，欢迎老板来电。个人到团队",
	"",
	"水电十年已上，有证，点包均可，广州范围内强弱水电均可做。有事请联系",
	"",
	"专业承包各种吊顶包工包料可免费上门测量定制经验丰富有意请致电非诚勿扰",
	"",
	"本人敬业务实，踏实肯干，工作认真负责",
	"",
	"本人诚实可靠，勤劳肯干，吃苦耐劳",
	"",
	"本施工队承接集成精装修，展厅、展柜，KTV，酒店大型厂房木工天花吊顶工程。施工队现有常期木工队伍三十余人。如有木工需求，可包可点",
	"",
	"专业团队、放心可靠、全广东接水电分包包工活",
	"",
	"装修木工，专业天花造型吊顶，隔墙，等木工活老板有活的请下单",
	"",
	"本人专接瓷砖美缝清缝瓷砖美缝翻新",
	"",
	"做拉毛挂网8年了，外墙内墙都可做，最好是承包",
	"",
	"从事本行业十余年，对工作负责、细心、非常能吃苦、本人性格开朗乐观、人与人之间也很容易相处",
	"",
	"我干了15年油漆工，我手艺不错了",
	"",
	"本人做事积极，能按时完成老板分配的工作。",
	"",
	"氩弧焊，电焊，气保焊都会，平时是做扶手栏杆防盗窗的，做事实在，探伤焊不会，不上高，工资低于三百勿扰，有两个人。欢迎老板来电领走～",
	"",
	"我是一名钢筋工，服从管理，听领导安排",
	"",
	"小工之类;在深圳之类;搬家居丶搞卫生丶清理垃圾丶工地小工都可以;欢迎老板来电",
	"",
	"焊工找活电焊氩弧焊，有焊工证！可包、可点。最好深圳的活",
	"",
	"木工装修，天花吊顶，招牌，焊接工，阁楼龙骨隔墙，铝塑板，防火板，各种衣柜酒柜，鞋柜，吧台，收银台，铺位装修，办公楼装修",
	"",
	"专业无尘水电开槽，承接广东各地，地槽，水槽，线槽，混泥土开槽，量大从优，欢迎各位老板来电",
	"",
	"不怕苦累，服从安排，工资需要日结，工期不长的话可以完工结",
	"",
	"拆除上材料日结工 工地日结小工搬家 提供工人力工拆除砸墙打墙砖地板搬运工地打扫搞卫生垃圾清运打地板砖拆墙打墙打瓷片清理装修建筑垃圾都可以工地上材料搬水泥运沙袋",
	"",
}

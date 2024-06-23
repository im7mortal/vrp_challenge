package solvers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

var testSolution = [][]int{
	{79, 157, 104, 117, 53, 86, 188, 66, 146, 56, 108, 41, 165, 152},
	{118, 80, 38, 74, 112, 47, 60, 199, 135, 1, 193, 87},
	{113, 156, 49, 6, 12, 111, 186, 140, 162, 167, 132},
	{30, 121, 4, 176, 99, 9, 161, 160, 82, 164},
	{93, 147, 91, 159, 81, 22, 180, 198, 182},
	{85, 89, 107, 10, 45, 75, 51, 58},
	{54, 106, 105, 57, 7, 90, 166, 83},
	{28, 122, 20, 39, 129, 138, 25, 123},
	{190, 19, 149, 63, 88, 43, 170},
	{46, 78, 173, 183, 29, 72, 27},
	{139, 5, 120, 17, 23, 96, 18},
	{48, 115, 154, 55, 68, 76, 70},
	{65, 168, 197, 64, 95, 110},
	{33, 109, 16, 150, 40, 92},
	{84, 94, 179, 143, 32, 125},
	{134, 185, 102, 13, 131, 196},
	{42, 69, 151, 142, 141, 61},
	{130, 192, 194, 0, 177},
	{14, 184, 8, 3, 171},
	{52, 126, 145, 155, 195},
	{114, 62, 127, 124, 37},
	{24, 97, 67, 73, 11},
	{163, 15, 144, 36},
	{172, 26, 178, 31},
	{103, 98, 119, 153},
	{158, 35, 148, 169},
	{44, 21, 174, 34},
	{191, 2, 77},
	{133, 189, 59},
	{101, 137, 136},
	{50, 100, 187},
	{116, 71},
	{128, 175},
	{181},
}

var expectedResults = []float64{
	682.116071,
	719.364849,
	702.269054,
	687.175791,
	661.472427,
	655.328771,
	678.105105,
	704.463512,
	639.366636,
	669.756183,
	661.044197,
	700.785609,
	627.319307,
	650.262311,
	679.055468,
	706.058718,
	708.299660,
	631.470627,
	627.858911,
	687.579419,
	682.279133,
	708.453835,
	611.203858,
	626.568710,
	631.016515,
	656.439227,
	683.593468,
	567.832205,
	569.666758,
	701.173426,
	708.179153,
	541.355893,
	608.610337,
	454.457756,
}

var vectors = []*Vector{
	{Start: Point{X: 31.947803041218002, Y: 70.29735528044782}, End: Point{X: 85.13446863063243, Y: 4.4188878672117085}},
	{Start: Point{X: -17.367721860086466, Y: -51.65738157372827}, End: Point{X: -63.42562977104248, Y: -49.391281778769844}},
	{Start: Point{X: -3.479267837661599, Y: 98.63345745418567}, End: Point{X: -69.32768265755861, Y: 153.38002073322536}},
	{Start: Point{X: -58.63269874634831, Y: 172.58079475820227}, End: Point{X: 4.104920392892289, Y: 194.9561180896896}},
	{Start: Point{X: -27.25070855405738, Y: 17.601965645555197}, End: Point{X: -49.32712547336135, Y: 27.469249526565797}},
	{Start: Point{X: 52.62656450446505, Y: -18.12843766049885}, End: Point{X: 79.23516602833925, Y: 41.25232661248464}},
	{Start: Point{X: -6.861186469138947, Y: 68.25586240098038}, End: Point{X: 33.62876549272124, Y: 36.973904537998806}},
	{Start: Point{X: 70.48648269523191, Y: 100.17125629616591}, End: Point{X: 134.06364880682543, Y: 100.74975654035062}},
	{Start: Point{X: 27.991050126382213, Y: 152.3191842613466}, End: Point{X: -49.63692965010855, Y: 158.69788387906533}},
	{Start: Point{X: -100.2769624293266, Y: -69.28503941658697}, End: Point{X: -43.76008411912163, Y: -59.90887811827783}},
	{Start: Point{X: 7.266015968391283, Y: -36.7840838015574}, End: Point{X: -52.877289963555654, Y: 17.135003772033087}},
	{Start: Point{X: 184.56158738201506, Y: 32.92869378373049}, End: Point{X: 141.01470205743678, Y: 118.58594819722165}},
	{Start: Point{X: 20.668807971581206, Y: 15.621076318641864}, End: Point{X: 22.781521595629226, Y: 21.17422180168927}},
	{Start: Point{X: 49.44779834319345, Y: -208.48051631679826}, End: Point{X: 44.5461595010324, Y: -215.9886252988606}},
	{Start: Point{X: -13.964067766492713, Y: 33.43939473673973}, End: Point{X: 61.172977945699195, Y: 96.1318160958225}},
	{Start: Point{X: -60.340180398391006, Y: 94.65554507432293}, End: Point{X: -15.380726289758982, Y: 177.19975879379268}},
	{Start: Point{X: 14.968459029428857, Y: -42.21347263668738}, End: Point{X: 76.74684187510047, Y: -30.429683333914145}},
	{Start: Point{X: 113.80902895277073, Y: 29.95084068420485}, End: Point{X: 128.96311918222506, Y: 6.0280902810286285}},
	{Start: Point{X: 99.37232232534916, Y: -84.69665558214267}, End: Point{X: 75.47354244783773, Y: -12.629729424364086}},
	{Start: Point{X: -26.508032262946806, Y: 16.024867927985763}, End: Point{X: -8.955048504444488, Y: 94.84539012279834}},
	{Start: Point{X: -54.18652794222152, Y: -70.09340360669798}, End: Point{X: -72.18477493675479, Y: -26.953671149476044}},
	{Start: Point{X: -77.24377154236072, Y: -66.17128053031108}, End: Point{X: -123.07275282378849, Y: 5.305335168209297}},
	{Start: Point{X: -12.686306464864682, Y: 89.7517149732545}, End: Point{X: -42.59532074641959, Y: 98.58973371991385}},
	{Start: Point{X: 109.9259806174794, Y: 29.69637795089274}, End: Point{X: 74.38829972380199, Y: -42.707208397784456}},
	{Start: Point{X: -16.87744877098201, Y: -17.897473636587687}, End: Point{X: 60.84496132048125, Y: 41.47614126487437}},
	{Start: Point{X: -7.293097654493448, Y: -27.568263632328836}, End: Point{X: 12.673583245769898, Y: -85.82165935487309}},
	{Start: Point{X: -92.42488497440381, Y: -30.61253201436191}, End: Point{X: -118.97149689074662, Y: 22.645266300857323}},
	{Start: Point{X: 24.304087472460612, Y: -24.803214436316892}, End: Point{X: -21.706547761097408, Y: 6.913048578157252}},
	{Start: Point{X: -13.700354346338845, Y: -5.836727854518779}, End: Point{X: -22.416441230314533, Y: -85.50932895674318}},
	{Start: Point{X: -34.236027460367985, Y: -13.627536013521127}, End: Point{X: 56.466861199966154, Y: 18.49867182366989}},
	{Start: Point{X: -1.0005398114361128, Y: -14.258921887716818}, End: Point{X: 57.83124894581467, Y: -24.118133633221376}},
	{Start: Point{X: 19.82004609852699, Y: 106.01689418908367}, End: Point{X: -65.62278458341325, Y: 130.31306121590066}},
	{Start: Point{X: 73.48025002351925, Y: -100.49761395522637}, End: Point{X: 129.43533788751884, Y: -66.08203667421208}},
	{Start: Point{X: 5.338533002326251, Y: 7.926839245747189}, End: Point{X: -29.0319484963402, Y: 84.43954078432496}},
	{Start: Point{X: -105.47826208303113, Y: 66.22966795474188}, End: Point{X: -127.3407928831663, Y: 123.46309417162607}},
	{Start: Point{X: 97.05191180085566, Y: -35.37832746568934}, End: Point{X: 145.9001936885202, Y: -103.90982863491199}},
	{Start: Point{X: -15.417678726572426, Y: 264.10469447126314}, End: Point{X: -29.217037184481562, Y: 171.41222268712818}},
	{Start: Point{X: 101.12243568658045, Y: -36.32339359447427}, End: Point{X: 85.20325322853388, Y: 56.28995098297048}},
	{Start: Point{X: 29.85428582544904, Y: -49.78602439568257}, End: Point{X: 64.07647353185601, Y: -65.5286832010756}},
	{Start: Point{X: -69.3148378435486, Y: -42.393113034976565}, End: Point{X: -8.059902601881397, Y: -76.04327827283791}},
	{Start: Point{X: 151.27055444826334, Y: -33.69899416959527}, End: Point{X: 82.99982062675585, Y: -31.852378380339996}},
	{Start: Point{X: 99.24830174810121, Y: -7.947095781843899}, End: Point{X: 27.94069341192069, Y: -2.2776996747625233}},
	{Start: Point{X: 4.059619661038467, Y: -28.429096860533768}, End: Point{X: 82.59016292765747, Y: 13.070309805412052}},
	{Start: Point{X: 28.3022149470777, Y: 113.28544175221423}, End: Point{X: 72.24725361873666, Y: 96.19400958957311}},
	{Start: Point{X: -2.658789737295521, Y: 38.25007189405323}, End: Point{X: -81.01938458107695, Y: -7.278868131479591}},
	{Start: Point{X: -66.02980716058082, Y: 30.306977500988133}, End: Point{X: -97.21160944439391, Y: -25.8641688752462}},
	{Start: Point{X: 4.829460539781201, Y: -17.33367123185737}, End: Point{X: 9.58823599077283, Y: -105.88752042721916}},
	{Start: Point{X: 23.26958630697847, Y: -64.32849937249672}, End: Point{X: 62.03344878992365, Y: -6.4360762534501745}},
	{Start: Point{X: -20.922968665639537, Y: -6.106899056694644}, End: Point{X: -77.83798290551294, Y: 15.262439785274097}},
	{Start: Point{X: 13.001501252176158, Y: 33.87445071900927}, End: Point{X: 14.303143251729761, Y: 69.36496011841768}},
	{Start: Point{X: -49.6150636005733, Y: -76.42871267332029}, End: Point{X: -66.16555413581126, Y: -152.96846643474174}},
	{Start: Point{X: -154.13634328398132, Y: 15.037232992598998}, End: Point{X: -89.41423084280484, Y: 4.182635402449494}},
	{Start: Point{X: 31.283751132872023, Y: 1.240566013253531}, End: Point{X: -60.67900200466559, Y: 23.34381260444188}},
	{Start: Point{X: 3.392375743497875, Y: 23.505616750551646}, End: Point{X: 49.86963169377092, Y: 74.42987558786841}},
	{Start: Point{X: 8.840807225211448, Y: 13.825132183705254}, End: Point{X: -56.42341569845867, Y: 55.644439788323524}},
	{Start: Point{X: -204.94319622814356, Y: -90.15492854515581}, End: Point{X: -170.91285203677393, Y: -49.82470073249992}},
	{Start: Point{X: 50.1185309394165, Y: 23.858507418996062}, End: Point{X: 79.91422287564727, Y: -40.076475782860946}},
	{Start: Point{X: 54.98114845189468, Y: 111.2321972488704}, End: Point{X: 81.99795209487473, Y: 84.1336603588553}},
	{Start: Point{X: -93.69521320726365, Y: -6.951258328391589}, End: Point{X: -26.243751811700378, Y: -21.233006871114984}},
	{Start: Point{X: 68.36044245295233, Y: 77.7753559258765}, End: Point{X: 139.92654350277428, Y: 77.43687765849035}},
	{Start: Point{X: 62.147952281199686, Y: -8.04223889916301}, End: Point{X: -24.631130339944, Y: -16.685514313077093}},
	{Start: Point{X: 124.73960314070779, Y: 163.5383602689155}, End: Point{X: 126.84778832927626, Y: 169.1898448060043}},
	{Start: Point{X: -74.72195165361386, Y: -61.494932924980745}, End: Point{X: 7.488759290842083, Y: -26.725362330955797}},
	{Start: Point{X: -33.935147394315365, Y: 129.14193367657418}, End: Point{X: 17.02661237901077, Y: 180.4223922431187}},
	{Start: Point{X: -105.89686578657702, Y: 58.53552937178067}, End: Point{X: -59.27692872199851, Y: 104.16802494335303}},
	{Start: Point{X: 22.158511347398754, Y: 4.174434455538961}, End: Point{X: -52.022588313713825, Y: 7.367107545317071}},
	{Start: Point{X: 46.066202580365925, Y: 79.73366726706347}, End: Point{X: 43.82649044888412, Y: 42.02668372143584}},
	{Start: Point{X: 175.74911876148235, Y: 102.88069582446795}, End: Point{X: 178.54220356868512, Y: 41.51759630657598}},
	{Start: Point{X: -119.2852826649211, Y: -72.6006180022914}, End: Point{X: -94.45451862780958, Y: -29.87280260112321}},
	{Start: Point{X: 80.11334816179016, Y: -25.18176113208251}, End: Point{X: 113.69151228559161, Y: 22.356033333237068}},
	{Start: Point{X: -10.45684233879357, Y: -26.675011047642432}, End: Point{X: 7.934294032086102, Y: 40.14610695700755}},
	{Start: Point{X: 153.54073667144655, Y: -65.28531226132421}, End: Point{X: 202.67197684797458, Y: -151.0078549948535}},
	{Start: Point{X: 38.60401477350934, Y: 5.022353738745579}, End: Point{X: 5.266439230977468, Y: -20.696816840476124}},
	{Start: Point{X: 168.81031413934087, Y: 12.118209736837947}, End: Point{X: 211.19704374013094, Y: 26.052622606110155}},
	{Start: Point{X: 68.8575135313865, Y: -82.46779717035872}, End: Point{X: 0.9523218140451064, Y: -52.298318217622864}},
	{Start: Point{X: -99.82353207205999, Y: -23.21285555965197}, End: Point{X: -138.30025290268316, Y: -4.51127057275086}},
	{Start: Point{X: -93.52772230060553, Y: -13.310736970761244}, End: Point{X: -7.397360307875104, Y: -39.77700379784662}},
	{Start: Point{X: -60.74895171616467, Y: 165.31599917303413}, End: Point{X: -59.89158811282101, Y: 225.9128555784808}},
	{Start: Point{X: 6.4573268528523515, Y: -111.33382393843377}, End: Point{X: -1.0084371185843084, Y: -18.698131051787527}},
	{Start: Point{X: -6.166534570713066, Y: 9.881302790649695}, End: Point{X: 38.86093306171427, Y: 3.6684953562938336}},
	{Start: Point{X: 21.537566819371452, Y: -27.487082274313465}, End: Point{X: 22.160186412748615, Y: -49.260892533949466}},
	{Start: Point{X: -63.79711297303081, Y: 117.43217167079581}, End: Point{X: 1.5591078040868993, Y: 104.41298554806932}},
	{Start: Point{X: -30.67954469476345, Y: -15.406578500075838}, End: Point{X: -66.71514535990028, Y: -43.898035065646745}},
	{Start: Point{X: 59.64677657382323, Y: 60.60564252825312}, End: Point{X: -8.679781600614362, Y: 54.03088518939775}},
	{Start: Point{X: 20.561334527766434, Y: 10.782376090667507}, End: Point{X: 118.4152568260253, Y: 28.827506036112382}},
	{Start: Point{X: 11.124112902091674, Y: -12.614160569129119}, End: Point{X: 25.07421835928634, Y: -77.05599218798585}},
	{Start: Point{X: 44.23801015015006, Y: 77.51814284067694}, End: Point{X: 78.82157914166767, Y: 85.83903873165686}},
	{Start: Point{X: -66.67974055716095, Y: -5.582981185834683}, End: Point{X: -14.863482139459983, Y: -12.499257003371012}},
	{Start: Point{X: 19.4184720730983, Y: 191.66865496078609}, End: Point{X: 16.469850419794486, Y: 111.61665021834752}},
	{Start: Point{X: 35.40084442879584, Y: -91.53340172158546}, End: Point{X: 15.004500051390782, Y: -117.09950292487095}},
	{Start: Point{X: 135.54976435446844, Y: 115.0728403093562}, End: Point{X: 104.698834287767, Y: 103.28805396162349}},
	{Start: Point{X: -57.82185429377773, Y: 87.3276312533097}, End: Point{X: -94.17608348455303, Y: 89.1721956406227}},
	{Start: Point{X: 65.15063308418584, Y: -54.86875410300183}, End: Point{X: -31.00623708250285, Y: -32.77582978670159}},
	{Start: Point{X: -13.56211995529468, Y: 2.0509810701147733}, End: Point{X: -6.023351702844494, Y: 58.201496514208195}},
	{Start: Point{X: 119.32924488077079, Y: 19.59537716602831}, End: Point{X: 174.04986676303545, Y: 27.027479525722086}},
	{Start: Point{X: -23.974452535440278, Y: 100.0375109638304}, End: Point{X: -76.80555671696925, Y: 108.95702113242045}},
	{Start: Point{X: 94.91467456861784, Y: -37.52756086431977}, End: Point{X: 113.84062275637764, Y: -88.9065401897073}},
	{Start: Point{X: 76.50801459435009, Y: 62.426764529442046}, End: Point{X: 143.55622002700818, Y: 109.80432611642442}},
	{Start: Point{X: 28.44421903497674, Y: -107.60684195097507}, End: Point{X: 98.72039432319364, Y: -167.32840629790286}},
	{Start: Point{X: -107.65212105904904, Y: -22.613300115404137}, End: Point{X: -105.45636630109354, Y: -55.34014284723868}},
	{Start: Point{X: -52.77690211942523, Y: -166.55625251369955}, End: Point{X: -7.725272440631365, Y: -243.76111091423434}},
	{Start: Point{X: -4.074310357585901, Y: 45.74717237800441}, End: Point{X: 2.373882059629918, Y: 140.86210560618485}},
	{Start: Point{X: -84.28449712143937, Y: -184.00031252068501}, End: Point{X: -70.18015333635596, Y: -209.95051141864582}},
	{Start: Point{X: 32.22852208696037, Y: -29.077757326111385}, End: Point{X: 2.2066093846814496, Y: -115.9215033928129}},
	{Start: Point{X: 11.947819668020603, Y: -11.310867692457212}, End: Point{X: -11.771140451142884, Y: -15.42152627575973}},
	{Start: Point{X: -1.738557294963039, Y: 114.23446747552502}, End: Point{X: 73.88112559984162, Y: 104.29733538246371}},
	{Start: Point{X: -29.122803160808303, Y: 53.08934128399356}, End: Point{X: -8.608383649368605, Y: 114.93217622780085}},
	{Start: Point{X: -3.733037204591411, Y: -115.38438400202703}, End: Point{X: 8.384721182723844, Y: -32.74321484582731}},
	{Start: Point{X: 89.3691825314011, Y: -22.957502669185228}, End: Point{X: 96.14678334560166, Y: -8.092711538901556}},
	{Start: Point{X: -8.562836427429039, Y: 37.599331383235246}, End: Point{X: 4.130213192248826, Y: -54.082919991410215}},
	{Start: Point{X: -43.51842884524492, Y: 89.27563566020682}, End: Point{X: 27.54340895893197, Y: 41.24612288204945}},
	{Start: Point{X: 8.7447612497409, Y: 14.51046711038025}, End: Point{X: -15.9818656477374, Y: -37.37710932445882}},
	{Start: Point{X: 4.393027844455958, Y: -45.9041536134625}, End: Point{X: 24.577461945173365, Y: -63.14284767561695}},
	{Start: Point{X: 7.6545102577808075, Y: 4.174194768496207}, End: Point{X: 79.45454279115177, Y: 9.108743301172703}},
	{Start: Point{X: 13.91133374940284, Y: -35.246366729362904}, End: Point{X: -70.89069800494892, Y: -70.12226920651386}},
	{Start: Point{X: -90.10601042079372, Y: -13.503246485480537}, End: Point{X: -180.3279506860456, Y: -34.00611855979239}},
	{Start: Point{X: 39.54577664328555, Y: -55.197431438926685}, End: Point{X: 133.2757785806252, Y: -79.88934579902754}},
	{Start: Point{X: -10.075219189728287, Y: -26.0910126497455}, End: Point{X: -3.7377383371760713, Y: 13.26162558517349}},
	{Start: Point{X: 1.8803780554555791, Y: -0.8198387210737218}, End: Point{X: 41.594516910015734, Y: -29.16495662694606}},
	{Start: Point{X: 77.60526393516567, Y: -158.02695933625162}, End: Point{X: 46.378634471940416, Y: -189.4518099599049}},
	{Start: Point{X: 115.08010949731286, Y: 49.010459844981}, End: Point{X: 97.02612637799804, Y: 26.44859680363294}},
	{Start: Point{X: 50.4482423221341, Y: -0.14875333733161877}, End: Point{X: -23.541731285567188, Y: 14.811447529184601}},
	{Start: Point{X: -9.379911988945198, Y: -70.30481984893407}, End: Point{X: -63.249879499197206, Y: -79.76833689579202}},
	{Start: Point{X: 19.578356032914517, Y: -85.47818007216057}, End: Point{X: -27.741374697156058, Y: -39.25836808031268}},
	{Start: Point{X: 43.30212973826779, Y: -83.39738102396943}, End: Point{X: 108.45823400228622, Y: -20.61645135907181}},
	{Start: Point{X: 148.31422547799576, Y: -55.166191405525744}, End: Point{X: 100.90960796961025, Y: -73.13003545775949}},
	{Start: Point{X: -57.95506692490703, Y: 23.746134200453962}, End: Point{X: -66.94439566551722, Y: 97.32430034285883}},
	{Start: Point{X: -10.511950924538999, Y: -37.73592911289537}, End: Point{X: 76.70748398728334, Y: -71.40096016446955}},
	{Start: Point{X: -140.2993984316661, Y: -17.30624096975834}, End: Point{X: -186.11874518825016, Y: -93.41666586847413}},
	{Start: Point{X: -11.982437922376983, Y: -59.53889187404705}, End: Point{X: -58.184346706243396, Y: 18.698128497509884}},
	{Start: Point{X: 31.959736271527714, Y: -3.535979488146353}, End: Point{X: 31.820542247214274, Y: -64.74881404032115}},
	{Start: Point{X: 36.548524384106685, Y: -194.79637893429273}, End: Point{X: 91.90617836742607, Y: -163.10302899806896}},
	{Start: Point{X: 6.846087877046128, Y: -30.659252014982897}, End: Point{X: -25.93226378519692, Y: -33.58173353708797}},
	{Start: Point{X: 52.438936048961565, Y: 29.771563281813386}, End: Point{X: 145.26077092420468, Y: 28.09319066026811}},
	{Start: Point{X: 17.85059481569371, Y: -7.971965777586929}, End: Point{X: -17.54511755522318, Y: -86.1327838858759}},
	{Start: Point{X: -21.80264392239333, Y: -63.019201733725524}, End: Point{X: -16.33607915905759, Y: -71.86707267334859}},
	{Start: Point{X: -84.08666776983904, Y: 35.927165271635566}, End: Point{X: -100.40923266676977, Y: 122.06320376044542}},
	{Start: Point{X: 42.3449640311903, Y: 58.94997797761134}, End: Point{X: -7.758035267078121, Y: 139.21117323797674}},
	{Start: Point{X: -65.48492242220017, Y: 3.6301156888929738}, End: Point{X: -16.722369723824073, Y: -40.53796944235563}},
	{Start: Point{X: -19.709157445355913, Y: 4.968854886065927}, End: Point{X: 53.296229917727175, Y: -41.64503219552405}},
	{Start: Point{X: 36.603582286496355, Y: -89.80691249432934}, End: Point{X: 51.295558581549116, Y: -105.78437198613017}},
	{Start: Point{X: 173.59230455183396, Y: 153.79690771103182}, End: Point{X: 98.74007727617844, Y: 187.26510536602387}},
	{Start: Point{X: 179.16221228883543, Y: 152.37243336383048}, End: Point{X: 180.42715507627727, Y: 156.0830832349235}},
	{Start: Point{X: 149.6000431038263, Y: -79.2826218060568}, End: Point{X: 68.94132455477705, Y: -86.79772409226774}},
	{Start: Point{X: -16.782071789641616, Y: 170.48688789282784}, End: Point{X: -45.411583985084775, Y: 248.57938291976848}},
	{Start: Point{X: -53.9201149920261, Y: 131.20482385550957}, End: Point{X: -150.53032198666455, Y: 147.47432126058567}},
	{Start: Point{X: 37.71493236536315, Y: 19.22857510648856}, End: Point{X: 57.455377659861554, Y: 28.35479408343454}},
	{Start: Point{X: -6.527259333747034, Y: 75.93518722831334}, End: Point{X: -66.99590896124548, Y: 89.29343879456695}},
	{Start: Point{X: 146.04615363415152, Y: -51.247033811637834}, End: Point{X: 207.6863249666918, Y: -26.514487922223985}},
	{Start: Point{X: -7.199207960066548, Y: 69.2064718956988}, End: Point{X: -34.60855816075672, Y: 115.79116270817053}},
	{Start: Point{X: 86.29521343845175, Y: -27.070006612702727}, End: Point{X: 136.00856056345958, Y: -8.223546853207552}},
	{Start: Point{X: 123.28335233531804, Y: 38.43434446804336}, End: Point{X: 159.43258199387523, Y: 82.01337451441978}},
	{Start: Point{X: 21.14110817858685, Y: 1.894938698290256}, End: Point{X: -16.46019059829136, Y: -6.457354190494399}},
	{Start: Point{X: -27.8776149718327, Y: -61.09842004259902}, End: Point{X: -47.75039172112308, Y: 36.88402601832031}},
	{Start: Point{X: -195.3553535965549, Y: -90.4634468845427}, End: Point{X: -176.7521053464178, Y: -92.40517186376617}},
	{Start: Point{X: -169.1704677780649, Y: 174.1402689233438}, End: Point{X: -172.38332345181075, Y: 128.85219356335617}},
	{Start: Point{X: 66.5661666661102, Y: 23.14987376483479}, End: Point{X: 33.47653097084161, Y: 44.267994272289585}},
	{Start: Point{X: 35.81393447565969, Y: -3.5400668862124713}, End: Point{X: 22.665989725105536, Y: -14.106599741559402}},
	{Start: Point{X: -10.412467125160108, Y: -36.02665496725593}, End: Point{X: 87.41322504009511, Y: -30.983685278205524}},
	{Start: Point{X: -54.06972482683144, Y: 82.77379626643939}, End: Point{X: -58.70181300507925, Y: 108.57237274653019}},
	{Start: Point{X: -32.4767855517984, Y: -4.512752746987837}, End: Point{X: -28.965754317287033, Y: -29.717914832122055}},
	{Start: Point{X: -48.95213431850546, Y: -68.76895299168612}, End: Point{X: -51.017615052941885, Y: -13.82345764006854}},
	{Start: Point{X: 58.93213310144723, Y: -90.60377384120471}, End: Point{X: 3.589651630277004, Y: -99.00449358476061}},
	{Start: Point{X: -31.443442475203465, Y: 23.676841517746812}, End: Point{X: -70.53591942680346, Y: 78.30101903713586}},
	{Start: Point{X: -64.72209705445684, Y: -44.61917016550051}, End: Point{X: -59.328527486539286, Y: 11.854617754835118}},
	{Start: Point{X: 19.04355017679213, Y: -5.786591361260477}, End: Point{X: 19.827112612865978, Y: 0.9649405649592531}},
	{Start: Point{X: 90.05632573320811, Y: 89.16067828718272}, End: Point{X: 40.19488699073238, Y: 73.39529999113269}},
	{Start: Point{X: 8.281494884294165, Y: -81.1743338167131}, End: Point{X: 13.669070898909649, Y: -29.352269432688708}},
	{Start: Point{X: -14.238104537678295, Y: 19.75088106343165}, End: Point{X: -55.07597193596416, Y: 61.253460879887825}},
	{Start: Point{X: 213.2035720427486, Y: -23.799352812576306}, End: Point{X: 223.22241812447328, Y: -85.0950858058802}},
	{Start: Point{X: 62.118241081867545, Y: 100.23127028576306}, End: Point{X: 45.60254330987445, Y: 56.299231871137756}},
	{Start: Point{X: 15.739790546723126, Y: 190.2774939246857}, End: Point{X: 61.523898309863434, Y: 144.72185148096455}},
	{Start: Point{X: 15.326272942796162, Y: 2.0402168902647784}, End: Point{X: -84.19179426757196, Y: 7.747172836577189}},
	{Start: Point{X: 15.192894867691873, Y: -30.388526097262513}, End: Point{X: 42.8320498179215, Y: 39.49066247525103}},
	{Start: Point{X: -169.28188146647335, Y: -17.2413340405536}, End: Point{X: -134.44617416607315, Y: 67.53737455571768}},
	{Start: Point{X: -189.2816639880738, Y: -68.86686708660096}, End: Point{X: -223.92470610688125, Y: -147.2211350902536}},
	{Start: Point{X: -68.35750724984082, Y: 16.379686672198368}, End: Point{X: -108.87228341498883, Y: -1.0629358744488684}},
	{Start: Point{X: 42.52372005990183, Y: 14.344726769752246}, End: Point{X: 3.6982284609134624, Y: 86.86372851012484}},
	{Start: Point{X: -48.80964578478638, Y: 23.474837078015867}, End: Point{X: 12.422711137005017, Y: 101.24084683207082}},
	{Start: Point{X: 161.98523273930445, Y: -8.79998883018347}, End: Point{X: 164.47230941093986, Y: -87.52911507833709}},
	{Start: Point{X: -35.01387469452675, Y: 86.92469386031766}, End: Point{X: -0.8385600050763458, Y: 47.37615992386601}},
	{Start: Point{X: 108.77603872407613, Y: 82.65438680742962}, End: Point{X: 153.56790696237084, Y: 164.18904826134337}},
	{Start: Point{X: 67.84447190679673, Y: 65.89259330526806}, End: Point{X: 10.62875793688577, Y: 30.96982957489471}},
	{Start: Point{X: 50.697206201715275, Y: 24.924000089008857}, End: Point{X: -21.96974434835593, Y: -6.696163348556084}},
	{Start: Point{X: 39.78666797461843, Y: 104.59243560762715}, End: Point{X: 6.825559773329701, Y: 146.639498160015}},
	{Start: Point{X: -19.55827103696719, Y: -107.2847237602036}, End: Point{X: -87.70802211705853, Y: -159.47503930099552}},
	{Start: Point{X: -11.92379726657719, Y: -41.59523457611821}, End: Point{X: 40.67665464468057, Y: -89.35150642541342}},
	{Start: Point{X: -110.01985029748892, Y: -207.82833932195098}, End: Point{X: -170.73134473851962, Y: -196.13763232682498}},
	{Start: Point{X: 74.30192894819217, Y: 96.0635933457765}, End: Point{X: 60.09992629117977, Y: 90.50670454415442}},
	{Start: Point{X: 97.51948675981666, Y: 4.283520225972573}, End: Point{X: 100.91822283279004, Y: 97.72578974787206}},
	{Start: Point{X: 18.536636565384185, Y: 1.9042274071242162}, End: Point{X: -44.64526583565592, Y: 31.463946871088766}},
	{Start: Point{X: -14.684854052276568, Y: 16.473282651415474}, End: Point{X: -55.52496987242657, Y: 106.13729454645423}},
	{Start: Point{X: 32.692304839192765, Y: -48.41222364365557}, End: Point{X: 95.96834668652876, Y: 28.932709142981757}},
	{Start: Point{X: -53.1266213366773, Y: -46.655765645011826}, End: Point{X: -61.24478869213057, Y: 1.1480950500320262}},
	{Start: Point{X: 51.03204816134715, Y: 2.180201984042277}, End: Point{X: 34.515566730887876, Y: 69.32928023244129}},
	{Start: Point{X: -203.92291089597026, Y: 124.53552937136942}, End: Point{X: -152.91517277722286, Y: 77.70112663467893}},
	{Start: Point{X: 114.25177021558147, Y: -128.20868747204113}, End: Point{X: 31.94786528640404, Y: -83.01002656287437}},
	{Start: Point{X: -29.70842909358281, Y: 62.55451623681357}, End: Point{X: -87.98271690744748, Y: 51.18507722374556}},
	{Start: Point{X: -17.45627414119525, Y: 48.148750151679764}, End: Point{X: 44.71771546350095, Y: 56.29555186794275}},
	{Start: Point{X: -21.70319059222168, Y: -11.074397996874394}, End: Point{X: -33.21339692931606, Y: -69.51384505733981}},
}
var vectorDistances = []float64{
	84.668730,
	46.113622,
	85.634105,
	66.608287,
	24.181222,
	65.069907,
	51.166367,
	63.579798,
	77.889608,
	57.289353,
	80.774286,
	96.091084,
	5.941463,
	8.966480,
	97.856606,
	93.994147,
	62.892180,
	28.318624,
	75.926237,
	80.751359,
	46.743699,
	84.907020,
	31.187493,
	80.654858,
	97.805926,
	61.580244,
	59.507274,
	55.882912,
	80.147948,
	96.224255,
	59.652187,
	88.830068,
	65.691733,
	83.878027,
	61.266919,
	84.158905,
	93.714015,
	93.971549,
	37.669476,
	69.889203,
	68.295703,
	71.532629,
	88.821433,
	47.151707,
	90.627078,
	64.245642,
	88.681622,
	69.671872,
	60.794469,
	35.514371,
	78.308701,
	65.626017,
	94.581719,
	68.945018,
	77.513052,
	52.769230,
	70.536978,
	38.265368,
	68.946849,
	71.566901,
	87.208459,
	6.031892,
	89.260988,
	72.296517,
	65.236057,
	74.249772,
	37.773442,
	61.426634,
	49.418954,
	58.200816,
	69.305813,
	98.804014,
	42.105459,
	44.618412,
	74.305535,
	42.780922,
	90.104953,
	60.602921,
	92.936049,
	45.454063,
	21.782710,
	66.640339,
	45.938302,
	68.642159,
	99.503853,
	65.934476,
	35.570501,
	52.275802,
	80.106291,
	32.705296,
	33.025158,
	36.400994,
	98.662257,
	56.654333,
	55.223026,
	53.578757,
	54.753913,
	82.098083,
	92.224758,
	32.800420,
	89.388140,
	95.333256,
	29.535493,
	91.886623,
	24.072528,
	76.269804,
	65.156563,
	83.524864,
	16.337010,
	92.556732,
	85.770734,
	57.478054,
	26.544000,
	71.969399,
	91.693583,
	92.522247,
	96.927828,
	39.859676,
	48.791993,
	44.301508,
	28.896089,
	75.487243,
	54.694895,
	66.147018,
	90.480733,
	50.694156,
	74.125263,
	93.490984,
	88.838107,
	90.860595,
	61.212993,
	63.788230,
	32.908376,
	92.837008,
	85.801923,
	10.400392,
	87.668941,
	94.615908,
	65.792145,
	86.617787,
	21.705607,
	81.993757,
	3.920332,
	81.008059,
	83.175037,
	97.970550,
	21.747944,
	61.926573,
	66.416937,
	54.050031,
	53.165834,
	56.620656,
	38.517768,
	99.977432,
	18.704308,
	45.401897,
	39.254286,
	16.867723,
	97.955590,
	26.211120,
	25.448527,
	54.984304,
	55.976445,
	67.171590,
	56.730762,
	6.796849,
	52.294457,
	52.101366,
	58.225386,
	62.109131,
	46.933914,
	64.587159,
	99.681568,
	75.146682,
	91.656729,
	85.671067,
	44.110001,
	82.258279,
	98.979563,
	78.768400,
	52.268907,
	93.028020,
	67.031615,
	79.248473,
	53.426493,
	85.838322,
	71.045542,
	61.826840,
	15.250439,
	93.504059,
	69.754783,
	98.526900,
	99.930456,
	48.488285,
	69.150509,
	69.247748,
	93.898092,
	59.373031,
	62.705465,
	59.562185,
}

func TestDistance(t *testing.T) {
	for i := range vectors {
		d := Distance(vectors[i].Start, vectors[i].End)
		assert.True(t, closeEnough(vectorDistances[i], d, 0.001),
			fmt.Sprintf("Expected %f got %f", vectorDistances[i], d))
	}
}

func closeEnough(a, b, e float64) bool {
	return math.Abs(a-b) <= e
}

func TestTotalDistance(t *testing.T) {
	for i := range testSolution {
		d := TotalDistance(testSolution[i], vectors)
		assert.True(t, closeEnough(expectedResults[i], d, 0.001),
			fmt.Sprintf("Expected %f got %f", expectedResults[i], d))
	}

}

func TestCost(t *testing.T) {
	expectedCost := math.Round(39229.98289616)
	resCost := math.Round(Cost(testSolution, vectors))
	assert.Equal(t, expectedCost, resCost,
		fmt.Sprintf("Expected %f got %f", expectedCost, resCost))

}
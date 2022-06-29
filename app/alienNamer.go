package app

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// AlienNamer decides how newborn aliens should be known to others, regardless of their innate "DNA fingerprint" [Alien.ID]
type AlienNamer interface {
	NameAlien() string
}

type ordinalNamer struct {
	seq int
}

// NewAlienOrdinalNamer instantiates an AlienNamer that simply names Aliens sequentially, they are just cannon fodder at the end... the alien hive mind doesn't want heroes to be remembered when using this implementation
func NewAlienOrdinalNamer() AlienNamer {
	return &ordinalNamer{
		seq: 0,
	}
}

func (n *ordinalNamer) NameAlien() string {
	defer func() { n.seq++ }()
	return fmt.Sprintf("Alien %d", n.seq)
}

// this is just for fun :)
type famousAliensNamer struct {
	left  []string
	right []string
}

// NewFamousAliensNamer instantiates an AlienNamer that instills some glory into the Aliens in order to terrify humans a bit more
func NewFamousAliensNamer() AlienNamer {
	return &famousAliensNamer{
		left: []string{
			"admiring",
			"adoring",
			"affectionate",
			"agitated",
			"amazing",
			"angry",
			"awesome",
			"beautiful",
			"blissful",
			"bold",
			"boring",
			"brave",
			"busy",
			"charming",
			"clever",
			"cool",
			"compassionate",
			"competent",
			"condescending",
			"confident",
			"cranky",
			"crazy",
			"dazzling",
			"determined",
			"distracted",
			"dreamy",
			"eager",
			"ecstatic",
			"elastic",
			"elated",
			"elegant",
			"eloquent",
			"epic",
			"exciting",
			"fervent",
			"festive",
			"flamboyant",
			"focused",
			"friendly",
			"frosty",
			"funny",
			"gallant",
			"gifted",
			"goofy",
			"gracious",
			"great",
			"happy",
			"hardcore",
			"heuristic",
			"hopeful",
			"hungry",
			"infallible",
			"inspiring",
			"intelligent",
			"interesting",
			"jolly",
			"jovial",
			"keen",
			"kind",
			"laughing",
			"loving",
			"lucid",
			"magical",
			"mystifying",
			"modest",
			"musing",
			"naughty",
			"nervous",
			"nice",
			"nifty",
			"nostalgic",
			"objective",
			"optimistic",
			"peaceful",
			"pedantic",
			"pensive",
			"practical",
			"priceless",
			"quirky",
			"quizzical",
			"recursing",
			"relaxed",
			"reverent",
			"romantic",
			"sad",
			"serene",
			"sharp",
			"silly",
			"sleepy",
			"stoic",
			"strange",
			"stupefied",
			"suspicious",
			"sweet",
			"tender",
			"thirsty",
			"trusting",
			"unruffled",
			"upbeat",
			"vibrant",
			"vigilant",
			"vigorous",
			"wizardly",
			"wonderful",
			"xenodochial",
			"youthful",
			"zealous",
			"zen",
		},
		//extracted from https://en.wikipedia.org/wiki/List_of_fictional_alien_species:_A...Z with some DOM parsing, console.logging and formatting here
		/*
			var a =[];

			for (let t of $('#mw-content-text > div.mw-parser-output > table > tbody > tr > td:nth-child(1)')) {
				a.push(t.innerText);
			}
			console.log(a);
		*/
		right: []string{
			//A
			`Aaamazzarite`, `Aalaag`, `Abh`, `Abductor`, `Abzorbalof`, `Abyormenite`, `Acamarian`, `Acquaran`, `Adipose`, `Advent`, `Advents (aka Uranus or the Creators)`, `Aenar`, `Aeodronian`, `Aerophibian`, `Affront`, `Agorian`, `Akaali`, `Akis`, `Akrennian`, `Akritirian`, `Albategna`, `The Alien`, `Alien`, `Alien Sex Goddesse`, `Alkari`, `Allasomorph`, `Amaut`, `Altarian`, `Amazonian`, `Amnioni`, `Amorph`, `Anabi`, `Ancient`, `Ancient`, `Ancient`, `Andalite`, `Andorian`, `Andromeni`, `Angol Moi`, `Angosian`, `Annari`, `Anodite`, `Antaran`, `Antarian`, `Antedean`, `Anterian`, `Antican`, `The Anti-Monitor`, `Appoplexian`, `Aqualish`, `Aquatoid`, `Arachnichimp`, `Arachnid`, `Arachnoid`, `Ara`, `Arbryl`, `Arburian Pelarota`, `Arcadian`, `Arcean`, `Arcturian`, `Argolin`, "Arilou Lalee`lay", `Arisians and Eddorian`, `Ark Megaform`, `Arkonide`, `Armada of Annihilation`, `Arquillian`, `Airlia`, `Arnor`, `Asari`, `Aschen`, `Asgard`, `Aslan`, `Asuran`, `Atavu`, `Atevi`, `Aurelian`, `Autobot`, `Auronar`, `Auton`, `Axanar`, `Axon`, `Azathoth`, `Azgonian`, `Aziam`, `Azwaca`,
			//B
			`Baalol`, `Babel fish`, `Badoon`, `Bahmi`, `Bailie`, `Bajoran`, "Ba`ku", `Baliflid`, `Ballchinnian`, `Balmarian`, `Baltan`, `Bane`, `Bandersnatchi`, `Bandi`, `Banik`, `Bannermen`, `Bantha`, `Barkonide`, `Barzan`, `Battle Cupcake`, `Baufrin`, `Batarian`, `Beffel`, `Beings of the Extra Terrestrial origin which is Adversary of human race`, `Benjari`, `Benzite`, `Berellian`, `Berserker`, `Berubelan`, `Besalisk`, `Betazoid`, `Bgztl`, `Bioroid`, `Bith`, `Bismollian`, `Bjorn`, `Black Arm`, `Black Cloud`, `Black Oil`, `Blastaar`, `Blathereen`, `Blisk`, `Blob`, `Bolian`, `Boolite`, `Boomalope`, `Boomrat`, `Borg`, `Bothan`, `Braalian`, `Bradicor`, `Brain Dog`, `Brainiac`, `Brain`, `Brakiri`, `Breen`, `Brikar`, `Briori`, `Brood`, `Brunnen G`, `Brunali`, `Brutes (Jiralhanae)`, `Budong`, `Bugs from Klendathu`, `Bulrathi`, `Bugger`, `Bunless People`, `Bynar`,
			//C
			`Caitian`, `Calamarain`, `Caleban`, `Calcinite`, `Callinean`, `Calvin`, `Caliban`, `Capelon`, `Catnipian`, `Conductoid`, `Cascan`, `Cardassian`, `Carggite`, `Carrionite`, `Cat`, `Catbug`, `Cathar`, `Catkind`, `Cat People`, `Catalyte`, `Catteni`, `Cavalier`, `Cazare`, `Celareon`, `Celatid`, `Celestial`, `Celestialsapien`, `Centauri`, `Centran`, `Cerebrocrustacean`, `Ceti eel`, `Chaethe`, `Chalnoth`, `Chamachie`, `Changeling`, `Chao`, `Charrid`, `Chasch`, `Chatilian`, `Cheblon`, `Cheela`, `Chelonian`, `Chenjesu`, `Cherub`, `Chevanno`, `Chi`, `Chig`, `Chimera`, `Chimera Sui Generi`, `Chinger`, `Chis`, `Chmmr`, `Cho-choi`, `Chronovore`, `Chozo`, `Chronian`, `Chronomyst`, `Chryssalid`, `Cinnrusskin`, `Cirronian`, `Cizerack`, `Clairconctlar`, `Clutch Turtle`, `Cocytan`, `Coeurl`, `Cognocenti`, `Colata`, `Cole`, `Colonist`, `Colour out of space`, `Coluan`, `Combine`, `Conehead`, `Control Brain`, `Coreeshi`, `Coridanite`, `Corporal Giroro`, `Corvallen`, `Covenant`, `Crystal`, `Cragmite`, `Creator`, `Crite`, `Cryon`, `Ctarl-Ctarl`, `Cthulhi`, `Cthulhu`, `Cybermen`, `Cybertronian`, `Cybyota`, `Cyclon`, `Cylon`, `Cynoid`, `Cyrollan`,
			//D
			`Daemonite`, `Daimon`, `Dakkamite`, `Daktaklakpak`, `Dalek`, `The Dance`, `Dantari`, "Darjakr`Ul", `Dark Eldar`, `Darlok`, `Darrian`, `Darzok`, `Daxamite`, `Decapodian`, `Decepticon`, `Dear`, `Deep One`, `Deep One`, `Defiance`, `Delphon`, `Deltan`, `Delvian`, `Demiurg`, `Denean`, `Denebian`, `Deng`, `Denobulan`, `Dentic`, `Dentrassi`, `Deoxy`, `Derkuhr`, `Dessarian`, `Detrovite`, `Devaronian`, `Devastator`, `Devilukean`, `Dilbian`, `Dirdir`, `Dnyarri`, `Dog`, `Dom Kavash`, `Dominator`, `Dominator`, `DomZ`, `Doog`, `the Doubler`, `Douwd`, `Draaknaar`, `Drac`, `Draconian`, `Draenei`, `Drahvin`, `Draic Kin`, `Drak`, `Drakh`, `Dralasite`, `Drath`, `Drayan`, `Dread Lord`, `Drej`, `Dremer`, `Drengin`, "Drones (Yanme`e)", `Drophyd`, `Druuge`, `Dug`, `Duos from Uranu`, `Dubtak`, `Durlan`, `Duro`, `Dyson Alien`,
			//E
			`Eber`, `Ectonurite`, `Edestekai`, `Edo`, `Ekhonide`, `Ekoplektoid`, `El-Aurian`, `Eldar`, `Elder God`, `Elder Thing`, `Eldorian`, `Electrogoomba`, `Elerian`, `Elites (Sangheili)`, `Ellimist`, `Elve`, `Elvi`, `Elgyem`, `Engineers (Huragok)`, `Engineer`, `Eoladi`, `Eorna`, `Eosapien`, `Ep-Hoger`, `Eppori`, `Erg`, `Eridani`, `Eschiff`, `Esmer`, `Esperian`, `E.T`, `Ethereal`, `Exobot`, `Extraterrestrial Biological Entities (EBE)`, `Exquivan`, `Extraterrestrial Living-metal Shape-shifter (ELS)`, `Evon`, `Evroniani`, `Ewok`, `Experiment`, `Eywa`,
			//F
			`Face Dancer`, `Falleen`, `Felysian`, `Felinu`, "Fek`Ihri", `Fendahl`, `Ferengi`, `Ferronan`, `Festival`, `Firebug`, `Fithp`, `Flat cat`, `Fleeblebroxian`, `Floater`, `Flood`, `Florauna`, `Fludentri`, `Foamasi`, `Foralbo`, `Forerunner`, `Forerunner`, `The Forest of Cheem`, `Formic`, `Fotiallian`, `Fott`, `Frieza`, `Frutmaka`, "F`sherl-Ganni", `Fulmini`, `Furbl`, `Furling (one of the Four Great Races)`, `Furon`, `Futar`,
			//G
			`Gadmeer`, `Gaim`, `Galactoid`, `Galactu`, `Galaxoid`, `Gallaxhar`, `Galvan`, `Galvanic Mechomorph`, `Gamilons/Gamilu`, `Gamorrean`, `Gand`, `Ganymean`, `Gaoul`, `Garthling`, `Gashlai`, `The Gelth`, `Gedd`, `GELF`, `Gelgamek`, `Gemini`, `Gem`, `Geochelone Aerio`, `Geonosian`, `Geth`, `Gethenian`, `Ghaman`, `Ghost`, `Gibi`, `Gigan`, `Giyga`, `Gigglepie`, `Gill Men`, "G`kek", `Gladifer`, "Glapum`tian", `Glendalien`, `Gloarft`, `Gnaar`, `Gnolam`, `Gnosi`, "Goa`uld", `Godan`, `Gonknoid`, `Gonzo`, `Goola-Goola`, `Gorlock`, `Gorg`, `Gorn`, `Gort`, `Goszul`, `Gourmand`, `Govorom`, `Gowachin`, `G.R.A.I.S.E.`, `Gran`, `The Graske`, `Great Race of Yith`, `Grendarl`, `Grendler`, `Gretchin/Grotz`, `Grey alien`, `Grog`, `Grudek`, `Grue`, `Grund`, `Grundo`, `Grunts (Unggoy)`, `Guardians of the Universe`, `Gubru`, `Gungan`,
			//H
			"H`Harn", `Hacan`, `Haggunenon`, `Halfling`, `Hallessi`, `Hallucinoid`, `Halosian`, `Hangi`, `Hani`, `Hanshak`, `Harammin`, `Hardcore Hill Midget`, `Harika`, `Harmonia`, `Haydonite`, `Headcrab`, `Headie`, `Heechee`, `Helmacron`, `Heptapod`, `Herald`, `Hermat`, `Hi-Five Alien`, `Hierarchy`, `Highbreed`, `High One`, `Hiigaran`, `Hirogen`, `Hisa`, `Hive, The`, `Hiver`, `Hiver`, `Hobo`, `Hoka`, `Hoix`, `Hooloovoo`, `Hoofonoggle`, `Hoon`, `Horda`, `Hork-Bajir`, `Horta`, `Howler`, `Hrossa`, `Hrud`, `Hunters (Lekgolo)`, `Hunter`, `Humma`, "Hur`q", `Hutt`, `Husnock`, `Hydran`, `Hydrogue`, `Hykraiu`, `Hylar`, `Hynerian`,
			//I
			`Iberon`, `Ib`, `Ice Warrior`, `Iconian`, `Iconian`, `Ifshnit`, `Ikrini`, `Ilanic`, `Ilwrath`, `Imsaei`, `Imskian`, `Ing`, `Insect`, `Interion`, `Invid/Inbit`, `Irken`, `Isanian`, `Ishtarian`, `Ithanite`, `Ithkul`, `Ixtl`,
			//J
			`Jaffa`, `Jaridian`, `Jawa`, "Jem`Hadar", `Jenova`, `Jezzedaic Priest`, `Jiralhanae`, `Jjaro`, "J`naii", `Jocacean`, `Jophur`, `Jorenian`, `Jotoki`, `Joozian`, `Judoon`, `Junk squid`,
			//K
			`Kafer`, `Kaled`, `Kalish`, `Kaalium`, `Kal-Kriken`, `Kalliran`, `Kambuchka`, `Kanamit`, `Kang and Kodo`, `Karemma`, `Kariak`, `Kat`, `Kazon`, `Kdatlyno`, `Keron`, "Kes`Rith", `Key-Guardian`, `King Ghidorah`, `Kineceleran`, `Kisshu`, `Kharaa`, `Kherubim`, `Khund`, `Khurtarnan`, `Kif`, `Kig-yar`, `Kilaak`, `Kilrathi`, `Kizanti`, `Kimera`, `King Cold`, `Kivar`, `Klaestronian`, `Klackon`, `Klaxun`, `Kleer Skeleton`, `Klingon`, `Knnn`, `Koozbanian`, `Korath`, `Korbinite`, `Korvax`, `Korx`, `K-PAXian`, "Kra`hen", `Kraked`, `Kraylor`, `Kree`, `Kreel`, `Kreely`, `Kreetassan`, `Krell`, `Krellepem`, `Krenim`, `Kriken`, `Krishnan`, `Krith`, `Krogan`, `Krolp`, `Kromagg`, `Kronn`, `Kroot`, `Krynn`, `Krynn`, `Krynoid`, `Kryptonian`, `Kryten`, `Kssthrata`, "K`tang", `Ktarian`, `Kulturan`, `Kweltikwan`, `Kymellian`, `Kymera`, `Kymnar`, `Kyo`, `Kyrie`, `Kyulek`, `Kzinti`,
			//L
			`L1Z1X Mindnet`, `Lalloran`, `Lance Corporal Dororo`, `Large Nostril People of Boodie-Nen`, `Lavo`, `Laxidasian`, `Lectroid`, `Leeran`, `Lekgolo`, `Leonider`, `Lepidopterran`, `Leviathan`, `Life Fiber`, `Liir`, `Limax`, `Lipul`, `Lithian`, `Little Guy`, `Livrai`, `Lk (fungi type alien)`, `Llorn`, `Loboan`, `Lobster Men`, `Locust`, `Logrian`, `Lohvo`, `Lombaxe`, `Loomi`, `Loric`, `Lorwardian`, `Lucratian`, `Luma`, `Luminoth`, `Lunarian`, `Lunarian`, `Lunatone`, `Lurian`, `Lurman`, `Luxan`, `Ly-Cilph`, `Lycocian`, `Lyran`,
			//M
			`M-113 Creature`, `Macra`, `Magma`, `Magog`, `Magolor`, "Mahendo`sat", `Maian`, `Majat`, `Malon`, `Mangalore`, `Manti`, `Mantrin`, `Marklar`, `Marmosian`, `Marro`, `Martian`, `Martian`, `Martian`, `Martian`, `Martian`, `Martian`, `Martian`, `Martian`, `Martian`, `Martian Manhunter`, `Marzoid`, `Masari`, `Master`, `Mazian`, `Mebe`, `Mechanon`, `Medusan`, `Meehook`, `Meekrob`, `Meklar`, `Melconian`, `Melmacian`, `Melnorme`, `Melotian`, `Memory Donk`, `Menoptra`, `Mentor`, `Merseian`, `Merewif`, `Mesklinite`, `Methorian`, `Metroid`, `Metron`, `Micro`, `Microman`, `Micronoid`, `Minbari`, `Minion`, `Minosian`, `Mintakan`, `Miradorn`, `Misha`, `Mmrnmhrm`, `Mon Calamari`, `Mondoshawan`, `Monoid`, `Mogadorian`, `Moon Dragon`, `Moonflower`, `Mooninite`, `Mork`, `Morlock`, `Morlock`, `Morok`, `Mor-Taxan`, `Morthren`, `Movellan`, `Moxx of Balhoon`, `Mri`, `Mr. Saturn`, `Mrrshan`, `Mudokon`, `Muffalo`, `Mugato`, `Muton`, `Mutzachan`, `Muuh`, `Mycon`, `The Mysteron`,
			//N
			`Naalu`, `Naglon`, `Nairnama`, `Namekian`, `Naram`, `Narn`, `Nebulon`, "Nhar-Gh`Ok", "Na`vi", `Nausicaan`, `Neadlehead`, `Nosedeenian`, `Nebari`, `Nebular`, `Necri`, `Necrofriggian`, `Necron`, `Necroton`, `Nemet`, `Nephilim`, `Neptunian`, `Neutral`, `New God`, `New Orion`, `Nibblonian`, `Nicassar`, `Niea`, `Nietzschean`, `Nihilanth`, `Nimbuloid`, "N`Kull", `Nitros Oxide`, `Nodulian`, `Nomad`, `Nome`, `Nommo`, "N`Orr", `Novu`, `Nox`, `Nyrond`,
			//O
			`Oan`, `Oanne`, `Oankali`, `Oasian`, "Ob`enn", `Ocampa`, `Oculon`, `Ogri`, `Ogron`, `Old One`, `Omicronians / Poppler`, `Ood`, `Optera`, `Orandoan`, `Orfa`, `Organian`, `Ori`, `Orion`, `Orion Rogue`, `Orion`, `Orishan`, `Orkan`, `Ork`, `Orsian`, `Orthean`, `Ortog`, `Orz`, `Oscar`, `Osirian`, `Outsider`, `Overlord`, `Owa`,
			//P
			"Paan`uri", `Pah Wraith`, `Pak, aka Protector`, `Pakled`, `Parillatian`, `Pascalene`, `Pascian`, "Pa`utu`ril", `Pemalite`, `Pentapod`, `The People from the Home`, `Pequenino`, `Petrosapien`, `Pfhor`, `Pfifltrig`, `Phagor`, `Phalanx`, `Phentari`, `Phibian`, `Phleebhutinski`, `Phosphorescent Maze Midget`, `Phtagur`, "Pierson`s Puppeteer", `Pilot`, `Pinkunz`, `Pisciss Premann`, `Pisciss Volann`, `Pkunk`, `Planet Jacker`, `Plerkappi`, "P`lod", `Plookesian`, `Plorgonarian`, `Ploxi`, `Plutonian`, `Pnume`, `Polionation Tech`, `Polymorph`, `Porquinho`, `Posleen`, `Prʔ*tans`, `Prawn`, `Precursor`, `Predator alien`, `Prime`, `The Prin`, `Private Tamama`, `Progenitor`, `Promethian`, "Prophets (San `Shyuum)", `Prophet`, `Protean`, `Prothean`, `Protoculture`, `Protos`, `Psilon`, `Psiren`, `Psychlo`, `Psychon`, `The Puppet Master`, `Purple Alien`, `Pyrian`, `Pyronite`, `Pyrovile`, `Python Lizard`,
			//Q
			`Q`, "Qou`thala", `Quarian`, `Quark`, `Quarren, aka "Squid Head"`, `Quintaglio`, `Quintesson`, `Qwardian`,
			//R
			`Raa`, `Rabotev`, `The Race`, `The Race`, `Rachni`, `Raiel`, `Rako-Gorda`, `Ram Python`, `Rancor`, `Raptor`, `Raxacoricofallapatorian`, `Reaper`, `Reaper`, `Reaper`, `Regul`, `Relgarian`, `Reman`, "Re`ol", `Replicator`, "Re`tu", `Rigellian`, `Rigellians (including Kang and Kodos)`, `Riim`, `Rill`, `Rimerunner`, `Rimmer`, `Risian`, `Robotech Master`, `Rodian`, `Roger the Alien`, `Rogue Simulant`, `Romulan`, `Rull`, `Rutan`, `Ryouko Asakura`, `Ryqril`,
			//S
			`Saiyan`, `Sakkra`, `Salayan`, `Sal-Kadeem`, `Salarian`, `Samaan`, "San `Shyuum", `Sandworm`, `Sangheili`, `Santeeian`, `Sarien`, `Sarturian`, `Sathar`, `Sau-Bau`, `Scarran`, `Schniarfeur`, `Scorvian`, `Screwhead`, `Scrin`, `Scwozzwort`, `Sdanli`, `Sebacean`, `Second through Last Men`, `Sectoid`, `Seeker`, `Selay`, `Selkath`, `Seraphim`, `Séroni`, `Sevod`, `Shadok`, `Shadow`, `Shalka`, `Sharrh`, `Sheeda`, `Sheliak`, `Shevar`, `Sheyang`, "Shi`ar", `Shinari`, `Shingouz`, `Shivan`, `Shofixti`, `Shoggoth`, `Sholan`, `Shonunin`, `Shroob`, `Silacoid`, `Silfen`, `Silicoid`, `Sirian`, `Skaarj`, `Skedar`, `Skinnie`, `Skolarian`, `Skroderider`, `Skrull`, `Slaughtering Rat People`, `Slavers (see Thrint)`, `Slitheen`, `Slylandro`, `Snakemen`, `Snark`, `Snotling`, `Snovemdoma`, `SoI-002`, `Solari`, `Solomon Family`, `Solon`, `Solrock`, "Son`a", `Sonorosian`, `Sontaran`, `Soomanii`, `Sorn`, `Soro`, `Sosiqui`, `Southern Alien`, `Space Chicken`, `Space Pirate`, `Spathi`, `Specie`, `Species 8472`, "`pht", `Spibbley`, `Spiridon`, `Spirit`, `Splixson`, `Squirp`, `Squiz-Quijy`, `Ssora`, `Ssvapi`, `Starmen`, `Strogg`, `Stsho`, `Stu`, `Sugarbellie`, `Suliban`, `Sun-Dog`, `Sung`, "Supox", "Serrakin", "Swaparaman", "Sycorax", "Sye-Men", "Sykarian", "Symbiote", "Syreen",
			//T
			`Taalo`, `Tachidi`, `Taelon`, `Tagorian`, `Taiidan`, `Tak Tak`, `Talan`, `Talarian`, `Talaxian`, `Tallerian`, `Talokian`, `Talpaedan`, `Talosian`, `Tamaranean`, `Tamarian`, `Tandaran`, `Tandu`, `Tangean`, `Tanndai Techknight`, `Tarellian`, `Tarentatek`, `Targum`, `Tarkan`, `Taronyu`, `Tarka`, `Tasoth`, `Tatanga`, `Tau`, `Tavlek`, `Taxxon`, "Tc`a", `Tchoung`, `Technarchy`, `Tecresean`, `Tellarite`, `Temarkian`, `Tenctonese`, `Tenebrian`, `Tentaculat`, `Terileptil`, `Terra Novan`, `Terran`, `Terrellian`, `Terrian`, `Tetramand`, `Tetrap`, `Thalan`, `Thalian`, `Thalonian`, `Thal`, `Than-Thre-Kull`, `Thanagarian`, `Thargoid`, `Thasian`, `Thep Khufan`, `Therbian (see Aaamazzarite)`, `Thermian`, `Theron`, `The Other`, `The Thing`, `The Stranger`, `Tholian`, `Thraddash`, `Thranx`, `Thrint`, `Thrumbo`, `Thuranin`, `Tiberian`, `Tilikanthua`, `Time Lords of Gallifrey`, `Tine`, `Titanian`, `Titanide`, `Titan`, "T`Lani", `Tleilaxu`, `Tnuctip`, `Toclafane`, "Tok`ra", "To`kustar", `Tollan`, `Toluen`, `Tony`, `Torian`, "To`ul`h", `Trabe`, `Tractator`, `Traeki`, `Tralfamadorian`, `Tran`, `Transylian`, `Transylvanian`, `Transformer`, `Trandoshan`, `Traskan`, `Treecat`, `Treen`, `Tribble`, `Triceraton`, `Trilarian`, "Trillion", "Trill", "Trinoc", "Triscene", "Troll", "Tromite", "Tryvuulian", "Tsiongi", "Tsufurujin", "Tunnel Maker", "Tusken Raider", "Turian", "Twi`lek", "Twinsunian", "Twonkie", "Tyranid", "Tzenkethi",
			//U
			`Ul-Mor`, `Umgah`, `Una`, `Uncreated`, `Unggoy`, `Ungooma`, `Unioc`, `Unity`, `Ur-Quan Kohr-Ah`, `Ur-Quan Kzer-Za`, `Urpney`, "U`tani", `Utrom`, `Utwig`, `Uxorite`,
			//V
			`Vaadwaur`, `Vademon`, `Vanacancia`, `Vanryn`, `Vardian`, `Varga plant`, `Vardrag`, `Vargr`, `Vasari`, `Ventrexian`, `Vashta Nerada`, `Vasudan`, `Vauvusar`, `Vaxasaurian`, `Velantian`, `Venek`, `Venom and Carnage`, `Venom grub`, `Verga`, `Vervoid`, `Vespid`, `Vhorwed`, `Vidiian`, `Vilani`, `Viltrumite`, `Vinean`, `Visitor`, `Visitor`, `Vogon`, `Void whale`, `Vok`, `Vorc`, `Vorcarian bloodtracker`, `Vorlon`, `Vorta`, `Vortex life form`, `Vortian`, `Vorticon`, `Vortigaunt`, `Vortisaur`, `Voth`, `Vroarscan`, `Vrusk`, `Vulcan`, `Vullard`, `Vulpimancer`, `VUX`, `Vyrium`, `Vyro-Ingo`, "Vy`keen",
			//W
			`Wadi`, `Waldahudin`, `Wanderer`, "Wang`s Carpet", `Wankh`, `The Watcher`, `Waterseeker`, `Weeping Angel`, `Weevil`, `Whrloo`, `Willis the Bouncer`, `Winathian`, `Wirrn`, `Wisp`, `Wogneer`, `Wolfweed`, `Wookiee`, `Wraith`, "W`rkncacnter",
			//X
			`X-ist`, `X-Naut`, `X Parasite`, `Xandarian`, `Xanthuan`, `Xarian`, `Xarquid`, `Xchagger`, "Xel`Naga", `Xenexian`, `Xenomorph`, `Xenu`, `Xiang`, `Xilian`, `Xindi`, `Xorda`, `Xxcha`, `Xyrillian`,
			//Y
			`Yacatisma`, `Yag-Kosha`, "Yanme`e", `Yautja`, `Yazirian`, `Yeerk`, `Yehat`, `Yeti`, `Yilane`, `Yip-Yip`, `Ylii`, `Yolkian`, `Yomingan`, `Yor`, `Yorn`, `Ythrian`, `Yridian`, `Yugopotamian`, `Yuuzhan Vong`,
			//Z
			`Zabrak`, `Zakdorn`, `Zaldan`, `Zalkonian`, `Zarbi`, `Zebesian`, `Zen Rigeln`, `Zen-Whoberi`, `Zenetan`, `Zenn-Lavian`, `Zentradi`, `Zerg`, `Zhodani`, `Zin`, `Zinoboppian`, `Zirkonian`, `Zisuili`, `Zog`, `Zoni`, `Zoq-Fot-Pik`, `Zorgon`, `Zorgon`, `Zygon`,
		},
	}
}

// NameAlien generates a random name
// credits: https://github.com/moby/moby/blob/master/pkg/namesgenerator/names-generator.go#L852 with modifications
func (n *famousAliensNamer) NameAlien() string {
	name := cases.Title(language.Und).String(n.left[Random.Intn(len(n.left))]) + " " + n.right[Random.Intn(len(n.right))]
	return name
}

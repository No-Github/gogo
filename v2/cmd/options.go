package cmd

type InputOption struct {
	IP                string `short:"i" long:"ip" description:"IP/CIDR, support comma-split ip/cidr, e.g. 192.168.1.1/24,172.16.1.1/24"`
	Ports             string `short:"p" long:"port" description:"Port, support comma-split preset('-P port' show all preset), range, alias port, e.g. top2,mysql,12345,10000-10100,oxid,smb"`
	JsonFile          string `short:"j" long:"json" description:"File,  list of IP/CIDR"`
	ListFile          string `short:"l" long:"list" description:"File,  previous results file e.g. -j 1.dat1 or list of colon-split ip:port, e.g. 123.123.123.123:123"`
	IsListInput       bool   `short:"L" description:"same as -l, input from stdin"`
	IsJsonInput       bool   `short:"J" description:"same as -j, input from stdin"`
	WorkFlowName      string `short:"w" long:"workflow" description:"String, workflow name('-P workflow' show all workflow)"`
	IsWorkFlow        bool   `short:"W" description:"same as -w, input from stdin"`
	FormatterFilename string `short:"F" long:"format" description:"File, to be formatted result file"` // 待格式化文件名
}

type OutputOption struct {
	Filename    string `short:"f" long:"file" description:"String, output filename"`
	FilePath    string `long:"path" description:"String, output file path"`
	Outputf     string `short:"o" long:"output" default:"default" description:"String,cmdline output format"`
	FileOutputf string `short:"O" long:"file-output" default:"default" description:"String, file output format"` // 输出格式
	AutoFile    bool   `long:"af" description:"Bool, auto choice filename"`                                      // 自动生成格式化文件名
	HiddenFile  bool   `long:"hf" description:"Bool, auto choice hidden filename"`                               // 启用自动隐藏文件
	Compress    bool   `short:"C" long:"compress" description:"Bool, close compress output file"`
	Clean       bool   `short:"c" long:"clean" description:"Bool, close stdout output"` // 是否开启命令行输出扫描结果
	Quiet       bool   `short:"q" long:"quiet" description:"Bool, close log output"`    // 是否开启命令行输出日志
}

type SmartOption struct {
	Mod       string `short:"m" long:"mod" choice:"s" choice:"ss" choice:"default" choice:"sc" default:"default" description:"String, smart mod"` // 扫描模式
	Ping      bool   `long:"ping" description:"Bool, alive pre-scan"`
	NoScan    bool   `long:"no" description:"Bool, no-plugin, only smart scan"`
	SmartPort string `long:"sp" default:"default" description:"String, smart-port-probe"` // 启发式扫描预设探针
	IpProbe   string `long:"ipp"  default:"default" description:"String, IP-probe"`
}

type MiscOption struct {
	Key         string `short:"k" long:"key" description:"String, file encrypt key"`
	Ver         bool   `long:"version" description:"Bool, show version"`                                                                     // 输出版本号
	Printer     string `short:"P" choice:"port" choice:"workflow" choice:"nuclei" choice:"extract" description:"String, show preset config"` // 输出特定的预设
	Debug       bool   `long:"debug" description:"Bool, show debug info"`
	PluginDebug bool   `long:"plugin-debug" description:"Bool, show plugin debug stack"`
	Proxy       string `long:"proxy" description:"String, socks5 proxy url, e.g. socks5://127.0.0.1:11111"`
}

type ConfigOption struct {
	Threads     int      `short:"t" long:"thread" description:"Int, concurrent thread number"`    // 线程数
	Exploit     bool     `short:"e" long:"exploit" description:"Bool,enable nuclei exploit scan"` // 启用漏洞扫描
	Version     bool     `short:"v" long:"verbose" description:"Bool, enable active finger scan"` // version level1
	PortSpray   bool     `short:"s" long:"spray" description:"Bool, enable port-first spray generator. if ports number > 500, auto enable"`
	NoSpray     bool     `long:"no-spray" description:"Bool, force to close spray"`
	ExploitName string   `short:"E" long:"exploit-name" description:"String, specify nuclei template name"` // 指定漏扫poc名字
	ExploitFile string   `long:"ef" description:"String, load specified templates file "`                   // 指定漏扫文件
	filters     []string `long:"filter" description:"String, filter formatting(-F) results "`
	payloads    []string `long:"payload" description:"String, specify nuclei payload"`
	extract     []string `long:"extract" description:"String, custom extract regexp"`
	extracts    string   `long:"extracts" description:"String, choice preset extract regexp, e.g. -extracts ip,url"`
	Delay       int      `short:"d" long:"timeout" default:"2" description:"Int, socket and http timeout"`
	HttpsDelay  int      `short:"D" long:"ssl-timeout" default:"2" description:"Int, ssl and https timeout"`
	SuffixStr   string   `long:"suffix" description:"String, url path"`
}
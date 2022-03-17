package datamarker

var DataMarker2007 = map[string]float64{
	//***********电能量数据精度***********
	"00000000": 100, //组合有功总电能
	"00010000": 100, //正向有功总电能
	"00020000": 100, //反向有功总电能
	"00150000": 100, //A相正向有功总电能
	"00160000": 100, //A相反向有功总电能
	"00290000": 100, //B相正向有功总电能
	"002A0000": 100, //B相反向有功总电能
	"003D0000": 100, //C相正向有功总电能
	"003E0000": 100, //C相反向有功总电能
	//***********变量数据***********
	"02010100": 10,    //A相电压
	"02010200": 10,    //B相电压
	"02010300": 10,    //C相电压
	"02020100": 1000,  //A相电流
	"02020200": 1000,  //B相电流
	"02020300": 1000,  //C相电流
	"02030000": 10000, //瞬时总有功功率
	"02040000": 10000, //瞬时总无功功率
	"02060000": 1000,  //总功率因数
	"02800002": 100,   //电网频率

}

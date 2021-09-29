package main

import (
	"fmt"
	"flag"
	"os"
	"regexp"
	"strings"

	"github.com/gookit/color"
)


var VERSION = "colorview 0.1.1"


/* https://gitlab.freedesktop.org/xorg/app/rgb/raw/master/rgb.txt */
var x11Colors map[string]color.RGBColor = map[string]color.RGBColor{
	"snow"                 : color.RGB(255,250,250,true),
	"ghostwhite"           : color.RGB(248,248,255,true),
	"whitesmoke"           : color.RGB(245,245,245,true),
	"gainsboro"            : color.RGB(220,220,220,true),
	"floralwhite"          : color.RGB(255,250,240,true),
	"oldlace"              : color.RGB(253,245,230,true),
	"linen"                : color.RGB(250,240,230,true),
	"antiquewhite"         : color.RGB(250,235,215,true),
	"papayawhip"           : color.RGB(255,239,213,true),
	"blanchedalmond"       : color.RGB(255,235,205,true),
	"bisque"               : color.RGB(255,228,196,true),
	"peachpuff"            : color.RGB(255,218,185,true),
	"navajowhite"          : color.RGB(255,222,173,true),
	"moccasin"             : color.RGB(255,228,181,true),
	"cornsilk"             : color.RGB(255,248,220,true),
	"ivory"                : color.RGB(255,255,240,true),
	"lemonchiffon"         : color.RGB(255,250,205,true),
	"seashell"             : color.RGB(255,245,238,true),
	"honeydew"             : color.RGB(240,255,240,true),
	"mintcream"            : color.RGB(245,255,250,true),
	"azure"                : color.RGB(240,255,255,true),
	"aliceblue"            : color.RGB(240,248,255,true),
	"lavender"             : color.RGB(230,230,250,true),
	"lavenderblush"        : color.RGB(255,240,245,true),
	"mistyrose"            : color.RGB(255,228,225,true),
	"white"                : color.RGB(255,255,255,true),
	"black"                : color.RGB(0,0,0,true),
	"darkslate gray"       : color.RGB(47,79,79,true),
	"darkslategray"        : color.RGB(47,79,79,true),
	"darkslate grey"       : color.RGB(47,79,79,true),
	"darkslategrey"        : color.RGB(47,79,79,true),
	"dimgray"              : color.RGB(105,105,105,true),
	"dimgrey"              : color.RGB(105,105,105,true),
	"slategray"            : color.RGB(112,128,144,true),
	"slategrey"            : color.RGB(112,128,144,true),
	"lightslate gray"      : color.RGB(119,136,153,true),
	"lightslategray"       : color.RGB(119,136,153,true),
	"lightslate grey"      : color.RGB(119,136,153,true),
	"lightslategrey"       : color.RGB(119,136,153,true),
	"gray"                 : color.RGB(190,190,190,true),
	"grey"                 : color.RGB(190,190,190,true),
	"x11gray"              : color.RGB(190,190,190,true),
	"x11grey"              : color.RGB(190,190,190,true),
	"webgray"              : color.RGB(128,128,128,true),
	"webgrey"              : color.RGB(128,128,128,true),
	"lightgrey"            : color.RGB(211,211,211,true),
	"lightgray"            : color.RGB(211,211,211,true),
	"midnightblue"         : color.RGB(25,25,112,true),
	"navy"                 : color.RGB(0,0,128,true),
	"navyblue"             : color.RGB(0,0,128,true),
	"cornflowerblue"       : color.RGB(100,149,237,true),
	"darkslate blue"       : color.RGB(72,61,139,true),
	"darkslateblue"        : color.RGB(72,61,139,true),
	"slateblue"            : color.RGB(106,90,205,true),
	"mediumslate blue"     : color.RGB(123,104,238,true),
	"mediumslateblue"      : color.RGB(123,104,238,true),
	"lightslate blue"      : color.RGB(132,112,255,true),
	"lightslateblue"       : color.RGB(132,112,255,true),
	"mediumblue"           : color.RGB(0,0,205,true),
	"royalblue"            : color.RGB(65,105,225,true),
	"blue"                 : color.RGB(0,0,255,true),
	"dodgerblue"           : color.RGB(30,144,255,true),
	"deepsky blue"         : color.RGB(0,191,255,true),
	"deepskyblue"          : color.RGB(0,191,255,true),
	"skyblue"              : color.RGB(135,206,235,true),
	"lightsky blue"        : color.RGB(135,206,250,true),
	"lightskyblue"         : color.RGB(135,206,250,true),
	"steelblue"            : color.RGB(70,130,180,true),
	"lightsteel blue"      : color.RGB(176,196,222,true),
	"lightsteelblue"       : color.RGB(176,196,222,true),
	"lightblue"            : color.RGB(173,216,230,true),
	"powderblue"           : color.RGB(176,224,230,true),
	"paleturquoise"        : color.RGB(175,238,238,true),
	"darkturquoise"        : color.RGB(0,206,209,true),
	"mediumturquoise"      : color.RGB(72,209,204,true),
	"turquoise"            : color.RGB(64,224,208,true),
	"cyan"                 : color.RGB(0,255,255,true),
	"aqua"                 : color.RGB(0,255,255,true),
	"lightcyan"            : color.RGB(224,255,255,true),
	"cadetblue"            : color.RGB(95,158,160,true),
	"mediumaquamarine"     : color.RGB(102,205,170,true),
	"aquamarine"           : color.RGB(127,255,212,true),
	"darkgreen"            : color.RGB(0,100,0,true),
	"darkolive green"      : color.RGB(85,107,47,true),
	"darkolivegreen"       : color.RGB(85,107,47,true),
	"darksea green"        : color.RGB(143,188,143,true),
	"darkseagreen"         : color.RGB(143,188,143,true),
	"seagreen"             : color.RGB(46,139,87,true),
	"mediumsea green"      : color.RGB(60,179,113,true),
	"mediumseagreen"       : color.RGB(60,179,113,true),
	"lightsea green"       : color.RGB(32,178,170,true),
	"lightseagreen"        : color.RGB(32,178,170,true),
	"palegreen"            : color.RGB(152,251,152,true),
	"springgreen"          : color.RGB(0,255,127,true),
	"lawngreen"            : color.RGB(124,252,0,true),
	"green"                : color.RGB(0,255,0,true),
	"lime"                 : color.RGB(0,255,0,true),
	"x11green"             : color.RGB(0,255,0,true),
	"webgreen"             : color.RGB(0,128,0,true),
	"chartreuse"           : color.RGB(127,255,0,true),
	"mediumspring green"   : color.RGB(0,250,154,true),
	"mediumspringgreen"    : color.RGB(0,250,154,true),
	"greenyellow"          : color.RGB(173,255,47,true),
	"limegreen"            : color.RGB(50,205,50,true),
	"yellowgreen"          : color.RGB(154,205,50,true),
	"forestgreen"          : color.RGB(34,139,34,true),
	"olivedrab"            : color.RGB(107,142,35,true),
	"darkkhaki"            : color.RGB(189,183,107,true),
	"khaki"                : color.RGB(240,230,140,true),
	"palegoldenrod"        : color.RGB(238,232,170,true),
	"lightgoldenrod yellow": color.RGB(250,250,210,true),
	"lightgoldenrodyellow" : color.RGB(250,250,210,true),
	"lightyellow"          : color.RGB(255,255,224,true),
	"yellow"               : color.RGB(255,255,0,true),
	"gold"                 : color.RGB(255,215,0,true),
	"lightgoldenrod"       : color.RGB(238,221,130,true),
	"goldenrod"            : color.RGB(218,165,32,true),
	"darkgoldenrod"        : color.RGB(184,134,11,true),
	"rosybrown"            : color.RGB(188,143,143,true),
	"indianred"            : color.RGB(205,92,92,true),
	"saddlebrown"          : color.RGB(139,69,19,true),
	"sienna"               : color.RGB(160,82,45,true),
	"peru"                 : color.RGB(205,133,63,true),
	"burlywood"            : color.RGB(222,184,135,true),
	"beige"                : color.RGB(245,245,220,true),
	"wheat"                : color.RGB(245,222,179,true),
	"sandybrown"           : color.RGB(244,164,96,true),
	"tan"                  : color.RGB(210,180,140,true),
	"chocolate"            : color.RGB(210,105,30,true),
	"firebrick"            : color.RGB(178,34,34,true),
	"brown"                : color.RGB(165,42,42,true),
	"darksalmon"           : color.RGB(233,150,122,true),
	"salmon"               : color.RGB(250,128,114,true),
	"lightsalmon"          : color.RGB(255,160,122,true),
	"orange"               : color.RGB(255,165,0,true),
	"darkorange"           : color.RGB(255,140,0,true),
	"coral"                : color.RGB(255,127,80,true),
	"lightcoral"           : color.RGB(240,128,128,true),
	"tomato"               : color.RGB(255,99,71,true),
	"orangered"            : color.RGB(255,69,0,true),
	"red"                  : color.RGB(255,0,0,true),
	"hotpink"              : color.RGB(255,105,180,true),
	"deeppink"             : color.RGB(255,20,147,true),
	"pink"                 : color.RGB(255,192,203,true),
	"lightpink"            : color.RGB(255,182,193,true),
	"paleviolet red"       : color.RGB(219,112,147,true),
	"palevioletred"        : color.RGB(219,112,147,true),
	"maroon"               : color.RGB(176,48,96,true),
	"x11maroon"            : color.RGB(176,48,96,true),
	"webmaroon"            : color.RGB(128,0,0,true),
	"mediumviolet red"     : color.RGB(199,21,133,true),
	"mediumvioletred"      : color.RGB(199,21,133,true),
	"violetred"            : color.RGB(208,32,144,true),
	"magenta"              : color.RGB(255,0,255,true),
	"fuchsia"              : color.RGB(255,0,255,true),
	"violet"               : color.RGB(238,130,238,true),
	"plum"                 : color.RGB(221,160,221,true),
	"orchid"               : color.RGB(218,112,214,true),
	"mediumorchid"         : color.RGB(186,85,211,true),
	"darkorchid"           : color.RGB(153,50,204,true),
	"darkviolet"           : color.RGB(148,0,211,true),
	"blueviolet"           : color.RGB(138,43,226,true),
	"purple"               : color.RGB(160,32,240,true),
	"x11purple"            : color.RGB(160,32,240,true),
	"webpurple"            : color.RGB(128,0,128,true),
	"mediumpurple"         : color.RGB(147,112,219,true),
	"thistle"              : color.RGB(216,191,216,true),
	"snow1"                : color.RGB(255,250,250,true),
	"snow2"                : color.RGB(238,233,233,true),
	"snow3"                : color.RGB(205,201,201,true),
	"snow4"                : color.RGB(139,137,137,true),
	"seashell1"            : color.RGB(255,245,238,true),
	"seashell2"            : color.RGB(238,229,222,true),
	"seashell3"            : color.RGB(205,197,191,true),
	"seashell4"            : color.RGB(139,134,130,true),
	"antiquewhite1"        : color.RGB(255,239,219,true),
	"antiquewhite2"        : color.RGB(238,223,204,true),
	"antiquewhite3"        : color.RGB(205,192,176,true),
	"antiquewhite4"        : color.RGB(139,131,120,true),
	"bisque1"              : color.RGB(255,228,196,true),
	"bisque2"              : color.RGB(238,213,183,true),
	"bisque3"              : color.RGB(205,183,158,true),
	"bisque4"              : color.RGB(139,125,107,true),
	"peachpuff1"           : color.RGB(255,218,185,true),
	"peachpuff2"           : color.RGB(238,203,173,true),
	"peachpuff3"           : color.RGB(205,175,149,true),
	"peachpuff4"           : color.RGB(139,119,101,true),
	"navajowhite1"         : color.RGB(255,222,173,true),
	"navajowhite2"         : color.RGB(238,207,161,true),
	"navajowhite3"         : color.RGB(205,179,139,true),
	"navajowhite4"         : color.RGB(139,121,94,true),
	"lemonchiffon1"        : color.RGB(255,250,205,true),
	"lemonchiffon2"        : color.RGB(238,233,191,true),
	"lemonchiffon3"        : color.RGB(205,201,165,true),
	"lemonchiffon4"        : color.RGB(139,137,112,true),
	"cornsilk1"            : color.RGB(255,248,220,true),
	"cornsilk2"            : color.RGB(238,232,205,true),
	"cornsilk3"            : color.RGB(205,200,177,true),
	"cornsilk4"            : color.RGB(139,136,120,true),
	"ivory1"               : color.RGB(255,255,240,true),
	"ivory2"               : color.RGB(238,238,224,true),
	"ivory3"               : color.RGB(205,205,193,true),
	"ivory4"               : color.RGB(139,139,131,true),
	"honeydew1"            : color.RGB(240,255,240,true),
	"honeydew2"            : color.RGB(224,238,224,true),
	"honeydew3"            : color.RGB(193,205,193,true),
	"honeydew4"            : color.RGB(131,139,131,true),
	"lavenderblush1"       : color.RGB(255,240,245,true),
	"lavenderblush2"       : color.RGB(238,224,229,true),
	"lavenderblush3"       : color.RGB(205,193,197,true),
	"lavenderblush4"       : color.RGB(139,131,134,true),
	"mistyrose1"           : color.RGB(255,228,225,true),
	"mistyrose2"           : color.RGB(238,213,210,true),
	"mistyrose3"           : color.RGB(205,183,181,true),
	"mistyrose4"           : color.RGB(139,125,123,true),
	"azure1"               : color.RGB(240,255,255,true),
	"azure2"               : color.RGB(224,238,238,true),
	"azure3"               : color.RGB(193,205,205,true),
	"azure4"               : color.RGB(131,139,139,true),
	"slateblue1"           : color.RGB(131,111,255,true),
	"slateblue2"           : color.RGB(122,103,238,true),
	"slateblue3"           : color.RGB(105,89,205,true),
	"slateblue4"           : color.RGB(71,60,139,true),
	"royalblue1"           : color.RGB(72,118,255,true),
	"royalblue2"           : color.RGB(67,110,238,true),
	"royalblue3"           : color.RGB(58,95,205,true),
	"royalblue4"           : color.RGB(39,64,139,true),
	"blue1"                : color.RGB(0,0,255,true),
	"blue2"                : color.RGB(0,0,238,true),
	"blue3"                : color.RGB(0,0,205,true),
	"blue4"                : color.RGB(0,0,139,true),
	"dodgerblue1"          : color.RGB(30,144,255,true),
	"dodgerblue2"          : color.RGB(28,134,238,true),
	"dodgerblue3"          : color.RGB(24,116,205,true),
	"dodgerblue4"          : color.RGB(16,78,139,true),
	"steelblue1"           : color.RGB(99,184,255,true),
	"steelblue2"           : color.RGB(92,172,238,true),
	"steelblue3"           : color.RGB(79,148,205,true),
	"steelblue4"           : color.RGB(54,100,139,true),
	"deepskyblue1"         : color.RGB(0,191,255,true),
	"deepskyblue2"         : color.RGB(0,178,238,true),
	"deepskyblue3"         : color.RGB(0,154,205,true),
	"deepskyblue4"         : color.RGB(0,104,139,true),
	"skyblue1"             : color.RGB(135,206,255,true),
	"skyblue2"             : color.RGB(126,192,238,true),
	"skyblue3"             : color.RGB(108,166,205,true),
	"skyblue4"             : color.RGB(74,112,139,true),
	"lightskyblue1"        : color.RGB(176,226,255,true),
	"lightskyblue2"        : color.RGB(164,211,238,true),
	"lightskyblue3"        : color.RGB(141,182,205,true),
	"lightskyblue4"        : color.RGB(96,123,139,true),
	"slategray1"           : color.RGB(198,226,255,true),
	"slategray2"           : color.RGB(185,211,238,true),
	"slategray3"           : color.RGB(159,182,205,true),
	"slategray4"           : color.RGB(108,123,139,true),
	"lightsteelblue1"      : color.RGB(202,225,255,true),
	"lightsteelblue2"      : color.RGB(188,210,238,true),
	"lightsteelblue3"      : color.RGB(162,181,205,true),
	"lightsteelblue4"      : color.RGB(110,123,139,true),
	"lightblue1"           : color.RGB(191,239,255,true),
	"lightblue2"           : color.RGB(178,223,238,true),
	"lightblue3"           : color.RGB(154,192,205,true),
	"lightblue4"           : color.RGB(104,131,139,true),
	"lightcyan1"           : color.RGB(224,255,255,true),
	"lightcyan2"           : color.RGB(209,238,238,true),
	"lightcyan3"           : color.RGB(180,205,205,true),
	"lightcyan4"           : color.RGB(122,139,139,true),
	"paleturquoise1"       : color.RGB(187,255,255,true),
	"paleturquoise2"       : color.RGB(174,238,238,true),
	"paleturquoise3"       : color.RGB(150,205,205,true),
	"paleturquoise4"       : color.RGB(102,139,139,true),
	"cadetblue1"           : color.RGB(152,245,255,true),
	"cadetblue2"           : color.RGB(142,229,238,true),
	"cadetblue3"           : color.RGB(122,197,205,true),
	"cadetblue4"           : color.RGB(83,134,139,true),
	"turquoise1"           : color.RGB(0,245,255,true),
	"turquoise2"           : color.RGB(0,229,238,true),
	"turquoise3"           : color.RGB(0,197,205,true),
	"turquoise4"           : color.RGB(0,134,139,true),
	"cyan1"                : color.RGB(0,255,255,true),
	"cyan2"                : color.RGB(0,238,238,true),
	"cyan3"                : color.RGB(0,205,205,true),
	"cyan4"                : color.RGB(0,139,139,true),
	"darkslategray1"       : color.RGB(151,255,255,true),
	"darkslategray2"       : color.RGB(141,238,238,true),
	"darkslategray3"       : color.RGB(121,205,205,true),
	"darkslategray4"       : color.RGB(82,139,139,true),
	"aquamarine1"          : color.RGB(127,255,212,true),
	"aquamarine2"          : color.RGB(118,238,198,true),
	"aquamarine3"          : color.RGB(102,205,170,true),
	"aquamarine4"          : color.RGB(69,139,116,true),
	"darkseagreen1"        : color.RGB(193,255,193,true),
	"darkseagreen2"        : color.RGB(180,238,180,true),
	"darkseagreen3"        : color.RGB(155,205,155,true),
	"darkseagreen4"        : color.RGB(105,139,105,true),
	"seagreen1"            : color.RGB(84,255,159,true),
	"seagreen2"            : color.RGB(78,238,148,true),
	"seagreen3"            : color.RGB(67,205,128,true),
	"seagreen4"            : color.RGB(46,139,87,true),
	"palegreen1"           : color.RGB(154,255,154,true),
	"palegreen2"           : color.RGB(144,238,144,true),
	"palegreen3"           : color.RGB(124,205,124,true),
	"palegreen4"           : color.RGB(84,139,84,true),
	"springgreen1"         : color.RGB(0,255,127,true),
	"springgreen2"         : color.RGB(0,238,118,true),
	"springgreen3"         : color.RGB(0,205,102,true),
	"springgreen4"         : color.RGB(0,139,69,true),
	"green1"               : color.RGB(0,255,0,true),
	"green2"               : color.RGB(0,238,0,true),
	"green3"               : color.RGB(0,205,0,true),
	"green4"               : color.RGB(0,139,0,true),
	"chartreuse1"          : color.RGB(127,255,0,true),
	"chartreuse2"          : color.RGB(118,238,0,true),
	"chartreuse3"          : color.RGB(102,205,0,true),
	"chartreuse4"          : color.RGB(69,139,0,true),
	"olivedrab1"           : color.RGB(192,255,62,true),
	"olivedrab2"           : color.RGB(179,238,58,true),
	"olivedrab3"           : color.RGB(154,205,50,true),
	"olivedrab4"           : color.RGB(105,139,34,true),
	"darkolivegreen1"      : color.RGB(202,255,112,true),
	"darkolivegreen2"      : color.RGB(188,238,104,true),
	"darkolivegreen3"      : color.RGB(162,205,90,true),
	"darkolivegreen4"      : color.RGB(110,139,61,true),
	"khaki1"               : color.RGB(255,246,143,true),
	"khaki2"               : color.RGB(238,230,133,true),
	"khaki3"               : color.RGB(205,198,115,true),
	"khaki4"               : color.RGB(139,134,78,true),
	"lightgoldenrod1"      : color.RGB(255,236,139,true),
	"lightgoldenrod2"      : color.RGB(238,220,130,true),
	"lightgoldenrod3"      : color.RGB(205,190,112,true),
	"lightgoldenrod4"      : color.RGB(139,129,76,true),
	"lightyellow1"         : color.RGB(255,255,224,true),
	"lightyellow2"         : color.RGB(238,238,209,true),
	"lightyellow3"         : color.RGB(205,205,180,true),
	"lightyellow4"         : color.RGB(139,139,122,true),
	"yellow1"              : color.RGB(255,255,0,true),
	"yellow2"              : color.RGB(238,238,0,true),
	"yellow3"              : color.RGB(205,205,0,true),
	"yellow4"              : color.RGB(139,139,0,true),
	"gold1"                : color.RGB(255,215,0,true),
	"gold2"                : color.RGB(238,201,0,true),
	"gold3"                : color.RGB(205,173,0,true),
	"gold4"                : color.RGB(139,117,0,true),
	"goldenrod1"           : color.RGB(255,193,37,true),
	"goldenrod2"           : color.RGB(238,180,34,true),
	"goldenrod3"           : color.RGB(205,155,29,true),
	"goldenrod4"           : color.RGB(139,105,20,true),
	"darkgoldenrod1"       : color.RGB(255,185,15,true),
	"darkgoldenrod2"       : color.RGB(238,173,14,true),
	"darkgoldenrod3"       : color.RGB(205,149,12,true),
	"darkgoldenrod4"       : color.RGB(139,101,8,true),
	"rosybrown1"           : color.RGB(255,193,193,true),
	"rosybrown2"           : color.RGB(238,180,180,true),
	"rosybrown3"           : color.RGB(205,155,155,true),
	"rosybrown4"           : color.RGB(139,105,105,true),
	"indianred1"           : color.RGB(255,106,106,true),
	"indianred2"           : color.RGB(238,99,99,true),
	"indianred3"           : color.RGB(205,85,85,true),
	"indianred4"           : color.RGB(139,58,58,true),
	"sienna1"              : color.RGB(255,130,71,true),
	"sienna2"              : color.RGB(238,121,66,true),
	"sienna3"              : color.RGB(205,104,57,true),
	"sienna4"              : color.RGB(139,71,38,true),
	"burlywood1"           : color.RGB(255,211,155,true),
	"burlywood2"           : color.RGB(238,197,145,true),
	"burlywood3"           : color.RGB(205,170,125,true),
	"burlywood4"           : color.RGB(139,115,85,true),
	"wheat1"               : color.RGB(255,231,186,true),
	"wheat2"               : color.RGB(238,216,174,true),
	"wheat3"               : color.RGB(205,186,150,true),
	"wheat4"               : color.RGB(139,126,102,true),
	"tan1"                 : color.RGB(255,165,79,true),
	"tan2"                 : color.RGB(238,154,73,true),
	"tan3"                 : color.RGB(205,133,63,true),
	"tan4"                 : color.RGB(139,90,43,true),
	"chocolate1"           : color.RGB(255,127,36,true),
	"chocolate2"           : color.RGB(238,118,33,true),
	"chocolate3"           : color.RGB(205,102,29,true),
	"chocolate4"           : color.RGB(139,69,19,true),
	"firebrick1"           : color.RGB(255,48,48,true),
	"firebrick2"           : color.RGB(238,44,44,true),
	"firebrick3"           : color.RGB(205,38,38,true),
	"firebrick4"           : color.RGB(139,26,26,true),
	"brown1"               : color.RGB(255,64,64,true),
	"brown2"               : color.RGB(238,59,59,true),
	"brown3"               : color.RGB(205,51,51,true),
	"brown4"               : color.RGB(139,35,35,true),
	"salmon1"              : color.RGB(255,140,105,true),
	"salmon2"              : color.RGB(238,130,98,true),
	"salmon3"              : color.RGB(205,112,84,true),
	"salmon4"              : color.RGB(139,76,57,true),
	"lightsalmon1"         : color.RGB(255,160,122,true),
	"lightsalmon2"         : color.RGB(238,149,114,true),
	"lightsalmon3"         : color.RGB(205,129,98,true),
	"lightsalmon4"         : color.RGB(139,87,66,true),
	"orange1"              : color.RGB(255,165,0,true),
	"orange2"              : color.RGB(238,154,0,true),
	"orange3"              : color.RGB(205,133,0,true),
	"orange4"              : color.RGB(139,90,0,true),
	"darkorange1"          : color.RGB(255,127,0,true),
	"darkorange2"          : color.RGB(238,118,0,true),
	"darkorange3"          : color.RGB(205,102,0,true),
	"darkorange4"          : color.RGB(139,69,0,true),
	"coral1"               : color.RGB(255,114,86,true),
	"coral2"               : color.RGB(238,106,80,true),
	"coral3"               : color.RGB(205,91,69,true),
	"coral4"               : color.RGB(139,62,47,true),
	"tomato1"              : color.RGB(255,99,71,true),
	"tomato2"              : color.RGB(238,92,66,true),
	"tomato3"              : color.RGB(205,79,57,true),
	"tomato4"              : color.RGB(139,54,38,true),
	"orangered1"           : color.RGB(255,69,0,true),
	"orangered2"           : color.RGB(238,64,0,true),
	"orangered3"           : color.RGB(205,55,0,true),
	"orangered4"           : color.RGB(139,37,0,true),
	"red1"                 : color.RGB(255,0,0,true),
	"red2"                 : color.RGB(238,0,0,true),
	"red3"                 : color.RGB(205,0,0,true),
	"red4"                 : color.RGB(139,0,0,true),
	"deeppink1"            : color.RGB(255,20,147,true),
	"deeppink2"            : color.RGB(238,18,137,true),
	"deeppink3"            : color.RGB(205,16,118,true),
	"deeppink4"            : color.RGB(139,10,80,true),
	"hotpink1"             : color.RGB(255,110,180,true),
	"hotpink2"             : color.RGB(238,106,167,true),
	"hotpink3"             : color.RGB(205,96,144,true),
	"hotpink4"             : color.RGB(139,58,98,true),
	"pink1"                : color.RGB(255,181,197,true),
	"pink2"                : color.RGB(238,169,184,true),
	"pink3"                : color.RGB(205,145,158,true),
	"pink4"                : color.RGB(139,99,108,true),
	"lightpink1"           : color.RGB(255,174,185,true),
	"lightpink2"           : color.RGB(238,162,173,true),
	"lightpink3"           : color.RGB(205,140,149,true),
	"lightpink4"           : color.RGB(139,95,101,true),
	"palevioletred1"       : color.RGB(255,130,171,true),
	"palevioletred2"       : color.RGB(238,121,159,true),
	"palevioletred3"       : color.RGB(205,104,137,true),
	"palevioletred4"       : color.RGB(139,71,93,true),
	"maroon1"              : color.RGB(255,52,179,true),
	"maroon2"              : color.RGB(238,48,167,true),
	"maroon3"              : color.RGB(205,41,144,true),
	"maroon4"              : color.RGB(139,28,98,true),
	"violetred1"           : color.RGB(255,62,150,true),
	"violetred2"           : color.RGB(238,58,140,true),
	"violetred3"           : color.RGB(205,50,120,true),
	"violetred4"           : color.RGB(139,34,82,true),
	"magenta1"             : color.RGB(255,0,255,true),
	"magenta2"             : color.RGB(238,0,238,true),
	"magenta3"             : color.RGB(205,0,205,true),
	"magenta4"             : color.RGB(139,0,139,true),
	"orchid1"              : color.RGB(255,131,250,true),
	"orchid2"              : color.RGB(238,122,233,true),
	"orchid3"              : color.RGB(205,105,201,true),
	"orchid4"              : color.RGB(139,71,137,true),
	"plum1"                : color.RGB(255,187,255,true),
	"plum2"                : color.RGB(238,174,238,true),
	"plum3"                : color.RGB(205,150,205,true),
	"plum4"                : color.RGB(139,102,139,true),
	"mediumorchid1"        : color.RGB(224,102,255,true),
	"mediumorchid2"        : color.RGB(209,95,238,true),
	"mediumorchid3"        : color.RGB(180,82,205,true),
	"mediumorchid4"        : color.RGB(122,55,139,true),
	"darkorchid1"          : color.RGB(191,62,255,true),
	"darkorchid2"          : color.RGB(178,58,238,true),
	"darkorchid3"          : color.RGB(154,50,205,true),
	"darkorchid4"          : color.RGB(104,34,139,true),
	"purple1"              : color.RGB(155,48,255,true),
	"purple2"              : color.RGB(145,44,238,true),
	"purple3"              : color.RGB(125,38,205,true),
	"purple4"              : color.RGB(85,26,139,true),
	"mediumpurple1"        : color.RGB(171,130,255,true),
	"mediumpurple2"        : color.RGB(159,121,238,true),
	"mediumpurple3"        : color.RGB(137,104,205,true),
	"mediumpurple4"        : color.RGB(93,71,139,true),
	"thistle1"             : color.RGB(255,225,255,true),
	"thistle2"             : color.RGB(238,210,238,true),
	"thistle3"             : color.RGB(205,181,205,true),
	"thistle4"             : color.RGB(139,123,139,true),
	"gray0"                : color.RGB(0,0,0,true),
	"grey0"                : color.RGB(0,0,0,true),
	"gray1"                : color.RGB(3,3,3,true),
	"grey1"                : color.RGB(3,3,3,true),
	"gray2"                : color.RGB(5,5,5,true),
	"grey2"                : color.RGB(5,5,5,true),
	"gray3"                : color.RGB(8,8,8,true),
	"grey3"                : color.RGB(8,8,8,true),
	"gray4"                : color.RGB(10,10,10,true),
	"grey4"                : color.RGB(10,10,10,true),
	"gray5"                : color.RGB(13,13,13,true),
	"grey5"                : color.RGB(13,13,13,true),
	"gray6"                : color.RGB(15,15,15,true),
	"grey6"                : color.RGB(15,15,15,true),
	"gray7"                : color.RGB(18,18,18,true),
	"grey7"                : color.RGB(18,18,18,true),
	"gray8"                : color.RGB(20,20,20,true),
	"grey8"                : color.RGB(20,20,20,true),
	"gray9"                : color.RGB(23,23,23,true),
	"grey9"                : color.RGB(23,23,23,true),
	"gray10"               : color.RGB(26,26,26,true),
	"grey10"               : color.RGB(26,26,26,true),
	"gray11"               : color.RGB(28,28,28,true),
	"grey11"               : color.RGB(28,28,28,true),
	"gray12"               : color.RGB(31,31,31,true),
	"grey12"               : color.RGB(31,31,31,true),
	"gray13"               : color.RGB(33,33,33,true),
	"grey13"               : color.RGB(33,33,33,true),
	"gray14"               : color.RGB(36,36,36,true),
	"grey14"               : color.RGB(36,36,36,true),
	"gray15"               : color.RGB(38,38,38,true),
	"grey15"               : color.RGB(38,38,38,true),
	"gray16"               : color.RGB(41,41,41,true),
	"grey16"               : color.RGB(41,41,41,true),
	"gray17"               : color.RGB(43,43,43,true),
	"grey17"               : color.RGB(43,43,43,true),
	"gray18"               : color.RGB(46,46,46,true),
	"grey18"               : color.RGB(46,46,46,true),
	"gray19"               : color.RGB(48,48,48,true),
	"grey19"               : color.RGB(48,48,48,true),
	"gray20"               : color.RGB(51,51,51,true),
	"grey20"               : color.RGB(51,51,51,true),
	"gray21"               : color.RGB(54,54,54,true),
	"grey21"               : color.RGB(54,54,54,true),
	"gray22"               : color.RGB(56,56,56,true),
	"grey22"               : color.RGB(56,56,56,true),
	"gray23"               : color.RGB(59,59,59,true),
	"grey23"               : color.RGB(59,59,59,true),
	"gray24"               : color.RGB(61,61,61,true),
	"grey24"               : color.RGB(61,61,61,true),
	"gray25"               : color.RGB(64,64,64,true),
	"grey25"               : color.RGB(64,64,64,true),
	"gray26"               : color.RGB(66,66,66,true),
	"grey26"               : color.RGB(66,66,66,true),
	"gray27"               : color.RGB(69,69,69,true),
	"grey27"               : color.RGB(69,69,69,true),
	"gray28"               : color.RGB(71,71,71,true),
	"grey28"               : color.RGB(71,71,71,true),
	"gray29"               : color.RGB(74,74,74,true),
	"grey29"               : color.RGB(74,74,74,true),
	"gray30"               : color.RGB(77,77,77,true),
	"grey30"               : color.RGB(77,77,77,true),
	"gray31"               : color.RGB(79,79,79,true),
	"grey31"               : color.RGB(79,79,79,true),
	"gray32"               : color.RGB(82,82,82,true),
	"grey32"               : color.RGB(82,82,82,true),
	"gray33"               : color.RGB(84,84,84,true),
	"grey33"               : color.RGB(84,84,84,true),
	"gray34"               : color.RGB(87,87,87,true),
	"grey34"               : color.RGB(87,87,87,true),
	"gray35"               : color.RGB(89,89,89,true),
	"grey35"               : color.RGB(89,89,89,true),
	"gray36"               : color.RGB(92,92,92,true),
	"grey36"               : color.RGB(92,92,92,true),
	"gray37"               : color.RGB(94,94,94,true),
	"grey37"               : color.RGB(94,94,94,true),
	"gray38"               : color.RGB(97,97,97,true),
	"grey38"               : color.RGB(97,97,97,true),
	"gray39"               : color.RGB(99,99,99,true),
	"grey39"               : color.RGB(99,99,99,true),
	"gray40"               : color.RGB(102,102,102,true),
	"grey40"               : color.RGB(102,102,102,true),
	"gray41"               : color.RGB(105,105,105,true),
	"grey41"               : color.RGB(105,105,105,true),
	"gray42"               : color.RGB(107,107,107,true),
	"grey42"               : color.RGB(107,107,107,true),
	"gray43"               : color.RGB(110,110,110,true),
	"grey43"               : color.RGB(110,110,110,true),
	"gray44"               : color.RGB(112,112,112,true),
	"grey44"               : color.RGB(112,112,112,true),
	"gray45"               : color.RGB(115,115,115,true),
	"grey45"               : color.RGB(115,115,115,true),
	"gray46"               : color.RGB(117,117,117,true),
	"grey46"               : color.RGB(117,117,117,true),
	"gray47"               : color.RGB(120,120,120,true),
	"grey47"               : color.RGB(120,120,120,true),
	"gray48"               : color.RGB(122,122,122,true),
	"grey48"               : color.RGB(122,122,122,true),
	"gray49"               : color.RGB(125,125,125,true),
	"grey49"               : color.RGB(125,125,125,true),
	"gray50"               : color.RGB(127,127,127,true),
	"grey50"               : color.RGB(127,127,127,true),
	"gray51"               : color.RGB(130,130,130,true),
	"grey51"               : color.RGB(130,130,130,true),
	"gray52"               : color.RGB(133,133,133,true),
	"grey52"               : color.RGB(133,133,133,true),
	"gray53"               : color.RGB(135,135,135,true),
	"grey53"               : color.RGB(135,135,135,true),
	"gray54"               : color.RGB(138,138,138,true),
	"grey54"               : color.RGB(138,138,138,true),
	"gray55"               : color.RGB(140,140,140,true),
	"grey55"               : color.RGB(140,140,140,true),
	"gray56"               : color.RGB(143,143,143,true),
	"grey56"               : color.RGB(143,143,143,true),
	"gray57"               : color.RGB(145,145,145,true),
	"grey57"               : color.RGB(145,145,145,true),
	"gray58"               : color.RGB(148,148,148,true),
	"grey58"               : color.RGB(148,148,148,true),
	"gray59"               : color.RGB(150,150,150,true),
	"grey59"               : color.RGB(150,150,150,true),
	"gray60"               : color.RGB(153,153,153,true),
	"grey60"               : color.RGB(153,153,153,true),
	"gray61"               : color.RGB(156,156,156,true),
	"grey61"               : color.RGB(156,156,156,true),
	"gray62"               : color.RGB(158,158,158,true),
	"grey62"               : color.RGB(158,158,158,true),
	"gray63"               : color.RGB(161,161,161,true),
	"grey63"               : color.RGB(161,161,161,true),
	"gray64"               : color.RGB(163,163,163,true),
	"grey64"               : color.RGB(163,163,163,true),
	"gray65"               : color.RGB(166,166,166,true),
	"grey65"               : color.RGB(166,166,166,true),
	"gray66"               : color.RGB(168,168,168,true),
	"grey66"               : color.RGB(168,168,168,true),
	"gray67"               : color.RGB(171,171,171,true),
	"grey67"               : color.RGB(171,171,171,true),
	"gray68"               : color.RGB(173,173,173,true),
	"grey68"               : color.RGB(173,173,173,true),
	"gray69"               : color.RGB(176,176,176,true),
	"grey69"               : color.RGB(176,176,176,true),
	"gray70"               : color.RGB(179,179,179,true),
	"grey70"               : color.RGB(179,179,179,true),
	"gray71"               : color.RGB(181,181,181,true),
	"grey71"               : color.RGB(181,181,181,true),
	"gray72"               : color.RGB(184,184,184,true),
	"grey72"               : color.RGB(184,184,184,true),
	"gray73"               : color.RGB(186,186,186,true),
	"grey73"               : color.RGB(186,186,186,true),
	"gray74"               : color.RGB(189,189,189,true),
	"grey74"               : color.RGB(189,189,189,true),
	"gray75"               : color.RGB(191,191,191,true),
	"grey75"               : color.RGB(191,191,191,true),
	"gray76"               : color.RGB(194,194,194,true),
	"grey76"               : color.RGB(194,194,194,true),
	"gray77"               : color.RGB(196,196,196,true),
	"grey77"               : color.RGB(196,196,196,true),
	"gray78"               : color.RGB(199,199,199,true),
	"grey78"               : color.RGB(199,199,199,true),
	"gray79"               : color.RGB(201,201,201,true),
	"grey79"               : color.RGB(201,201,201,true),
	"gray80"               : color.RGB(204,204,204,true),
	"grey80"               : color.RGB(204,204,204,true),
	"gray81"               : color.RGB(207,207,207,true),
	"grey81"               : color.RGB(207,207,207,true),
	"gray82"               : color.RGB(209,209,209,true),
	"grey82"               : color.RGB(209,209,209,true),
	"gray83"               : color.RGB(212,212,212,true),
	"grey83"               : color.RGB(212,212,212,true),
	"gray84"               : color.RGB(214,214,214,true),
	"grey84"               : color.RGB(214,214,214,true),
	"gray85"               : color.RGB(217,217,217,true),
	"grey85"               : color.RGB(217,217,217,true),
	"gray86"               : color.RGB(219,219,219,true),
	"grey86"               : color.RGB(219,219,219,true),
	"gray87"               : color.RGB(222,222,222,true),
	"grey87"               : color.RGB(222,222,222,true),
	"gray88"               : color.RGB(224,224,224,true),
	"grey88"               : color.RGB(224,224,224,true),
	"gray89"               : color.RGB(227,227,227,true),
	"grey89"               : color.RGB(227,227,227,true),
	"gray90"               : color.RGB(229,229,229,true),
	"grey90"               : color.RGB(229,229,229,true),
	"gray91"               : color.RGB(232,232,232,true),
	"grey91"               : color.RGB(232,232,232,true),
	"gray92"               : color.RGB(235,235,235,true),
	"grey92"               : color.RGB(235,235,235,true),
	"gray93"               : color.RGB(237,237,237,true),
	"grey93"               : color.RGB(237,237,237,true),
	"gray94"               : color.RGB(240,240,240,true),
	"grey94"               : color.RGB(240,240,240,true),
	"gray95"               : color.RGB(242,242,242,true),
	"grey95"               : color.RGB(242,242,242,true),
	"gray96"               : color.RGB(245,245,245,true),
	"grey96"               : color.RGB(245,245,245,true),
	"gray97"               : color.RGB(247,247,247,true),
	"grey97"               : color.RGB(247,247,247,true),
	"gray98"               : color.RGB(250,250,250,true),
	"grey98"               : color.RGB(250,250,250,true),
	"gray99"               : color.RGB(252,252,252,true),
	"grey99"               : color.RGB(252,252,252,true),
	"gray100"              : color.RGB(255,255,255,true),
	"grey100"              : color.RGB(255,255,255,true),
	"darkgrey"             : color.RGB(169,169,169,true),
	"darkgray"             : color.RGB(169,169,169,true),
	"darkblue"             : color.RGB(0,0,139,true),
	"darkcyan"             : color.RGB(0,139,139,true),
	"darkmagenta"          : color.RGB(139,0,139,true),
	"darkred"              : color.RGB(139,0,0,true),
	"lightgreen"           : color.RGB(144,238,144,true),
	"crimson"              : color.RGB(220,20,60,true),
	"indigo"               : color.RGB(75,0,130,true),
	"olive"                : color.RGB(128,128,0,true),
	"rebeccapurple"        : color.RGB(102,51,153,true),
	"silver"               : color.RGB(192,192,192,true),
	"teal"                 : color.RGB(0,128,128,true),
}


var whitespace = regexp.MustCompile(`\s+`)
func cleanString(s string) string {
	s = strings.ToLower(s)
	s_bytes := []byte(s)
	s_bytes = whitespace.ReplaceAll(s_bytes, []byte(""))
	return string(s_bytes)
}


func colorNameToRGB(colorName string) (rgbColor color.RGBColor, colorOutputType string, isValid bool) {
	// TODO: this is stupid, i think this is passed directly to .RGB where it overflows.. yikes
	rgbColor = color.RGBFromString(colorName, true)
	isValid = ! rgbColor.IsEmpty()  // TODO FIXME
	colorOutputType = "rgb"
	return
}


func colorNameToX11(colorName string) (rgbColor color.RGBColor, colorOutputType string, isValid bool) {
	rgbColor, isValid = x11Colors[colorName]
	colorOutputType = "rgb"
	return
}


func colorNameToHex(colorName string) (rgbColor color.RGBColor, colorOutputType string, isValid bool) {
	rgbColor = color.HEX(colorName, true)
	isValid = ! rgbColor.IsEmpty()
	colorOutputType = "256"
	return
}


func dieImmediate(status int, message... string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(status)
}


func printVersion() {
	fmt.Println(VERSION)
}


var STATUS_UNKNOWN_COLORTYPE = 1
var STATUS_INVALID_COLOR = 2
var STATUS_NOT_IMPLEMENTED = 99


func main() {
	/* Logic
	 *   If flag given, use that specific color set
	 *     --x11
	 *     --web
	 *     --hex
	 *     --rgb
	 *     --hsv
	 *     --hsl
	 *     --lab
	 *     etc...
	 *   Otherwise, search in this order:
	 *     1. hex
	 *     3. rgb
	 *     2. x11
	 *
	 * X11 and web colors can be written in any case and with any whitespace - they will be "cleaned" to lower-case and no whitespace
	 */

	var colorName, colorNameClean, colorType string

	var colorTypeFlag = flag.String("type", "", "Color type. Must be one of: 'x11', 'hex', 'rgb'.")
	var versionFlag = flag.Bool("version", false, "Print version and exit.")
	//flag_x11 = flag.Bool("x11", false, "Use X11 colors")
	//flag_web = flag.Bool("web", false, "Use web colors")
	//flag_hex = flag.Bool("hex", false, "Use hexadecimal colors")
	//flag_rgb = flag.Bool("rgb", false, "Use RGB colors")
	//flag_hsv = flag.Bool("hsv", false, "Use HSV colors")
	//flag_hsl = flag.Bool("hsl", false, "Use HSL colors")
	//flag_lab = flag.Bool("lab", false, "Use LAB colors")

	flag.Parse()

	colorName = flag.Arg(0)
	if len(colorName) == 0 {
		dieImmediate(STATUS_INVALID_COLOR, "Color name is required")
	}

	if *versionFlag {
		printVersion()
		os.Exit(0)
	}

	// TODO: --fg and --bg options

	// TODO: use iota instead of magic strings? see https://stackoverflow.com/q/14426366

	if len(*colorTypeFlag) != 0 {
		colorType = *colorTypeFlag
	}

	colorType = cleanString(colorType)
	colorNameClean = cleanString(colorName)

	//fmt.Println("colorType", colorType)
	//fmt.Println("colorName", colorName)
	//fmt.Println("colorName", colorNameClean)

	//fmt.Println("Bad rgb (overflow)", color.RGBFromString("300,300,300"))
	//fmt.Println("Bad rgb (invalid)", color.RGBFromString("oogey"))
	//fmt.Println("Bad rgb (both)", color.RGBFromString("300,300,asdfsaf"))
	//fmt.Println("Bad hex", color.HEX("oogabooga"))

	var rgbColor        color.RGBColor
	var colorOutputType string
	var isValid         bool

	if len(colorType) > 0 {
		switch {
		case colorType == "x11":
			rgbColor, colorOutputType, isValid = colorNameToX11(colorNameClean)
		case colorType == "web":
			dieImmediate(STATUS_NOT_IMPLEMENTED, "Web color not implemented")
		case colorType == "hex":
			rgbColor, colorOutputType, isValid = colorNameToHex(colorNameClean)
		case colorType == "rgb":
			rgbColor, colorOutputType, isValid = colorNameToRGB(colorNameClean)
		case colorType == "hsv":
			dieImmediate(STATUS_NOT_IMPLEMENTED, "HSV color not implemented")
		case colorType == "hsl":
			dieImmediate(STATUS_NOT_IMPLEMENTED, "HSL color not implemented")
		case colorType == "lab":
			dieImmediate(STATUS_NOT_IMPLEMENTED, "LAB color not implemented")
		default:
			dieImmediate(STATUS_UNKNOWN_COLORTYPE, "Unknown color type")
		}
	} else {
		colorTransformers := [](func(string) (color.RGBColor, string, bool)) {colorNameToHex, colorNameToRGB, colorNameToX11}
		for _, transformer := range colorTransformers {
			rgbColor, colorOutputType, isValid = transformer(colorNameClean)
			//fmt.Println("rgbColor", rgbColor)
			//fmt.Println("colorOutputType", colorOutputType)
			//fmt.Println("isValid", isValid)
			if isValid {
				break
			}
		}
		if ! isValid {
			dieImmediate(STATUS_UNKNOWN_COLORTYPE, "Could not detect colortype")
		}
	}

	if ! isValid {
		dieImmediate(STATUS_INVALID_COLOR, "Invalid color:", colorName)
	}

	if colorOutputType == "256" {
		// rgbToC256(rgbColor).Print(colorNameClean)
		rgbColor.Println(colorNameClean)
	} else {
		rgbColor.Println(colorNameClean)
	}
}

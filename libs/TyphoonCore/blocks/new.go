package blocks

type NewMapping map[uint16]int

type NewBlock struct {
	Name     string
	Fallback *string
	Ids      NewMapping
}

var newBlocks = []NewBlock{
	{
		"minecraft:blackstone_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16153,
		},
	},
	{
		"minecraft:crimson_sign[rotation=5,waterlogged=true]",
		nil,
		NewMapping{
			718: 15665,
		},
	},
	{
		"minecraft:potted_fern",
		nil,
		NewMapping{
			404: 5273,
			477: 5776,
			573: 5776,
			718: 6312,
			393: 5272,
		},
	},
	{
		"minecraft:spruce_fence[east=true,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7521,
			404: 7522,
			477: 8046,
			573: 8046,
			718: 8582,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 6093,
			477: 6599,
			573: 6599,
			718: 7135,
			393: 6092,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 6346,
			718: 6882,
			393: 5839,
			404: 5840,
			477: 6346,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12392,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=none,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12920,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12989,
		},
	},
	{
		"minecraft:green_banner[rotation=10]",
		nil,
		NewMapping{
			393: 7072,
			404: 7073,
			477: 7579,
			573: 7579,
			718: 8115,
		},
	},
	{
		"minecraft:command_block[conditional=true,facing=east]",
		nil,
		NewMapping{
			573: 5629,
			718: 5645,
			4:   2205,
			393: 5125,
			404: 5126,
			477: 5629,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10949,
			573: 10949,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10360,
			573: 10360,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=none,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14215,
		},
	},
	{
		"minecraft:fire[age=5,east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1319,
			404: 1320,
			477: 1623,
			573: 1623,
			718: 1624,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 3488,
			404: 3489,
			477: 3992,
			573: 3992,
			718: 3994,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9179,
			573: 9179,
			718: 9715,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13645,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10663,
			573: 10663,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=none,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13589,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=north,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			404: 3673,
			477: 4176,
			573: 4176,
			718: 4190,
			393: 3672,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16959,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16617,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=low,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11579,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=low,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16927,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16103,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13317,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 7187,
			718: 7723,
			393: 6680,
			404: 6681,
			477: 7187,
		},
	},
	{
		"minecraft:light_blue_wall_banner[facing=east]",
		nil,
		NewMapping{
			393: 7125,
			404: 7126,
			477: 7632,
			573: 7632,
			718: 8168,
		},
	},
	{
		"minecraft:jungle_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 5584,
			573: 5584,
			718: 5600,
			393: 5080,
			404: 5081,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10403,
			573: 10403,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=4,south=none,west=none]",
		nil,
		NewMapping{
			393: 2660,
			404: 2661,
			477: 2964,
			573: 2964,
			718: 2966,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11025,
			573: 11025,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=tall,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12158,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=tall,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13866,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=tall,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13128,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=low,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16191,
		},
	},
	{
		"minecraft:player_head[rotation=13]",
		nil,
		NewMapping{
			404: 5525,
			477: 6027,
			573: 6027,
			718: 6563,
			393: 5524,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			477: 4617,
			573: 4617,
			718: 4631,
			393: 4113,
			404: 4114,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9300,
			573: 9300,
			718: 9836,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9559,
			573: 9559,
			718: 10095,
		},
	},
	{
		"minecraft:granite_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9906,
			573: 9906,
			718: 10442,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=low,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13848,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=none,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 15925,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=none,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12813,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=none,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11001,
		},
	},
	{
		"minecraft:light_blue_bed[facing=west,occupied=true,part=head]",
		nil,
		NewMapping{
			718: 1105,
			393: 804,
			404: 804,
			477: 1104,
			573: 1104,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6160,
			404: 6161,
			477: 6667,
			573: 6667,
			718: 7203,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=14,powered=true]",
		nil,
		NewMapping{
			477: 876,
			573: 876,
			718: 877,
		},
	},
	{
		"minecraft:crimson_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15366,
		},
	},
	{
		"minecraft:purple_glazed_terracotta[facing=east]",
		nil,
		NewMapping{
			718: 9417,
			4:   3923,
			393: 8356,
			404: 8357,
			477: 8881,
			573: 8881,
		},
	},
	{
		"minecraft:spruce_button[face=floor,facing=east,powered=false]",
		nil,
		NewMapping{
			477: 5841,
			573: 5841,
			718: 6377,
			393: 5334,
			404: 5335,
		},
	},
	{
		"minecraft:iron_door[facing=west,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			573: 3852,
			718: 3854,
			4:   1142,
			393: 3348,
			404: 3349,
			477: 3852,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10864,
			573: 10864,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9510,
			573: 9510,
			718: 10046,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=none,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5659,
		},
	},
	{
		"minecraft:fire[age=10,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1470,
			404: 1471,
			477: 1774,
			573: 1774,
			718: 1775,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=west,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3701,
			404: 3702,
			477: 4205,
			573: 4205,
			718: 4219,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=none,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16145,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=1,powered=true]",
		nil,
		NewMapping{
			477: 750,
			573: 750,
			718: 751,
		},
	},
	{
		"minecraft:acacia_sign[rotation=11,waterlogged=false]",
		nil,
		NewMapping{
			477: 3498,
			573: 3498,
			718: 3500,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15853,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=none,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14357,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12956,
		},
	},
	{
		"minecraft:acacia_door[facing=south,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			573: 8412,
			718: 8948,
			393: 7887,
			404: 7888,
			477: 8412,
		},
	},
	{
		"minecraft:piston_head[facing=south,short=true,type=normal]",
		nil,
		NewMapping{
			393: 1067,
			404: 1067,
			477: 1367,
			573: 1367,
			718: 1368,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=low,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12745,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=none,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12830,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6173,
			404: 6174,
			477: 6680,
			573: 6680,
			718: 7216,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=5,south=side,west=side]",
		nil,
		NewMapping{
			718: 2539,
			393: 2233,
			404: 2234,
			477: 2537,
			573: 2537,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=low,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11992,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10463,
			573: 10463,
		},
	},
	{
		"minecraft:birch_sign[rotation=6,waterlogged=true]",
		nil,
		NewMapping{
			477: 3455,
			573: 3455,
			718: 3457,
		},
	},
	{
		"minecraft:fire[age=2,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1206,
			404: 1207,
			477: 1510,
			573: 1510,
			718: 1511,
		},
	},
	{
		"minecraft:repeater[delay=1,facing=north,locked=true,powered=true]",
		nil,
		NewMapping{
			718: 4031,
			393: 3513,
			404: 3514,
			477: 4017,
			573: 4017,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 3234,
			477: 3697,
			573: 3697,
			718: 3699,
			393: 3233,
		},
	},
	{
		"minecraft:magenta_bed[facing=east,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 794,
			404: 794,
			477: 1094,
			573: 1094,
			718: 1095,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16824,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=south,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7434,
			404: 7435,
			477: 7959,
			573: 7959,
			718: 8495,
			4:   2964,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12135,
		},
	},
	{
		"minecraft:crimson_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			718: 15047,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5889,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6110,
			404: 6111,
			477: 6617,
			573: 6617,
			718: 7153,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 10243,
			573: 10243,
			718: 10779,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=north,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 15127,
		},
	},
	{
		"minecraft:jungle_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 5098,
			404: 5099,
			477: 5602,
			573: 5602,
			718: 5618,
		},
	},
	{
		"minecraft:jungle_door[facing=west,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			477: 8376,
			573: 8376,
			718: 8912,
			393: 7851,
			404: 7852,
		},
	},
	{
		"minecraft:sand",
		nil,
		NewMapping{
			393: 66,
			404: 66,
			477: 66,
			573: 66,
			718: 66,
			4:   192,
		},
	},
	{
		"minecraft:quartz_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 5760,
			477: 6266,
			573: 6266,
			718: 6802,
			393: 5759,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 6447,
			718: 6983,
			393: 5940,
			404: 5941,
			477: 6447,
		},
	},
	{
		"minecraft:fire[age=12,east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			573: 1842,
			718: 1843,
			393: 1538,
			404: 1539,
			477: 1842,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=none,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13056,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=none,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 6105,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=5,south=side,west=up]",
		nil,
		NewMapping{
			393: 1800,
			404: 1801,
			477: 2104,
			573: 2104,
			718: 2106,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12195,
		},
	},
	{
		"minecraft:cyan_banner[rotation=6]",
		nil,
		NewMapping{
			718: 8047,
			393: 7004,
			404: 7005,
			477: 7511,
			573: 7511,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 8291,
			4:   2880,
			393: 7248,
			404: 7249,
			477: 7755,
			573: 7755,
		},
	},
	{
		"minecraft:weeping_vines[age=12]",
		nil,
		NewMapping{
			718: 15002,
		},
	},
	{
		"minecraft:twisting_vines[age=9]",
		nil,
		NewMapping{
			718: 15026,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11592,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=low,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13090,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=2,south=none,west=none]",
		nil,
		NewMapping{
			393: 2066,
			404: 2067,
			477: 2370,
			573: 2370,
			718: 2372,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=low,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11136,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=none,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12167,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=low,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5715,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=11,south=up,west=side]",
		nil,
		NewMapping{
			477: 3020,
			573: 3020,
			718: 3022,
			393: 2716,
			404: 2717,
		},
	},
	{
		"minecraft:cyan_bed[facing=south,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 896,
			404: 896,
			477: 1196,
			573: 1196,
			718: 1197,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=low,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11123,
		},
	},
	{
		"minecraft:fire[age=6,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			573: 1633,
			718: 1634,
			393: 1329,
			404: 1330,
			477: 1633,
		},
	},
	{
		"minecraft:spruce_sign[rotation=14,waterlogged=false]",
		nil,
		NewMapping{
			477: 3440,
			573: 3440,
			718: 3442,
		},
	},
	{
		"minecraft:lectern[facing=west,has_book=true,powered=true]",
		nil,
		NewMapping{
			718: 14841,
			477: 11185,
			573: 11185,
		},
	},
	{
		"minecraft:brick_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4404,
			404: 4405,
			477: 4908,
			573: 4908,
			718: 4924,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 5084,
			573: 5084,
			718: 5100,
			393: 4580,
			404: 4581,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9373,
			573: 9373,
			718: 9909,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=tall,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12687,
		},
	},
	{
		"minecraft:kelp_plant",
		nil,
		NewMapping{
			393: 8435,
			404: 8436,
			477: 8960,
			573: 8960,
			718: 9496,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5149,
			404: 5150,
			477: 5653,
			573: 5653,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 1953,
			718: 1955,
			4:   855,
			393: 1649,
			404: 1650,
			477: 1953,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 10115,
			718: 10651,
			477: 10115,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=north,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			718: 7548,
			4:   2676,
			393: 6505,
			404: 6506,
			477: 7012,
			573: 7012,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=6,south=side,west=none]",
		nil,
		NewMapping{
			393: 2531,
			404: 2532,
			477: 2835,
			573: 2835,
			718: 2837,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10839,
			573: 10839,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14403,
		},
	},
	{
		"minecraft:warped_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15452,
		},
	},
	{
		"minecraft:green_concrete_powder",
		nil,
		NewMapping{
			477: 8931,
			573: 8931,
			718: 9467,
			4:   4045,
			393: 8406,
			404: 8407,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=17,powered=false]",
		nil,
		NewMapping{
			573: 933,
			718: 934,
			477: 933,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9843,
			573: 9843,
			718: 10379,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=low,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13853,
		},
	},
	{
		"minecraft:orange_wall_banner[facing=west]",
		nil,
		NewMapping{
			477: 7623,
			573: 7623,
			718: 8159,
			393: 7116,
			404: 7117,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11080,
			573: 11080,
		},
	},
	{
		"minecraft:melon_stem[age=6]",
		nil,
		NewMapping{
			718: 4786,
			4:   1686,
			393: 4266,
			404: 4267,
			477: 4770,
			573: 4770,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=low,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16613,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16992,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=true,north=false,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			477: 5269,
			573: 5269,
			718: 5285,
			393: 4765,
			404: 4766,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=10,south=up,west=up]",
		nil,
		NewMapping{
			393: 2850,
			404: 2851,
			477: 3154,
			573: 3154,
			718: 3156,
		},
	},
	{
		"minecraft:campfire[facing=north,lit=true,signal_fire=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 11217,
			573: 11233,
			718: 14891,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=none,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11546,
		},
	},
	{
		"minecraft:acacia_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6406,
			404: 6407,
			477: 6913,
			573: 6913,
			718: 7449,
		},
	},
	{
		"minecraft:dark_oak_button[face=ceiling,facing=west,powered=true]",
		nil,
		NewMapping{
			404: 5444,
			477: 5950,
			573: 5950,
			718: 6486,
			393: 5443,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13087,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 16743,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6081,
			404: 6082,
			477: 6588,
			573: 6588,
			718: 7124,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 6479,
			477: 6985,
			573: 6985,
			718: 7521,
			393: 6478,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=low,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13932,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=east,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			718: 15181,
		},
	},
	{
		"minecraft:spruce_sign[rotation=12,waterlogged=false]",
		nil,
		NewMapping{
			477: 3436,
			573: 3436,
			718: 3438,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=low,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16921,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=10,south=side,west=none]",
		nil,
		NewMapping{
			393: 2711,
			404: 2712,
			477: 3015,
			573: 3015,
			718: 3017,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9139,
			573: 9139,
			718: 9675,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10099,
			573: 10099,
			718: 10635,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9402,
			573: 9402,
			718: 9938,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13640,
		},
	},
	{
		"minecraft:soul_campfire[facing=south,lit=true,signal_fire=true,waterlogged=false]",
		nil,
		NewMapping{
			718: 14931,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=west,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3822,
			404: 3823,
			477: 4326,
			573: 4326,
			718: 4340,
		},
	},
	{
		"minecraft:acacia_door[facing=east,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7917,
			404: 7918,
			477: 8442,
			573: 8442,
			718: 8978,
		},
	},
	{
		"minecraft:jungle_wall_sign[facing=west,waterlogged=true]",
		nil,
		NewMapping{
			477: 3769,
			573: 3769,
			718: 3771,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=none,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11736,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=low,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16931,
		},
	},
	{
		"minecraft:acacia_leaves[distance=7,persistent=true]",
		nil,
		NewMapping{
			573: 212,
			718: 213,
			393: 212,
			404: 212,
			477: 212,
		},
	},
	{
		"minecraft:spruce_leaves[distance=3,persistent=true]",
		nil,
		NewMapping{
			393: 162,
			404: 162,
			477: 162,
			573: 162,
			718: 163,
		},
	},
	{
		"minecraft:light_weighted_pressure_plate[power=2]",
		nil,
		NewMapping{
			477: 6112,
			573: 6112,
			718: 6648,
			4:   2354,
			393: 5605,
			404: 5606,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 5750,
		},
	},
	{
		"minecraft:soul_campfire[facing=south,lit=true,signal_fire=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 14930,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=tall,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13006,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=low,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11342,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 7519,
			393: 6476,
			404: 6477,
			477: 6983,
			573: 6983,
		},
	},
	{
		"minecraft:glass_pane[east=false,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 4743,
			573: 4743,
			718: 4759,
			393: 4239,
			404: 4240,
		},
	},
	{
		"minecraft:zombie_wall_head[facing=east]",
		nil,
		NewMapping{
			573: 6013,
			718: 6549,
			393: 5490,
			404: 5491,
			477: 6013,
		},
	},
	{
		"minecraft:smooth_stone_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			573: 7809,
			718: 8345,
			404: 7297,
			477: 7809,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=tall,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16130,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=15,south=up,west=side]",
		nil,
		NewMapping{
			404: 2609,
			477: 2912,
			573: 2912,
			718: 2914,
			393: 2608,
		},
	},
	{
		"minecraft:jungle_log[axis=x]",
		nil,
		NewMapping{
			573: 81,
			718: 82,
			4:   279,
			393: 81,
			404: 81,
			477: 81,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4684,
			404: 4685,
			477: 5188,
			573: 5188,
			718: 5204,
		},
	},
	{
		"minecraft:warped_sign[rotation=5,waterlogged=false]",
		nil,
		NewMapping{
			718: 15698,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 7754,
			573: 7754,
			718: 8290,
			393: 7247,
			404: 7248,
		},
	},
	{
		"minecraft:stripped_oak_log[axis=y]",
		nil,
		NewMapping{
			477: 106,
			573: 106,
			718: 107,
			393: 106,
			404: 106,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=21,powered=true]",
		nil,
		NewMapping{
			477: 840,
			573: 840,
			718: 841,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16809,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=tall,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13561,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6295,
			404: 6296,
			477: 6802,
			573: 6802,
			718: 7338,
		},
	},
	{
		"minecraft:scaffolding[bottom=false,distance=7,waterlogged=true]",
		nil,
		NewMapping{
			477: 11129,
			573: 11129,
			718: 14785,
		},
	},
	{
		"minecraft:polished_andesite_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			477: 10321,
			573: 10321,
			718: 10857,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=none,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12811,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=low,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6263,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=none,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13165,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12563,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13501,
		},
	},
	{
		"minecraft:spruce_door[facing=east,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7740,
			404: 7741,
			477: 8265,
			573: 8265,
			718: 8801,
			4:   3088,
		},
	},
	{
		"minecraft:fire[age=1,east=true,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			404: 1181,
			477: 1484,
			573: 1484,
			718: 1485,
			393: 1180,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16098,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11346,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=1,south=none,west=up]",
		nil,
		NewMapping{
			393: 2199,
			404: 2200,
			477: 2503,
			573: 2503,
			718: 2505,
		},
	},
	{
		"minecraft:repeater[delay=1,facing=south,locked=true,powered=true]",
		nil,
		NewMapping{
			393: 3517,
			404: 3518,
			477: 4021,
			573: 4021,
			718: 4035,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=east,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			404: 7485,
			477: 8009,
			573: 8009,
			718: 8545,
			4:   2995,
			393: 7484,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=none,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 6108,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=15,powered=true]",
		nil,
		NewMapping{
			477: 1028,
			573: 1028,
			718: 1029,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=none,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5682,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6794,
			404: 6795,
			477: 7301,
			573: 7301,
			718: 7837,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=11,powered=true]",
		nil,
		NewMapping{
			477: 620,
			573: 620,
			718: 621,
			393: 620,
			404: 620,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6752,
			573: 6752,
			718: 7288,
			393: 6245,
			404: 6246,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10808,
			477: 10808,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10423,
			573: 10423,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=none,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13377,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=none,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11104,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=tall,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12894,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=8,south=none,west=up]",
		nil,
		NewMapping{
			573: 3142,
			718: 3144,
			393: 2838,
			404: 2839,
			477: 3142,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=9,south=none,west=none]",
		nil,
		NewMapping{
			718: 2867,
			393: 2561,
			404: 2562,
			477: 2865,
			573: 2865,
		},
	},
	{
		"minecraft:oak_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 2032,
			393: 1726,
			404: 1727,
			477: 2030,
			573: 2030,
		},
	},
	{
		"minecraft:cyan_banner[rotation=7]",
		nil,
		NewMapping{
			393: 7005,
			404: 7006,
			477: 7512,
			573: 7512,
			718: 8048,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=east,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			718: 4164,
			393: 3646,
			404: 3647,
			477: 4150,
			573: 4150,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=9,waterlogged=true]",
		nil,
		NewMapping{
			477: 3557,
			573: 3557,
			718: 3559,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=none,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6227,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 16686,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=20,powered=false]",
		nil,
		NewMapping{
			718: 440,
			393: 439,
			404: 439,
			477: 439,
			573: 439,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14634,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=low,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5826,
		},
	},
	{
		"minecraft:dirt",
		nil,
		NewMapping{
			4:   48,
			393: 10,
			404: 10,
			477: 10,
			573: 10,
			718: 10,
		},
	},
	{
		"minecraft:iron_door[facing=west,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 3344,
			404: 3345,
			477: 3848,
			573: 3848,
			718: 3850,
		},
	},
	{
		"minecraft:purple_banner[rotation=8]",
		nil,
		NewMapping{
			393: 7022,
			404: 7023,
			477: 7529,
			573: 7529,
			718: 8065,
		},
	},
	{
		"minecraft:fire[age=3,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1242,
			404: 1243,
			477: 1546,
			573: 1546,
			718: 1547,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=10,powered=false]",
		nil,
		NewMapping{
			718: 470,
			393: 469,
			404: 469,
			477: 469,
			573: 469,
		},
	},
	{
		"minecraft:gray_bed[facing=west,occupied=false,part=foot]",
		nil,
		NewMapping{
			404: 871,
			477: 1171,
			573: 1171,
			718: 1172,
			393: 871,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6597,
			404: 6598,
			477: 7104,
			573: 7104,
			718: 7640,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10105,
			573: 10105,
			718: 10641,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=19,powered=true]",
		nil,
		NewMapping{
			393: 536,
			404: 536,
			477: 536,
			573: 536,
			718: 537,
		},
	},
	{
		"minecraft:oak_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			718: 8303,
			4:   2016,
			393: 7260,
			404: 7261,
			477: 7767,
			573: 7767,
		},
	},
	{
		"minecraft:dark_oak_leaves[distance=3,persistent=true]",
		nil,
		NewMapping{
			477: 218,
			573: 218,
			718: 219,
			393: 218,
			404: 218,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 4034,
			404: 4035,
			477: 4538,
			573: 4538,
			718: 4552,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=low,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11448,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13799,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10070,
			573: 10070,
			718: 10606,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=none,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14323,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=none,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13492,
		},
	},
	{
		"minecraft:red_shulker_box[facing=north]",
		nil,
		NewMapping{
			393: 8301,
			404: 8302,
			477: 8826,
			573: 8826,
			718: 9362,
			4:   3730,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=south,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			477: 4443,
			573: 4443,
			718: 4457,
			393: 3939,
			404: 3940,
		},
	},
	{
		"minecraft:yellow_bed[facing=east,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 826,
			404: 826,
			477: 1126,
			573: 1126,
			718: 1127,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10919,
			573: 10919,
		},
	},
	{
		"minecraft:fire[age=7,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1389,
			404: 1390,
			477: 1693,
			573: 1693,
			718: 1694,
		},
	},
	{
		"minecraft:trapped_chest[facing=west,type=left,waterlogged=true]",
		nil,
		NewMapping{
			393: 5593,
			404: 5594,
			477: 6100,
			573: 6100,
			718: 6636,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=east,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			573: 4160,
			718: 4174,
			4:   1539,
			393: 3656,
			404: 3657,
			477: 4160,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13582,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=none,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16450,
		},
	},
	{
		"minecraft:zombie_wall_head[facing=north]",
		nil,
		NewMapping{
			393: 5487,
			404: 5488,
			477: 6010,
			573: 6010,
			718: 6546,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=14,south=none,west=up]",
		nil,
		NewMapping{
			477: 2476,
			573: 2476,
			718: 2478,
			393: 2172,
			404: 2173,
		},
	},
	{
		"minecraft:oak_door[facing=west,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 3148,
			404: 3149,
			477: 3612,
			573: 3612,
			718: 3614,
		},
	},
	{
		"minecraft:infested_cobblestone",
		nil,
		NewMapping{
			393: 3978,
			404: 3979,
			477: 4486,
			573: 4486,
			718: 4500,
			4:   1553,
		},
	},
	{
		"minecraft:jungle_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 5574,
			393: 5054,
			404: 5055,
			477: 5558,
			573: 5558,
		},
	},
	{
		"minecraft:gray_bed[facing=east,occupied=true,part=head]",
		nil,
		NewMapping{
			718: 1173,
			393: 872,
			404: 872,
			477: 1172,
			573: 1172,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=none,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12816,
		},
	},
	{
		"minecraft:acacia_sign[rotation=3,waterlogged=false]",
		nil,
		NewMapping{
			573: 3482,
			718: 3484,
			477: 3482,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5741,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12451,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			573: 8571,
			718: 9107,
			393: 8046,
			404: 8047,
			477: 8571,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=west,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			404: 7443,
			477: 7967,
			573: 7967,
			718: 8503,
			4:   2965,
			393: 7442,
		},
	},
	{
		"minecraft:acacia_fence[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			718: 8687,
			393: 7626,
			404: 7627,
			477: 8151,
			573: 8151,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10992,
			573: 10992,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=north,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3595,
			404: 3596,
			477: 4099,
			573: 4099,
			718: 4113,
		},
	},
	{
		"minecraft:polished_blackstone_button[face=ceiling,facing=north,powered=false]",
		nil,
		NewMapping{
			718: 16770,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			718: 9075,
			393: 8014,
			404: 8015,
			477: 8539,
			573: 8539,
		},
	},
	{
		"minecraft:birch_door[facing=north,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			477: 8271,
			573: 8271,
			718: 8807,
			393: 7746,
			404: 7747,
		},
	},
	{
		"minecraft:birch_sign[rotation=15,waterlogged=true]",
		nil,
		NewMapping{
			477: 3473,
			573: 3473,
			718: 3475,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=tall,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12479,
		},
	},
	{
		"minecraft:lime_shulker_box[facing=north]",
		nil,
		NewMapping{
			404: 8248,
			477: 8772,
			573: 8772,
			718: 9308,
			4:   3586,
			393: 8247,
		},
	},
	{
		"minecraft:fire[age=2,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1230,
			404: 1231,
			477: 1534,
			573: 1534,
			718: 1535,
			4:   818,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10742,
			573: 10742,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=tall,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12682,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=tall,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11610,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14151,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 6945,
			573: 6945,
			718: 7481,
			393: 6438,
			404: 6439,
		},
	},
	{
		"minecraft:spruce_sign[rotation=5,waterlogged=true]",
		nil,
		NewMapping{
			477: 3421,
			573: 3421,
			718: 3423,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10120,
			573: 10120,
			718: 10656,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=low,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 5914,
		},
	},
	{
		"minecraft:lever[face=ceiling,facing=south,powered=false]",
		nil,
		NewMapping{
			404: 3297,
			477: 3800,
			573: 3800,
			718: 3802,
			393: 3296,
		},
	},
	{
		"minecraft:dead_horn_coral_fan[waterlogged=false]",
		nil,
		NewMapping{
			393: 8553,
			404: 8569,
			477: 9013,
			573: 9013,
			718: 9549,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12364,
		},
	},
	{
		"minecraft:dark_oak_door[facing=west,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7971,
			404: 7972,
			477: 8496,
			573: 8496,
			718: 9032,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9314,
			573: 9314,
			718: 9850,
		},
	},
	{
		"minecraft:water[level=1]",
		nil,
		NewMapping{
			718: 35,
			4:   129,
			393: 35,
			404: 35,
			477: 35,
			573: 35,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5841,
		},
	},
	{
		"minecraft:warped_trapdoor[facing=north,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 15199,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=none,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12929,
		},
	},
	{
		"minecraft:crimson_sign[rotation=15,waterlogged=true]",
		nil,
		NewMapping{
			718: 15685,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 1954,
			573: 1954,
			718: 1956,
			393: 1650,
			404: 1651,
		},
	},
	{
		"minecraft:nether_brick_fence[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			718: 5030,
			393: 4510,
			404: 4511,
			477: 5014,
			573: 5014,
		},
	},
	{
		"minecraft:fire[age=4,east=true,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			573: 1572,
			718: 1573,
			393: 1268,
			404: 1269,
			477: 1572,
		},
	},
	{
		"minecraft:spruce_sign[rotation=12,waterlogged=true]",
		nil,
		NewMapping{
			477: 3435,
			573: 3435,
			718: 3437,
		},
	},
	{
		"minecraft:spruce_fence[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7523,
			404: 7524,
			477: 8048,
			573: 8048,
			718: 8584,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=21,powered=true]",
		nil,
		NewMapping{
			477: 890,
			573: 890,
			718: 891,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=low,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12986,
		},
	},
	{
		"minecraft:gray_wall_banner[facing=south]",
		nil,
		NewMapping{
			477: 7646,
			573: 7646,
			718: 8182,
			393: 7139,
			404: 7140,
		},
	},
	{
		"minecraft:fire[age=10,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1462,
			404: 1463,
			477: 1766,
			573: 1766,
			718: 1767,
		},
	},
	{
		"minecraft:jungle_door[facing=west,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			718: 8905,
			393: 7844,
			404: 7845,
			477: 8369,
			573: 8369,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=14,powered=true]",
		nil,
		NewMapping{
			477: 526,
			573: 526,
			718: 527,
			393: 526,
			404: 526,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9357,
			573: 9357,
			718: 9893,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=none,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11519,
		},
	},
	{
		"minecraft:warped_fence_gate[facing=east,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			718: 15315,
		},
	},
	{
		"minecraft:spruce_wood[axis=z]",
		nil,
		NewMapping{
			393: 113,
			404: 113,
			477: 113,
			573: 113,
			718: 114,
		},
	},
	{
		"minecraft:creeper_head[rotation=12]",
		nil,
		NewMapping{
			718: 6582,
			393: 5543,
			404: 5544,
			477: 6046,
			573: 6046,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 4003,
			404: 4004,
			477: 4507,
			573: 4507,
			718: 4521,
		},
	},
	{
		"minecraft:acacia_sign[rotation=1,waterlogged=false]",
		nil,
		NewMapping{
			477: 3478,
			573: 3478,
			718: 3480,
		},
	},
	{
		"minecraft:warped_fence_gate[facing=west,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			718: 15304,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16460,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=4,south=up,west=side]",
		nil,
		NewMapping{
			404: 2078,
			477: 2381,
			573: 2381,
			718: 2383,
			393: 2077,
		},
	},
	{
		"minecraft:dark_oak_fence[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7669,
			404: 7670,
			477: 8194,
			573: 8194,
			718: 8730,
		},
	},
	{
		"minecraft:jungle_door[facing=west,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7848,
			404: 7849,
			477: 8373,
			573: 8373,
			718: 8909,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 15846,
		},
	},
	{
		"minecraft:twisting_vines[age=23]",
		nil,
		NewMapping{
			718: 15040,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 4139,
			404: 4140,
			477: 4643,
			573: 4643,
			718: 4657,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14254,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=none,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12617,
		},
	},
	{
		"minecraft:weeping_vines[age=0]",
		nil,
		NewMapping{
			718: 14990,
		},
	},
	{
		"minecraft:quartz_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 6270,
			573: 6270,
			718: 6806,
			393: 5763,
			404: 5764,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=none,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14112,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=none,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12619,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13747,
		},
	},
	{
		"minecraft:structure_void",
		nil,
		NewMapping{
			4:   3472,
			393: 8198,
			404: 8199,
			477: 8723,
			573: 8723,
			718: 9259,
		},
	},
	{
		"minecraft:dark_oak_fence[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7655,
			404: 7656,
			477: 8180,
			573: 8180,
			718: 8716,
		},
	},
	{
		"minecraft:spruce_door[facing=north,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7686,
			404: 7687,
			477: 8211,
			573: 8211,
			718: 8747,
		},
	},
	{
		"minecraft:birch_door[facing=west,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			404: 7785,
			477: 8309,
			573: 8309,
			718: 8845,
			393: 7784,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10779,
			573: 10779,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10345,
			573: 10345,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=none,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11630,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=tall,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11170,
		},
	},
	{
		"minecraft:wall_torch[facing=east]",
		nil,
		NewMapping{
			4:   801,
			393: 1134,
			404: 1135,
			477: 1438,
			573: 1438,
			718: 1439,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=south,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 4117,
			573: 4117,
			718: 4131,
			393: 3613,
			404: 3614,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10145,
			573: 10145,
			718: 10681,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9425,
			573: 9425,
			718: 9961,
		},
	},
	{
		"minecraft:andesite_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9998,
			573: 9998,
			718: 10534,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11127,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=9,south=side,west=side]",
		nil,
		NewMapping{
			477: 2861,
			573: 2861,
			718: 2863,
			393: 2557,
			404: 2558,
		},
	},
	{
		"minecraft:jungle_door[facing=north,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			573: 8332,
			718: 8868,
			393: 7807,
			404: 7808,
			477: 8332,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=north,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			477: 4164,
			573: 4164,
			718: 4178,
			393: 3660,
			404: 3661,
		},
	},
	{
		"minecraft:farmland[moisture=4]",
		nil,
		NewMapping{
			404: 3064,
			477: 3367,
			573: 3367,
			718: 3369,
			4:   964,
			393: 3063,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12467,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=south,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			718: 8430,
			4:   2940,
			393: 7369,
			404: 7370,
			477: 7894,
			573: 7894,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=none,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16049,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 16285,
		},
	},
	{
		"minecraft:fire[age=9,east=false,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			718: 1748,
			393: 1443,
			404: 1444,
			477: 1747,
			573: 1747,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9707,
			573: 9707,
			718: 10243,
		},
	},
	{
		"minecraft:diorite_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 10736,
			477: 10200,
			573: 10200,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9204,
			573: 9204,
			718: 9740,
		},
	},
	{
		"minecraft:creeper_head[rotation=10]",
		nil,
		NewMapping{
			404: 5542,
			477: 6044,
			573: 6044,
			718: 6580,
			393: 5541,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=0,south=side,west=none]",
		nil,
		NewMapping{
			477: 2493,
			573: 2493,
			718: 2495,
			393: 2189,
			404: 2190,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=0,south=side,west=up]",
		nil,
		NewMapping{
			404: 2764,
			477: 3067,
			573: 3067,
			718: 3069,
			393: 2763,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=north,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			404: 6497,
			477: 7003,
			573: 7003,
			718: 7539,
			393: 6496,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=tall,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13772,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=low,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 6157,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=east,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7413,
			404: 7414,
			477: 7938,
			573: 7938,
			718: 8474,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=low,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14158,
		},
	},
	{
		"minecraft:quartz_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 5748,
			404: 5749,
			477: 6255,
			573: 6255,
			718: 6791,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=7,south=none,west=up]",
		nil,
		NewMapping{
			404: 2686,
			477: 2989,
			573: 2989,
			718: 2991,
			393: 2685,
		},
	},
	{
		"minecraft:warped_fence_gate[facing=south,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			718: 15299,
		},
	},
	{
		"minecraft:wall_sign[facing=south,waterlogged=false]",
		nil,
		NewMapping{
			393: 3272,
			4:   1091,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=tall,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13116,
		},
	},
	{
		"minecraft:yellow_banner[rotation=3]",
		nil,
		NewMapping{
			393: 6921,
			404: 6922,
			477: 7428,
			573: 7428,
			718: 7964,
		},
	},
	{
		"minecraft:orange_banner[rotation=0]",
		nil,
		NewMapping{
			573: 7377,
			718: 7913,
			393: 6870,
			404: 6871,
			477: 7377,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=tall,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16545,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 6601,
			404: 6602,
			477: 7108,
			573: 7108,
			718: 7644,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6027,
			404: 6028,
			477: 6534,
			573: 6534,
			718: 7070,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 9719,
			477: 9183,
			573: 9183,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=tall,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16983,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=low,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11995,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=6,south=up,west=side]",
		nil,
		NewMapping{
			393: 1807,
			404: 1808,
			477: 2111,
			573: 2111,
			718: 2113,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=6,south=none,west=up]",
		nil,
		NewMapping{
			718: 2838,
			393: 2532,
			404: 2533,
			477: 2836,
			573: 2836,
		},
	},
	{
		"minecraft:dragon_head[rotation=7]",
		nil,
		NewMapping{
			718: 6597,
			393: 5558,
			404: 5559,
			477: 6061,
			573: 6061,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=low,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13940,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11967,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11372,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6317,
			404: 6318,
			477: 6824,
			573: 6824,
			718: 7360,
		},
	},
	{
		"minecraft:birch_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 5035,
			404: 5036,
			477: 5539,
			573: 5539,
			718: 5555,
			4:   2160,
		},
	},
	{
		"minecraft:poppy",
		nil,
		NewMapping{
			573: 1412,
			718: 1413,
			4:   608,
			393: 1112,
			404: 1112,
			477: 1412,
		},
	},
	{
		"minecraft:lectern[facing=north,has_book=true,powered=false]",
		nil,
		NewMapping{
			477: 11178,
			573: 11178,
			718: 14834,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6128,
			404: 6129,
			477: 6635,
			573: 6635,
			718: 7171,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9546,
			573: 9546,
			718: 10082,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10385,
			573: 10385,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 6288,
		},
	},
	{
		"minecraft:acacia_button[face=wall,facing=north,powered=false]",
		nil,
		NewMapping{
			393: 5408,
			404: 5409,
			477: 5915,
			573: 5915,
			718: 6451,
		},
	},
	{
		"minecraft:wither_skeleton_skull[rotation=10]",
		nil,
		NewMapping{
			573: 5984,
			718: 6520,
			393: 5481,
			404: 5482,
			477: 5984,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=13,south=up,west=side]",
		nil,
		NewMapping{
			573: 2318,
			718: 2320,
			393: 2014,
			404: 2015,
			477: 2318,
		},
	},
	{
		"minecraft:water[level=15]",
		nil,
		NewMapping{
			573: 49,
			718: 49,
			4:   159,
			393: 49,
			404: 49,
			477: 49,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=low,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11239,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 5198,
			393: 4678,
			404: 4679,
			477: 5182,
			573: 5182,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=tall,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14723,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			404: 7884,
			477: 8408,
			573: 8408,
			718: 8944,
			393: 7883,
		},
	},
	{
		"minecraft:kelp[age=2]",
		nil,
		NewMapping{
			393: 8411,
			404: 8412,
			477: 8936,
			573: 8936,
			718: 9472,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=tall,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12462,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15857,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=low,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11560,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=none,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5796,
		},
	},
	{
		"minecraft:spruce_fence[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 7535,
			477: 8059,
			573: 8059,
			718: 8595,
			393: 7534,
		},
	},
	{
		"minecraft:pink_terracotta",
		nil,
		NewMapping{
			393: 5810,
			404: 5811,
			477: 6317,
			573: 6317,
			718: 6853,
			4:   2550,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 9921,
			477: 9385,
			573: 9385,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=tall,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14743,
		},
	},
	{
		"minecraft:diorite_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10209,
			573: 10209,
			718: 10745,
		},
	},
	{
		"minecraft:crimson_door[facing=east,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			718: 15576,
		},
	},
	{
		"minecraft:orange_bed[facing=east,occupied=false,part=head]",
		nil,
		NewMapping{
			573: 1078,
			718: 1079,
			393: 778,
			404: 778,
			477: 1078,
		},
	},
	{
		"minecraft:andesite_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9939,
			573: 9939,
			718: 10475,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=21,powered=true]",
		nil,
		NewMapping{
			477: 990,
			573: 990,
			718: 991,
		},
	},
	{
		"minecraft:jungle_sign[rotation=8,waterlogged=true]",
		nil,
		NewMapping{
			718: 3525,
			477: 3523,
			573: 3523,
		},
	},
	{
		"minecraft:cocoa[age=1,facing=north]",
		nil,
		NewMapping{
			393: 4642,
			404: 4643,
			477: 5146,
			573: 5146,
			718: 5162,
			4:   2038,
		},
	},
	{
		"minecraft:diorite_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 10217,
			718: 10753,
			477: 10217,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=none,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12596,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12683,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=6,south=none,west=none]",
		nil,
		NewMapping{
			573: 3270,
			718: 3272,
			4:   886,
			393: 2966,
			404: 2967,
			477: 3270,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=4,south=side,west=up]",
		nil,
		NewMapping{
			393: 1935,
			404: 1936,
			477: 2239,
			573: 2239,
			718: 2241,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 5249,
			393: 4729,
			404: 4730,
			477: 5233,
			573: 5233,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=15,south=side,west=none]",
		nil,
		NewMapping{
			718: 2774,
			393: 2468,
			404: 2469,
			477: 2772,
			573: 2772,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9702,
			573: 9702,
			718: 10238,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14691,
		},
	},
	{
		"minecraft:polished_blackstone_button[face=wall,facing=east,powered=true]",
		nil,
		NewMapping{
			718: 16767,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=tall,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13667,
		},
	},
	{
		"minecraft:lava[level=0]",
		nil,
		NewMapping{
			393: 50,
			404: 50,
			477: 50,
			573: 50,
			718: 50,
			4:   160,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9650,
			573: 9650,
			718: 10186,
		},
	},
	{
		"minecraft:brown_bed[facing=east,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 953,
			404: 953,
			477: 1253,
			573: 1253,
			718: 1254,
		},
	},
	{
		"minecraft:light_blue_banner[rotation=5]",
		nil,
		NewMapping{
			393: 6907,
			404: 6908,
			477: 7414,
			573: 7414,
			718: 7950,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 6666,
			718: 7202,
			393: 6159,
			404: 6160,
			477: 6666,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=0,powered=true]",
		nil,
		NewMapping{
			718: 699,
			393: 698,
			404: 698,
			477: 698,
			573: 698,
		},
	},
	{
		"minecraft:fire[age=13,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			477: 1867,
			573: 1867,
			718: 1868,
			393: 1563,
			404: 1564,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=0,powered=true]",
		nil,
		NewMapping{
			477: 998,
			573: 998,
			718: 999,
		},
	},
	{
		"minecraft:warped_door[facing=north,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			718: 15594,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=none,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16785,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=true,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			4:   1614,
			393: 4051,
			404: 4052,
			477: 4555,
			573: 4555,
			718: 4569,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6281,
			404: 6282,
			477: 6788,
			573: 6788,
			718: 7324,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=14,south=none,west=side]",
		nil,
		NewMapping{
			718: 2767,
			393: 2461,
			404: 2462,
			477: 2765,
			573: 2765,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 3207,
			404: 3208,
			477: 3671,
			573: 3671,
			718: 3673,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14360,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12634,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6732,
			573: 6732,
			718: 7268,
			393: 6225,
			404: 6226,
		},
	},
	{
		"minecraft:spruce_sign[rotation=2,waterlogged=false]",
		nil,
		NewMapping{
			477: 3416,
			573: 3416,
			718: 3418,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=low,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 6031,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=low,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13821,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=low,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11237,
		},
	},
	{
		"minecraft:dark_oak_door[facing=south,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7954,
			404: 7955,
			477: 8479,
			573: 8479,
			718: 9015,
		},
	},
	{
		"minecraft:player_head[rotation=0]",
		nil,
		NewMapping{
			573: 6014,
			718: 6550,
			393: 5511,
			404: 5512,
			477: 6014,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=6,powered=false]",
		nil,
		NewMapping{
			477: 1011,
			573: 1011,
			718: 1012,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=none,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13467,
		},
	},
	{
		"minecraft:glass_pane[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 4737,
			573: 4737,
			718: 4753,
			393: 4233,
			404: 4234,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=none,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13675,
		},
	},
	{
		"minecraft:crimson_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15362,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=none,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14010,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 4996,
			393: 4476,
			404: 4477,
			477: 4980,
			573: 4980,
		},
	},
	{
		"minecraft:fire[age=1,east=true,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			477: 1481,
			573: 1481,
			718: 1482,
			393: 1177,
			404: 1178,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=2,south=up,west=side]",
		nil,
		NewMapping{
			477: 2795,
			573: 2795,
			718: 2797,
			393: 2491,
			404: 2492,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=13,south=up,west=up]",
		nil,
		NewMapping{
			718: 2463,
			393: 2157,
			404: 2158,
			477: 2461,
			573: 2461,
		},
	},
	{
		"minecraft:ancient_debris",
		nil,
		NewMapping{
			718: 15827,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 7221,
			404: 7222,
			477: 7728,
			573: 7728,
			718: 8264,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=3,south=side,west=none]",
		nil,
		NewMapping{
			393: 1784,
			404: 1785,
			477: 2088,
			573: 2088,
			718: 2090,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=east,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			477: 4404,
			573: 4404,
			718: 4418,
			393: 3900,
			404: 3901,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=none,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 6006,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=13,south=side,west=up]",
		nil,
		NewMapping{
			573: 2320,
			718: 2322,
			393: 2016,
			404: 2017,
			477: 2320,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10592,
			573: 10592,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10607,
			573: 10607,
		},
	},
	{
		"minecraft:warped_button[face=ceiling,facing=east,powered=false]",
		nil,
		NewMapping{
			718: 15526,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			573: 4610,
			718: 4624,
			393: 4106,
			404: 4107,
			477: 4610,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=north,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			573: 4240,
			718: 4254,
			393: 3736,
			404: 3737,
			477: 4240,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 6086,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=tall,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12883,
		},
	},
	{
		"minecraft:fire[age=7,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1385,
			404: 1386,
			477: 1689,
			573: 1689,
			718: 1690,
		},
	},
	{
		"minecraft:polished_granite_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			477: 10258,
			573: 10258,
			718: 10794,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=none,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12194,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13258,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13120,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13309,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			4:   2225,
			393: 5260,
			404: 5261,
			477: 5764,
			573: 5764,
		},
	},
	{
		"minecraft:purple_shulker_box[facing=down]",
		nil,
		NewMapping{
			477: 8807,
			573: 8807,
			718: 9343,
			4:   3664,
			393: 8282,
			404: 8283,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4686,
			404: 4687,
			477: 5190,
			573: 5190,
			718: 5206,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10168,
			573: 10168,
			718: 10704,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 6136,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=9,powered=true]",
		nil,
		NewMapping{
			718: 567,
			393: 566,
			404: 566,
			477: 566,
			573: 566,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5938,
			404: 5939,
			477: 6445,
			573: 6445,
			718: 6981,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 5966,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13306,
		},
	},
	{
		"minecraft:iron_door[facing=north,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			573: 3819,
			718: 3821,
			393: 3315,
			404: 3316,
			477: 3819,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10172,
			573: 10172,
			718: 10708,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10147,
			573: 10147,
			718: 10683,
		},
	},
	{
		"minecraft:twisting_vines[age=4]",
		nil,
		NewMapping{
			718: 15021,
		},
	},
	{
		"minecraft:blue_glazed_terracotta[facing=south]",
		nil,
		NewMapping{
			573: 8883,
			718: 9419,
			4:   3936,
			393: 8358,
			404: 8359,
			477: 8883,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=23,powered=true]",
		nil,
		NewMapping{
			393: 594,
			404: 594,
			477: 594,
			573: 594,
			718: 595,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9271,
			573: 9271,
			718: 9807,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12666,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=tall,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14204,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10588,
			573: 10588,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9755,
			573: 9755,
			718: 10291,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=low,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12635,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=3,south=side,west=up]",
		nil,
		NewMapping{
			573: 2230,
			718: 2232,
			393: 1926,
			404: 1927,
			477: 2230,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 4564,
			477: 5067,
			573: 5067,
			718: 5083,
			393: 4563,
		},
	},
	{
		"minecraft:warped_door[facing=north,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			718: 15603,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=none,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11437,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=none,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 5678,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 10946,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 17070,
		},
	},
	{
		"minecraft:warped_door[facing=north,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			718: 15604,
		},
	},
	{
		"minecraft:brick_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 4876,
			718: 4892,
			393: 4372,
			404: 4373,
			477: 4876,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10141,
			573: 10141,
			718: 10677,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=tall,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12680,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=low,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16614,
		},
	},
	{
		"minecraft:daylight_detector[inverted=false,power=12]",
		nil,
		NewMapping{
			718: 6722,
			4:   2428,
			393: 5679,
			404: 5680,
			477: 6186,
			573: 6186,
		},
	},
	{
		"minecraft:birch_fence[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7565,
			404: 7566,
			477: 8090,
			573: 8090,
			718: 8626,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12346,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=none,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11848,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=low,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 6237,
		},
	},
	{
		"minecraft:fire[age=15,east=false,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1642,
			404: 1643,
			477: 1946,
			573: 1946,
			718: 1947,
		},
	},
	{
		"minecraft:orange_banner[rotation=5]",
		nil,
		NewMapping{
			393: 6875,
			404: 6876,
			477: 7382,
			573: 7382,
			718: 7918,
		},
	},
	{
		"minecraft:fire[age=7,east=false,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			404: 1385,
			477: 1688,
			573: 1688,
			718: 1689,
			393: 1384,
		},
	},
	{
		"minecraft:lime_shulker_box[facing=west]",
		nil,
		NewMapping{
			718: 9311,
			4:   3588,
			393: 8250,
			404: 8251,
			477: 8775,
			573: 8775,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=tall,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16018,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=tall,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12244,
		},
	},
	{
		"minecraft:warped_wall_sign[facing=east,waterlogged=false]",
		nil,
		NewMapping{
			718: 15734,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=south,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			573: 7027,
			718: 7563,
			393: 6520,
			404: 6521,
			477: 7027,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=10,south=side,west=side]",
		nil,
		NewMapping{
			393: 1990,
			404: 1991,
			477: 2294,
			573: 2294,
			718: 2296,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10729,
			573: 10729,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=low,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16077,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=none,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14574,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9570,
			573: 9570,
			718: 10106,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=tall,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16637,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11530,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=low,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11781,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=none,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16788,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14341,
		},
	},
	{
		"minecraft:detector_rail[powered=false,shape=ascending_west]",
		nil,
		NewMapping{
			393: 1025,
			404: 1025,
			477: 1325,
			573: 1325,
			718: 1326,
			4:   451,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=19,powered=false]",
		nil,
		NewMapping{
			477: 787,
			573: 787,
			718: 788,
		},
	},
	{
		"minecraft:stone_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9629,
			573: 9629,
			718: 10165,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=low,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12319,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=low,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 10934,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=8,south=side,west=up]",
		nil,
		NewMapping{
			393: 2979,
			404: 2980,
			477: 3283,
			573: 3283,
			718: 3285,
		},
	},
	{
		"minecraft:acacia_fence[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 7640,
			404: 7641,
			477: 8165,
			573: 8165,
			718: 8701,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 9817,
			718: 10353,
			477: 9817,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=none,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14224,
		},
	},
	{
		"minecraft:warped_door[facing=west,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			718: 15636,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=13,powered=true]",
		nil,
		NewMapping{
			393: 324,
			404: 324,
			477: 324,
			573: 324,
			718: 325,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=12,powered=false]",
		nil,
		NewMapping{
			393: 473,
			404: 473,
			477: 473,
			573: 473,
			718: 474,
		},
	},
	{
		"minecraft:birch_sign[rotation=5,waterlogged=false]",
		nil,
		NewMapping{
			477: 3454,
			573: 3454,
			718: 3456,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 17055,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14447,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 6521,
			573: 6521,
			718: 7057,
			393: 6014,
			404: 6015,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10142,
			573: 10142,
			718: 10678,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=none,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12171,
		},
	},
	{
		"minecraft:warped_trapdoor[facing=west,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			718: 15229,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 6057,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11383,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=none,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 6201,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=low,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11043,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=none,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11842,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=none,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16165,
		},
	},
	{
		"minecraft:warped_fence_gate[facing=south,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15297,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=tall,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16242,
		},
	},
	{
		"minecraft:oak_leaves[distance=7,persistent=true]",
		nil,
		NewMapping{
			393: 156,
			404: 156,
			477: 156,
			573: 156,
			718: 157,
		},
	},
	{
		"minecraft:light_blue_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 805,
			404: 805,
			477: 1105,
			573: 1105,
			718: 1106,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10497,
			573: 10497,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12472,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=none,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13681,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=none,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12077,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=none,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14431,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=7,south=up,west=none]",
		nil,
		NewMapping{
			393: 1817,
			404: 1818,
			477: 2121,
			573: 2121,
			718: 2123,
		},
	},
	{
		"minecraft:granite_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9929,
			573: 9929,
			718: 10465,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10109,
			573: 10109,
			718: 10645,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=none,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 10871,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6125,
			404: 6126,
			477: 6632,
			573: 6632,
			718: 7168,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=south,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7367,
			404: 7368,
			477: 7892,
			573: 7892,
			718: 8428,
		},
	},
	{
		"minecraft:vine[east=true,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			573: 4777,
			718: 4793,
			393: 4273,
			404: 4274,
			477: 4777,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14616,
		},
	},
	{
		"minecraft:fire[age=3,east=true,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1237,
			404: 1238,
			477: 1541,
			573: 1541,
			718: 1542,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 5194,
			393: 4674,
			404: 4675,
			477: 5178,
			573: 5178,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16576,
		},
	},
	{
		"minecraft:birch_leaves[distance=5,persistent=true]",
		nil,
		NewMapping{
			404: 180,
			477: 180,
			573: 180,
			718: 181,
			393: 180,
		},
	},
	{
		"minecraft:dead_tube_coral_wall_fan[facing=north,waterlogged=false]",
		nil,
		NewMapping{
			393: 8465,
			404: 8481,
			477: 9025,
			573: 9025,
			718: 9561,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=west,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3959,
			404: 3960,
			477: 4463,
			573: 4463,
			718: 4477,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=north,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			477: 8011,
			573: 8011,
			718: 8547,
			393: 7486,
			404: 7487,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5911,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=none,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13570,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=none,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5995,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			573: 8558,
			718: 9094,
			393: 8033,
			404: 8034,
			477: 8558,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=15,powered=false]",
		nil,
		NewMapping{
			404: 729,
			477: 729,
			573: 729,
			718: 730,
			393: 729,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11019,
			573: 11019,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10524,
			477: 10524,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=low,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12010,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14184,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=low,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16177,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=west,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3895,
			404: 3896,
			477: 4399,
			573: 4399,
			718: 4413,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=5,south=none,west=up]",
		nil,
		NewMapping{
			393: 2235,
			404: 2236,
			477: 2539,
			573: 2539,
			718: 2541,
		},
	},
	{
		"minecraft:blue_bed[facing=north,occupied=true,part=head]",
		nil,
		NewMapping{
			477: 1224,
			573: 1224,
			718: 1225,
			393: 924,
			404: 924,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9511,
			573: 9511,
			718: 10047,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14519,
		},
	},
	{
		"minecraft:fire[age=6,east=false,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1348,
			404: 1349,
			477: 1652,
			573: 1652,
			718: 1653,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=tall,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13670,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=tall,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 6296,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=none,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14020,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=low,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14044,
		},
	},
	{
		"minecraft:twisting_vines[age=24]",
		nil,
		NewMapping{
			718: 15041,
		},
	},
	{
		"minecraft:fire[age=9,east=true,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1433,
			404: 1434,
			477: 1737,
			573: 1737,
			718: 1738,
		},
	},
	{
		"minecraft:fire[age=2,east=false,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			404: 1227,
			477: 1530,
			573: 1530,
			718: 1531,
			393: 1226,
		},
	},
	{
		"minecraft:gold_ore",
		nil,
		NewMapping{
			573: 69,
			718: 69,
			4:   224,
			393: 69,
			404: 69,
			477: 69,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=5,south=up,west=none]",
		nil,
		NewMapping{
			393: 2663,
			404: 2664,
			477: 2967,
			573: 2967,
			718: 2969,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12441,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10852,
			573: 10852,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=4,south=side,west=side]",
		nil,
		NewMapping{
			404: 2081,
			477: 2384,
			573: 2384,
			718: 2386,
			393: 2080,
		},
	},
	{
		"minecraft:iron_door[facing=east,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			477: 3864,
			573: 3864,
			718: 3866,
			393: 3360,
			404: 3361,
		},
	},
	{
		"minecraft:fire[age=15,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			718: 1921,
			393: 1616,
			404: 1617,
			477: 1920,
			573: 1920,
		},
	},
	{
		"minecraft:fire[age=5,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			573: 1615,
			718: 1616,
			393: 1311,
			404: 1312,
			477: 1615,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 9451,
			718: 9987,
			477: 9451,
		},
	},
	{
		"minecraft:granite_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9877,
			573: 9877,
			718: 10413,
		},
	},
	{
		"minecraft:spruce_sign[rotation=10,waterlogged=true]",
		nil,
		NewMapping{
			477: 3431,
			573: 3431,
			718: 3433,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=tall,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11078,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6143,
			404: 6144,
			477: 6650,
			573: 6650,
			718: 7186,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=5,south=up,west=none]",
		nil,
		NewMapping{
			393: 2375,
			404: 2376,
			477: 2679,
			573: 2679,
			718: 2681,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 6960,
			718: 7496,
			4:   2629,
			393: 6453,
			404: 6454,
			477: 6960,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=8,powered=false]",
		nil,
		NewMapping{
			393: 365,
			404: 365,
			477: 365,
			573: 365,
			718: 366,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=tall,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 6294,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12344,
		},
	},
	{
		"minecraft:end_stone_brick_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			477: 10283,
			573: 10283,
			718: 10819,
		},
	},
	{
		"minecraft:pumpkin",
		nil,
		NewMapping{
			477: 3996,
			573: 3996,
			718: 3998,
			393: 3492,
			404: 3493,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=14,powered=true]",
		nil,
		NewMapping{
			393: 576,
			404: 576,
			477: 576,
			573: 576,
			718: 577,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 4701,
			404: 4702,
			477: 5205,
			573: 5205,
			718: 5221,
			4:   2049,
		},
	},
	{
		"minecraft:lime_bed[facing=west,occupied=true,part=head]",
		nil,
		NewMapping{
			404: 836,
			477: 1136,
			573: 1136,
			718: 1137,
			393: 836,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12350,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11662,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=none,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13493,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9484,
			573: 9484,
			718: 10020,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 16263,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11443,
		},
	},
	{
		"minecraft:crimson_door[facing=south,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15557,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14083,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=3,south=up,west=side]",
		nil,
		NewMapping{
			393: 1924,
			404: 1925,
			477: 2228,
			573: 2228,
			718: 2230,
		},
	},
	{
		"minecraft:fire[age=12,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1523,
			404: 1524,
			477: 1827,
			573: 1827,
			718: 1828,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			718: 7237,
			393: 6194,
			404: 6195,
			477: 6701,
			573: 6701,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12885,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=6,south=side,west=side]",
		nil,
		NewMapping{
			404: 2531,
			477: 2834,
			573: 2834,
			718: 2836,
			393: 2530,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=5,powered=true]",
		nil,
		NewMapping{
			393: 658,
			404: 658,
			477: 658,
			573: 658,
			718: 659,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 17008,
		},
	},
	{
		"minecraft:fire[age=2,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			718: 1522,
			393: 1217,
			404: 1218,
			477: 1521,
			573: 1521,
		},
	},
	{
		"minecraft:redstone_ore[lit=false]",
		nil,
		NewMapping{
			573: 3884,
			718: 3886,
			4:   1168,
			393: 3380,
			404: 3381,
			477: 3884,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14229,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14292,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=none,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11657,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=none,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 6008,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=none,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14028,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=low,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11673,
		},
	},
	{
		"minecraft:brown_bed[facing=north,occupied=true,part=foot]",
		nil,
		NewMapping{
			404: 941,
			477: 1241,
			573: 1241,
			718: 1242,
			393: 941,
		},
	},
	{
		"minecraft:acacia_door[facing=west,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			404: 7912,
			477: 8436,
			573: 8436,
			718: 8972,
			393: 7911,
		},
	},
	{
		"minecraft:pink_bed[facing=west,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 854,
			404: 854,
			477: 1154,
			573: 1154,
			718: 1155,
		},
	},
	{
		"minecraft:acacia_door[facing=east,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			477: 8457,
			573: 8457,
			718: 8993,
			4:   3136,
			393: 7932,
			404: 7933,
		},
	},
	{
		"minecraft:iron_bars[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			718: 4722,
			393: 4204,
			404: 4205,
			477: 4708,
			573: 4708,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5207,
			404: 5208,
			477: 5711,
			573: 5711,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 6488,
			477: 6994,
			573: 6994,
			718: 7530,
			393: 6487,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14289,
		},
	},
	{
		"minecraft:crimson_fence[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			718: 15088,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=7,south=up,west=none]",
		nil,
		NewMapping{
			393: 2825,
			404: 2826,
			477: 3129,
			573: 3129,
			718: 3131,
		},
	},
	{
		"minecraft:spruce_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 5470,
			393: 4950,
			404: 4951,
			477: 5454,
			573: 5454,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=15,powered=true]",
		nil,
		NewMapping{
			573: 428,
			718: 429,
			393: 428,
			404: 428,
			477: 428,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 15905,
		},
	},
	{
		"minecraft:blue_wool",
		nil,
		NewMapping{
			393: 1094,
			404: 1094,
			477: 1394,
			573: 1394,
			718: 1395,
			4:   571,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=low,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5928,
		},
	},
	{
		"minecraft:granite_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9876,
			573: 9876,
			718: 10412,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=tall,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16986,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9558,
			573: 9558,
			718: 10094,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12895,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=tall,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11395,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12349,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12608,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 5636,
			393: 5116,
			404: 5117,
			477: 5620,
			573: 5620,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=north,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7421,
			404: 7422,
			477: 7946,
			573: 7946,
			718: 8482,
		},
	},
	{
		"minecraft:jungle_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 5606,
			573: 5606,
			718: 5622,
			393: 5102,
			404: 5103,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=low,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12661,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5928,
			404: 5929,
			477: 6435,
			573: 6435,
			718: 6971,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=west,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3640,
			404: 3641,
			477: 4144,
			573: 4144,
			718: 4158,
			4:   1538,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9366,
			718: 9902,
			477: 9366,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=none,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14133,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 6169,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=low,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12234,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=none,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13574,
		},
	},
	{
		"minecraft:brown_wool",
		nil,
		NewMapping{
			4:   572,
			393: 1095,
			404: 1095,
			477: 1395,
			573: 1395,
			718: 1396,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=east,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7386,
			404: 7387,
			477: 7911,
			573: 7911,
			718: 8447,
			4:   2935,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9387,
			573: 9387,
			718: 9923,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=none,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11867,
		},
	},
	{
		"minecraft:birch_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			393: 7270,
			404: 7271,
			477: 7777,
			573: 7777,
			718: 8313,
			4:   2026,
		},
	},
	{
		"minecraft:orange_banner[rotation=12]",
		nil,
		NewMapping{
			718: 7925,
			393: 6882,
			404: 6883,
			477: 7389,
			573: 7389,
		},
	},
	{
		"minecraft:acacia_door[facing=south,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7892,
			404: 7893,
			477: 8417,
			573: 8417,
			718: 8953,
		},
	},
	{
		"minecraft:oak_door[facing=west,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 3140,
			404: 3141,
			477: 3604,
			573: 3604,
			718: 3606,
		},
	},
	{
		"minecraft:fire[age=4,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1270,
			404: 1271,
			477: 1574,
			573: 1574,
			718: 1575,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=none,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 6112,
		},
	},
	{
		"minecraft:red_nether_bricks",
		nil,
		NewMapping{
			718: 9255,
			4:   3440,
			393: 8194,
			404: 8195,
			477: 8719,
			573: 8719,
		},
	},
	{
		"minecraft:diorite_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10229,
			573: 10229,
			718: 10765,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=low,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12746,
		},
	},
	{
		"minecraft:acacia_button[face=ceiling,facing=south,powered=false]",
		nil,
		NewMapping{
			718: 6461,
			393: 5418,
			404: 5419,
			477: 5925,
			573: 5925,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 4048,
			404: 4049,
			477: 4552,
			573: 4552,
			718: 4566,
			4:   1589,
		},
	},
	{
		"minecraft:dark_oak_door[facing=north,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			477: 8472,
			573: 8472,
			718: 9008,
			393: 7947,
			404: 7948,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=west,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 6540,
			404: 6541,
			477: 7047,
			573: 7047,
			718: 7583,
		},
	},
	{
		"minecraft:glass_pane[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 4717,
			573: 4717,
			718: 4733,
			393: 4213,
			404: 4214,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=low,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12221,
		},
	},
	{
		"minecraft:warped_fence_gate[facing=east,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			718: 15316,
		},
	},
	{
		"minecraft:cyan_shulker_box[facing=east]",
		nil,
		NewMapping{
			404: 8273,
			477: 8797,
			573: 8797,
			718: 9333,
			4:   3653,
			393: 8272,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=low,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16178,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=none,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 15929,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=low,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13396,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=none,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13060,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=10,south=none,west=side]",
		nil,
		NewMapping{
			393: 3001,
			404: 3002,
			477: 3305,
			573: 3305,
			718: 3307,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6308,
			404: 6309,
			477: 6815,
			573: 6815,
			718: 7351,
		},
	},
	{
		"minecraft:spruce_door[facing=east,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7728,
			404: 7729,
			477: 8253,
			573: 8253,
			718: 8789,
			4:   3096,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=7,south=none,west=none]",
		nil,
		NewMapping{
			393: 2255,
			404: 2256,
			477: 2559,
			573: 2559,
			718: 2561,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=4,powered=false]",
		nil,
		NewMapping{
			477: 807,
			573: 807,
			718: 808,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14275,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=low,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16096,
		},
	},
	{
		"minecraft:fire[age=12,east=true,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1525,
			404: 1526,
			477: 1829,
			573: 1829,
			718: 1830,
		},
	},
	{
		"minecraft:jungle_door[facing=south,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			404: 7834,
			477: 8358,
			573: 8358,
			718: 8894,
			393: 7833,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=23,powered=false]",
		nil,
		NewMapping{
			477: 945,
			573: 945,
			718: 946,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9150,
			573: 9150,
			718: 9686,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9761,
			573: 9761,
			718: 10297,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 6413,
			404: 6414,
			477: 6920,
			573: 6920,
			718: 7456,
			4:   2631,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5217,
			404: 5218,
			477: 5721,
			573: 5721,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5182,
			404: 5183,
			477: 5686,
			573: 5686,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=3,south=none,west=side]",
		nil,
		NewMapping{
			393: 2650,
			404: 2651,
			477: 2954,
			573: 2954,
			718: 2956,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=east,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3844,
			404: 3845,
			477: 4348,
			573: 4348,
			718: 4362,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=5,south=side,west=none]",
		nil,
		NewMapping{
			404: 2235,
			477: 2538,
			573: 2538,
			718: 2540,
			393: 2234,
		},
	},
	{
		"minecraft:warped_trapdoor[facing=east,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			718: 15254,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=none,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 5798,
		},
	},
	{
		"minecraft:brick_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 4909,
			393: 4389,
			404: 4390,
			477: 4893,
			573: 4893,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6041,
			404: 6042,
			477: 6548,
			573: 6548,
			718: 7084,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=tall,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13004,
		},
	},
	{
		"minecraft:warped_fence[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			718: 15098,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12939,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5186,
			404: 5187,
			477: 5690,
			573: 5690,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=14,south=up,west=none]",
		nil,
		NewMapping{
			573: 2760,
			718: 2762,
			393: 2456,
			404: 2457,
			477: 2760,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12331,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=tall,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5748,
		},
	},
	{
		"minecraft:acacia_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 6411,
			477: 6917,
			573: 6917,
			718: 7453,
			393: 6410,
		},
	},
	{
		"minecraft:jungle_fence[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 7584,
			404: 7585,
			477: 8109,
			573: 8109,
			718: 8645,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=south,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			573: 7957,
			718: 8493,
			393: 7432,
			404: 7433,
			477: 7957,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=14,south=none,west=none]",
		nil,
		NewMapping{
			718: 2624,
			393: 2318,
			404: 2319,
			477: 2622,
			573: 2622,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9200,
			573: 9200,
			718: 9736,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13064,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=low,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14042,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=low,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11661,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=none,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12063,
		},
	},
	{
		"minecraft:crimson_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			718: 15049,
		},
	},
	{
		"minecraft:hay_block[axis=z]",
		nil,
		NewMapping{
			4:   2728,
			393: 6822,
			404: 6823,
			477: 7329,
			573: 7329,
			718: 7865,
		},
	},
	{
		"minecraft:brown_bed[facing=west,occupied=true,part=head]",
		nil,
		NewMapping{
			573: 1248,
			718: 1249,
			393: 948,
			404: 948,
			477: 1248,
		},
	},
	{
		"minecraft:jungle_leaves[distance=2,persistent=true]",
		nil,
		NewMapping{
			4:   303,
			393: 188,
			404: 188,
			477: 188,
			573: 188,
			718: 189,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=none,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11733,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=none,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16920,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6441,
			404: 6442,
			477: 6948,
			573: 6948,
			718: 7484,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 6213,
		},
	},
	{
		"minecraft:birch_leaves[distance=2,persistent=false]",
		nil,
		NewMapping{
			4:   298,
			393: 175,
			404: 175,
			477: 175,
			573: 175,
			718: 176,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			573: 8406,
			718: 8942,
			393: 7881,
			404: 7882,
			477: 8406,
		},
	},
	{
		"minecraft:jungle_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 5616,
			393: 5096,
			404: 5097,
			477: 5600,
			573: 5600,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=none,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 6016,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6297,
			404: 6298,
			477: 6804,
			573: 6804,
			718: 7340,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=7,south=none,west=side]",
		nil,
		NewMapping{
			573: 2846,
			718: 2848,
			393: 2542,
			404: 2543,
			477: 2846,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10858,
			573: 10858,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10535,
			573: 10535,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11424,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=none,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13278,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16536,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 6473,
			404: 6474,
			477: 6980,
			573: 6980,
			718: 7516,
			4:   2628,
		},
	},
	{
		"minecraft:sign[rotation=15,waterlogged=true]",
		nil,
		NewMapping{
			393: 3105,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10653,
			573: 10653,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 6067,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=9,powered=false]",
		nil,
		NewMapping{
			718: 1018,
			477: 1017,
			573: 1017,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=11,waterlogged=true]",
		nil,
		NewMapping{
			718: 3563,
			477: 3561,
			573: 3561,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12632,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=tall,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11164,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=2,south=none,west=up]",
		nil,
		NewMapping{
			404: 2065,
			477: 2368,
			573: 2368,
			718: 2370,
			393: 2064,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 5660,
			573: 5660,
			393: 5156,
			404: 5157,
		},
	},
	{
		"minecraft:chain_command_block[conditional=false,facing=north]",
		nil,
		NewMapping{
			718: 9243,
			4:   3378,
			393: 8182,
			404: 8183,
			477: 8707,
			573: 8707,
		},
	},
	{
		"minecraft:purple_glazed_terracotta[facing=south]",
		nil,
		NewMapping{
			718: 9415,
			4:   3920,
			393: 8354,
			404: 8355,
			477: 8879,
			573: 8879,
		},
	},
	{
		"minecraft:spruce_door[facing=east,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			573: 8256,
			718: 8792,
			4:   3099,
			393: 7731,
			404: 7732,
			477: 8256,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=6,powered=false]",
		nil,
		NewMapping{
			477: 911,
			573: 911,
			718: 912,
		},
	},
	{
		"minecraft:crimson_pressure_plate[powered=false]",
		nil,
		NewMapping{
			718: 15060,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13984,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=none,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13253,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13977,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6186,
			404: 6187,
			477: 6693,
			573: 6693,
			718: 7229,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=0,south=up,west=side]",
		nil,
		NewMapping{
			393: 2761,
			404: 2762,
			477: 3065,
			573: 3065,
			718: 3067,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=east,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			477: 8034,
			573: 8034,
			718: 8570,
			393: 7509,
			404: 7510,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=south,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 4378,
			573: 4378,
			718: 4392,
			393: 3874,
			404: 3875,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10917,
			573: 10917,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12126,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=low,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13851,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=low,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5705,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=11,south=none,west=up]",
		nil,
		NewMapping{
			393: 2577,
			404: 2578,
			477: 2881,
			573: 2881,
			718: 2883,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=9,south=side,west=none]",
		nil,
		NewMapping{
			393: 2558,
			404: 2559,
			477: 2862,
			573: 2862,
			718: 2864,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=15,south=none,west=side]",
		nil,
		NewMapping{
			393: 2182,
			404: 2183,
			477: 2486,
			573: 2486,
			718: 2488,
		},
	},
	{
		"minecraft:acacia_door[facing=west,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			404: 7915,
			477: 8439,
			573: 8439,
			718: 8975,
			4:   3142,
			393: 7914,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=low,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11895,
		},
	},
	{
		"minecraft:dark_oak_fence[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 7670,
			404: 7671,
			477: 8195,
			573: 8195,
			718: 8731,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=low,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 10925,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11986,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=7,powered=true]",
		nil,
		NewMapping{
			393: 562,
			404: 562,
			477: 562,
			573: 562,
			718: 563,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16838,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6257,
		},
	},
	{
		"minecraft:jigsaw[orientation=down_west]",
		nil,
		NewMapping{
			718: 15742,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=none,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11516,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 3194,
			404: 3195,
			477: 3658,
			573: 3658,
			718: 3660,
		},
	},
	{
		"minecraft:daylight_detector[inverted=false,power=13]",
		nil,
		NewMapping{
			393: 5680,
			404: 5681,
			477: 6187,
			573: 6187,
			718: 6723,
			4:   2429,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			718: 6978,
			393: 5935,
			404: 5936,
			477: 6442,
			573: 6442,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=10,powered=true]",
		nil,
		NewMapping{
			393: 418,
			404: 418,
			477: 418,
			573: 418,
			718: 419,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=none,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12302,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12843,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=7,powered=true]",
		nil,
		NewMapping{
			477: 312,
			573: 312,
			718: 313,
			393: 312,
			404: 312,
		},
	},
	{
		"minecraft:bell[attachment=floor,facing=north]",
		nil,
		NewMapping{
			477: 11198,
		},
	},
	{
		"minecraft:campfire[facing=south,lit=false,signal_fire=false,waterlogged=false]",
		nil,
		NewMapping{
			477: 11231,
			573: 11247,
			718: 14905,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=none,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16560,
		},
	},
	{
		"minecraft:fire[age=0,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1166,
			404: 1167,
			477: 1470,
			573: 1470,
			718: 1471,
			4:   816,
		},
	},
	{
		"minecraft:birch_fence[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 8088,
			573: 8088,
			718: 8624,
			393: 7563,
			404: 7564,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=low,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14172,
		},
	},
	{
		"minecraft:diorite_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			477: 10329,
			573: 10329,
			718: 10865,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13014,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14060,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=none,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13048,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16962,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4445,
			404: 4446,
			477: 4949,
			573: 4949,
			718: 4965,
		},
	},
	{
		"minecraft:wither_skeleton_wall_skull[facing=west]",
		nil,
		NewMapping{
			393: 5469,
			404: 5470,
			477: 5992,
			573: 5992,
			718: 6528,
		},
	},
	{
		"minecraft:andesite_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10002,
			573: 10002,
			718: 10538,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=low,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16948,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=north,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3593,
			404: 3594,
			477: 4097,
			573: 4097,
			718: 4111,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=tall,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11820,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11122,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=1,powered=false]",
		nil,
		NewMapping{
			404: 351,
			477: 351,
			573: 351,
			718: 352,
			393: 351,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=8,south=none,west=side]",
		nil,
		NewMapping{
			393: 2407,
			404: 2408,
			477: 2711,
			573: 2711,
			718: 2713,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=5,south=none,west=none]",
		nil,
		NewMapping{
			393: 2669,
			404: 2670,
			477: 2973,
			573: 2973,
			718: 2975,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 6845,
			573: 6845,
			718: 7381,
			393: 6338,
			404: 6339,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=low,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14602,
		},
	},
	{
		"minecraft:cauldron[level=1]",
		nil,
		NewMapping{
			573: 5126,
			718: 5142,
			4:   1889,
			393: 4622,
			404: 4623,
			477: 5126,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=2,south=none,west=side]",
		nil,
		NewMapping{
			718: 2803,
			393: 2497,
			404: 2498,
			477: 2801,
			573: 2801,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10775,
			573: 10775,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 5175,
			573: 5175,
			718: 5191,
			4:   2054,
			393: 4671,
			404: 4672,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12357,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13350,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=none,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16373,
		},
	},
	{
		"minecraft:fire[age=15,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			573: 1949,
			718: 1950,
			393: 1645,
			404: 1646,
			477: 1949,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			404: 6291,
			477: 6797,
			573: 6797,
			718: 7333,
			393: 6290,
		},
	},
	{
		"minecraft:birch_door[facing=west,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7787,
			404: 7788,
			477: 8312,
			573: 8312,
			718: 8848,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=21,powered=false]",
		nil,
		NewMapping{
			477: 991,
			573: 991,
			718: 992,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 4097,
			404: 4098,
			477: 4601,
			573: 4601,
			718: 4615,
		},
	},
	{
		"minecraft:birch_fence[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 7556,
			404: 7557,
			477: 8081,
			573: 8081,
			718: 8617,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=false,north=true,powered=false,south=true,west=false]",
		nil,
		NewMapping{
			718: 5392,
			393: 4872,
			404: 4873,
			477: 5376,
			573: 5376,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=none,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13162,
		},
	},
	{
		"minecraft:birch_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 5484,
			573: 5484,
			718: 5500,
			393: 4980,
			404: 4981,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=6,south=side,west=none]",
		nil,
		NewMapping{
			573: 2115,
			718: 2117,
			393: 1811,
			404: 1812,
			477: 2115,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 10887,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 6043,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=none,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13571,
		},
	},
	{
		"minecraft:fire[age=14,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			404: 1586,
			477: 1889,
			573: 1889,
			718: 1890,
			393: 1585,
		},
	},
	{
		"minecraft:acacia_wood[axis=y]",
		nil,
		NewMapping{
			404: 121,
			477: 121,
			573: 121,
			718: 122,
			4:   2604,
			393: 121,
		},
	},
	{
		"minecraft:birch_fence[east=false,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			718: 8640,
			393: 7579,
			404: 7580,
			477: 8104,
			573: 8104,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9802,
			573: 9802,
			718: 10338,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9339,
			573: 9339,
			718: 9875,
		},
	},
	{
		"minecraft:target[power=11]",
		nil,
		NewMapping{
			718: 15771,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=tall,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11057,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13365,
		},
	},
	{
		"minecraft:repeater[delay=2,facing=north,locked=false,powered=true]",
		nil,
		NewMapping{
			404: 3532,
			477: 4035,
			573: 4035,
			718: 4049,
			4:   1510,
			393: 3531,
		},
	},
	{
		"minecraft:nether_brick_fence[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 4498,
			404: 4499,
			477: 5002,
			573: 5002,
			718: 5018,
		},
	},
	{
		"minecraft:fire[age=7,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			404: 1378,
			477: 1681,
			573: 1681,
			718: 1682,
			393: 1377,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 3984,
			718: 3986,
			393: 3480,
			404: 3481,
			477: 3984,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11444,
		},
	},
	{
		"minecraft:soul_wall_torch[facing=north]",
		nil,
		NewMapping{
			718: 4009,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=none,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5897,
		},
	},
	{
		"minecraft:potted_warped_roots",
		nil,
		NewMapping{
			718: 15837,
		},
	},
	{
		"minecraft:crimson_door[facing=east,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			718: 15590,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 10059,
			718: 10595,
			477: 10059,
		},
	},
	{
		"minecraft:smoker[facing=west,lit=true]",
		nil,
		NewMapping{
			477: 11151,
			573: 11151,
			718: 14807,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=none,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11954,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14072,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10623,
			573: 10623,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11778,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=low,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12109,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 15912,
		},
	},
	{
		"minecraft:black_banner[rotation=0]",
		nil,
		NewMapping{
			404: 7095,
			477: 7601,
			573: 7601,
			718: 8137,
			393: 7094,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=9,south=none,west=none]",
		nil,
		NewMapping{
			393: 2705,
			404: 2706,
			477: 3009,
			573: 3009,
			718: 3011,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11092,
			573: 11092,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10578,
			573: 10578,
		},
	},
	{
		"minecraft:crimson_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 15374,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 6871,
			573: 6871,
			718: 7407,
			393: 6364,
			404: 6365,
		},
	},
	{
		"minecraft:lime_banner[rotation=2]",
		nil,
		NewMapping{
			393: 6936,
			404: 6937,
			477: 7443,
			573: 7443,
			718: 7979,
		},
	},
	{
		"minecraft:mossy_stone_brick_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			477: 10269,
			573: 10269,
			718: 10805,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12395,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9782,
			573: 9782,
			718: 10318,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12005,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 10882,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=low,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12874,
		},
	},
	{
		"minecraft:fire[age=14,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1597,
			404: 1598,
			477: 1901,
			573: 1901,
			718: 1902,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 4955,
			393: 4435,
			404: 4436,
			477: 4939,
			573: 4939,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=16,powered=true]",
		nil,
		NewMapping{
			477: 630,
			573: 630,
			718: 631,
			393: 630,
			404: 630,
		},
	},
	{
		"minecraft:quartz_pillar[axis=z]",
		nil,
		NewMapping{
			573: 6206,
			718: 6742,
			4:   2484,
			393: 5699,
			404: 5700,
			477: 6206,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=none,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16057,
		},
	},
	{
		"minecraft:granite_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9911,
			573: 9911,
			718: 10447,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=none,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13461,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=none,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12819,
		},
	},
	{
		"minecraft:red_concrete_powder",
		nil,
		NewMapping{
			393: 8407,
			404: 8408,
			477: 8932,
			573: 8932,
			718: 9468,
			4:   4046,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 6301,
			477: 6807,
			573: 6807,
			718: 7343,
			393: 6300,
		},
	},
	{
		"minecraft:orange_banner[rotation=2]",
		nil,
		NewMapping{
			393: 6872,
			404: 6873,
			477: 7379,
			573: 7379,
			718: 7915,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=west,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3955,
			404: 3956,
			477: 4459,
			573: 4459,
			718: 4473,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12740,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=none,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14354,
		},
	},
	{
		"minecraft:crimson_fence[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			718: 15066,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=none,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12081,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=14,powered=true]",
		nil,
		NewMapping{
			404: 476,
			477: 476,
			573: 476,
			718: 477,
			393: 476,
		},
	},
	{
		"minecraft:farmland[moisture=3]",
		nil,
		NewMapping{
			573: 3366,
			718: 3368,
			4:   963,
			393: 3062,
			404: 3063,
			477: 3366,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6749,
			404: 6750,
			477: 7256,
			573: 7256,
			718: 7792,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=low,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13198,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12765,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=0,south=none,west=up]",
		nil,
		NewMapping{
			393: 2046,
			404: 2047,
			477: 2350,
			573: 2350,
			718: 2352,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=north,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			573: 4301,
			718: 4315,
			393: 3797,
			404: 3798,
			477: 4301,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10765,
			477: 10765,
		},
	},
	{
		"minecraft:barrel[facing=up,open=true]",
		nil,
		NewMapping{
			477: 11143,
			573: 11143,
			718: 14799,
		},
	},
	{
		"minecraft:iron_door[facing=south,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			718: 3832,
			393: 3326,
			404: 3327,
			477: 3830,
			573: 3830,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 4977,
			718: 4993,
			4:   1748,
			393: 4473,
			404: 4474,
			477: 4977,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=14,south=none,west=none]",
		nil,
		NewMapping{
			393: 2894,
			404: 2895,
			477: 3198,
			573: 3198,
			718: 3200,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 7694,
			393: 6651,
			404: 6652,
			477: 7158,
			573: 7158,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=none,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11948,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 8063,
			404: 8064,
			477: 8588,
			573: 8588,
			718: 9124,
		},
	},
	{
		"minecraft:tripwire_hook[attached=true,facing=south,powered=false]",
		nil,
		NewMapping{
			393: 4742,
			404: 4743,
			477: 5246,
			573: 5246,
			718: 5262,
			4:   2100,
		},
	},
	{
		"minecraft:spruce_door[facing=south,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			404: 7709,
			477: 8233,
			573: 8233,
			718: 8769,
			4:   3089,
			393: 7708,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16526,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=10,south=none,west=none]",
		nil,
		NewMapping{
			393: 1994,
			404: 1995,
			477: 2298,
			573: 2298,
			718: 2300,
		},
	},
	{
		"minecraft:lime_wall_banner[facing=south]",
		nil,
		NewMapping{
			477: 7638,
			573: 7638,
			718: 8174,
			393: 7131,
			404: 7132,
		},
	},
	{
		"minecraft:nether_brick_fence[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 4516,
			404: 4517,
			477: 5020,
			573: 5020,
			718: 5036,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=1,south=up,west=none]",
		nil,
		NewMapping{
			718: 2933,
			393: 2627,
			404: 2628,
			477: 2931,
			573: 2931,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=1,south=none,west=none]",
		nil,
		NewMapping{
			393: 1913,
			404: 1914,
			477: 2217,
			573: 2217,
			718: 2219,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10928,
			573: 10928,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 6438,
			718: 6974,
			393: 5931,
			404: 5932,
			477: 6438,
		},
	},
	{
		"minecraft:jungle_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 5581,
			393: 5061,
			404: 5062,
			477: 5565,
			573: 5565,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 5104,
			393: 4584,
			404: 4585,
			477: 5088,
			573: 5088,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=tall,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14624,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=5,waterlogged=true]",
		nil,
		NewMapping{
			718: 3551,
			477: 3549,
			573: 3549,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10482,
			573: 10482,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=tall,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12701,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13979,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=true,north=true,powered=false,south=false,west=true]",
		nil,
		NewMapping{
			477: 5265,
			573: 5265,
			718: 5281,
			393: 4761,
			404: 4762,
		},
	},
	{
		"minecraft:sticky_piston[extended=false,facing=east]",
		nil,
		NewMapping{
			573: 1335,
			718: 1336,
			4:   469,
			393: 1035,
			404: 1035,
			477: 1335,
		},
	},
	{
		"minecraft:brain_coral_wall_fan[facing=east,waterlogged=true]",
		nil,
		NewMapping{
			404: 8534,
			477: 9078,
			573: 9078,
			718: 9614,
			393: 8518,
		},
	},
	{
		"minecraft:birch_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4971,
			404: 4972,
			477: 5475,
			573: 5475,
			718: 5491,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=tall,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16529,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=none,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11755,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=none,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 17004,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=low,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11125,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 10021,
			718: 10557,
			477: 10021,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=18,powered=false]",
		nil,
		NewMapping{
			477: 785,
			573: 785,
			718: 786,
		},
	},
	{
		"minecraft:cut_sandstone_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			573: 7821,
			718: 8357,
			477: 7821,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 17006,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 17054,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=7,south=up,west=up]",
		nil,
		NewMapping{
			718: 3273,
			393: 2967,
			404: 2968,
			477: 3271,
			573: 3271,
		},
	},
	{
		"minecraft:cut_sandstone",
		nil,
		NewMapping{
			718: 248,
			4:   386,
			393: 247,
			404: 247,
			477: 247,
			573: 247,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13280,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=tall,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14749,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13638,
		},
	},
	{
		"minecraft:orange_banner[rotation=7]",
		nil,
		NewMapping{
			393: 6877,
			404: 6878,
			477: 7384,
			573: 7384,
			718: 7920,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5934,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14613,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6181,
			404: 6182,
			477: 6688,
			573: 6688,
			718: 7224,
		},
	},
	{
		"minecraft:fire[age=14,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			404: 1587,
			477: 1890,
			573: 1890,
			718: 1891,
			393: 1586,
		},
	},
	{
		"minecraft:warped_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15416,
		},
	},
	{
		"minecraft:stripped_warped_hyphae[axis=z]",
		nil,
		NewMapping{
			718: 14969,
		},
	},
	{
		"minecraft:brick_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 4375,
			477: 4878,
			573: 4878,
			718: 4894,
			393: 4374,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14722,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11791,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13294,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=9,south=none,west=up]",
		nil,
		NewMapping{
			573: 2143,
			718: 2145,
			393: 1839,
			404: 1840,
			477: 2143,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9671,
			573: 9671,
			718: 10207,
		},
	},
	{
		"minecraft:acacia_sign[rotation=12,waterlogged=true]",
		nil,
		NewMapping{
			718: 3501,
			477: 3499,
			573: 3499,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10892,
			573: 10892,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 5156,
			718: 5172,
			393: 4652,
			404: 4653,
			477: 5156,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4586,
			404: 4587,
			477: 5090,
			573: 5090,
			718: 5106,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=none,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11870,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=none,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 5884,
		},
	},
	{
		"minecraft:spruce_sign[rotation=1,waterlogged=false]",
		nil,
		NewMapping{
			477: 3414,
			573: 3414,
			718: 3416,
		},
	},
	{
		"minecraft:fire[age=3,east=false,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1258,
			404: 1259,
			477: 1562,
			573: 1562,
			718: 1563,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6743,
			404: 6744,
			477: 7250,
			573: 7250,
			718: 7786,
		},
	},
	{
		"minecraft:oak_button[face=wall,facing=east,powered=false]",
		nil,
		NewMapping{
			393: 5318,
			404: 5319,
			477: 5825,
			573: 5825,
			718: 6361,
			4:   2289,
		},
	},
	{
		"minecraft:glass_pane[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 4733,
			718: 4749,
			393: 4229,
			404: 4230,
			477: 4733,
		},
	},
	{
		"minecraft:twisting_vines[age=17]",
		nil,
		NewMapping{
			718: 15034,
		},
	},
	{
		"minecraft:purple_shulker_box[facing=up]",
		nil,
		NewMapping{
			4:   3665,
			393: 8281,
			404: 8282,
			477: 8806,
			573: 8806,
			718: 9342,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			718: 4598,
			393: 4080,
			404: 4081,
			477: 4584,
			573: 4584,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 9257,
			718: 9793,
			477: 9257,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12858,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=none,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13925,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14397,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=low,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13183,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9278,
			573: 9278,
			718: 9814,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 9389,
			718: 9925,
			477: 9389,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=low,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14582,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13113,
		},
	},
	{
		"minecraft:acacia_wall_sign[facing=west,waterlogged=false]",
		nil,
		NewMapping{
			718: 3764,
			477: 3762,
			573: 3762,
		},
	},
	{
		"minecraft:warped_trapdoor[facing=south,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 15207,
		},
	},
	{
		"minecraft:lava[level=14]",
		nil,
		NewMapping{
			573: 64,
			718: 64,
			4:   190,
			393: 64,
			404: 64,
			477: 64,
		},
	},
	{
		"minecraft:dark_oak_button[face=wall,facing=north,powered=false]",
		nil,
		NewMapping{
			573: 5939,
			718: 6475,
			393: 5432,
			404: 5433,
			477: 5939,
		},
	},
	{
		"minecraft:brown_bed[facing=north,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 943,
			404: 943,
			477: 1243,
			573: 1243,
			718: 1244,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10361,
			573: 10361,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=5,south=side,west=none]",
		nil,
		NewMapping{
			404: 2523,
			477: 2826,
			573: 2826,
			718: 2828,
			393: 2522,
		},
	},
	{
		"minecraft:dark_oak_door[facing=west,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			573: 8503,
			718: 9039,
			4:   3158,
			393: 7978,
			404: 7979,
			477: 8503,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=14,powered=false]",
		nil,
		NewMapping{
			477: 877,
			573: 877,
			718: 878,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13367,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11491,
		},
	},
	{
		"minecraft:dark_oak_button[face=ceiling,facing=north,powered=true]",
		nil,
		NewMapping{
			404: 5440,
			477: 5946,
			573: 5946,
			718: 6482,
			393: 5439,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 7193,
			404: 7194,
			477: 7700,
			573: 7700,
			718: 8236,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9434,
			573: 9434,
			718: 9970,
		},
	},
	{
		"minecraft:spruce_sign[rotation=3,waterlogged=false]",
		nil,
		NewMapping{
			477: 3418,
			573: 3418,
			718: 3420,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 3482,
			404: 3483,
			477: 3986,
			573: 3986,
			718: 3988,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 10058,
			477: 9522,
			573: 9522,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11698,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=low,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12539,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=low,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13612,
		},
	},
	{
		"minecraft:melon_stem[age=0]",
		nil,
		NewMapping{
			573: 4764,
			718: 4780,
			4:   1680,
			393: 4260,
			404: 4261,
			477: 4764,
		},
	},
	{
		"minecraft:moving_piston[facing=east,type=sticky]",
		nil,
		NewMapping{
			573: 1402,
			718: 1403,
			4:   589,
			393: 1102,
			404: 1102,
			477: 1402,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=14,south=none,west=none]",
		nil,
		NewMapping{
			404: 2031,
			477: 2334,
			573: 2334,
			718: 2336,
			393: 2030,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16102,
		},
	},
	{
		"minecraft:fire[age=4,east=true,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1269,
			404: 1270,
			477: 1573,
			573: 1573,
			718: 1574,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9338,
			573: 9338,
			718: 9874,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 17066,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=10,waterlogged=true]",
		nil,
		NewMapping{
			477: 3559,
			573: 3559,
			718: 3561,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=tall,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13226,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6275,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=8,south=side,west=up]",
		nil,
		NewMapping{
			718: 2277,
			393: 1971,
			404: 1972,
			477: 2275,
			573: 2275,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=5,south=up,west=none]",
		nil,
		NewMapping{
			393: 2519,
			404: 2520,
			477: 2823,
			573: 2823,
			718: 2825,
		},
	},
	{
		"minecraft:oak_door[facing=east,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 3164,
			404: 3165,
			477: 3628,
			573: 3628,
			718: 3630,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 9245,
			718: 9781,
			477: 9245,
		},
	},
	{
		"minecraft:oak_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 1720,
			404: 1721,
			477: 2024,
			573: 2024,
			718: 2026,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=23,powered=true]",
		nil,
		NewMapping{
			404: 494,
			477: 494,
			573: 494,
			718: 495,
			393: 494,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=none,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16800,
		},
	},
	{
		"minecraft:lime_bed[facing=east,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 842,
			404: 842,
			477: 1142,
			573: 1142,
			718: 1143,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=none,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14442,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14610,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 5054,
			718: 5070,
			393: 4550,
			404: 4551,
			477: 5054,
		},
	},
	{
		"minecraft:fire[age=13,east=false,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			477: 1876,
			573: 1876,
			718: 1877,
			393: 1572,
			404: 1573,
		},
	},
	{
		"minecraft:granite_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9888,
			573: 9888,
			718: 10424,
		},
	},
	{
		"minecraft:blue_shulker_box[facing=east]",
		nil,
		NewMapping{
			477: 8809,
			573: 8809,
			718: 9345,
			4:   3685,
			393: 8284,
			404: 8285,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=12,south=side,west=side]",
		nil,
		NewMapping{
			393: 2728,
			404: 2729,
			477: 3032,
			573: 3032,
			718: 3034,
		},
	},
	{
		"minecraft:purpur_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 8145,
			404: 8146,
			477: 8670,
			573: 8670,
			718: 9206,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16813,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9230,
			573: 9230,
			718: 9766,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=tall,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 10948,
		},
	},
	{
		"minecraft:red_bed[facing=west,occupied=true,part=head]",
		nil,
		NewMapping{
			404: 980,
			477: 1280,
			573: 1280,
			718: 1281,
			4:   429,
			393: 980,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=4,powered=false]",
		nil,
		NewMapping{
			477: 757,
			573: 757,
			718: 758,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=low,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14268,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=none,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16363,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12033,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=2,south=side,west=up]",
		nil,
		NewMapping{
			718: 3087,
			393: 2781,
			404: 2782,
			477: 3085,
			573: 3085,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6639,
			404: 6640,
			477: 7146,
			573: 7146,
			718: 7682,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=north,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3607,
			404: 3608,
			477: 4111,
			573: 4111,
			718: 4125,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=21,powered=false]",
		nil,
		NewMapping{
			573: 1041,
			718: 1042,
			477: 1041,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=none,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 15922,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=low,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16389,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=18,powered=false]",
		nil,
		NewMapping{
			393: 735,
			404: 735,
			477: 735,
			573: 735,
			718: 736,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=6,south=none,west=up]",
		nil,
		NewMapping{
			393: 1956,
			404: 1957,
			477: 2260,
			573: 2260,
			718: 2262,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 4031,
			404: 4032,
			477: 4535,
			573: 4535,
			718: 4549,
		},
	},
	{
		"minecraft:bell[attachment=single_wall,facing=north,powered=false]",
		nil,
		NewMapping{
			573: 11215,
			718: 14871,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=1,powered=false]",
		nil,
		NewMapping{
			477: 1001,
			573: 1001,
			718: 1002,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=low,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11258,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=none,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12705,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=west,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7506,
			404: 7507,
			477: 8031,
			573: 8031,
			718: 8567,
			4:   2981,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=east,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			573: 4278,
			718: 4292,
			393: 3774,
			404: 3775,
			477: 4278,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 6402,
			573: 6402,
			718: 6938,
			393: 5895,
			404: 5896,
		},
	},
	{
		"minecraft:diorite_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 10221,
			718: 10757,
			477: 10221,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=low,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14706,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9282,
			573: 9282,
			718: 9818,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10081,
			573: 10081,
			718: 10617,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=none,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14351,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13599,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 7677,
			393: 6634,
			404: 6635,
			477: 7141,
			573: 7141,
		},
	},
	{
		"minecraft:jungle_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5087,
			404: 5088,
			477: 5591,
			573: 5591,
			718: 5607,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=none,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11633,
		},
	},
	{
		"minecraft:warped_sign[rotation=6,waterlogged=false]",
		nil,
		NewMapping{
			718: 15700,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=none,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13266,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=tall,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11931,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=7,powered=true]",
		nil,
		NewMapping{
			393: 462,
			404: 462,
			477: 462,
			573: 462,
			718: 463,
		},
	},
	{
		"minecraft:hopper[enabled=true,facing=east]",
		nil,
		NewMapping{
			573: 6196,
			718: 6732,
			4:   2469,
			393: 5689,
			404: 5690,
			477: 6196,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 5864,
			477: 6370,
			573: 6370,
			718: 6906,
			393: 5863,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=8,south=up,west=none]",
		nil,
		NewMapping{
			393: 2546,
			404: 2547,
			477: 2850,
			573: 2850,
			718: 2852,
		},
	},
	{
		"minecraft:green_bed[facing=north,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 956,
			404: 956,
			477: 1256,
			573: 1256,
			718: 1257,
		},
	},
	{
		"minecraft:jungle_fence[east=true,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 7582,
			404: 7583,
			477: 8107,
			573: 8107,
			718: 8643,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 10883,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13007,
		},
	},
	{
		"minecraft:iron_bars[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 4210,
			404: 4211,
			477: 4714,
			573: 4714,
			718: 4728,
			4:   1616,
		},
	},
	{
		"minecraft:oak_button[face=floor,facing=north,powered=false]",
		nil,
		NewMapping{
			573: 5811,
			718: 6347,
			4:   2293,
			393: 5304,
			404: 5305,
			477: 5811,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6050,
			404: 6051,
			477: 6557,
			573: 6557,
			718: 7093,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10924,
			477: 10924,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=tall,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11489,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=tall,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13764,
		},
	},
	{
		"minecraft:dragon_head[rotation=10]",
		nil,
		NewMapping{
			393: 5561,
			404: 5562,
			477: 6064,
			573: 6064,
			718: 6600,
		},
	},
	{
		"minecraft:stripped_oak_wood[axis=z]",
		nil,
		NewMapping{
			393: 128,
			404: 128,
			477: 128,
			573: 128,
			718: 129,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9252,
			573: 9252,
			718: 9788,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=south,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 15155,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4467,
			404: 4468,
			477: 4971,
			573: 4971,
			718: 4987,
		},
	},
	{
		"minecraft:zombie_wall_head[facing=south]",
		nil,
		NewMapping{
			477: 6011,
			573: 6011,
			718: 6547,
			393: 5488,
			404: 5489,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=12,powered=false]",
		nil,
		NewMapping{
			477: 973,
			573: 973,
			718: 974,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11001,
			573: 11001,
		},
	},
	{
		"minecraft:dropper[facing=west,triggered=false]",
		nil,
		NewMapping{
			404: 5800,
			477: 6306,
			573: 6306,
			718: 6842,
			4:   2532,
			393: 5799,
		},
	},
	{
		"minecraft:gray_bed[facing=south,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 865,
			404: 865,
			477: 1165,
			573: 1165,
			718: 1166,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13043,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5255,
			404: 5256,
			477: 5759,
			573: 5759,
		},
	},
	{
		"minecraft:sweet_berry_bush[age=2]",
		nil,
		NewMapping{
			718: 14956,
			477: 11250,
			573: 11266,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16844,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=tall,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11187,
		},
	},
	{
		"minecraft:polished_andesite_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			477: 10322,
			573: 10322,
			718: 10858,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=none,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12413,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=none,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11205,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=tall,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16239,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=11,south=up,west=side]",
		nil,
		NewMapping{
			477: 2732,
			573: 2732,
			718: 2734,
			393: 2428,
			404: 2429,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=8,south=up,west=side]",
		nil,
		NewMapping{
			393: 2689,
			404: 2690,
			477: 2993,
			573: 2993,
			718: 2995,
		},
	},
	{
		"minecraft:cyan_banner[rotation=3]",
		nil,
		NewMapping{
			393: 7001,
			404: 7002,
			477: 7508,
			573: 7508,
			718: 8044,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 10139,
			718: 10675,
			477: 10139,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=low,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11475,
		},
	},
	{
		"minecraft:quartz_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 5751,
			404: 5752,
			477: 6258,
			573: 6258,
			718: 6794,
			4:   2497,
		},
	},
	{
		"minecraft:wheat[age=2]",
		nil,
		NewMapping{
			477: 3357,
			573: 3357,
			718: 3359,
			4:   946,
			393: 3053,
			404: 3054,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11075,
			573: 11075,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13204,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16416,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=9,powered=false]",
		nil,
		NewMapping{
			393: 367,
			404: 367,
			477: 367,
			573: 367,
			718: 368,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=8,south=none,west=side]",
		nil,
		NewMapping{
			573: 2423,
			718: 2425,
			393: 2119,
			404: 2120,
			477: 2423,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=none,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 5692,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=none,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11008,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 16297,
		},
	},
	{
		"minecraft:sticky_piston[extended=true,facing=west]",
		nil,
		NewMapping{
			477: 1331,
			573: 1331,
			718: 1332,
			4:   476,
			393: 1031,
			404: 1031,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=tall,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12374,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13066,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=low,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12436,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=east,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3711,
			404: 3712,
			477: 4215,
			573: 4215,
			718: 4229,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6344,
			404: 6345,
			477: 6851,
			573: 6851,
			718: 7387,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=false,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			404: 4079,
			477: 4582,
			573: 4582,
			718: 4596,
			393: 4078,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=tall,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11282,
		},
	},
	{
		"minecraft:blue_bed[facing=east,occupied=true,part=foot]",
		nil,
		NewMapping{
			573: 1237,
			718: 1238,
			393: 937,
			404: 937,
			477: 1237,
		},
	},
	{
		"minecraft:purpur_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			477: 7875,
			573: 7875,
			718: 8411,
			4:   3280,
			393: 7350,
			404: 7351,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 17062,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=13,south=up,west=none]",
		nil,
		NewMapping{
			393: 2159,
			404: 2160,
			477: 2463,
			573: 2463,
			718: 2465,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=low,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11353,
		},
	},
	{
		"minecraft:acacia_leaves[distance=3,persistent=true]",
		nil,
		NewMapping{
			393: 204,
			404: 204,
			477: 204,
			573: 204,
			718: 205,
		},
	},
	{
		"minecraft:acacia_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 6382,
			404: 6383,
			477: 6889,
			573: 6889,
			718: 7425,
		},
	},
	{
		"minecraft:spruce_leaves[distance=7,persistent=true]",
		nil,
		NewMapping{
			718: 171,
			393: 170,
			404: 170,
			477: 170,
			573: 170,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=true,powered=false,south=true,west=false]",
		nil,
		NewMapping{
			393: 4808,
			404: 4809,
			477: 5312,
			573: 5312,
			718: 5328,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 6706,
			718: 7242,
			393: 6199,
			404: 6200,
			477: 6706,
		},
	},
	{
		"minecraft:warped_button[face=wall,facing=east,powered=true]",
		nil,
		NewMapping{
			718: 15517,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=none,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16142,
		},
	},
	{
		"minecraft:jungle_button[face=ceiling,facing=east,powered=true]",
		nil,
		NewMapping{
			718: 6440,
			393: 5397,
			404: 5398,
			477: 5904,
			573: 5904,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=false,north=true,powered=false,south=true,west=true]",
		nil,
		NewMapping{
			477: 5343,
			573: 5343,
			718: 5359,
			393: 4839,
			404: 4840,
		},
	},
	{
		"minecraft:jungle_sign[rotation=11,waterlogged=true]",
		nil,
		NewMapping{
			477: 3529,
			573: 3529,
			718: 3531,
		},
	},
	{
		"minecraft:granite_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9907,
			573: 9907,
			718: 10443,
		},
	},
	{
		"minecraft:trapped_chest[facing=north,type=right,waterlogged=true]",
		nil,
		NewMapping{
			718: 6626,
			393: 5583,
			404: 5584,
			477: 6090,
			573: 6090,
		},
	},
	{
		"minecraft:warped_sign[rotation=4,waterlogged=true]",
		nil,
		NewMapping{
			718: 15695,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=low,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14054,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16506,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=none,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 6226,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 16676,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13208,
		},
	},
	{
		"minecraft:oak_door[facing=east,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			573: 3632,
			718: 3634,
			4:   1028,
			393: 3168,
			404: 3169,
			477: 3632,
		},
	},
	{
		"minecraft:granite_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9884,
			573: 9884,
			718: 10420,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12087,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=low,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14271,
		},
	},
	{
		"minecraft:dark_oak_door[facing=east,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			718: 9045,
			4:   3160,
			393: 7984,
			404: 7985,
			477: 8509,
			573: 8509,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=10,south=up,west=none]",
		nil,
		NewMapping{
			477: 2580,
			573: 2580,
			718: 2582,
			393: 2276,
			404: 2277,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6147,
			404: 6148,
			477: 6654,
			573: 6654,
			718: 7190,
		},
	},
	{
		"minecraft:brown_glazed_terracotta[facing=south]",
		nil,
		NewMapping{
			404: 8363,
			477: 8887,
			573: 8887,
			718: 9423,
			4:   3952,
			393: 8362,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12340,
		},
	},
	{
		"minecraft:acacia_fence[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 8138,
			573: 8138,
			718: 8674,
			393: 7613,
			404: 7614,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=low,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16484,
		},
	},
	{
		"minecraft:crimson_fence[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			718: 15090,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10516,
			573: 10516,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=tall,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11616,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=12,south=none,west=up]",
		nil,
		NewMapping{
			393: 2010,
			404: 2011,
			477: 2314,
			573: 2314,
			718: 2316,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=north,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3918,
			404: 3919,
			477: 4422,
			573: 4422,
			718: 4436,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 10053,
			573: 10053,
			718: 10589,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=low,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12955,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11961,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=true,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 5673,
			573: 5673,
			393: 5169,
			404: 5170,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=12,south=side,west=up]",
		nil,
		NewMapping{
			393: 2295,
			404: 2296,
			477: 2599,
			573: 2599,
			718: 2601,
		},
	},
	{
		"minecraft:crimson_fence[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			718: 15083,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=none,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13250,
		},
	},
	{
		"minecraft:lime_bed[facing=north,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 831,
			404: 831,
			477: 1131,
			573: 1131,
			718: 1132,
		},
	},
	{
		"minecraft:bubble_coral_wall_fan[facing=south,waterlogged=true]",
		nil,
		NewMapping{
			393: 8522,
			404: 8538,
			477: 9082,
			573: 9082,
			718: 9618,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6307,
			404: 6308,
			477: 6814,
			573: 6814,
			718: 7350,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 7214,
			477: 7720,
			573: 7720,
			718: 8256,
			393: 7213,
		},
	},
	{
		"minecraft:twisting_vines[age=2]",
		nil,
		NewMapping{
			718: 15019,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=9,south=side,west=side]",
		nil,
		NewMapping{
			718: 2719,
			393: 2413,
			404: 2414,
			477: 2717,
			573: 2717,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12006,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13442,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=true,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 4125,
			404: 4126,
			477: 4629,
			573: 4629,
			718: 4643,
		},
	},
	{
		"minecraft:soul_campfire[facing=north,lit=true,signal_fire=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 14922,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=none,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13816,
		},
	},
	{
		"minecraft:dark_oak_fence[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 7672,
			404: 7673,
			477: 8197,
			573: 8197,
			718: 8733,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=north,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3734,
			404: 3735,
			477: 4238,
			573: 4238,
			718: 4252,
		},
	},
	{
		"minecraft:birch_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 5526,
			573: 5526,
			718: 5542,
			393: 5022,
			404: 5023,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 9671,
			477: 9135,
			573: 9135,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=tall,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14748,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6247,
			404: 6248,
			477: 6754,
			573: 6754,
			718: 7290,
		},
	},
	{
		"minecraft:birch_door[facing=south,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7764,
			404: 7765,
			477: 8289,
			573: 8289,
			718: 8825,
		},
	},
	{
		"minecraft:oak_fence[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 3474,
			404: 3475,
			477: 3978,
			573: 3978,
			718: 3980,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=south,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 6751,
			404: 6752,
			477: 7258,
			573: 7258,
			718: 7794,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6451,
			404: 6452,
			477: 6958,
			573: 6958,
			718: 7494,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16817,
		},
	},
	{
		"minecraft:warped_fence_gate[facing=east,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15317,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6131,
			404: 6132,
			477: 6638,
			573: 6638,
			718: 7174,
		},
	},
	{
		"minecraft:acacia_fence[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7619,
			404: 7620,
			477: 8144,
			573: 8144,
			718: 8680,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=2,south=none,west=none]",
		nil,
		NewMapping{
			393: 2642,
			404: 2643,
			477: 2946,
			573: 2946,
			718: 2948,
		},
	},
	{
		"minecraft:purpur_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 8147,
			477: 8671,
			573: 8671,
			718: 9207,
			393: 8146,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=15,south=up,west=side]",
		nil,
		NewMapping{
			393: 2176,
			404: 2177,
			477: 2480,
			573: 2480,
			718: 2482,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 8257,
			393: 7214,
			404: 7215,
			477: 7721,
			573: 7721,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=none,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5877,
		},
	},
	{
		"minecraft:warped_door[facing=west,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			718: 15638,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6796,
			404: 6797,
			477: 7303,
			573: 7303,
			718: 7839,
		},
	},
	{
		"minecraft:chest[facing=north,type=right,waterlogged=true]",
		nil,
		NewMapping{
			404: 1733,
			477: 2036,
			573: 2036,
			718: 2038,
			393: 1732,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=6,south=none,west=up]",
		nil,
		NewMapping{
			477: 2548,
			573: 2548,
			718: 2550,
			393: 2244,
			404: 2245,
		},
	},
	{
		"minecraft:sign[rotation=8,waterlogged=true]",
		nil,
		NewMapping{
			393: 3091,
		},
	},
	{
		"minecraft:dark_oak_door[facing=west,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7965,
			404: 7966,
			477: 8490,
			573: 8490,
			718: 9026,
		},
	},
	{
		"minecraft:light_weighted_pressure_plate[power=11]",
		nil,
		NewMapping{
			404: 5615,
			477: 6121,
			573: 6121,
			718: 6657,
			4:   2363,
			393: 5614,
		},
	},
	{
		"minecraft:warped_planks",
		nil,
		NewMapping{
			718: 15046,
		},
	},
	{
		"minecraft:brown_bed[facing=west,occupied=false,part=foot]",
		nil,
		NewMapping{
			404: 951,
			477: 1251,
			573: 1251,
			718: 1252,
			393: 951,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10964,
			573: 10964,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11793,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=14,south=none,west=none]",
		nil,
		NewMapping{
			404: 1887,
			477: 2190,
			573: 2190,
			718: 2192,
			393: 1886,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=1,powered=true]",
		nil,
		NewMapping{
			573: 300,
			718: 301,
			393: 300,
			404: 300,
			477: 300,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=east,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7387,
			404: 7388,
			477: 7912,
			573: 7912,
			718: 8448,
			4:   2939,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=6,powered=false]",
		nil,
		NewMapping{
			393: 361,
			404: 361,
			477: 361,
			573: 361,
			718: 362,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6589,
			404: 6590,
			477: 7096,
			573: 7096,
			718: 7632,
		},
	},
	{
		"minecraft:light_gray_bed[facing=north,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 878,
			404: 878,
			477: 1178,
			573: 1178,
			718: 1179,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=false,north=true,powered=true,south=true,west=false]",
		nil,
		NewMapping{
			718: 5356,
			393: 4836,
			404: 4837,
			477: 5340,
			573: 5340,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=west,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3885,
			404: 3886,
			477: 4389,
			573: 4389,
			718: 4403,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			718: 7040,
			393: 5997,
			404: 5998,
			477: 6504,
			573: 6504,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=22,powered=true]",
		nil,
		NewMapping{
			718: 943,
			477: 942,
			573: 942,
		},
	},
	{
		"minecraft:bell[attachment=single_wall,facing=north,powered=true]",
		nil,
		NewMapping{
			573: 11214,
			718: 14870,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=south,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3876,
			404: 3877,
			477: 4380,
			573: 4380,
			718: 4394,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=none,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12060,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11017,
			573: 11017,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=tall,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13216,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 5726,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=low,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5801,
		},
	},
	{
		"minecraft:heavy_weighted_pressure_plate[power=7]",
		nil,
		NewMapping{
			404: 5627,
			477: 6133,
			573: 6133,
			718: 6669,
			4:   2375,
			393: 5626,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=south,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 6511,
			404: 6512,
			477: 7018,
			573: 7018,
			718: 7554,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 5836,
			477: 6342,
			573: 6342,
			718: 6878,
			393: 5835,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 5761,
			573: 5761,
			393: 5257,
			404: 5258,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 9702,
			477: 9166,
			573: 9166,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=none,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14121,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=6,south=up,west=none]",
		nil,
		NewMapping{
			393: 2672,
			404: 2673,
			477: 2976,
			573: 2976,
			718: 2978,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=low,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12120,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=10,south=side,west=up]",
		nil,
		NewMapping{
			404: 2422,
			477: 2725,
			573: 2725,
			718: 2727,
			393: 2421,
		},
	},
	{
		"minecraft:jungle_wall_sign[facing=east,waterlogged=false]",
		nil,
		NewMapping{
			477: 3772,
			573: 3772,
			718: 3774,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16108,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=18,powered=false]",
		nil,
		NewMapping{
			393: 385,
			404: 385,
			477: 385,
			573: 385,
			718: 386,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11004,
			573: 11004,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 11059,
			477: 11059,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15854,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=low,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11124,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4588,
			404: 4589,
			477: 5092,
			573: 5092,
			718: 5108,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=false,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			477: 4614,
			573: 4614,
			718: 4628,
			393: 4110,
			404: 4111,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4692,
			404: 4693,
			477: 5196,
			573: 5196,
			718: 5212,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=south,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			573: 4245,
			718: 4259,
			393: 3741,
			404: 3742,
			477: 4245,
		},
	},
	{
		"minecraft:nether_brick_fence[east=true,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 4497,
			404: 4498,
			477: 5001,
			573: 5001,
			718: 5017,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			718: 7302,
			393: 6259,
			404: 6260,
			477: 6766,
			573: 6766,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=low,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11339,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=low,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13176,
		},
	},
	{
		"minecraft:warped_door[facing=south,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			718: 15620,
		},
	},
	{
		"minecraft:glass_pane[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 4740,
			718: 4756,
			393: 4236,
			404: 4237,
			477: 4740,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10805,
			573: 10805,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 10232,
			477: 9696,
			573: 9696,
		},
	},
	{
		"minecraft:crimson_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 15324,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			404: 6733,
			477: 7239,
			573: 7239,
			718: 7775,
			393: 6732,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9415,
			573: 9415,
			718: 9951,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=low,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12108,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=none,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12522,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5160,
			404: 5161,
			477: 5664,
			573: 5664,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9401,
			573: 9401,
			718: 9937,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=tall,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13554,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5939,
			404: 5940,
			477: 6446,
			573: 6446,
			718: 6982,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 9750,
			477: 9214,
			573: 9214,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 16331,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 10884,
		},
	},
	{
		"minecraft:orange_shulker_box[facing=south]",
		nil,
		NewMapping{
			718: 9286,
			4:   3523,
			393: 8225,
			404: 8226,
			477: 8750,
			573: 8750,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=9,south=none,west=none]",
		nil,
		NewMapping{
			393: 2417,
			404: 2418,
			477: 2721,
			573: 2721,
			718: 2723,
		},
	},
	{
		"minecraft:fire_coral_fan[waterlogged=true]",
		nil,
		NewMapping{
			393: 8560,
			404: 8576,
			477: 9020,
			573: 9020,
			718: 9556,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=0,powered=false]",
		nil,
		NewMapping{
			393: 249,
			404: 249,
			477: 249,
			573: 249,
			718: 250,
			4:   400,
		},
	},
	{
		"minecraft:spruce_sign[rotation=9,waterlogged=false]",
		nil,
		NewMapping{
			477: 3430,
			573: 3430,
			718: 3432,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 15940,
		},
	},
	{
		"minecraft:smooth_sandstone_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			477: 10294,
			573: 10294,
			718: 10830,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=none,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13917,
		},
	},
	{
		"minecraft:green_banner[rotation=4]",
		nil,
		NewMapping{
			393: 7066,
			404: 7067,
			477: 7573,
			573: 7573,
			718: 8109,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=9,powered=true]",
		nil,
		NewMapping{
			393: 466,
			404: 466,
			477: 466,
			573: 466,
			718: 467,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=none,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 17014,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=low,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14471,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=south,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3675,
			404: 3676,
			477: 4179,
			573: 4179,
			718: 4193,
		},
	},
	{
		"minecraft:andesite_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9960,
			718: 10496,
			477: 9960,
		},
	},
	{
		"minecraft:red_nether_brick_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			477: 10317,
			573: 10317,
			718: 10853,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10387,
			573: 10387,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=none,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 6124,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5849,
		},
	},
	{
		"minecraft:birch_door[facing=west,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			718: 8840,
			393: 7779,
			404: 7780,
			477: 8304,
			573: 8304,
		},
	},
	{
		"minecraft:piston_head[facing=west,short=true,type=normal]",
		nil,
		NewMapping{
			573: 1371,
			718: 1372,
			393: 1071,
			404: 1071,
			477: 1371,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=14,south=side,west=none]",
		nil,
		NewMapping{
			404: 1884,
			477: 2187,
			573: 2187,
			718: 2189,
			393: 1883,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			718: 4605,
			393: 4087,
			404: 4088,
			477: 4591,
			573: 4591,
		},
	},
	{
		"minecraft:red_sandstone_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			404: 7345,
			477: 7863,
			573: 7863,
			718: 8399,
			4:   2912,
			393: 7344,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9541,
			573: 9541,
			718: 10077,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=none,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5983,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11932,
		},
	},
	{
		"minecraft:spruce_fence[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 7532,
			404: 7533,
			477: 8057,
			573: 8057,
			718: 8593,
		},
	},
	{
		"minecraft:cactus[age=3]",
		nil,
		NewMapping{
			477: 3932,
			573: 3932,
			718: 3934,
			4:   1299,
			393: 3428,
			404: 3429,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=14,south=up,west=side]",
		nil,
		NewMapping{
			393: 2743,
			404: 2744,
			477: 3047,
			573: 3047,
			718: 3049,
		},
	},
	{
		"minecraft:smooth_sandstone_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			477: 10293,
			573: 10293,
			718: 10829,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11859,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=tall,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14404,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=tall,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12693,
		},
	},
	{
		"minecraft:water[level=0]",
		nil,
		NewMapping{
			477: 34,
			573: 34,
			718: 34,
			4:   144,
			393: 34,
			404: 34,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=10,south=side,west=up]",
		nil,
		NewMapping{
			393: 2133,
			404: 2134,
			477: 2437,
			573: 2437,
			718: 2439,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=north,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 6497,
			404: 6498,
			477: 7004,
			573: 7004,
			718: 7540,
			4:   2684,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=low,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14165,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=0,south=up,west=up]",
		nil,
		NewMapping{
			404: 2905,
			477: 3208,
			573: 3208,
			718: 3210,
			393: 2904,
		},
	},
	{
		"minecraft:acacia_fence[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			404: 7616,
			477: 8140,
			573: 8140,
			718: 8676,
			393: 7615,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=tall,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5869,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 10032,
			477: 9496,
			573: 9496,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 17007,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=low,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11245,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13424,
		},
	},
	{
		"minecraft:horn_coral_wall_fan[facing=south,waterlogged=true]",
		nil,
		NewMapping{
			477: 9098,
			573: 9098,
			718: 9634,
			393: 8538,
			404: 8554,
		},
	},
	{
		"minecraft:brick_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			393: 7322,
			404: 7323,
			477: 7841,
			573: 7841,
			718: 8377,
			4:   692,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6138,
			404: 6139,
			477: 6645,
			573: 6645,
			718: 7181,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16629,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6175,
			404: 6176,
			477: 6682,
			573: 6682,
			718: 7218,
		},
	},
	{
		"minecraft:green_stained_glass",
		nil,
		NewMapping{
			393: 3590,
			404: 3591,
			477: 4094,
			573: 4094,
			718: 4108,
			4:   1533,
		},
	},
	{
		"minecraft:warped_door[facing=south,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			718: 15610,
		},
	},
	{
		"minecraft:spruce_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4957,
			404: 4958,
			477: 5461,
			573: 5461,
			718: 5477,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=true,north=true,powered=false,south=true,west=false]",
		nil,
		NewMapping{
			477: 5328,
			573: 5328,
			718: 5344,
			393: 4824,
			404: 4825,
		},
	},
	{
		"minecraft:quartz_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 5766,
			404: 5767,
			477: 6273,
			573: 6273,
			718: 6809,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 16292,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9360,
			573: 9360,
			718: 9896,
		},
	},
	{
		"minecraft:jungle_sign[rotation=1,waterlogged=true]",
		nil,
		NewMapping{
			477: 3509,
			573: 3509,
			718: 3511,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15864,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=6,south=none,west=side]",
		nil,
		NewMapping{
			393: 2821,
			404: 2822,
			477: 3125,
			573: 3125,
			718: 3127,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=8,powered=false]",
		nil,
		NewMapping{
			477: 1015,
			573: 1015,
			718: 1016,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=3,powered=true]",
		nil,
		NewMapping{
			477: 804,
			573: 804,
			718: 805,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=south,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3878,
			404: 3879,
			477: 4382,
			573: 4382,
			718: 4396,
		},
	},
	{
		"minecraft:crimson_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 15325,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=9,south=up,west=up]",
		nil,
		NewMapping{
			393: 2121,
			404: 2122,
			477: 2425,
			573: 2425,
			718: 2427,
		},
	},
	{
		"minecraft:jukebox[has_record=false]",
		nil,
		NewMapping{
			393: 3459,
			404: 3460,
			477: 3963,
			573: 3963,
			718: 3965,
			4:   1344,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6274,
			404: 6275,
			477: 6781,
			573: 6781,
			718: 7317,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6314,
			404: 6315,
			477: 6821,
			573: 6821,
			718: 7357,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9831,
			573: 9831,
			718: 10367,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14721,
		},
	},
	{
		"minecraft:yellow_shulker_box[facing=east]",
		nil,
		NewMapping{
			404: 8243,
			477: 8767,
			573: 8767,
			718: 9303,
			4:   3573,
			393: 8242,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=none,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16159,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16000,
		},
	},
	{
		"minecraft:warped_door[facing=west,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15625,
		},
	},
	{
		"minecraft:fire[age=12,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1545,
			404: 1546,
			477: 1849,
			573: 1849,
			718: 1850,
		},
	},
	{
		"minecraft:spruce_pressure_plate[powered=true]",
		nil,
		NewMapping{
			573: 3873,
			718: 3875,
			393: 3369,
			404: 3370,
			477: 3873,
		},
	},
	{
		"minecraft:spruce_door[facing=west,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			477: 8242,
			573: 8242,
			718: 8778,
			393: 7717,
			404: 7718,
		},
	},
	{
		"minecraft:light_gray_banner[rotation=6]",
		nil,
		NewMapping{
			393: 6988,
			404: 6989,
			477: 7495,
			573: 7495,
			718: 8031,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10732,
			573: 10732,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=tall,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13775,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=7,south=up,west=side]",
		nil,
		NewMapping{
			393: 2968,
			404: 2969,
			477: 3272,
			573: 3272,
			718: 3274,
		},
	},
	{
		"minecraft:acacia_door[facing=south,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7899,
			404: 7900,
			477: 8424,
			573: 8424,
			718: 8960,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6241,
			404: 6242,
			477: 6748,
			573: 6748,
			718: 7284,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 4439,
			477: 4942,
			573: 4942,
			718: 4958,
			393: 4438,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=low,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14693,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 6001,
			477: 6507,
			573: 6507,
			718: 7043,
			393: 6000,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=0,powered=false]",
		nil,
		NewMapping{
			404: 499,
			477: 499,
			573: 499,
			718: 500,
			393: 499,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9865,
			573: 9865,
			718: 10401,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12020,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=none,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 6204,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=none,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14558,
		},
	},
	{
		"minecraft:twisting_vines[age=19]",
		nil,
		NewMapping{
			718: 15036,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=10,south=up,west=up]",
		nil,
		NewMapping{
			404: 1843,
			477: 2146,
			573: 2146,
			718: 2148,
			393: 1842,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=west,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3960,
			404: 3961,
			477: 4464,
			573: 4464,
			718: 4478,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=13,powered=false]",
		nil,
		NewMapping{
			718: 926,
			477: 925,
			573: 925,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 6060,
		},
	},
	{
		"minecraft:birch_fence[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 7567,
			477: 8091,
			573: 8091,
			718: 8627,
			393: 7566,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14415,
		},
	},
	{
		"minecraft:crimson_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 15355,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=low,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12757,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 6743,
			718: 7279,
			393: 6236,
			404: 6237,
			477: 6743,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 5188,
			393: 4668,
			404: 4669,
			477: 5172,
			573: 5172,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=low,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12863,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 16268,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 7661,
			393: 6618,
			404: 6619,
			477: 7125,
			573: 7125,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16917,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=tall,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12577,
		},
	},
	{
		"minecraft:dead_brain_coral_fan[waterlogged=true]",
		nil,
		NewMapping{
			573: 9006,
			718: 9542,
			393: 8546,
			404: 8562,
			477: 9006,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11143,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=none,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 5894,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=none,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12714,
		},
	},
	{
		"minecraft:brick_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4407,
			404: 4408,
			477: 4911,
			573: 4911,
			718: 4927,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=tall,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 6279,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16816,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=low,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12754,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 4664,
			477: 5167,
			573: 5167,
			718: 5183,
			393: 4663,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10511,
			477: 10511,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 9758,
			718: 10294,
			477: 9758,
		},
	},
	{
		"minecraft:beehive[facing=east,honey_level=0]",
		nil,
		NewMapping{
			718: 15818,
			573: 11329,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=south,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			718: 4194,
			393: 3676,
			404: 3677,
			477: 4180,
			573: 4180,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 9664,
			718: 10200,
			477: 9664,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12248,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=tall,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 10967,
		},
	},
	{
		"minecraft:dark_oak_door[facing=south,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			477: 8487,
			573: 8487,
			718: 9023,
			4:   3157,
			393: 7962,
			404: 7963,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=5,powered=true]",
		nil,
		NewMapping{
			573: 358,
			718: 359,
			393: 358,
			404: 358,
			477: 358,
		},
	},
	{
		"minecraft:infested_chiseled_stone_bricks",
		nil,
		NewMapping{
			393: 3982,
			404: 3983,
			477: 4490,
			573: 4490,
			718: 4504,
			4:   1557,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10632,
			477: 10632,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10478,
			573: 10478,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=none,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 6198,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=tall,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13656,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 15989,
		},
	},
	{
		"minecraft:warped_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 15441,
		},
	},
	{
		"minecraft:fire[age=14,east=false,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1603,
			404: 1604,
			477: 1907,
			573: 1907,
			718: 1908,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5253,
			404: 5254,
			477: 5757,
			573: 5757,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 4047,
			404: 4048,
			477: 4551,
			573: 4551,
			718: 4565,
			4:   1588,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=low,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13936,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=tall,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16423,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=none,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13920,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=9,south=none,west=up]",
		nil,
		NewMapping{
			477: 3007,
			573: 3007,
			718: 3009,
			393: 2703,
			404: 2704,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12828,
		},
	},
	{
		"minecraft:birch_fence[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 7555,
			477: 8079,
			573: 8079,
			718: 8615,
			393: 7554,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13890,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=low,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12767,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=3,south=side,west=side]",
		nil,
		NewMapping{
			573: 2951,
			718: 2953,
			393: 2647,
			404: 2648,
			477: 2951,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10543,
			573: 10543,
		},
	},
	{
		"minecraft:warped_sign[rotation=15,waterlogged=true]",
		nil,
		NewMapping{
			718: 15717,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10661,
			477: 10661,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=none,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14675,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=none,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11951,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13441,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=low,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16497,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 17013,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 17086,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16521,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=15,south=none,west=side]",
		nil,
		NewMapping{
			477: 3062,
			573: 3062,
			718: 3064,
			393: 2758,
			404: 2759,
		},
	},
	{
		"minecraft:cactus[age=1]",
		nil,
		NewMapping{
			4:   1297,
			393: 3426,
			404: 3427,
			477: 3930,
			573: 3930,
			718: 3932,
		},
	},
	{
		"minecraft:acacia_sign[rotation=0,waterlogged=false]",
		nil,
		NewMapping{
			477: 3476,
			573: 3476,
			718: 3478,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9397,
			573: 9397,
			718: 9933,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5694,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=14,powered=false]",
		nil,
		NewMapping{
			393: 427,
			404: 427,
			477: 427,
			573: 427,
			718: 428,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=west,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 4138,
			573: 4138,
			718: 4152,
			393: 3634,
			404: 3635,
		},
	},
	{
		"minecraft:dragon_head[rotation=15]",
		nil,
		NewMapping{
			477: 6069,
			573: 6069,
			718: 6605,
			393: 5566,
			404: 5567,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=low,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14469,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9603,
			573: 9603,
			718: 10139,
		},
	},
	{
		"minecraft:birch_button[face=floor,facing=north,powered=false]",
		nil,
		NewMapping{
			404: 5353,
			477: 5859,
			573: 5859,
			718: 6395,
			393: 5352,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=15,south=up,west=up]",
		nil,
		NewMapping{
			393: 2175,
			404: 2176,
			477: 2479,
			573: 2479,
			718: 2481,
		},
	},
	{
		"minecraft:purple_banner[rotation=10]",
		nil,
		NewMapping{
			573: 7531,
			718: 8067,
			393: 7024,
			404: 7025,
			477: 7531,
		},
	},
	{
		"minecraft:fire[age=15,east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			404: 1633,
			477: 1936,
			573: 1936,
			718: 1937,
			393: 1632,
		},
	},
	{
		"minecraft:frosted_ice[age=2]",
		nil,
		NewMapping{
			393: 8190,
			404: 8191,
			477: 8715,
			573: 8715,
			718: 9251,
			4:   3394,
		},
	},
	{
		"minecraft:spruce_fence[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7527,
			404: 7528,
			477: 8052,
			573: 8052,
			718: 8588,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 8017,
			404: 8018,
			477: 8542,
			573: 8542,
			718: 9078,
		},
	},
	{
		"minecraft:acacia_door[facing=west,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7909,
			404: 7910,
			477: 8434,
			573: 8434,
			718: 8970,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=west,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3817,
			404: 3818,
			477: 4321,
			573: 4321,
			718: 4335,
		},
	},
	{
		"minecraft:fire[age=0,east=false,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1162,
			404: 1163,
			477: 1466,
			573: 1466,
			718: 1467,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 5231,
			477: 5734,
			573: 5734,
			393: 5230,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11087,
			573: 11087,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=low,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11888,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=none,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16787,
		},
	},
	{
		"minecraft:cocoa[age=1,facing=east]",
		nil,
		NewMapping{
			393: 4645,
			404: 4646,
			477: 5149,
			573: 5149,
			718: 5165,
			4:   2039,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10789,
			573: 10789,
		},
	},
	{
		"minecraft:andesite_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9940,
			573: 9940,
			718: 10476,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=low,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 5918,
		},
	},
	{
		"minecraft:quartz_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5757,
			404: 5758,
			477: 6264,
			573: 6264,
			718: 6800,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=7,south=up,west=up]",
		nil,
		NewMapping{
			404: 2392,
			477: 2695,
			573: 2695,
			718: 2697,
			393: 2391,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 3989,
			718: 3991,
			393: 3485,
			404: 3486,
			477: 3989,
		},
	},
	{
		"minecraft:fire[age=0,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1153,
			404: 1154,
			477: 1457,
			573: 1457,
			718: 1458,
		},
	},
	{
		"minecraft:spruce_door[facing=north,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7688,
			404: 7689,
			477: 8213,
			573: 8213,
			718: 8749,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 4109,
			404: 4110,
			477: 4613,
			573: 4613,
			718: 4627,
		},
	},
	{
		"minecraft:iron_bars[east=true,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 4695,
			573: 4695,
			718: 4709,
			393: 4191,
			404: 4192,
		},
	},
	{
		"minecraft:acacia_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 6909,
			573: 6909,
			718: 7445,
			393: 6402,
			404: 6403,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=low,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12548,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=10,powered=false]",
		nil,
		NewMapping{
			393: 269,
			404: 269,
			477: 269,
			573: 269,
			718: 270,
		},
	},
	{
		"minecraft:pink_banner[rotation=15]",
		nil,
		NewMapping{
			393: 6965,
			404: 6966,
			477: 7472,
			573: 7472,
			718: 8008,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13544,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5245,
			404: 5246,
			477: 5749,
			573: 5749,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9661,
			573: 9661,
			718: 10197,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=tall,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14424,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=none,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13700,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=true,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			393: 4805,
			404: 4806,
			477: 5309,
			573: 5309,
			718: 5325,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6244,
			404: 6245,
			477: 6751,
			573: 6751,
			718: 7287,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12504,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14192,
		},
	},
	{
		"minecraft:dropper[facing=north,triggered=false]",
		nil,
		NewMapping{
			718: 6836,
			4:   2530,
			393: 5793,
			404: 5794,
			477: 6300,
			573: 6300,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=west,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3696,
			404: 3697,
			477: 4200,
			573: 4200,
			718: 4214,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=true,north=false,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			573: 5365,
			718: 5381,
			393: 4861,
			404: 4862,
			477: 5365,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=none,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5879,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			404: 7218,
			477: 7724,
			573: 7724,
			718: 8260,
			393: 7217,
		},
	},
	{
		"minecraft:dispenser[facing=north,triggered=false]",
		nil,
		NewMapping{
			573: 234,
			718: 235,
			4:   370,
			393: 234,
			404: 234,
			477: 234,
		},
	},
	{
		"minecraft:hay_block[axis=y]",
		nil,
		NewMapping{
			718: 7864,
			4:   2720,
			393: 6821,
			404: 6822,
			477: 7328,
			573: 7328,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=2,powered=false]",
		nil,
		NewMapping{
			477: 853,
			573: 853,
			718: 854,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11103,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13285,
		},
	},
	{
		"minecraft:birch_door[facing=east,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			573: 8327,
			718: 8863,
			4:   3108,
			393: 7802,
			404: 7803,
			477: 8327,
		},
	},
	{
		"minecraft:red_banner[rotation=0]",
		nil,
		NewMapping{
			393: 7078,
			404: 7079,
			477: 7585,
			573: 7585,
			718: 8121,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 3691,
			393: 3225,
			404: 3226,
			477: 3689,
			573: 3689,
		},
	},
	{
		"minecraft:player_wall_head[facing=south]",
		nil,
		NewMapping{
			718: 6567,
			393: 5508,
			404: 5509,
			477: 6031,
			573: 6031,
		},
	},
	{
		"minecraft:water[level=5]",
		nil,
		NewMapping{
			477: 39,
			573: 39,
			718: 39,
			4:   133,
			393: 39,
			404: 39,
		},
	},
	{
		"minecraft:blue_bed[facing=east,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 939,
			404: 939,
			477: 1239,
			573: 1239,
			718: 1240,
		},
	},
	{
		"minecraft:jungle_button[face=wall,facing=west,powered=false]",
		nil,
		NewMapping{
			718: 6431,
			393: 5388,
			404: 5389,
			477: 5895,
			573: 5895,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=tall,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12266,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=tall,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13133,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=none,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11976,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 6130,
			477: 6636,
			573: 6636,
			718: 7172,
			393: 6129,
		},
	},
	{
		"minecraft:light_blue_bed[facing=east,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 809,
			404: 809,
			477: 1109,
			573: 1109,
			718: 1110,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=2,powered=false]",
		nil,
		NewMapping{
			477: 803,
			573: 803,
			718: 804,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=none,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12823,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 9700,
			477: 9164,
			573: 9164,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11098,
		},
	},
	{
		"minecraft:acacia_door[facing=south,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7886,
			404: 7887,
			477: 8411,
			573: 8411,
			718: 8947,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=4,south=none,west=side]",
		nil,
		NewMapping{
			393: 2371,
			404: 2372,
			477: 2675,
			573: 2675,
			718: 2677,
		},
	},
	{
		"minecraft:jungle_sign[rotation=2,waterlogged=true]",
		nil,
		NewMapping{
			477: 3511,
			573: 3511,
			718: 3513,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 10113,
			477: 9577,
			573: 9577,
		},
	},
	{
		"minecraft:spruce_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 5472,
			393: 4952,
			404: 4953,
			477: 5456,
			573: 5456,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4491,
			404: 4492,
			477: 4995,
			573: 4995,
			718: 5011,
		},
	},
	{
		"minecraft:cyan_banner[rotation=15]",
		nil,
		NewMapping{
			393: 7013,
			404: 7014,
			477: 7520,
			573: 7520,
			718: 8056,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=false,north=false,powered=false,south=false,west=true]",
		nil,
		NewMapping{
			718: 5401,
			393: 4881,
			404: 4882,
			477: 5385,
			573: 5385,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11037,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9184,
			573: 9184,
			718: 9720,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=6,powered=false]",
		nil,
		NewMapping{
			718: 962,
			477: 961,
			573: 961,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10975,
			573: 10975,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9395,
			573: 9395,
			718: 9931,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=5,south=none,west=none]",
		nil,
		NewMapping{
			393: 2957,
			404: 2958,
			477: 3261,
			573: 3261,
			718: 3263,
			4:   885,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5243,
			404: 5244,
			477: 5747,
			573: 5747,
		},
	},
	{
		"minecraft:gray_bed[facing=south,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 867,
			404: 867,
			477: 1167,
			573: 1167,
			718: 1168,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6448,
			404: 6449,
			477: 6955,
			573: 6955,
			718: 7491,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=none,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16558,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12399,
		},
	},
	{
		"minecraft:light_gray_glazed_terracotta[facing=south]",
		nil,
		NewMapping{
			718: 9407,
			4:   3888,
			393: 8346,
			404: 8347,
			477: 8871,
			573: 8871,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12268,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=none,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14438,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13985,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12564,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=low,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5925,
		},
	},
	{
		"minecraft:repeater[delay=2,facing=east,locked=true,powered=false]",
		nil,
		NewMapping{
			393: 3542,
			404: 3543,
			477: 4046,
			573: 4046,
			718: 4060,
		},
	},
	{
		"minecraft:bee_nest[facing=south,honey_level=4]",
		nil,
		NewMapping{
			573: 11297,
			718: 15786,
		},
	},
	{
		"minecraft:warped_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 15431,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5977,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 6216,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13179,
		},
	},
	{
		"minecraft:fire[age=6,east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1346,
			404: 1347,
			477: 1650,
			573: 1650,
			718: 1651,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=8,south=none,west=none]",
		nil,
		NewMapping{
			718: 2858,
			393: 2552,
			404: 2553,
			477: 2856,
			573: 2856,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9820,
			573: 9820,
			718: 10356,
		},
	},
	{
		"minecraft:bee_nest[facing=west,honey_level=4]",
		nil,
		NewMapping{
			573: 11303,
			718: 15792,
		},
	},
	{
		"minecraft:trapped_chest[facing=north,type=right,waterlogged=false]",
		nil,
		NewMapping{
			393: 5584,
			404: 5585,
			477: 6091,
			573: 6091,
			718: 6627,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10429,
			477: 10429,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=low,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13734,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16861,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=7,south=up,west=none]",
		nil,
		NewMapping{
			477: 2841,
			573: 2841,
			718: 2843,
			393: 2537,
			404: 2538,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			404: 6295,
			477: 6801,
			573: 6801,
			718: 7337,
			393: 6294,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 8015,
			404: 8016,
			477: 8540,
			573: 8540,
			718: 9076,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=east,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			718: 4493,
			393: 3975,
			404: 3976,
			477: 4479,
			573: 4479,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16819,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 3682,
			718: 3684,
			393: 3218,
			404: 3219,
			477: 3682,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=none,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12295,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 15875,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12690,
		},
	},
	{
		"minecraft:warped_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15406,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14679,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=10,south=side,west=none]",
		nil,
		NewMapping{
			393: 2135,
			404: 2136,
			477: 2439,
			573: 2439,
			718: 2441,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10401,
			573: 10401,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10631,
			477: 10631,
		},
	},
	{
		"minecraft:stone_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9683,
			573: 9683,
			718: 10219,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9834,
			573: 9834,
			718: 10370,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=low,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14715,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=low,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 15968,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=tall,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12476,
		},
	},
	{
		"minecraft:spruce_door[facing=south,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7701,
			404: 7702,
			477: 8226,
			573: 8226,
			718: 8762,
		},
	},
	{
		"minecraft:prismarine",
		nil,
		NewMapping{
			393: 6558,
			404: 6559,
			477: 7065,
			573: 7065,
			718: 7601,
			4:   2688,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=17,powered=false]",
		nil,
		NewMapping{
			404: 383,
			477: 383,
			573: 383,
			718: 384,
			393: 383,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=21,powered=true]",
		nil,
		NewMapping{
			477: 1040,
			573: 1040,
			718: 1041,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 16711,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5906,
			404: 5907,
			477: 6413,
			573: 6413,
			718: 6949,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=12,powered=true]",
		nil,
		NewMapping{
			477: 872,
			573: 872,
			718: 873,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10369,
			573: 10369,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5689,
		},
	},
	{
		"minecraft:pink_shulker_box[facing=north]",
		nil,
		NewMapping{
			4:   3602,
			393: 8253,
			404: 8254,
			477: 8778,
			573: 8778,
			718: 9314,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			404: 4147,
			477: 4650,
			573: 4650,
			718: 4664,
			393: 4146,
		},
	},
	{
		"minecraft:fire[age=5,east=true,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1300,
			404: 1301,
			477: 1604,
			573: 1604,
			718: 1605,
		},
	},
	{
		"minecraft:lectern[facing=west,has_book=false,powered=true]",
		nil,
		NewMapping{
			573: 11187,
			718: 14843,
			477: 11187,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11728,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=tall,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11619,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=south,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3941,
			404: 3942,
			477: 4445,
			573: 4445,
			718: 4459,
		},
	},
	{
		"minecraft:purpur_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 8082,
			404: 8083,
			477: 8607,
			573: 8607,
			718: 9143,
		},
	},
	{
		"minecraft:fire[age=14,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			404: 1615,
			477: 1918,
			573: 1918,
			718: 1919,
			4:   830,
			393: 1614,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=tall,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16976,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=none,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11758,
		},
	},
	{
		"minecraft:brick_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 4885,
			573: 4885,
			718: 4901,
			393: 4381,
			404: 4382,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=1,south=none,west=none]",
		nil,
		NewMapping{
			393: 2489,
			404: 2490,
			477: 2793,
			573: 2793,
			718: 2795,
		},
	},
	{
		"minecraft:dead_horn_coral[waterlogged=false]",
		nil,
		NewMapping{
			404: 8469,
			477: 8993,
			573: 8993,
			718: 9529,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=none,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 15924,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=west,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 6537,
			404: 6538,
			477: 7044,
			573: 7044,
			718: 7580,
			4:   2678,
		},
	},
	{
		"minecraft:yellow_bed[facing=north,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 812,
			404: 812,
			477: 1112,
			573: 1112,
			718: 1113,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=2,south=none,west=side]",
		nil,
		NewMapping{
			573: 3089,
			718: 3091,
			393: 2785,
			404: 2786,
			477: 3089,
		},
	},
	{
		"minecraft:andesite_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9999,
			573: 9999,
			718: 10535,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=tall,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14528,
		},
	},
	{
		"minecraft:birch_door[facing=east,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			477: 8323,
			573: 8323,
			718: 8859,
			393: 7798,
			404: 7799,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4582,
			404: 4583,
			477: 5086,
			573: 5086,
			718: 5102,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=12,powered=true]",
		nil,
		NewMapping{
			393: 672,
			404: 672,
			477: 672,
			573: 672,
			718: 673,
		},
	},
	{
		"minecraft:jungle_button[face=wall,facing=east,powered=false]",
		nil,
		NewMapping{
			393: 5390,
			404: 5391,
			477: 5897,
			573: 5897,
			718: 6433,
		},
	},
	{
		"minecraft:bell[attachment=ceiling,facing=east,powered=true]",
		nil,
		NewMapping{
			573: 11212,
			718: 14868,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11751,
		},
	},
	{
		"minecraft:light_gray_shulker_box[facing=west]",
		nil,
		NewMapping{
			404: 8269,
			477: 8793,
			573: 8793,
			718: 9329,
			4:   3636,
			393: 8268,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=north,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 4431,
			393: 3913,
			404: 3914,
			477: 4417,
			573: 4417,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 4935,
			718: 4951,
			393: 4431,
			404: 4432,
			477: 4935,
		},
	},
	{
		"minecraft:rail[shape=south_east]",
		nil,
		NewMapping{
			393: 3185,
			404: 3186,
			477: 3649,
			573: 3649,
			718: 3651,
			4:   1062,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=12,south=side,west=up]",
		nil,
		NewMapping{
			573: 2311,
			718: 2313,
			393: 2007,
			404: 2008,
			477: 2311,
		},
	},
	{
		"minecraft:brick_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4356,
			404: 4357,
			477: 4860,
			573: 4860,
			718: 4876,
		},
	},
	{
		"minecraft:scaffolding[bottom=false,distance=2,waterlogged=true]",
		nil,
		NewMapping{
			477: 11119,
			573: 11119,
			718: 14775,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13005,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11478,
		},
	},
	{
		"minecraft:warped_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 15404,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=tall,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11724,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=low,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12867,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=0,powered=false]",
		nil,
		NewMapping{
			477: 749,
			573: 749,
			718: 750,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16842,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=tall,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11939,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=west,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 15167,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 5954,
			477: 6460,
			573: 6460,
			718: 6996,
			393: 5953,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 7224,
			404: 7225,
			477: 7731,
			573: 7731,
			718: 8267,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=tall,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13333,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=none,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16347,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 16721,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 4663,
			477: 5166,
			573: 5166,
			718: 5182,
			393: 4662,
		},
	},
	{
		"minecraft:jungle_door[facing=north,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7806,
			404: 7807,
			477: 8331,
			573: 8331,
			718: 8867,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 4069,
			404: 4070,
			477: 4573,
			573: 4573,
			718: 4587,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=low,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13946,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=north,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7389,
			404: 7390,
			477: 7914,
			573: 7914,
			718: 8450,
		},
	},
	{
		"minecraft:piston_head[facing=west,short=false,type=sticky]",
		nil,
		NewMapping{
			573: 1374,
			718: 1375,
			4:   556,
			393: 1074,
			404: 1074,
			477: 1374,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9242,
			573: 9242,
			718: 9778,
		},
	},
	{
		"minecraft:nether_wart[age=3]",
		nil,
		NewMapping{
			477: 5115,
			573: 5115,
			718: 5131,
			4:   1843,
			393: 4611,
			404: 4612,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			573: 4680,
			718: 4694,
			393: 4176,
			404: 4177,
			477: 4680,
		},
	},
	{
		"minecraft:spruce_door[facing=north,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7687,
			404: 7688,
			477: 8212,
			573: 8212,
			718: 8748,
		},
	},
	{
		"minecraft:conduit[waterlogged=false]",
		nil,
		NewMapping{
			404: 8590,
			477: 9114,
			573: 9114,
			718: 9650,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10495,
			573: 10495,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10658,
			573: 10658,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10127,
			573: 10127,
			718: 10663,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13713,
		},
	},
	{
		"minecraft:spruce_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 5435,
			4:   2146,
			393: 4915,
			404: 4916,
			477: 5419,
			573: 5419,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=9,powered=true]",
		nil,
		NewMapping{
			573: 416,
			718: 417,
			393: 416,
			404: 416,
			477: 416,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 4959,
			573: 4959,
			718: 4975,
			393: 4455,
			404: 4456,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10727,
			573: 10727,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=12,south=up,west=up]",
		nil,
		NewMapping{
			393: 1860,
			404: 1861,
			477: 2164,
			573: 2164,
			718: 2166,
		},
	},
	{
		"minecraft:fire[age=15,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1633,
			404: 1634,
			477: 1937,
			573: 1937,
			718: 1938,
		},
	},
	{
		"minecraft:jungle_wall_sign[facing=south,waterlogged=true]",
		nil,
		NewMapping{
			477: 3767,
			573: 3767,
			718: 3769,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13100,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12025,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=16,powered=false]",
		nil,
		NewMapping{
			393: 581,
			404: 581,
			477: 581,
			573: 581,
			718: 582,
		},
	},
	{
		"minecraft:farmland[moisture=6]",
		nil,
		NewMapping{
			4:   966,
			393: 3065,
			404: 3066,
			477: 3369,
			573: 3369,
			718: 3371,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5918,
			404: 5919,
			477: 6425,
			573: 6425,
			718: 6961,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10066,
			573: 10066,
			718: 10602,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10853,
			573: 10853,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=none,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12090,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11600,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15907,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=9,south=up,west=none]",
		nil,
		NewMapping{
			573: 2139,
			718: 2141,
			393: 1835,
			404: 1836,
			477: 2139,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6359,
			404: 6360,
			477: 6866,
			573: 6866,
			718: 7402,
		},
	},
	{
		"minecraft:grass_block[snowy=false]",
		nil,
		NewMapping{
			393: 9,
			404: 9,
			477: 9,
			573: 9,
			718: 9,
			4:   32,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10541,
			573: 10541,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11477,
		},
	},
	{
		"minecraft:acacia_wall_sign[facing=south,waterlogged=false]",
		nil,
		NewMapping{
			477: 3760,
			573: 3760,
			718: 3762,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 16719,
		},
	},
	{
		"minecraft:oak_door[facing=south,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 3124,
			404: 3125,
			477: 3588,
			573: 3588,
			718: 3590,
		},
	},
	{
		"minecraft:andesite_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9958,
			573: 9958,
			718: 10494,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 11014,
			573: 11014,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=tall,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13989,
		},
	},
	{
		"minecraft:mossy_cobblestone_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			573: 10281,
			718: 10817,
			477: 10281,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11019,
		},
	},
	{
		"minecraft:warped_door[facing=south,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			718: 15614,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13769,
		},
	},
	{
		"minecraft:orange_banner[rotation=11]",
		nil,
		NewMapping{
			393: 6881,
			404: 6882,
			477: 7388,
			573: 7388,
			718: 7924,
		},
	},
	{
		"minecraft:fire[age=13,east=true,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1559,
			404: 1560,
			477: 1863,
			573: 1863,
			718: 1864,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10525,
			477: 10525,
		},
	},
	{
		"minecraft:stripped_spruce_log[axis=x]",
		nil,
		NewMapping{
			573: 90,
			718: 91,
			393: 90,
			404: 90,
			477: 90,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=2,south=none,west=side]",
		nil,
		NewMapping{
			393: 2641,
			404: 2642,
			477: 2945,
			573: 2945,
			718: 2947,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=5,south=up,west=none]",
		nil,
		NewMapping{
			393: 1943,
			404: 1944,
			477: 2247,
			573: 2247,
			718: 2249,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16647,
		},
	},
	{
		"minecraft:warped_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15477,
		},
	},
	{
		"minecraft:spruce_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4900,
			404: 4901,
			477: 5404,
			573: 5404,
			718: 5420,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=23,powered=false]",
		nil,
		NewMapping{
			393: 345,
			404: 345,
			477: 345,
			573: 345,
			718: 346,
		},
	},
	{
		"minecraft:redstone_lamp[lit=false]",
		nil,
		NewMapping{
			477: 5141,
			573: 5141,
			718: 5157,
			4:   1968,
			393: 4637,
			404: 4638,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 7198,
			573: 7198,
			718: 7734,
			393: 6691,
			404: 6692,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=west,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3890,
			404: 3891,
			477: 4394,
			573: 4394,
			718: 4408,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11167,
		},
	},
	{
		"minecraft:polished_diorite",
		nil,
		NewMapping{
			4:   20,
			393: 5,
			404: 5,
			477: 5,
			573: 5,
			718: 5,
		},
	},
	{
		"minecraft:spruce_door[facing=north,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7682,
			404: 7683,
			477: 8207,
			573: 8207,
			718: 8743,
		},
	},
	{
		"minecraft:acacia_leaves[distance=6,persistent=true]",
		nil,
		NewMapping{
			573: 210,
			718: 211,
			393: 210,
			404: 210,
			477: 210,
		},
	},
	{
		"minecraft:andesite_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9981,
			573: 9981,
			718: 10517,
		},
	},
	{
		"minecraft:acacia_fence[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7633,
			404: 7634,
			477: 8158,
			573: 8158,
			718: 8694,
		},
	},
	{
		"minecraft:purpur_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 8143,
			404: 8144,
			477: 8668,
			573: 8668,
			718: 9204,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=5,south=up,west=up]",
		nil,
		NewMapping{
			573: 2101,
			718: 2103,
			393: 1797,
			404: 1798,
			477: 2101,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11061,
			573: 11061,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 7833,
			393: 6790,
			404: 6791,
			477: 7297,
			573: 7297,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=20,powered=false]",
		nil,
		NewMapping{
			477: 289,
			573: 289,
			718: 290,
			393: 289,
			404: 289,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6121,
			404: 6122,
			477: 6628,
			573: 6628,
			718: 7164,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 9841,
			477: 9305,
			573: 9305,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 5872,
			477: 6378,
			573: 6378,
			718: 6914,
			393: 5871,
		},
	},
	{
		"minecraft:yellow_bed[facing=west,occupied=false,part=head]",
		nil,
		NewMapping{
			718: 1123,
			393: 822,
			404: 822,
			477: 1122,
			573: 1122,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=none,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13140,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=none,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13373,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12486,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=1,south=side,west=side]",
		nil,
		NewMapping{
			477: 2357,
			573: 2357,
			718: 2359,
			393: 2053,
			404: 2054,
		},
	},
	{
		"minecraft:horn_coral_fan[waterlogged=false]",
		nil,
		NewMapping{
			404: 8579,
			477: 9023,
			573: 9023,
			718: 9559,
			393: 8563,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=none,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 10886,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=tall,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16968,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=9,south=side,west=up]",
		nil,
		NewMapping{
			404: 2701,
			477: 3004,
			573: 3004,
			718: 3006,
			393: 2700,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9447,
			573: 9447,
			718: 9983,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10456,
			573: 10456,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=low,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 17048,
		},
	},
	{
		"minecraft:brown_banner[rotation=0]",
		nil,
		NewMapping{
			477: 7553,
			573: 7553,
			718: 8089,
			393: 7046,
			404: 7047,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=1,powered=true]",
		nil,
		NewMapping{
			404: 400,
			477: 400,
			573: 400,
			718: 401,
			393: 400,
		},
	},
	{
		"minecraft:fire[age=4,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1293,
			404: 1294,
			477: 1597,
			573: 1597,
			718: 1598,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4694,
			404: 4695,
			477: 5198,
			573: 5198,
			718: 5214,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=low,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16936,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=tall,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12373,
		},
	},
	{
		"minecraft:warped_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15447,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14167,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 16667,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=none,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14244,
		},
	},
	{
		"minecraft:jungle_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 5607,
			573: 5607,
			718: 5623,
			393: 5103,
			404: 5104,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=tall,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11943,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=tall,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12370,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=low,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 5710,
		},
	},
	{
		"minecraft:dead_fire_coral_wall_fan[facing=west,waterlogged=false]",
		nil,
		NewMapping{
			393: 8493,
			404: 8509,
			477: 9053,
			573: 9053,
			718: 9589,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10618,
			573: 10618,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16355,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=none,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14454,
		},
	},
	{
		"minecraft:blue_bed[facing=east,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 936,
			404: 936,
			477: 1236,
			573: 1236,
			718: 1237,
		},
	},
	{
		"minecraft:mycelium[snowy=false]",
		nil,
		NewMapping{
			393: 4493,
			404: 4494,
			477: 4997,
			573: 4997,
			718: 5013,
			4:   1760,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=12,south=side,west=up]",
		nil,
		NewMapping{
			393: 1863,
			404: 1864,
			477: 2167,
			573: 2167,
			718: 2169,
		},
	},
	{
		"minecraft:brown_bed[facing=south,occupied=true,part=head]",
		nil,
		NewMapping{
			573: 1244,
			718: 1245,
			393: 944,
			404: 944,
			477: 1244,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=tall,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12686,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13753,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14154,
		},
	},
	{
		"minecraft:polished_diorite_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			477: 10271,
			573: 10271,
			718: 10807,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=tall,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11058,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=8,south=none,west=none]",
		nil,
		NewMapping{
			393: 2120,
			404: 2121,
			477: 2424,
			573: 2424,
			718: 2426,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5845,
			404: 5846,
			477: 6352,
			573: 6352,
			718: 6888,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=12,south=side,west=none]",
		nil,
		NewMapping{
			393: 2585,
			404: 2586,
			477: 2889,
			573: 2889,
			718: 2891,
		},
	},
	{
		"minecraft:iron_door[facing=south,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 3328,
			404: 3329,
			477: 3832,
			573: 3832,
			718: 3834,
		},
	},
	{
		"minecraft:brown_wall_banner[facing=east]",
		nil,
		NewMapping{
			404: 7162,
			477: 7668,
			573: 7668,
			718: 8204,
			393: 7161,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=true,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			404: 8012,
			477: 8536,
			573: 8536,
			718: 9072,
			393: 8011,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14120,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=tall,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13326,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4415,
			404: 4416,
			477: 4919,
			573: 4919,
			718: 4935,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			718: 7278,
			4:   2572,
			393: 6235,
			404: 6236,
			477: 6742,
			573: 6742,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5237,
			404: 5238,
			477: 5741,
			573: 5741,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=low,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16598,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14409,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=west,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			718: 4837,
			393: 4317,
			404: 4318,
			477: 4821,
			573: 4821,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=20,powered=true]",
		nil,
		NewMapping{
			393: 338,
			404: 338,
			477: 338,
			573: 338,
			718: 339,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=8,south=up,west=side]",
		nil,
		NewMapping{
			404: 2978,
			477: 3281,
			573: 3281,
			718: 3283,
			393: 2977,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=tall,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 6082,
		},
	},
	{
		"minecraft:warped_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 15458,
		},
	},
	{
		"minecraft:fire[age=5,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			718: 1628,
			393: 1323,
			404: 1324,
			477: 1627,
			573: 1627,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			404: 4123,
			477: 4626,
			573: 4626,
			718: 4640,
			393: 4122,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9612,
			573: 9612,
			718: 10148,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5975,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11175,
		},
	},
	{
		"minecraft:stone_button[face=floor,facing=north,powered=true]",
		nil,
		NewMapping{
			393: 3391,
			404: 3392,
			477: 3895,
			573: 3895,
			718: 3897,
			4:   1245,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=12,south=none,west=none]",
		nil,
		NewMapping{
			393: 1868,
			404: 1869,
			477: 2172,
			573: 2172,
			718: 2174,
		},
	},
	{
		"minecraft:birch_button[face=floor,facing=north,powered=true]",
		nil,
		NewMapping{
			393: 5351,
			404: 5352,
			477: 5858,
			573: 5858,
			718: 6394,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=none,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13901,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10566,
			573: 10566,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12673,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=false,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			477: 4672,
			573: 4672,
			718: 4686,
			393: 4168,
			404: 4169,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=1,south=side,west=up]",
		nil,
		NewMapping{
			393: 1908,
			404: 1909,
			477: 2212,
			573: 2212,
			718: 2214,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10569,
			477: 10569,
		},
	},
	{
		"minecraft:crimson_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15386,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=low,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14575,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=none,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12732,
		},
	},
	{
		"minecraft:daylight_detector[inverted=true,power=0]",
		nil,
		NewMapping{
			718: 6694,
			4:   2848,
			393: 5651,
			404: 5652,
			477: 6158,
			573: 6158,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 7093,
			718: 7629,
			393: 6586,
			404: 6587,
			477: 7093,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=9,powered=true]",
		nil,
		NewMapping{
			477: 816,
			573: 816,
			718: 817,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10998,
			573: 10998,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=low,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11904,
		},
	},
	{
		"minecraft:prismarine_brick_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			393: 6809,
			404: 6810,
			477: 7316,
			573: 7316,
			718: 7852,
		},
	},
	{
		"minecraft:iron_door[facing=west,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			404: 3351,
			477: 3854,
			573: 3854,
			718: 3856,
			4:   1138,
			393: 3350,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=19,powered=false]",
		nil,
		NewMapping{
			477: 587,
			573: 587,
			718: 588,
			393: 587,
			404: 587,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 7701,
			573: 7701,
			718: 8237,
			393: 7194,
			404: 7195,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=none,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5881,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13292,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=1,south=side,west=up]",
		nil,
		NewMapping{
			404: 2917,
			477: 3220,
			573: 3220,
			718: 3222,
			393: 2916,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5868,
			404: 5869,
			477: 6375,
			573: 6375,
			718: 6911,
		},
	},
	{
		"minecraft:stone_brick_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			477: 7843,
			573: 7843,
			718: 8379,
			4:   717,
			393: 7324,
			404: 7325,
		},
	},
	{
		"minecraft:cut_sandstone_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			477: 7819,
			573: 7819,
			718: 8355,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=none,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11644,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=13,south=none,west=up]",
		nil,
		NewMapping{
			393: 2883,
			404: 2884,
			477: 3187,
			573: 3187,
			718: 3189,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=none,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 6094,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=none,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16578,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=low,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14276,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=11,south=none,west=none]",
		nil,
		NewMapping{
			393: 2867,
			404: 2868,
			477: 3171,
			573: 3171,
			718: 3173,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=13,south=up,west=side]",
		nil,
		NewMapping{
			573: 2174,
			718: 2176,
			393: 1870,
			404: 1871,
			477: 2174,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=low,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 6240,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			477: 4508,
			573: 4508,
			718: 4522,
			393: 4004,
			404: 4005,
		},
	},
	{
		"minecraft:birch_fence[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 8100,
			573: 8100,
			718: 8636,
			393: 7575,
			404: 7576,
		},
	},
	{
		"minecraft:warped_door[facing=south,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			718: 15608,
		},
	},
	{
		"minecraft:repeater[delay=3,facing=west,locked=true,powered=false]",
		nil,
		NewMapping{
			393: 3554,
			404: 3555,
			477: 4058,
			573: 4058,
			718: 4072,
		},
	},
	{
		"minecraft:iron_door[facing=north,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 3317,
			404: 3318,
			477: 3821,
			573: 3821,
			718: 3823,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 10017,
			718: 10553,
			477: 10017,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16651,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11059,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=12,south=side,west=up]",
		nil,
		NewMapping{
			718: 2745,
			393: 2439,
			404: 2440,
			477: 2743,
			573: 2743,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=12,south=none,west=up]",
		nil,
		NewMapping{
			393: 2586,
			404: 2587,
			477: 2890,
			573: 2890,
			718: 2892,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10988,
			573: 10988,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=none,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14219,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=tall,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14302,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=low,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11132,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14524,
		},
	},
	{
		"minecraft:jungle_fence[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 7604,
			404: 7605,
			477: 8129,
			573: 8129,
			718: 8665,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 5861,
			477: 6367,
			573: 6367,
			718: 6903,
			393: 5860,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 6704,
			477: 7210,
			573: 7210,
			718: 7746,
			393: 6703,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14176,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14615,
		},
	},
	{
		"minecraft:fire[age=3,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			718: 1539,
			393: 1234,
			404: 1235,
			477: 1538,
			573: 1538,
		},
	},
	{
		"minecraft:repeater[delay=2,facing=south,locked=false,powered=true]",
		nil,
		NewMapping{
			393: 3535,
			404: 3536,
			477: 4039,
			573: 4039,
			718: 4053,
			4:   1508,
		},
	},
	{
		"minecraft:jungle_leaves[distance=1,persistent=true]",
		nil,
		NewMapping{
			4:   295,
			393: 186,
			404: 186,
			477: 186,
			573: 186,
			718: 187,
		},
	},
	{
		"minecraft:oak_door[facing=north,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			393: 3122,
			404: 3123,
			477: 3586,
			573: 3586,
			718: 3588,
			4:   1027,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=none,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11624,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=8,powered=false]",
		nil,
		NewMapping{
			393: 465,
			404: 465,
			477: 465,
			573: 465,
			718: 466,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=18,powered=true]",
		nil,
		NewMapping{
			477: 834,
			573: 834,
			718: 835,
		},
	},
	{
		"minecraft:spruce_sign[rotation=6,waterlogged=true]",
		nil,
		NewMapping{
			477: 3423,
			573: 3423,
			718: 3425,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=tall,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16110,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=none,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13254,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16568,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=21,powered=false]",
		nil,
		NewMapping{
			573: 541,
			718: 542,
			393: 541,
			404: 541,
			477: 541,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=12,south=none,west=up]",
		nil,
		NewMapping{
			393: 2298,
			404: 2299,
			477: 2602,
			573: 2602,
			718: 2604,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=east,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 7585,
			393: 6542,
			404: 6543,
			477: 7049,
			573: 7049,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 10114,
			477: 9578,
			573: 9578,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 10018,
			477: 9482,
			573: 9482,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16594,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=tall,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 17076,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=none,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 6114,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 7382,
			393: 6339,
			404: 6340,
			477: 6846,
			573: 6846,
		},
	},
	{
		"minecraft:diamond_block",
		nil,
		NewMapping{
			393: 3049,
			404: 3050,
			477: 3353,
			573: 3353,
			718: 3355,
			4:   912,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=true,powered=true,south=false,west=false]",
		nil,
		NewMapping{
			404: 4807,
			477: 5310,
			573: 5310,
			718: 5326,
			393: 4806,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=none,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11087,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=none,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12942,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5242,
			404: 5243,
			477: 5746,
			573: 5746,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=9,powered=false]",
		nil,
		NewMapping{
			718: 918,
			477: 917,
			573: 917,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13843,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=low,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 17061,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 5856,
			477: 6362,
			573: 6362,
			718: 6898,
			393: 5855,
		},
	},
	{
		"minecraft:iron_door[facing=west,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 3339,
			404: 3340,
			477: 3843,
			573: 3843,
			718: 3845,
		},
	},
	{
		"minecraft:lime_shulker_box[facing=up]",
		nil,
		NewMapping{
			393: 8251,
			404: 8252,
			477: 8776,
			573: 8776,
			718: 9312,
			4:   3585,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 5032,
			573: 5032,
			718: 5048,
			393: 4528,
			404: 4529,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=1,south=side,west=none]",
		nil,
		NewMapping{
			718: 2216,
			393: 1910,
			404: 1911,
			477: 2214,
			573: 2214,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10421,
			477: 10421,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6161,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=north,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			718: 15129,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=none,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 10894,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=tall,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14640,
		},
	},
	{
		"minecraft:gray_bed[facing=south,occupied=true,part=head]",
		nil,
		NewMapping{
			718: 1165,
			393: 864,
			404: 864,
			477: 1164,
			573: 1164,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=north,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			4:   2962,
			393: 7428,
			404: 7429,
			477: 7953,
			573: 7953,
			718: 8489,
		},
	},
	{
		"minecraft:kelp[age=7]",
		nil,
		NewMapping{
			393: 8416,
			404: 8417,
			477: 8941,
			573: 8941,
			718: 9477,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10247,
			573: 10247,
			718: 10783,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=none,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16906,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=none,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 15942,
		},
	},
	{
		"minecraft:melon_stem[age=1]",
		nil,
		NewMapping{
			393: 4261,
			404: 4262,
			477: 4765,
			573: 4765,
			718: 4781,
			4:   1681,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 7209,
			718: 7745,
			393: 6702,
			404: 6703,
			477: 7209,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10613,
			573: 10613,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9229,
			573: 9229,
			718: 9765,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=none,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 15944,
		},
	},
	{
		"minecraft:sugar_cane[age=11]",
		nil,
		NewMapping{
			393: 3453,
			404: 3454,
			477: 3957,
			573: 3957,
			718: 3959,
			4:   1339,
		},
	},
	{
		"minecraft:fire[age=11,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			404: 1500,
			477: 1803,
			573: 1803,
			718: 1804,
			393: 1499,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=22,powered=true]",
		nil,
		NewMapping{
			477: 892,
			573: 892,
			718: 893,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=east,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			718: 15189,
		},
	},
	{
		"minecraft:oak_sign[rotation=10,waterlogged=false]",
		nil,
		NewMapping{
			404: 3097,
			477: 3400,
			573: 3400,
			718: 3402,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13581,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=1,south=up,west=up]",
		nil,
		NewMapping{
			573: 2641,
			718: 2643,
			393: 2337,
			404: 2338,
			477: 2641,
		},
	},
	{
		"minecraft:rail[shape=north_west]",
		nil,
		NewMapping{
			4:   1064,
			393: 3187,
			404: 3188,
			477: 3651,
			573: 3651,
			718: 3653,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=tall,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12695,
		},
	},
	{
		"minecraft:crimson_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 15379,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16829,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 7825,
			393: 6782,
			404: 6783,
			477: 7289,
			573: 7289,
		},
	},
	{
		"minecraft:fire[age=4,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			404: 1268,
			477: 1571,
			573: 1571,
			718: 1572,
			393: 1267,
		},
	},
	{
		"minecraft:warped_trapdoor[facing=east,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 15251,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14339,
		},
	},
	{
		"minecraft:light_gray_bed[facing=south,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 883,
			404: 883,
			477: 1183,
			573: 1183,
			718: 1184,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			477: 4519,
			573: 4519,
			718: 4533,
			393: 4015,
			404: 4016,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=low,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 5720,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10251,
			573: 10251,
			718: 10787,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9421,
			573: 9421,
			718: 9957,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16040,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 16284,
		},
	},
	{
		"minecraft:lava[level=1]",
		nil,
		NewMapping{
			393: 51,
			404: 51,
			477: 51,
			573: 51,
			718: 51,
			4:   177,
		},
	},
	{
		"minecraft:vine[east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 4285,
			404: 4286,
			477: 4789,
			573: 4789,
			718: 4805,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=5,south=side,west=none]",
		nil,
		NewMapping{
			477: 2394,
			573: 2394,
			718: 2396,
			393: 2090,
			404: 2091,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6196,
			404: 6197,
			477: 6703,
			573: 6703,
			718: 7239,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13420,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=none,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13061,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11960,
		},
	},
	{
		"minecraft:stonecutter[facing=north]",
		nil,
		NewMapping{
			477: 11194,
			573: 11194,
			718: 14850,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10940,
			573: 10940,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9640,
			573: 9640,
			718: 10176,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=6,waterlogged=true]",
		nil,
		NewMapping{
			477: 3551,
			573: 3551,
			718: 3553,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9163,
			573: 9163,
			718: 9699,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12471,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=east,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3713,
			404: 3714,
			477: 4217,
			573: 4217,
			718: 4231,
		},
	},
	{
		"minecraft:gray_banner[rotation=12]",
		nil,
		NewMapping{
			393: 6978,
			404: 6979,
			477: 7485,
			573: 7485,
			718: 8021,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 7070,
			718: 7606,
			393: 6563,
			404: 6564,
			477: 7070,
		},
	},
	{
		"minecraft:cyan_bed[facing=west,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 903,
			404: 903,
			477: 1203,
			573: 1203,
			718: 1204,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5147,
			404: 5148,
			477: 5651,
			573: 5651,
		},
	},
	{
		"minecraft:fire[age=4,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1264,
			404: 1265,
			477: 1568,
			573: 1568,
			718: 1569,
		},
	},
	{
		"minecraft:jungle_fence[east=false,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7609,
			404: 7610,
			477: 8134,
			573: 8134,
			718: 8670,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=20,powered=false]",
		nil,
		NewMapping{
			573: 839,
			718: 840,
			477: 839,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=9,powered=false]",
		nil,
		NewMapping{
			477: 967,
			573: 967,
			718: 968,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=15,south=none,west=up]",
		nil,
		NewMapping{
			718: 2919,
			393: 2613,
			404: 2614,
			477: 2917,
			573: 2917,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4680,
			404: 4681,
			477: 5184,
			573: 5184,
			718: 5200,
		},
	},
	{
		"minecraft:oak_door[facing=north,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 3117,
			404: 3118,
			477: 3581,
			573: 3581,
			718: 3583,
		},
	},
	{
		"minecraft:sign[rotation=9,waterlogged=true]",
		nil,
		NewMapping{
			393: 3093,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4533,
			404: 4534,
			477: 5037,
			573: 5037,
			718: 5053,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10639,
			573: 10639,
		},
	},
	{
		"minecraft:crimson_fence_gate[facing=north,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			718: 15260,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=tall,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11189,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=tall,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14630,
		},
	},
	{
		"minecraft:warped_trapdoor[facing=east,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			718: 15252,
		},
	},
	{
		"minecraft:dispenser[facing=down,triggered=true]",
		nil,
		NewMapping{
			718: 244,
			4:   376,
			393: 243,
			404: 243,
			477: 243,
			573: 243,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4551,
			404: 4552,
			477: 5055,
			573: 5055,
			718: 5071,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10833,
			573: 10833,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5952,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 7200,
			404: 7201,
			477: 7707,
			573: 7707,
			718: 8243,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10514,
			573: 10514,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=tall,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11509,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10448,
			573: 10448,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15914,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=none,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11538,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=west,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			718: 15174,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=15,south=up,west=up]",
		nil,
		NewMapping{
			718: 2337,
			393: 2031,
			404: 2032,
			477: 2335,
			573: 2335,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 7215,
			573: 7215,
			718: 7751,
			393: 6708,
			404: 6709,
		},
	},
	{
		"minecraft:quartz_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			573: 7856,
			718: 8392,
			393: 7337,
			404: 7338,
			477: 7856,
		},
	},
	{
		"minecraft:bell[attachment=ceiling,facing=north]",
		nil,
		NewMapping{
			477: 11202,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 9575,
			718: 10111,
			477: 9575,
		},
	},
	{
		"minecraft:polished_blackstone_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			718: 16746,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15878,
		},
	},
	{
		"minecraft:sugar_cane[age=8]",
		nil,
		NewMapping{
			404: 3451,
			477: 3954,
			573: 3954,
			718: 3956,
			4:   1336,
			393: 3450,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=5,south=up,west=side]",
		nil,
		NewMapping{
			393: 2950,
			404: 2951,
			477: 3254,
			573: 3254,
			718: 3256,
		},
	},
	{
		"minecraft:campfire[facing=west,lit=false,signal_fire=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 11237,
			573: 11253,
			718: 14911,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 9633,
			718: 10169,
			477: 9633,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=tall,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 6195,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			718: 7005,
			393: 5962,
			404: 5963,
			477: 6469,
			573: 6469,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			718: 4616,
			393: 4098,
			404: 4099,
			477: 4602,
			573: 4602,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=11,south=up,west=up]",
		nil,
		NewMapping{
			393: 2427,
			404: 2428,
			477: 2731,
			573: 2731,
			718: 2733,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 6285,
		},
	},
	{
		"minecraft:spruce_door[facing=east,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			404: 7735,
			477: 8259,
			573: 8259,
			718: 8795,
			393: 7734,
		},
	},
	{
		"minecraft:red_banner[rotation=15]",
		nil,
		NewMapping{
			404: 7094,
			477: 7600,
			573: 7600,
			718: 8136,
			393: 7093,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10679,
			477: 10679,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=low,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16200,
		},
	},
	{
		"minecraft:diorite_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10178,
			573: 10178,
			718: 10714,
		},
	},
	{
		"minecraft:warped_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15426,
		},
	},
	{
		"minecraft:tripwire_hook[attached=true,facing=north,powered=false]",
		nil,
		NewMapping{
			393: 4740,
			404: 4741,
			477: 5244,
			573: 5244,
			718: 5260,
			4:   2102,
		},
	},
	{
		"minecraft:dark_oak_button[face=floor,facing=west,powered=true]",
		nil,
		NewMapping{
			404: 5428,
			477: 5934,
			573: 5934,
			718: 6470,
			393: 5427,
		},
	},
	{
		"minecraft:nether_brick_fence[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			718: 5047,
			4:   1808,
			393: 4527,
			404: 4528,
			477: 5031,
			573: 5031,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10688,
			573: 10688,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			404: 4043,
			477: 4546,
			573: 4546,
			718: 4560,
			393: 4042,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 7198,
			404: 7199,
			477: 7705,
			573: 7705,
			718: 8241,
			4:   2886,
		},
	},
	{
		"minecraft:oak_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 1677,
			404: 1678,
			477: 1981,
			573: 1981,
			718: 1983,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6140,
			404: 6141,
			477: 6647,
			573: 6647,
			718: 7183,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=tall,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14305,
		},
	},
	{
		"minecraft:dead_brain_coral_wall_fan[facing=west,waterlogged=false]",
		nil,
		NewMapping{
			393: 8477,
			404: 8493,
			477: 9037,
			573: 9037,
			718: 9573,
		},
	},
	{
		"minecraft:purpur_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 8136,
			477: 8660,
			573: 8660,
			718: 9196,
			393: 8135,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=4,south=up,west=none]",
		nil,
		NewMapping{
			393: 2078,
			404: 2079,
			477: 2382,
			573: 2382,
			718: 2384,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13282,
		},
	},
	{
		"minecraft:warped_door[facing=east,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			718: 15643,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			404: 5236,
			477: 5739,
			573: 5739,
			393: 5235,
		},
	},
	{
		"minecraft:crimson_button[face=floor,facing=west,powered=true]",
		nil,
		NewMapping{
			718: 15483,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=low,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12861,
		},
	},
	{
		"minecraft:stripped_warped_hyphae[axis=y]",
		nil,
		NewMapping{
			718: 14968,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=none,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 6206,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=1,south=none,west=side]",
		nil,
		NewMapping{
			477: 2216,
			573: 2216,
			718: 2218,
			393: 1912,
			404: 1913,
		},
	},
	{
		"minecraft:zombie_head[rotation=2]",
		nil,
		NewMapping{
			393: 5493,
			404: 5494,
			477: 5996,
			573: 5996,
			718: 6532,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6215,
			404: 6216,
			477: 6722,
			573: 6722,
			718: 7258,
		},
	},
	{
		"minecraft:oak_sign[rotation=7,waterlogged=false]",
		nil,
		NewMapping{
			718: 3396,
			404: 3091,
			477: 3394,
			573: 3394,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=20,powered=true]",
		nil,
		NewMapping{
			718: 439,
			393: 438,
			404: 438,
			477: 438,
			573: 438,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 5842,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16865,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=11,south=none,west=side]",
		nil,
		NewMapping{
			393: 2002,
			404: 2003,
			477: 2306,
			573: 2306,
			718: 2308,
		},
	},
	{
		"minecraft:pumpkin_stem[age=5]",
		nil,
		NewMapping{
			393: 4257,
			404: 4258,
			477: 4761,
			573: 4761,
			718: 4777,
			4:   1669,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=10,south=none,west=none]",
		nil,
		NewMapping{
			393: 2282,
			404: 2283,
			477: 2586,
			573: 2586,
			718: 2588,
		},
	},
	{
		"minecraft:light_gray_shulker_box[facing=down]",
		nil,
		NewMapping{
			477: 8795,
			573: 8795,
			718: 9331,
			4:   3632,
			393: 8270,
			404: 8271,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=low,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13397,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=tall,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16237,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=west,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 4279,
			393: 3761,
			404: 3762,
			477: 4265,
			573: 4265,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9248,
			573: 9248,
			718: 9784,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10710,
			573: 10710,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=none,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16170,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=south,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			404: 3618,
			477: 4121,
			573: 4121,
			718: 4135,
			393: 3617,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=1,south=none,west=side]",
		nil,
		NewMapping{
			573: 3080,
			718: 3082,
			393: 2776,
			404: 2777,
			477: 3080,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11079,
			573: 11079,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12162,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=none,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5687,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=low,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14255,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14068,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=south,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3877,
			404: 3878,
			477: 4381,
			573: 4381,
			718: 4395,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 8297,
			393: 7254,
			404: 7255,
			477: 7761,
			573: 7761,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 5199,
			718: 5215,
			393: 4695,
			404: 4696,
			477: 5199,
		},
	},
	{
		"minecraft:birch_sign[rotation=8,waterlogged=true]",
		nil,
		NewMapping{
			477: 3459,
			573: 3459,
			718: 3461,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10780,
			573: 10780,
		},
	},
	{
		"minecraft:warped_fence[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			718: 15097,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=none,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14441,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 6658,
			573: 6658,
			718: 7194,
			393: 6151,
			404: 6152,
		},
	},
	{
		"minecraft:bee_nest[facing=west,honey_level=3]",
		nil,
		NewMapping{
			573: 11302,
			718: 15791,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=low,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13180,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=none,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13684,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 9172,
			718: 9708,
			477: 9172,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5886,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5999,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16173,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=23,powered=true]",
		nil,
		NewMapping{
			404: 444,
			477: 444,
			573: 444,
			718: 445,
			393: 444,
		},
	},
	{
		"minecraft:yellow_wall_banner[facing=south]",
		nil,
		NewMapping{
			393: 7127,
			404: 7128,
			477: 7634,
			573: 7634,
			718: 8170,
		},
	},
	{
		"minecraft:acacia_sign[rotation=6,waterlogged=true]",
		nil,
		NewMapping{
			477: 3487,
			573: 3487,
			718: 3489,
		},
	},
	{
		"minecraft:campfire[facing=north,lit=false,signal_fire=false,waterlogged=false]",
		nil,
		NewMapping{
			477: 11223,
			573: 11239,
			718: 14897,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10548,
			573: 10548,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5965,
		},
	},
	{
		"minecraft:warped_trapdoor[facing=east,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			718: 15241,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=low,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16494,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=5,south=none,west=side]",
		nil,
		NewMapping{
			573: 2972,
			718: 2974,
			393: 2668,
			404: 2669,
			477: 2972,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=8,south=side,west=none]",
		nil,
		NewMapping{
			718: 2567,
			393: 2261,
			404: 2262,
			477: 2565,
			573: 2565,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 4939,
			477: 5442,
			573: 5442,
			718: 5458,
			393: 4938,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 4177,
			404: 4178,
			477: 4681,
			573: 4681,
			718: 4695,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=low,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 17045,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16507,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=west,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 4143,
			393: 3625,
			404: 3626,
			477: 4129,
			573: 4129,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=none,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14029,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=none,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11839,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15844,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 8246,
			393: 7203,
			404: 7204,
			477: 7710,
			573: 7710,
		},
	},
	{
		"minecraft:birch_door[facing=east,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			404: 7795,
			477: 8319,
			573: 8319,
			718: 8855,
			393: 7794,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=0,south=up,west=none]",
		nil,
		NewMapping{
			573: 2634,
			718: 2636,
			393: 2330,
			404: 2331,
			477: 2634,
		},
	},
	{
		"minecraft:kelp[age=9]",
		nil,
		NewMapping{
			393: 8418,
			404: 8419,
			477: 8943,
			573: 8943,
			718: 9479,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12473,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16072,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16107,
		},
	},
	{
		"minecraft:repeater[delay=3,facing=south,locked=true,powered=true]",
		nil,
		NewMapping{
			477: 4053,
			573: 4053,
			718: 4067,
			393: 3549,
			404: 3550,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10850,
			573: 10850,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=11,powered=true]",
		nil,
		NewMapping{
			718: 821,
			477: 820,
			573: 820,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=none,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11325,
		},
	},
	{
		"minecraft:lava[level=5]",
		nil,
		NewMapping{
			393: 55,
			404: 55,
			477: 55,
			573: 55,
			718: 55,
			4:   165,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10470,
			573: 10470,
		},
	},
	{
		"minecraft:polished_blackstone_button[face=wall,facing=west,powered=true]",
		nil,
		NewMapping{
			718: 16765,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=none,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13481,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=none,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13037,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=10,south=up,west=side]",
		nil,
		NewMapping{
			393: 2419,
			404: 2420,
			477: 2723,
			573: 2723,
			718: 2725,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4459,
			404: 4460,
			477: 4963,
			573: 4963,
			718: 4979,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=24,powered=false]",
		nil,
		NewMapping{
			573: 397,
			718: 398,
			393: 397,
			404: 397,
			477: 397,
		},
	},
	{
		"minecraft:acacia_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2612,
			393: 6393,
			404: 6394,
			477: 6900,
			573: 6900,
			718: 7436,
		},
	},
	{
		"minecraft:moving_piston[facing=north,type=normal]",
		nil,
		NewMapping{
			718: 1400,
			4:   578,
			393: 1099,
			404: 1099,
			477: 1399,
			573: 1399,
		},
	},
	{
		"minecraft:dark_oak_door[facing=south,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			404: 7950,
			477: 8474,
			573: 8474,
			718: 9010,
			393: 7949,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=3,powered=false]",
		nil,
		NewMapping{
			393: 655,
			404: 655,
			477: 655,
			573: 655,
			718: 656,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10681,
			573: 10681,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=11,south=up,west=up]",
		nil,
		NewMapping{
			404: 2716,
			477: 3019,
			573: 3019,
			718: 3021,
			393: 2715,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 7256,
			404: 7257,
			477: 7763,
			573: 7763,
			718: 8299,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=south,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			573: 4370,
			718: 4384,
			393: 3866,
			404: 3867,
			477: 4370,
		},
	},
	{
		"minecraft:stripped_acacia_log[axis=z]",
		nil,
		NewMapping{
			393: 101,
			404: 101,
			477: 101,
			573: 101,
			718: 102,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=tall,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13778,
		},
	},
	{
		"minecraft:daylight_detector[inverted=false,power=5]",
		nil,
		NewMapping{
			393: 5672,
			404: 5673,
			477: 6179,
			573: 6179,
			718: 6715,
			4:   2421,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=10,south=side,west=none]",
		nil,
		NewMapping{
			393: 2999,
			404: 3000,
			477: 3303,
			573: 3303,
			718: 3305,
		},
	},
	{
		"minecraft:honeycomb_block",
		nil,
		NewMapping{
			573: 11336,
			718: 15825,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=low,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12660,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=7,south=side,west=up]",
		nil,
		NewMapping{
			393: 2682,
			404: 2683,
			477: 2986,
			573: 2986,
			718: 2988,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9548,
			573: 9548,
			718: 10084,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=none,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 10979,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=none,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11110,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14393,
		},
	},
	{
		"minecraft:blue_ice",
		nil,
		NewMapping{
			393: 8572,
			404: 8588,
			477: 9112,
			573: 9112,
			718: 9648,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 5191,
			477: 5694,
			573: 5694,
			393: 5190,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=9,south=none,west=side]",
		nil,
		NewMapping{
			718: 3010,
			393: 2704,
			404: 2705,
			477: 3008,
			573: 3008,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 10126,
			718: 10662,
			477: 10126,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=11,powered=false]",
		nil,
		NewMapping{
			393: 571,
			404: 571,
			477: 571,
			573: 571,
			718: 572,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6090,
			404: 6091,
			477: 6597,
			573: 6597,
			718: 7133,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=5,powered=false]",
		nil,
		NewMapping{
			573: 609,
			718: 610,
			393: 609,
			404: 609,
			477: 609,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=13,south=up,west=up]",
		nil,
		NewMapping{
			393: 2301,
			404: 2302,
			477: 2605,
			573: 2605,
			718: 2607,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 10163,
			573: 10163,
			718: 10699,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=low,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12219,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=7,south=none,west=up]",
		nil,
		NewMapping{
			718: 2127,
			393: 1821,
			404: 1822,
			477: 2125,
			573: 2125,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 6366,
			477: 6872,
			573: 6872,
			718: 7408,
			393: 6365,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=none,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12950,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13672,
		},
	},
	{
		"minecraft:brown_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 949,
			404: 949,
			477: 1249,
			573: 1249,
			718: 1250,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=none,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14466,
		},
	},
	{
		"minecraft:activator_rail[powered=true,shape=north_south]",
		nil,
		NewMapping{
			393: 5780,
			404: 5781,
			477: 6287,
			573: 6287,
			718: 6823,
			4:   2520,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13240,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12594,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=north,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3794,
			404: 3795,
			477: 4298,
			573: 4298,
			718: 4312,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=north,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3796,
			404: 3797,
			477: 4300,
			573: 4300,
			718: 4314,
		},
	},
	{
		"minecraft:kelp[age=22]",
		nil,
		NewMapping{
			393: 8431,
			404: 8432,
			477: 8956,
			573: 8956,
			718: 9492,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10784,
			573: 10784,
		},
	},
	{
		"minecraft:gray_shulker_box[facing=south]",
		nil,
		NewMapping{
			718: 9322,
			4:   3619,
			393: 8261,
			404: 8262,
			477: 8786,
			573: 8786,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10603,
			573: 10603,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=tall,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14297,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5859,
			404: 5860,
			477: 6366,
			573: 6366,
			718: 6902,
		},
	},
	{
		"minecraft:oak_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 1698,
			404: 1699,
			477: 2002,
			573: 2002,
			718: 2004,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=none,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16783,
		},
	},
	{
		"minecraft:warped_door[facing=west,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			718: 15628,
		},
	},
	{
		"minecraft:lime_concrete_powder",
		nil,
		NewMapping{
			393: 8398,
			404: 8399,
			477: 8923,
			573: 8923,
			718: 9459,
			4:   4037,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=2,south=none,west=up]",
		nil,
		NewMapping{
			393: 2928,
			404: 2929,
			477: 3232,
			573: 3232,
			718: 3234,
		},
	},
	{
		"minecraft:warped_door[facing=south,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			718: 15607,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13297,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=false,north=true,powered=true,south=false,west=false]",
		nil,
		NewMapping{
			393: 4870,
			404: 4871,
			477: 5374,
			573: 5374,
			718: 5390,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9853,
			573: 9853,
			718: 10389,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11130,
		},
	},
	{
		"minecraft:warped_trapdoor[facing=west,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 15235,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=none,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12516,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=15,south=up,west=none]",
		nil,
		NewMapping{
			718: 2627,
			393: 2321,
			404: 2322,
			477: 2625,
			573: 2625,
		},
	},
	{
		"minecraft:birch_leaves[distance=6,persistent=true]",
		nil,
		NewMapping{
			393: 182,
			404: 182,
			477: 182,
			573: 182,
			718: 183,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11235,
		},
	},
	{
		"minecraft:birch_door[facing=east,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			718: 8857,
			4:   3113,
			393: 7796,
			404: 7797,
			477: 8321,
			573: 8321,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			573: 8579,
			718: 9115,
			393: 8054,
			404: 8055,
			477: 8579,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9864,
			573: 9864,
			718: 10400,
		},
	},
	{
		"minecraft:crimson_door[facing=north,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15537,
		},
	},
	{
		"minecraft:zombie_head[rotation=3]",
		nil,
		NewMapping{
			393: 5494,
			404: 5495,
			477: 5997,
			573: 5997,
			718: 6533,
		},
	},
	{
		"minecraft:brick_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 4837,
			573: 4837,
			718: 4853,
			4:   1735,
			393: 4333,
			404: 4334,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=tall,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13349,
		},
	},
	{
		"minecraft:red_banner[rotation=7]",
		nil,
		NewMapping{
			393: 7085,
			404: 7086,
			477: 7592,
			573: 7592,
			718: 8128,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=tall,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13550,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11962,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 16732,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=none,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11523,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=12,south=none,west=side]",
		nil,
		NewMapping{
			393: 2587,
			404: 2588,
			477: 2891,
			573: 2891,
			718: 2893,
		},
	},
	{
		"minecraft:spruce_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4922,
			404: 4923,
			477: 5426,
			573: 5426,
			718: 5442,
		},
	},
	{
		"minecraft:andesite_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9970,
			573: 9970,
			718: 10506,
		},
	},
	{
		"minecraft:crimson_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15397,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12150,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16823,
		},
	},
	{
		"minecraft:snow[layers=3]",
		nil,
		NewMapping{
			404: 3418,
			477: 3921,
			573: 3921,
			718: 3923,
			4:   1250,
			393: 3417,
		},
	},
	{
		"minecraft:birch_door[facing=north,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			718: 8809,
			393: 7748,
			404: 7749,
			477: 8273,
			573: 8273,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10752,
			477: 10752,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=none,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5665,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=3,powered=false]",
		nil,
		NewMapping{
			477: 855,
			573: 855,
			718: 856,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13476,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=low,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16501,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16832,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=tall,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16001,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 3667,
			573: 3667,
			718: 3669,
			393: 3203,
			404: 3204,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=north,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 6501,
			404: 6502,
			477: 7008,
			573: 7008,
			718: 7544,
			4:   2680,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 11056,
			477: 11056,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14644,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=20,powered=false]",
		nil,
		NewMapping{
			477: 889,
			573: 889,
			718: 890,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=none,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13372,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=tall,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12264,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=tall,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12804,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 3989,
			404: 3990,
			477: 4493,
			573: 4493,
			718: 4507,
		},
	},
	{
		"minecraft:andesite_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9949,
			573: 9949,
			718: 10485,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9452,
			573: 9452,
			718: 9988,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11083,
			573: 11083,
		},
	},
	{
		"minecraft:daylight_detector[inverted=false,power=7]",
		nil,
		NewMapping{
			393: 5674,
			404: 5675,
			477: 6181,
			573: 6181,
			718: 6717,
			4:   2423,
		},
	},
	{
		"minecraft:light_weighted_pressure_plate[power=7]",
		nil,
		NewMapping{
			393: 5610,
			404: 5611,
			477: 6117,
			573: 6117,
			718: 6653,
			4:   2359,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6569,
			404: 6570,
			477: 7076,
			573: 7076,
			718: 7612,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12393,
		},
	},
	{
		"minecraft:repeater[delay=3,facing=north,locked=true,powered=true]",
		nil,
		NewMapping{
			477: 4049,
			573: 4049,
			718: 4063,
			393: 3545,
			404: 3546,
		},
	},
	{
		"minecraft:jungle_fence[east=false,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7611,
			404: 7612,
			477: 8136,
			573: 8136,
			718: 8672,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=7,south=up,west=side]",
		nil,
		NewMapping{
			393: 2392,
			404: 2393,
			477: 2696,
			573: 2696,
			718: 2698,
		},
	},
	{
		"minecraft:magenta_bed[facing=west,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 791,
			404: 791,
			477: 1091,
			573: 1091,
			718: 1092,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=low,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16199,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=19,powered=false]",
		nil,
		NewMapping{
			393: 387,
			404: 387,
			477: 387,
			573: 387,
			718: 388,
		},
	},
	{
		"minecraft:pink_bed[facing=south,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 851,
			404: 851,
			477: 1151,
			573: 1151,
			718: 1152,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16231,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5953,
		},
	},
	{
		"minecraft:twisting_vines[age=0]",
		nil,
		NewMapping{
			718: 15017,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 6158,
		},
	},
	{
		"minecraft:polished_blackstone_brick_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			718: 16254,
		},
	},
	{
		"minecraft:spruce_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			404: 7268,
			477: 7774,
			573: 7774,
			718: 8310,
			393: 7267,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=7,south=none,west=side]",
		nil,
		NewMapping{
			573: 2558,
			718: 2560,
			393: 2254,
			404: 2255,
			477: 2558,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=true,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 4057,
			404: 4058,
			477: 4561,
			573: 4561,
			718: 4575,
		},
	},
	{
		"minecraft:crimson_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15393,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=east,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7448,
			404: 7449,
			477: 7973,
			573: 7973,
			718: 8509,
		},
	},
	{
		"minecraft:spruce_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4918,
			404: 4919,
			477: 5422,
			573: 5422,
			718: 5438,
		},
	},
	{
		"minecraft:birch_fence[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 8094,
			718: 8630,
			393: 7569,
			404: 7570,
			477: 8094,
		},
	},
	{
		"minecraft:dark_oak_fence[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 8188,
			718: 8724,
			393: 7663,
			404: 7664,
			477: 8188,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=1,powered=false]",
		nil,
		NewMapping{
			573: 951,
			718: 952,
			477: 951,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 16722,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=tall,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16020,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=tall,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11831,
		},
	},
	{
		"minecraft:crimson_wall_sign[facing=east,waterlogged=false]",
		nil,
		NewMapping{
			718: 15726,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6764,
			718: 7300,
			393: 6257,
			404: 6258,
			477: 6764,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=3,south=none,west=none]",
		nil,
		NewMapping{
			477: 2955,
			573: 2955,
			718: 2957,
			393: 2651,
			404: 2652,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=5,powered=false]",
		nil,
		NewMapping{
			393: 259,
			404: 259,
			477: 259,
			573: 259,
			718: 260,
		},
	},
	{
		"minecraft:diorite_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 10192,
			718: 10728,
			477: 10192,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=low,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13515,
		},
	},
	{
		"minecraft:fire[age=13,east=true,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			477: 1868,
			573: 1868,
			718: 1869,
			393: 1564,
			404: 1565,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=14,south=up,west=up]",
		nil,
		NewMapping{
			393: 1878,
			404: 1879,
			477: 2182,
			573: 2182,
			718: 2184,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=4,powered=false]",
		nil,
		NewMapping{
			477: 907,
			573: 907,
			718: 908,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14104,
		},
	},
	{
		"minecraft:jungle_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 5095,
			404: 5096,
			477: 5599,
			573: 5599,
			718: 5615,
			4:   2177,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=low,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13635,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16518,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=tall,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11812,
		},
	},
	{
		"minecraft:green_bed[facing=south,occupied=true,part=foot]",
		nil,
		NewMapping{
			477: 1261,
			573: 1261,
			718: 1262,
			393: 961,
			404: 961,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=north,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3665,
			404: 3666,
			477: 4169,
			573: 4169,
			718: 4183,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10090,
			573: 10090,
			718: 10626,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11034,
		},
	},
	{
		"minecraft:fire[age=8,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			718: 1708,
			393: 1403,
			404: 1404,
			477: 1707,
			573: 1707,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13132,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14308,
		},
	},
	{
		"minecraft:fire[age=14,east=true,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			477: 1892,
			573: 1892,
			718: 1893,
			393: 1588,
			404: 1589,
		},
	},
	{
		"minecraft:crimson_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 15320,
		},
	},
	{
		"minecraft:daylight_detector[inverted=true,power=14]",
		nil,
		NewMapping{
			573: 6172,
			718: 6708,
			4:   2862,
			393: 5665,
			404: 5666,
			477: 6172,
		},
	},
	{
		"minecraft:red_sandstone_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			393: 7346,
			404: 7347,
			477: 7865,
			573: 7865,
			718: 8401,
			4:   2896,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=none,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5903,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=tall,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11065,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=2,south=side,west=up]",
		nil,
		NewMapping{
			718: 2367,
			393: 2061,
			404: 2062,
			477: 2365,
			573: 2365,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=none,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12926,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=north,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 15131,
		},
	},
	{
		"minecraft:cactus[age=11]",
		nil,
		NewMapping{
			393: 3436,
			404: 3437,
			477: 3940,
			573: 3940,
			718: 3942,
			4:   1307,
		},
	},
	{
		"minecraft:light_gray_banner[rotation=3]",
		nil,
		NewMapping{
			393: 6985,
			404: 6986,
			477: 7492,
			573: 7492,
			718: 8028,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11100,
		},
	},
	{
		"minecraft:carved_pumpkin[facing=south]",
		nil,
		NewMapping{
			393: 3499,
			404: 3500,
			477: 4003,
			573: 4003,
			718: 4017,
			4:   1376,
		},
	},
	{
		"minecraft:sea_pickle[pickles=3,waterlogged=true]",
		nil,
		NewMapping{
			393: 8568,
			404: 8584,
			477: 9108,
			573: 9108,
			718: 9644,
		},
	},
	{
		"minecraft:nether_brick_fence[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 4520,
			404: 4521,
			477: 5024,
			573: 5024,
			718: 5040,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			718: 9100,
			393: 8039,
			404: 8040,
			477: 8564,
			573: 8564,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 3255,
			477: 3718,
			573: 3718,
			718: 3720,
			393: 3254,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=west,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3629,
			404: 3630,
			477: 4133,
			573: 4133,
			718: 4147,
		},
	},
	{
		"minecraft:birch_sign[rotation=2,waterlogged=false]",
		nil,
		NewMapping{
			573: 3448,
			718: 3450,
			477: 3448,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 10335,
			477: 9799,
			573: 9799,
		},
	},
	{
		"minecraft:fire[age=0,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1142,
			404: 1143,
			477: 1446,
			573: 1446,
			718: 1447,
		},
	},
	{
		"minecraft:brick_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 4898,
			718: 4914,
			393: 4394,
			404: 4395,
			477: 4898,
		},
	},
	{
		"minecraft:cyan_shulker_box[facing=north]",
		nil,
		NewMapping{
			718: 9332,
			4:   3650,
			393: 8271,
			404: 8272,
			477: 8796,
			573: 8796,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 5743,
			573: 5743,
			393: 5239,
			404: 5240,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=11,south=side,west=up]",
		nil,
		NewMapping{
			393: 2862,
			404: 2863,
			477: 3166,
			573: 3166,
			718: 3168,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=13,south=none,west=none]",
		nil,
		NewMapping{
			393: 2309,
			404: 2310,
			477: 2613,
			573: 2613,
			718: 2615,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13547,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13827,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=none,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12190,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 3698,
			573: 3698,
			718: 3700,
			393: 3234,
			404: 3235,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15883,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13714,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=low,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5913,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=low,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12651,
		},
	},
	{
		"minecraft:jungle_fence[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 8120,
			573: 8120,
			718: 8656,
			393: 7595,
			404: 7596,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=north,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 4302,
			573: 4302,
			718: 4316,
			393: 3798,
			404: 3799,
		},
	},
	{
		"minecraft:purple_wool",
		nil,
		NewMapping{
			718: 1394,
			4:   570,
			393: 1093,
			404: 1093,
			477: 1393,
			573: 1393,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=13,powered=false]",
		nil,
		NewMapping{
			718: 976,
			477: 975,
			573: 975,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=none,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13894,
		},
	},
	{
		"minecraft:magenta_glazed_terracotta[facing=south]",
		nil,
		NewMapping{
			4:   3792,
			393: 8322,
			404: 8323,
			477: 8847,
			573: 8847,
			718: 9383,
		},
	},
	{
		"minecraft:fire[age=11,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			404: 1495,
			477: 1798,
			573: 1798,
			718: 1799,
			393: 1494,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=north,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7453,
			404: 7454,
			477: 7978,
			573: 7978,
			718: 8514,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 9432,
			718: 9968,
			477: 9432,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5976,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12787,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=low,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13410,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=west,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3950,
			404: 3951,
			477: 4454,
			573: 4454,
			718: 4468,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=9,south=up,west=side]",
		nil,
		NewMapping{
			393: 2266,
			404: 2267,
			477: 2570,
			573: 2570,
			718: 2572,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 7493,
			393: 6450,
			404: 6451,
			477: 6957,
			573: 6957,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=none,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16058,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12871,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=15,south=none,west=side]",
		nil,
		NewMapping{
			573: 2918,
			718: 2920,
			393: 2614,
			404: 2615,
			477: 2918,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 6710,
			718: 7246,
			4:   2571,
			393: 6203,
			404: 6204,
			477: 6710,
		},
	},
	{
		"minecraft:lectern[facing=east,has_book=false,powered=false]",
		nil,
		NewMapping{
			477: 11192,
			573: 11192,
			718: 14848,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 15889,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=tall,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13222,
		},
	},
	{
		"minecraft:dark_oak_button[face=ceiling,facing=south,powered=true]",
		nil,
		NewMapping{
			477: 5948,
			573: 5948,
			718: 6484,
			393: 5441,
			404: 5442,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=13,south=up,west=up]",
		nil,
		NewMapping{
			718: 3183,
			393: 2877,
			404: 2878,
			477: 3181,
			573: 3181,
		},
	},
	{
		"minecraft:dead_brain_coral_wall_fan[facing=south,waterlogged=true]",
		nil,
		NewMapping{
			404: 8490,
			477: 9034,
			573: 9034,
			718: 9570,
			393: 8474,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=14,powered=false]",
		nil,
		NewMapping{
			393: 377,
			404: 377,
			477: 377,
			573: 377,
			718: 378,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11054,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=low,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 5714,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 6622,
			404: 6623,
			477: 7129,
			573: 7129,
			718: 7665,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10881,
			573: 10881,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=none,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12835,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 5998,
		},
	},
	{
		"minecraft:repeater[delay=2,facing=east,locked=false,powered=true]",
		nil,
		NewMapping{
			393: 3543,
			404: 3544,
			477: 4047,
			573: 4047,
			718: 4061,
			4:   1511,
		},
	},
	{
		"minecraft:dead_fire_coral_wall_fan[facing=south,waterlogged=false]",
		nil,
		NewMapping{
			477: 9051,
			573: 9051,
			718: 9587,
			393: 8491,
			404: 8507,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11091,
			573: 11091,
		},
	},
	{
		"minecraft:bee_nest[facing=west,honey_level=0]",
		nil,
		NewMapping{
			573: 11299,
			718: 15788,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=5,south=up,west=side]",
		nil,
		NewMapping{
			393: 2086,
			404: 2087,
			477: 2390,
			573: 2390,
			718: 2392,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9168,
			573: 9168,
			718: 9704,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12505,
		},
	},
	{
		"minecraft:warped_sign[rotation=3,waterlogged=true]",
		nil,
		NewMapping{
			718: 15693,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=tall,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13231,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 15886,
		},
	},
	{
		"minecraft:beacon",
		nil,
		NewMapping{
			393: 5136,
			404: 5137,
			477: 5640,
			573: 5640,
			718: 5656,
			4:   2208,
		},
	},
	{
		"minecraft:glowstone",
		nil,
		NewMapping{
			718: 4013,
			4:   1424,
			393: 3495,
			404: 3496,
			477: 3999,
			573: 3999,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=12,powered=false]",
		nil,
		NewMapping{
			573: 323,
			718: 324,
			393: 323,
			404: 323,
			477: 323,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=tall,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12696,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=low,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 6247,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14213,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=east,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3906,
			404: 3907,
			477: 4410,
			573: 4410,
			718: 4424,
		},
	},
	{
		"minecraft:birch_fence[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7559,
			404: 7560,
			477: 8084,
			573: 8084,
			718: 8620,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=18,powered=false]",
		nil,
		NewMapping{
			393: 585,
			404: 585,
			477: 585,
			573: 585,
			718: 586,
		},
	},
	{
		"minecraft:campfire[facing=north,lit=false,signal_fire=true,waterlogged=true]",
		nil,
		NewMapping{
			573: 11236,
			718: 14894,
			477: 11220,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=south,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			477: 7922,
			573: 7922,
			718: 8458,
			393: 7397,
			404: 7398,
		},
	},
	{
		"minecraft:seagrass",
		nil,
		NewMapping{
			393: 1044,
			404: 1044,
			477: 1344,
			573: 1344,
			718: 1345,
		},
	},
	{
		"minecraft:tall_seagrass[half=lower]",
		nil,
		NewMapping{
			573: 1346,
			718: 1347,
			393: 1046,
			404: 1046,
			477: 1346,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=none,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5663,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=15,south=none,west=up]",
		nil,
		NewMapping{
			573: 3205,
			718: 3207,
			393: 2901,
			404: 2902,
			477: 3205,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10608,
			573: 10608,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 16267,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12007,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9492,
			573: 9492,
			718: 10028,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11594,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=low,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11027,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13859,
		},
	},
	{
		"minecraft:oak_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 1705,
			404: 1706,
			477: 2009,
			573: 2009,
			718: 2011,
		},
	},
	{
		"minecraft:piston_head[facing=up,short=true,type=normal]",
		nil,
		NewMapping{
			393: 1075,
			404: 1075,
			477: 1375,
			573: 1375,
			718: 1376,
		},
	},
	{
		"minecraft:furnace[facing=west,lit=false]",
		nil,
		NewMapping{
			573: 3376,
			718: 3378,
			4:   980,
			393: 3072,
			404: 3073,
			477: 3376,
		},
	},
	{
		"minecraft:iron_door[facing=north,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			477: 3813,
			573: 3813,
			718: 3815,
			393: 3309,
			404: 3310,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=tall,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12367,
		},
	},
	{
		"minecraft:jungle_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 5083,
			404: 5084,
			477: 5587,
			573: 5587,
			718: 5603,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   1828,
			393: 4589,
			404: 4590,
			477: 5093,
			573: 5093,
			718: 5109,
		},
	},
	{
		"minecraft:fire[age=1,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1190,
			404: 1191,
			477: 1494,
			573: 1494,
			718: 1495,
		},
	},
	{
		"minecraft:cut_sandstone_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			477: 7823,
			573: 7823,
			718: 8359,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14446,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			718: 7097,
			393: 6054,
			404: 6055,
			477: 6561,
			573: 6561,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=11,south=none,west=side]",
		nil,
		NewMapping{
			573: 2738,
			718: 2740,
			393: 2434,
			404: 2435,
			477: 2738,
		},
	},
	{
		"minecraft:oak_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 2019,
			573: 2019,
			718: 2021,
			393: 1715,
			404: 1716,
		},
	},
	{
		"minecraft:spruce_door[facing=south,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7698,
			404: 7699,
			477: 8223,
			573: 8223,
			718: 8759,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=low,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11675,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=north,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			404: 6501,
			477: 7007,
			573: 7007,
			718: 7543,
			393: 6500,
		},
	},
	{
		"minecraft:oak_leaves[distance=3,persistent=false]",
		nil,
		NewMapping{
			404: 149,
			477: 149,
			573: 149,
			718: 150,
			393: 149,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10515,
			477: 10515,
		},
	},
	{
		"minecraft:scaffolding[bottom=false,distance=5,waterlogged=false]",
		nil,
		NewMapping{
			477: 11126,
			573: 11126,
			718: 14782,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6788,
			404: 6789,
			477: 7295,
			573: 7295,
			718: 7831,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13327,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=tall,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6071,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14185,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16414,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 7144,
			573: 7144,
			718: 7680,
			393: 6637,
			404: 6638,
		},
	},
	{
		"minecraft:glass_pane[east=true,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 4220,
			477: 4723,
			573: 4723,
			718: 4739,
			393: 4219,
		},
	},
	{
		"minecraft:jungle_door[facing=west,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			477: 8363,
			573: 8363,
			718: 8899,
			393: 7838,
			404: 7839,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=low,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12554,
		},
	},
	{
		"minecraft:trapped_chest[facing=west,type=left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5594,
			404: 5595,
			477: 6101,
			573: 6101,
			718: 6637,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=7,powered=false]",
		nil,
		NewMapping{
			573: 413,
			718: 414,
			393: 413,
			404: 413,
			477: 413,
		},
	},
	{
		"minecraft:potted_dead_bush",
		nil,
		NewMapping{
			573: 5792,
			718: 6328,
			393: 5285,
			404: 5286,
			477: 5792,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=none,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12499,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=21,powered=true]",
		nil,
		NewMapping{
			393: 390,
			404: 390,
			477: 390,
			573: 390,
			718: 391,
		},
	},
	{
		"minecraft:fire_coral_wall_fan[facing=east,waterlogged=false]",
		nil,
		NewMapping{
			393: 8535,
			404: 8551,
			477: 9095,
			573: 9095,
			718: 9631,
		},
	},
	{
		"minecraft:quartz_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 6282,
			573: 6282,
			718: 6818,
			393: 5775,
			404: 5776,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=south,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7466,
			404: 7467,
			477: 7991,
			573: 7991,
			718: 8527,
			4:   2996,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16520,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=low,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 10920,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=low,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12542,
		},
	},
	{
		"minecraft:lime_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			718: 1138,
			393: 837,
			404: 837,
			477: 1137,
			573: 1137,
		},
	},
	{
		"minecraft:blue_bed[facing=north,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 925,
			404: 925,
			477: 1225,
			573: 1225,
			718: 1226,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=tall,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12892,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12206,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=low,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11028,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=5,powered=true]",
		nil,
		NewMapping{
			393: 508,
			404: 508,
			477: 508,
			573: 508,
			718: 509,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9281,
			573: 9281,
			718: 9817,
		},
	},
	{
		"minecraft:warped_trapdoor[facing=west,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			718: 15238,
		},
	},
	{
		"minecraft:respawn_anchor[charges=1]",
		nil,
		NewMapping{
			718: 15830,
		},
	},
	{
		"minecraft:yellow_glazed_terracotta[facing=east]",
		nil,
		NewMapping{
			393: 8332,
			404: 8333,
			477: 8857,
			573: 8857,
			718: 9393,
			4:   3827,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=2,south=up,west=up]",
		nil,
		NewMapping{
			573: 3226,
			718: 3228,
			393: 2922,
			404: 2923,
			477: 3226,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=north,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3605,
			404: 3606,
			477: 4109,
			573: 4109,
			718: 4123,
		},
	},
	{
		"minecraft:jungle_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 5059,
			404: 5060,
			477: 5563,
			573: 5563,
			718: 5579,
		},
	},
	{
		"minecraft:crimson_button[face=wall,facing=west,powered=false]",
		nil,
		NewMapping{
			718: 15492,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=low,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11771,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13960,
		},
	},
	{
		"minecraft:warped_wall_sign[facing=south,waterlogged=true]",
		nil,
		NewMapping{
			718: 15729,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 7236,
			404: 7237,
			477: 7743,
			573: 7743,
			718: 8279,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10549,
			573: 10549,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10692,
			573: 10692,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11271,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=east,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			404: 3834,
			477: 4337,
			573: 4337,
			718: 4351,
			393: 3833,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 16339,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=low,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14256,
		},
	},
	{
		"minecraft:oak_sign[rotation=14,waterlogged=false]",
		nil,
		NewMapping{
			404: 3105,
			477: 3408,
			573: 3408,
			718: 3410,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=none,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13157,
		},
	},
	{
		"minecraft:warped_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 15411,
		},
	},
	{
		"minecraft:fire[age=8,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			573: 1721,
			718: 1722,
			393: 1417,
			404: 1418,
			477: 1721,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=4,powered=false]",
		nil,
		NewMapping{
			477: 707,
			573: 707,
			718: 708,
			393: 707,
			404: 707,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 5945,
			477: 6451,
			573: 6451,
			718: 6987,
			393: 5944,
		},
	},
	{
		"minecraft:acacia_door[facing=west,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7906,
			404: 7907,
			477: 8431,
			573: 8431,
			718: 8967,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=low,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13067,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13533,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=5,south=up,west=side]",
		nil,
		NewMapping{
			573: 3110,
			718: 3112,
			393: 2806,
			404: 2807,
			477: 3110,
		},
	},
	{
		"minecraft:blue_banner[rotation=0]",
		nil,
		NewMapping{
			718: 8073,
			393: 7030,
			404: 7031,
			477: 7537,
			573: 7537,
		},
	},
	{
		"minecraft:magenta_glazed_terracotta[facing=east]",
		nil,
		NewMapping{
			393: 8324,
			404: 8325,
			477: 8849,
			573: 8849,
			718: 9385,
			4:   3795,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=true,north=false,powered=false,south=false,west=true]",
		nil,
		NewMapping{
			393: 4769,
			404: 4770,
			477: 5273,
			573: 5273,
			718: 5289,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=south,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3805,
			404: 3806,
			477: 4309,
			573: 4309,
			718: 4323,
		},
	},
	{
		"minecraft:fire[age=9,east=false,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1448,
			404: 1449,
			477: 1752,
			573: 1752,
			718: 1753,
		},
	},
	{
		"minecraft:turtle_egg[eggs=4,hatch=0]",
		nil,
		NewMapping{
			404: 8447,
			477: 8971,
			573: 8971,
			718: 9507,
			393: 8446,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11692,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16208,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=low,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 10907,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=10,south=side,west=side]",
		nil,
		NewMapping{
			404: 2279,
			477: 2582,
			573: 2582,
			718: 2584,
			393: 2278,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=18,powered=true]",
		nil,
		NewMapping{
			393: 434,
			404: 434,
			477: 434,
			573: 434,
			718: 435,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=north,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			718: 4122,
			4:   1540,
			393: 3604,
			404: 3605,
			477: 4108,
			573: 4108,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10452,
			573: 10452,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=4,south=up,west=up]",
		nil,
		NewMapping{
			718: 2094,
			393: 1788,
			404: 1789,
			477: 2092,
			573: 2092,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9476,
			573: 9476,
			718: 10012,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14301,
		},
	},
	{
		"minecraft:chain[waterlogged=true]",
		nil,
		NewMapping{
			718: 4729,
		},
	},
	{
		"minecraft:warped_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			718: 15053,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 16265,
		},
	},
	{
		"minecraft:fire[age=11,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1488,
			404: 1489,
			477: 1792,
			573: 1792,
			718: 1793,
		},
	},
	{
		"minecraft:iron_bars[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 4705,
			718: 4719,
			393: 4201,
			404: 4202,
			477: 4705,
		},
	},
	{
		"minecraft:loom[facing=east]",
		nil,
		NewMapping{
			477: 11134,
			573: 11134,
			718: 14790,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10551,
			477: 10551,
		},
	},
	{
		"minecraft:repeater[delay=4,facing=west,locked=false,powered=false]",
		nil,
		NewMapping{
			4:   1501,
			393: 3572,
			404: 3573,
			477: 4076,
			573: 4076,
			718: 4090,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=3,south=side,west=none]",
		nil,
		NewMapping{
			718: 3242,
			393: 2936,
			404: 2937,
			477: 3240,
			573: 3240,
		},
	},
	{
		"minecraft:acacia_log[axis=y]",
		nil,
		NewMapping{
			393: 85,
			404: 85,
			477: 85,
			573: 85,
			718: 86,
			4:   2592,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 15910,
		},
	},
	{
		"minecraft:dragon_head[rotation=6]",
		nil,
		NewMapping{
			393: 5557,
			404: 5558,
			477: 6060,
			573: 6060,
			718: 6596,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=8,powered=true]",
		nil,
		NewMapping{
			573: 964,
			718: 965,
			477: 964,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11608,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=none,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14002,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 16671,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12526,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12465,
		},
	},
	{
		"minecraft:red_banner[rotation=2]",
		nil,
		NewMapping{
			573: 7587,
			718: 8123,
			393: 7080,
			404: 7081,
			477: 7587,
		},
	},
	{
		"minecraft:fire[age=10,east=true,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1468,
			404: 1469,
			477: 1772,
			573: 1772,
			718: 1773,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=8,powered=true]",
		nil,
		NewMapping{
			393: 264,
			404: 264,
			477: 264,
			573: 264,
			718: 265,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12180,
		},
	},
	{
		"minecraft:blackstone",
		nil,
		NewMapping{
			718: 15839,
		},
	},
	{
		"minecraft:rail[shape=east_west]",
		nil,
		NewMapping{
			393: 3180,
			404: 3181,
			477: 3644,
			573: 3644,
			718: 3646,
			4:   1057,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=false,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10650,
			573: 10650,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=low,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16195,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16618,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10464,
			477: 10464,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10698,
			573: 10698,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14342,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=low,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14052,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=5,powered=true]",
		nil,
		NewMapping{
			404: 708,
			477: 708,
			573: 708,
			718: 709,
			393: 708,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6030,
			404: 6031,
			477: 6537,
			573: 6537,
			718: 7073,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=22,powered=true]",
		nil,
		NewMapping{
			393: 292,
			404: 292,
			477: 292,
			573: 292,
			718: 293,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10643,
			573: 10643,
		},
	},
	{
		"minecraft:acacia_door[facing=east,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7929,
			404: 7930,
			477: 8454,
			573: 8454,
			718: 8990,
		},
	},
	{
		"minecraft:fire[age=11,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1516,
			404: 1517,
			477: 1820,
			573: 1820,
			718: 1821,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11260,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16048,
		},
	},
	{
		"minecraft:daylight_detector[inverted=false,power=15]",
		nil,
		NewMapping{
			718: 6725,
			4:   2431,
			393: 5682,
			404: 5683,
			477: 6189,
			573: 6189,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5946,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=8,powered=true]",
		nil,
		NewMapping{
			393: 414,
			404: 414,
			477: 414,
			573: 414,
			718: 415,
		},
	},
	{
		"minecraft:oak_sign[rotation=13,waterlogged=true]",
		nil,
		NewMapping{
			404: 3102,
			477: 3405,
			573: 3405,
			718: 3407,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=low,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13605,
		},
	},
	{
		"minecraft:spruce_door[facing=east,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7739,
			404: 7740,
			477: 8264,
			573: 8264,
			718: 8800,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=east,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			393: 4327,
			404: 4328,
			477: 4831,
			573: 4831,
			718: 4847,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=2,south=side,west=side]",
		nil,
		NewMapping{
			718: 2656,
			393: 2350,
			404: 2351,
			477: 2654,
			573: 2654,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 6519,
			573: 6519,
			718: 7055,
			393: 6012,
			404: 6013,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=none,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12837,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=none,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14673,
		},
	},
	{
		"minecraft:warped_trapdoor[facing=north,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 15195,
		},
	},
	{
		"minecraft:blue_glazed_terracotta[facing=west]",
		nil,
		NewMapping{
			393: 8359,
			404: 8360,
			477: 8884,
			573: 8884,
			718: 9420,
			4:   3937,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=3,south=none,west=up]",
		nil,
		NewMapping{
			393: 2073,
			404: 2074,
			477: 2377,
			573: 2377,
			718: 2379,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 6415,
			477: 6921,
			573: 6921,
			718: 7457,
			393: 6414,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11419,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16973,
		},
	},
	{
		"minecraft:pink_banner[rotation=14]",
		nil,
		NewMapping{
			393: 6964,
			404: 6965,
			477: 7471,
			573: 7471,
			718: 8007,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=7,south=side,west=none]",
		nil,
		NewMapping{
			573: 3132,
			718: 3134,
			393: 2828,
			404: 2829,
			477: 3132,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=south,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9803,
			573: 9803,
			718: 10339,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9720,
			573: 9720,
			718: 10256,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5241,
			404: 5242,
			477: 5745,
			573: 5745,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=low,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 15969,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=north,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			573: 4432,
			718: 4446,
			393: 3928,
			404: 3929,
			477: 4432,
		},
	},
	{
		"minecraft:oak_fence[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 3470,
			404: 3471,
			477: 3974,
			573: 3974,
			718: 3976,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=low,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 6052,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=none,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 15921,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=13,south=none,west=up]",
		nil,
		NewMapping{
			477: 2179,
			573: 2179,
			718: 2181,
			393: 1875,
			404: 1876,
		},
	},
	{
		"minecraft:end_gateway",
		nil,
		NewMapping{
			4:   3344,
			393: 8163,
			404: 8164,
			477: 8688,
			573: 8688,
			718: 9224,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=east,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3902,
			404: 3903,
			477: 4406,
			573: 4406,
			718: 4420,
		},
	},
	{
		"minecraft:fire[age=10,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1486,
			404: 1487,
			477: 1790,
			573: 1790,
			718: 1791,
			4:   826,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=none,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16031,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14517,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=none,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12168,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11590,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=none,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6221,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			718: 7210,
			393: 6167,
			404: 6168,
			477: 6674,
			573: 6674,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 8053,
			404: 8054,
			477: 8578,
			573: 8578,
			718: 9114,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=8,powered=false]",
		nil,
		NewMapping{
			477: 965,
			573: 965,
			718: 966,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=tall,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13873,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9752,
			573: 9752,
			718: 10288,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11279,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 6178,
			477: 6684,
			573: 6684,
			718: 7220,
			393: 6177,
		},
	},
	{
		"minecraft:light_blue_banner[rotation=13]",
		nil,
		NewMapping{
			393: 6915,
			404: 6916,
			477: 7422,
			573: 7422,
			718: 7958,
		},
	},
	{
		"minecraft:fire[age=8,east=false,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			404: 1417,
			477: 1720,
			573: 1720,
			718: 1721,
			393: 1416,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 7094,
			718: 7630,
			393: 6587,
			404: 6588,
			477: 7094,
		},
	},
	{
		"minecraft:pink_banner[rotation=9]",
		nil,
		NewMapping{
			404: 6960,
			477: 7466,
			573: 7466,
			718: 8002,
			393: 6959,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 6335,
			477: 6841,
			573: 6841,
			718: 7377,
			393: 6334,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=low,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12443,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=low,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13738,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=south,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 7565,
			393: 6522,
			404: 6523,
			477: 7029,
			573: 7029,
		},
	},
	{
		"minecraft:brown_bed[facing=east,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 952,
			404: 952,
			477: 1252,
			573: 1252,
			718: 1253,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=east,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			573: 7977,
			718: 8513,
			4:   2963,
			393: 7452,
			404: 7453,
			477: 7977,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=north,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			477: 7947,
			573: 7947,
			718: 8483,
			393: 7422,
			404: 7423,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9457,
			573: 9457,
			718: 9993,
		},
	},
	{
		"minecraft:warped_sign[rotation=10,waterlogged=false]",
		nil,
		NewMapping{
			718: 15708,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=none,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11417,
		},
	},
	{
		"minecraft:acacia_button[face=wall,facing=east,powered=true]",
		nil,
		NewMapping{
			477: 5920,
			573: 5920,
			718: 6456,
			393: 5413,
			404: 5414,
		},
	},
	{
		"minecraft:fire[age=9,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1449,
			404: 1450,
			477: 1753,
			573: 1753,
			718: 1754,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 4435,
			477: 4938,
			573: 4938,
			718: 4954,
			393: 4434,
		},
	},
	{
		"minecraft:rail[shape=ascending_south]",
		nil,
		NewMapping{
			573: 3648,
			718: 3650,
			4:   1061,
			393: 3184,
			404: 3185,
			477: 3648,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=tall,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16430,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=5,south=side,west=none]",
		nil,
		NewMapping{
			477: 2682,
			573: 2682,
			718: 2684,
			393: 2378,
			404: 2379,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=8,powered=false]",
		nil,
		NewMapping{
			477: 565,
			573: 565,
			718: 566,
			393: 565,
			404: 565,
		},
	},
	{
		"minecraft:bubble_coral[waterlogged=true]",
		nil,
		NewMapping{
			404: 8474,
			477: 8998,
			573: 8998,
			718: 9534,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=low,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16381,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13820,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=low,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16608,
		},
	},
	{
		"minecraft:magenta_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 789,
			404: 789,
			477: 1089,
			573: 1089,
			718: 1090,
		},
	},
	{
		"minecraft:andesite_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9975,
			573: 9975,
			718: 10511,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=tall,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12265,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=tall,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11942,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 10914,
		},
	},
	{
		"minecraft:stripped_jungle_log[axis=z]",
		nil,
		NewMapping{
			393: 98,
			404: 98,
			477: 98,
			573: 98,
			718: 99,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=false,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			477: 4545,
			573: 4545,
			718: 4559,
			393: 4041,
			404: 4042,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=4,south=none,west=up]",
		nil,
		NewMapping{
			393: 1938,
			404: 1939,
			477: 2242,
			573: 2242,
			718: 2244,
		},
	},
	{
		"minecraft:white_banner[rotation=1]",
		nil,
		NewMapping{
			718: 7898,
			4:   2817,
			393: 6855,
			404: 6856,
			477: 7362,
			573: 7362,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=none,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13243,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			718: 7306,
			393: 6263,
			404: 6264,
			477: 6770,
			573: 6770,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=low,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12751,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=tall,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 10966,
		},
	},
	{
		"minecraft:respawn_anchor[charges=2]",
		nil,
		NewMapping{
			718: 15831,
		},
	},
	{
		"minecraft:smoker[facing=north,lit=true]",
		nil,
		NewMapping{
			718: 14803,
			477: 11147,
			573: 11147,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9738,
			573: 9738,
			718: 10274,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12797,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=none,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14222,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 17100,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=1,powered=false]",
		nil,
		NewMapping{
			393: 401,
			404: 401,
			477: 401,
			573: 401,
			718: 402,
		},
	},
	{
		"minecraft:yellow_banner[rotation=10]",
		nil,
		NewMapping{
			393: 6928,
			404: 6929,
			477: 7435,
			573: 7435,
			718: 7971,
		},
	},
	{
		"minecraft:purple_banner[rotation=3]",
		nil,
		NewMapping{
			393: 7017,
			404: 7018,
			477: 7524,
			573: 7524,
			718: 8060,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=0,south=side,west=side]",
		nil,
		NewMapping{
			393: 2044,
			404: 2045,
			477: 2348,
			573: 2348,
			718: 2350,
		},
	},
	{
		"minecraft:gray_banner[rotation=10]",
		nil,
		NewMapping{
			393: 6976,
			404: 6977,
			477: 7483,
			573: 7483,
			718: 8019,
		},
	},
	{
		"minecraft:oak_leaves[distance=2,persistent=false]",
		nil,
		NewMapping{
			404: 147,
			477: 147,
			573: 147,
			718: 148,
			4:   296,
			393: 147,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=7,powered=true]",
		nil,
		NewMapping{
			477: 962,
			573: 962,
			718: 963,
		},
	},
	{
		"minecraft:dispenser[facing=south,triggered=true]",
		nil,
		NewMapping{
			573: 237,
			718: 238,
			4:   379,
			393: 237,
			404: 237,
			477: 237,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16965,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12285,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12905,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=0,south=none,west=none]",
		nil,
		NewMapping{
			404: 2337,
			477: 2640,
			573: 2640,
			718: 2642,
			393: 2336,
		},
	},
	{
		"minecraft:fire[age=12,east=false,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			477: 1848,
			573: 1848,
			718: 1849,
			393: 1544,
			404: 1545,
		},
	},
	{
		"minecraft:spruce_door[facing=east,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7736,
			404: 7737,
			477: 8261,
			573: 8261,
			718: 8797,
		},
	},
	{
		"minecraft:gray_banner[rotation=1]",
		nil,
		NewMapping{
			404: 6968,
			477: 7474,
			573: 7474,
			718: 8010,
			393: 6967,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=none,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11734,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=low,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11676,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=7,powered=false]",
		nil,
		NewMapping{
			393: 663,
			404: 663,
			477: 663,
			573: 663,
			718: 664,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 11037,
			477: 11037,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=17,powered=true]",
		nil,
		NewMapping{
			477: 882,
			573: 882,
			718: 883,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9772,
			573: 9772,
			718: 10308,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12700,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=none,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6113,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5905,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16214,
		},
	},
	{
		"minecraft:daylight_detector[inverted=false,power=1]",
		nil,
		NewMapping{
			718: 6711,
			4:   2417,
			393: 5668,
			404: 5669,
			477: 6175,
			573: 6175,
		},
	},
	{
		"minecraft:spruce_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 5424,
			718: 5440,
			393: 4920,
			404: 4921,
			477: 5424,
		},
	},
	{
		"minecraft:green_glazed_terracotta[facing=west]",
		nil,
		NewMapping{
			393: 8367,
			404: 8368,
			477: 8892,
			573: 8892,
			718: 9428,
			4:   3969,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=tall,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12792,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=none,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16451,
		},
	},
	{
		"minecraft:warped_fence_gate[facing=north,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15293,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6643,
			404: 6644,
			477: 7150,
			573: 7150,
			718: 7686,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10352,
			573: 10352,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12652,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=low,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12215,
		},
	},
	{
		"minecraft:brain_coral[waterlogged=true]",
		nil,
		NewMapping{
			404: 8472,
			477: 8996,
			573: 8996,
			718: 9532,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12238,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13536,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 10911,
		},
	},
	{
		"minecraft:fire[age=5,east=false,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1317,
			404: 1318,
			477: 1621,
			573: 1621,
			718: 1622,
		},
	},
	{
		"minecraft:command_block[conditional=false,facing=down]",
		nil,
		NewMapping{
			393: 5135,
			404: 5136,
			477: 5639,
			573: 5639,
			718: 5655,
			4:   2192,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=2,south=up,west=none]",
		nil,
		NewMapping{
			404: 2925,
			477: 3228,
			573: 3228,
			718: 3230,
			393: 2924,
		},
	},
	{
		"minecraft:oak_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 2017,
			573: 2017,
			718: 2019,
			393: 1713,
			404: 1714,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=none,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14445,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11064,
			573: 11064,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=low,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13959,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=tall,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14097,
		},
	},
	{
		"minecraft:spruce_door[facing=south,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7704,
			404: 7705,
			477: 8229,
			573: 8229,
			718: 8765,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=west,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			718: 4144,
			393: 3626,
			404: 3627,
			477: 4130,
			573: 4130,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=0,south=none,west=none]",
		nil,
		NewMapping{
			393: 2912,
			404: 2913,
			477: 3216,
			573: 3216,
			718: 3218,
			4:   880,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=west,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			573: 4204,
			718: 4218,
			393: 3700,
			404: 3701,
			477: 4204,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=15,south=up,west=none]",
		nil,
		NewMapping{
			393: 2033,
			404: 2034,
			477: 2337,
			573: 2337,
			718: 2339,
		},
	},
	{
		"minecraft:repeater[delay=1,facing=south,locked=false,powered=false]",
		nil,
		NewMapping{
			573: 4024,
			718: 4038,
			4:   1488,
			393: 3520,
			404: 3521,
			477: 4024,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12990,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14694,
		},
	},
	{
		"minecraft:crimson_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 15375,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4424,
			404: 4425,
			477: 4928,
			573: 4928,
			718: 4944,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10834,
			573: 10834,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=tall,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 10939,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5951,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=2,powered=false]",
		nil,
		NewMapping{
			393: 253,
			404: 253,
			477: 253,
			573: 253,
			718: 254,
		},
	},
	{
		"minecraft:bell[attachment=single_wall,facing=east]",
		nil,
		NewMapping{
			477: 11209,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=tall,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13109,
		},
	},
	{
		"minecraft:crimson_button[face=ceiling,facing=north,powered=false]",
		nil,
		NewMapping{
			718: 15496,
		},
	},
	{
		"minecraft:cyan_banner[rotation=14]",
		nil,
		NewMapping{
			393: 7012,
			404: 7013,
			477: 7519,
			573: 7519,
			718: 8055,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=none,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12401,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 5058,
			718: 5074,
			393: 4554,
			404: 4555,
			477: 5058,
		},
	},
	{
		"minecraft:fire[age=4,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			573: 1570,
			718: 1571,
			393: 1266,
			404: 1267,
			477: 1570,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=3,south=up,west=side]",
		nil,
		NewMapping{
			393: 1780,
			404: 1781,
			477: 2084,
			573: 2084,
			718: 2086,
		},
	},
	{
		"minecraft:jungle_button[face=ceiling,facing=east,powered=false]",
		nil,
		NewMapping{
			573: 5905,
			718: 6441,
			393: 5398,
			404: 5399,
			477: 5905,
		},
	},
	{
		"minecraft:brain_coral",
		nil,
		NewMapping{
			393: 8460,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=15,south=side,west=side]",
		nil,
		NewMapping{
			573: 3203,
			718: 3205,
			393: 2899,
			404: 2900,
			477: 3203,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=none,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12834,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=0,south=side,west=up]",
		nil,
		NewMapping{
			718: 2493,
			393: 2187,
			404: 2188,
			477: 2491,
			573: 2491,
		},
	},
	{
		"minecraft:bubble_coral_wall_fan[facing=south,waterlogged=false]",
		nil,
		NewMapping{
			718: 9619,
			393: 8523,
			404: 8539,
			477: 9083,
			573: 9083,
		},
	},
	{
		"minecraft:birch_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 5488,
			718: 5504,
			393: 4984,
			404: 4985,
			477: 5488,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14717,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=low,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14162,
		},
	},
	{
		"minecraft:weeping_vines[age=6]",
		nil,
		NewMapping{
			718: 14996,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 5764,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 7179,
			573: 7179,
			718: 7715,
			393: 6672,
			404: 6673,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=true,north=true,powered=false,south=true,west=true]",
		nil,
		NewMapping{
			573: 5263,
			718: 5279,
			393: 4759,
			404: 4760,
			477: 5263,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5238,
			404: 5239,
			477: 5742,
			573: 5742,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 7180,
			404: 7181,
			477: 7687,
			573: 7687,
			718: 8223,
		},
	},
	{
		"minecraft:jungle_fence[east=true,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7585,
			404: 7586,
			477: 8110,
			573: 8110,
			718: 8646,
		},
	},
	{
		"minecraft:oak_button[face=ceiling,facing=south,powered=false]",
		nil,
		NewMapping{
			404: 5323,
			477: 5829,
			573: 5829,
			718: 6365,
			393: 5322,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10137,
			573: 10137,
			718: 10673,
		},
	},
	{
		"minecraft:red_nether_brick_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			477: 10318,
			573: 10318,
			718: 10854,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=low,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12442,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=12,south=none,west=none]",
		nil,
		NewMapping{
			477: 2748,
			573: 2748,
			718: 2750,
			393: 2444,
			404: 2445,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 7216,
			477: 7722,
			573: 7722,
			718: 8258,
			393: 7215,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=true,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11003,
			573: 11003,
		},
	},
	{
		"minecraft:birch_sign[rotation=12,waterlogged=true]",
		nil,
		NewMapping{
			477: 3467,
			573: 3467,
			718: 3469,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=true,north=false,powered=false,south=true,west=false]",
		nil,
		NewMapping{
			404: 4769,
			477: 5272,
			573: 5272,
			718: 5288,
			393: 4768,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5981,
			404: 5982,
			477: 6488,
			573: 6488,
			718: 7024,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=low,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12445,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16851,
		},
	},
	{
		"minecraft:warped_fence_gate[facing=west,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			718: 15306,
		},
	},
	{
		"minecraft:purpur_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 8639,
			718: 9175,
			393: 8114,
			404: 8115,
			477: 8639,
		},
	},
	{
		"minecraft:yellow_banner[rotation=8]",
		nil,
		NewMapping{
			573: 7433,
			718: 7969,
			393: 6926,
			404: 6927,
			477: 7433,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			718: 7176,
			393: 6133,
			404: 6134,
			477: 6640,
			573: 6640,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9322,
			573: 9322,
			718: 9858,
		},
	},
	{
		"minecraft:quartz_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5717,
			404: 5718,
			477: 6224,
			573: 6224,
			718: 6760,
		},
	},
	{
		"minecraft:andesite_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9955,
			573: 9955,
			718: 10491,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9327,
			573: 9327,
			718: 9863,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=none,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13913,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=none,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12597,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=1,south=up,west=side]",
		nil,
		NewMapping{
			393: 2626,
			404: 2627,
			477: 2930,
			573: 2930,
			718: 2932,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 5446,
			718: 5462,
			393: 4942,
			404: 4943,
			477: 5446,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12176,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=7,south=none,west=up]",
		nil,
		NewMapping{
			477: 2701,
			573: 2701,
			718: 2703,
			393: 2397,
			404: 2398,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10248,
			573: 10248,
			718: 10784,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 16685,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14665,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12773,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=west,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7376,
			404: 7377,
			477: 7901,
			573: 7901,
			718: 8437,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=west,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3951,
			404: 3952,
			477: 4455,
			573: 4455,
			718: 4469,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9783,
			573: 9783,
			718: 10319,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11479,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=low,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11466,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=low,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14161,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12241,
		},
	},
	{
		"minecraft:heavy_weighted_pressure_plate[power=9]",
		nil,
		NewMapping{
			393: 5628,
			404: 5629,
			477: 6135,
			573: 6135,
			718: 6671,
			4:   2377,
		},
	},
	{
		"minecraft:birch_button[face=wall,facing=north,powered=true]",
		nil,
		NewMapping{
			393: 5359,
			404: 5360,
			477: 5866,
			573: 5866,
			718: 6402,
		},
	},
	{
		"minecraft:jungle_sign[rotation=8,waterlogged=false]",
		nil,
		NewMapping{
			477: 3524,
			573: 3524,
			718: 3526,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=low,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13829,
		},
	},
	{
		"minecraft:brick_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4349,
			404: 4350,
			477: 4853,
			573: 4853,
			718: 4869,
		},
	},
	{
		"minecraft:daylight_detector[inverted=false,power=0]",
		nil,
		NewMapping{
			393: 5667,
			404: 5668,
			477: 6174,
			573: 6174,
			718: 6710,
			4:   2416,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 4132,
			404: 4133,
			477: 4636,
			573: 4636,
			718: 4650,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=19,powered=true]",
		nil,
		NewMapping{
			404: 736,
			477: 736,
			573: 736,
			718: 737,
			393: 736,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12653,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=17,powered=true]",
		nil,
		NewMapping{
			477: 932,
			573: 932,
			718: 933,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12021,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=low,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16606,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=tall,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 5846,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13102,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=tall,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 6178,
		},
	},
	{
		"minecraft:oak_wall_sign[facing=east,waterlogged=true]",
		nil,
		NewMapping{
			404: 3276,
			477: 3739,
			573: 3739,
			718: 3741,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 9694,
			718: 10230,
			477: 9694,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=none,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14113,
		},
	},
	{
		"minecraft:jungle_log[axis=z]",
		nil,
		NewMapping{
			404: 83,
			477: 83,
			573: 83,
			718: 84,
			4:   283,
			393: 83,
		},
	},
	{
		"minecraft:quartz_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 6774,
			4:   2498,
			393: 5731,
			404: 5732,
			477: 6238,
			573: 6238,
		},
	},
	{
		"minecraft:spruce_fence[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			718: 8598,
			393: 7537,
			404: 7538,
			477: 8062,
			573: 8062,
		},
	},
	{
		"minecraft:dead_brain_coral_fan[waterlogged=false]",
		nil,
		NewMapping{
			393: 8547,
			404: 8563,
			477: 9007,
			573: 9007,
			718: 9543,
		},
	},
	{
		"minecraft:nether_brick_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			477: 7853,
			573: 7853,
			718: 8389,
			4:   702,
			393: 7334,
			404: 7335,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9399,
			573: 9399,
			718: 9935,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=low,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11031,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 5057,
			718: 5073,
			393: 4553,
			404: 4554,
			477: 5057,
		},
	},
	{
		"minecraft:spruce_fence[east=false,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7547,
			404: 7548,
			477: 8072,
			573: 8072,
			718: 8608,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10673,
			573: 10673,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=low,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16378,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=7,south=side,west=up]",
		nil,
		NewMapping{
			477: 2410,
			573: 2410,
			718: 2412,
			393: 2106,
			404: 2107,
		},
	},
	{
		"minecraft:fire[age=3,east=false,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			573: 1557,
			718: 1558,
			393: 1253,
			404: 1254,
			477: 1557,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14491,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11882,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11744,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 15855,
		},
	},
	{
		"minecraft:warped_door[facing=east,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15649,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 7751,
			573: 7751,
			718: 8287,
			393: 7244,
			404: 7245,
		},
	},
	{
		"minecraft:smooth_red_sandstone_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			477: 10261,
			573: 10261,
			718: 10797,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9400,
			573: 9400,
			718: 9936,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13219,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=low,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13500,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9815,
			573: 9815,
			718: 10351,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=none,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11299,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16902,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=none,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11418,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 6641,
			404: 6642,
			477: 7148,
			573: 7148,
			718: 7684,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=tall,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6083,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12794,
		},
	},
	{
		"minecraft:spruce_fence[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 8068,
			718: 8604,
			393: 7543,
			404: 7544,
			477: 8068,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=24,powered=false]",
		nil,
		NewMapping{
			404: 597,
			477: 597,
			573: 597,
			718: 598,
			393: 597,
		},
	},
	{
		"minecraft:birch_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 5505,
			573: 5505,
			718: 5521,
			393: 5001,
			404: 5002,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12674,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13935,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			718: 7267,
			393: 6224,
			404: 6225,
			477: 6731,
			573: 6731,
		},
	},
	{
		"minecraft:wall_sign[facing=east,waterlogged=false]",
		nil,
		NewMapping{
			393: 3276,
			4:   1093,
		},
	},
	{
		"minecraft:heavy_weighted_pressure_plate[power=15]",
		nil,
		NewMapping{
			573: 6141,
			718: 6677,
			4:   2383,
			393: 5634,
			404: 5635,
			477: 6141,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=low,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14594,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16541,
		},
	},
	{
		"minecraft:crimson_fence[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			718: 15069,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12993,
		},
	},
	{
		"minecraft:bookshelf",
		nil,
		NewMapping{
			573: 1431,
			718: 1432,
			4:   752,
			393: 1127,
			404: 1128,
			477: 1431,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=low,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13854,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14596,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=tall,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13652,
		},
	},
	{
		"minecraft:wall_sign[facing=east,waterlogged=true]",
		nil,
		NewMapping{
			393: 3275,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=low,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16187,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=4,powered=false]",
		nil,
		NewMapping{
			393: 557,
			404: 557,
			477: 557,
			573: 557,
			718: 558,
		},
	},
	{
		"minecraft:diorite_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 10194,
			573: 10194,
			718: 10730,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 15869,
		},
	},
	{
		"minecraft:gray_bed[facing=east,occupied=true,part=foot]",
		nil,
		NewMapping{
			573: 1173,
			718: 1174,
			393: 873,
			404: 873,
			477: 1173,
		},
	},
	{
		"minecraft:piston[extended=false,facing=east]",
		nil,
		NewMapping{
			4:   533,
			393: 1054,
			404: 1054,
			477: 1354,
			573: 1354,
			718: 1355,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=false,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 11034,
			573: 11034,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11914,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=low,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13184,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=2,powered=true]",
		nil,
		NewMapping{
			393: 252,
			404: 252,
			477: 252,
			573: 252,
			718: 253,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=4,south=side,west=up]",
		nil,
		NewMapping{
			393: 2799,
			404: 2800,
			477: 3103,
			573: 3103,
			718: 3105,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=low,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14712,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=low,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11042,
		},
	},
	{
		"minecraft:soul_campfire[facing=east,lit=false,signal_fire=false,waterlogged=true]",
		nil,
		NewMapping{
			718: 14952,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11643,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 3224,
			404: 3225,
			477: 3688,
			573: 3688,
			718: 3690,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=16,powered=true]",
		nil,
		NewMapping{
			477: 530,
			573: 530,
			718: 531,
			393: 530,
			404: 530,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=8,south=side,west=side]",
		nil,
		NewMapping{
			393: 2260,
			404: 2261,
			477: 2564,
			573: 2564,
			718: 2566,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10527,
			573: 10527,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9418,
			573: 9418,
			718: 9954,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10397,
			477: 10397,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=11,south=side,west=up]",
		nil,
		NewMapping{
			393: 2142,
			404: 2143,
			477: 2446,
			573: 2446,
			718: 2448,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=15,powered=true]",
		nil,
		NewMapping{
			393: 478,
			404: 478,
			477: 478,
			573: 478,
			718: 479,
		},
	},
	{
		"minecraft:warped_trapdoor[facing=north,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			718: 15201,
		},
	},
	{
		"minecraft:magenta_banner[rotation=8]",
		nil,
		NewMapping{
			393: 6894,
			404: 6895,
			477: 7401,
			573: 7401,
			718: 7937,
		},
	},
	{
		"minecraft:fire[age=15,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1646,
			404: 1647,
			477: 1950,
			573: 1950,
			718: 1951,
			4:   831,
		},
	},
	{
		"minecraft:cyan_bed[facing=south,occupied=true,part=foot]",
		nil,
		NewMapping{
			404: 897,
			477: 1197,
			573: 1197,
			718: 1198,
			393: 897,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=8,south=none,west=up]",
		nil,
		NewMapping{
			573: 2998,
			718: 3000,
			393: 2694,
			404: 2695,
			477: 2998,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 6028,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11562,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12978,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=0,powered=false]",
		nil,
		NewMapping{
			573: 849,
			718: 850,
			477: 849,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9361,
			573: 9361,
			718: 9897,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=none,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6119,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 15952,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9840,
			718: 10376,
			477: 9840,
		},
	},
	{
		"minecraft:stripped_warped_stem[axis=y]",
		nil,
		NewMapping{
			718: 14962,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=false,north=true,powered=true,south=true,west=false]",
		nil,
		NewMapping{
			404: 4773,
			477: 5276,
			573: 5276,
			718: 5292,
			393: 4772,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=west,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			404: 3821,
			477: 4324,
			573: 4324,
			718: 4338,
			393: 3820,
		},
	},
	{
		"minecraft:dark_oak_fence[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 8185,
			718: 8721,
			393: 7660,
			404: 7661,
			477: 8185,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=0,south=side,west=up]",
		nil,
		NewMapping{
			393: 2043,
			404: 2044,
			477: 2347,
			573: 2347,
			718: 2349,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=15,south=none,west=none]",
		nil,
		NewMapping{
			404: 2472,
			477: 2775,
			573: 2775,
			718: 2777,
			393: 2471,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 4922,
			718: 4938,
			393: 4418,
			404: 4419,
			477: 4922,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10909,
			573: 10909,
		},
	},
	{
		"minecraft:crimson_fence[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			718: 15072,
		},
	},
	{
		"minecraft:spruce_leaves[distance=3,persistent=false]",
		nil,
		NewMapping{
			393: 163,
			404: 163,
			477: 163,
			573: 163,
			718: 164,
		},
	},
	{
		"minecraft:light_gray_shulker_box[facing=north]",
		nil,
		NewMapping{
			718: 9326,
			4:   3634,
			393: 8265,
			404: 8266,
			477: 8790,
			573: 8790,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=1,south=side,west=up]",
		nil,
		NewMapping{
			477: 2068,
			573: 2068,
			718: 2070,
			393: 1764,
			404: 1765,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=tall,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11263,
		},
	},
	{
		"minecraft:bell[attachment=double_wall,facing=south]",
		nil,
		NewMapping{
			477: 11211,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12934,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=tall,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16129,
		},
	},
	{
		"minecraft:activator_rail[powered=false,shape=north_south]",
		nil,
		NewMapping{
			393: 5786,
			404: 5787,
			477: 6293,
			573: 6293,
			718: 6829,
			4:   2512,
		},
	},
	{
		"minecraft:petrified_oak_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			4:   698,
			393: 7310,
			404: 7311,
			477: 7829,
			573: 7829,
			718: 8365,
		},
	},
	{
		"minecraft:fire[age=8,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			573: 1709,
			718: 1710,
			393: 1405,
			404: 1406,
			477: 1709,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10737,
			573: 10737,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6417,
			404: 6418,
			477: 6924,
			573: 6924,
			718: 7460,
		},
	},
	{
		"minecraft:jungle_leaves[distance=4,persistent=true]",
		nil,
		NewMapping{
			718: 193,
			393: 192,
			404: 192,
			477: 192,
			573: 192,
		},
	},
	{
		"minecraft:comparator[facing=north,mode=compare,powered=false]",
		nil,
		NewMapping{
			573: 6143,
			718: 6679,
			4:   2386,
			393: 5636,
			404: 5637,
			477: 6143,
		},
	},
	{
		"minecraft:dark_oak_door[facing=south,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			718: 9022,
			393: 7961,
			404: 7962,
			477: 8486,
			573: 8486,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=low,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11150,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13998,
		},
	},
	{
		"minecraft:birch_planks",
		nil,
		NewMapping{
			393: 17,
			404: 17,
			477: 17,
			573: 17,
			718: 17,
			4:   82,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4660,
			404: 4661,
			477: 5164,
			573: 5164,
			718: 5180,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=9,powered=false]",
		nil,
		NewMapping{
			393: 267,
			404: 267,
			477: 267,
			573: 267,
			718: 268,
		},
	},
	{
		"minecraft:crimson_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 15360,
		},
	},
	{
		"minecraft:spruce_fence[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 7538,
			404: 7539,
			477: 8063,
			573: 8063,
			718: 8599,
		},
	},
	{
		"minecraft:oak_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 1708,
			404: 1709,
			477: 2012,
			573: 2012,
			718: 2014,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=low,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11474,
		},
	},
	{
		"minecraft:quartz_bricks",
		nil,
		NewMapping{
			718: 17103,
		},
	},
	{
		"minecraft:purpur_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			393: 7347,
			404: 7348,
			477: 7872,
			573: 7872,
			718: 8408,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13304,
		},
	},
	{
		"minecraft:fire[age=0,east=false,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1155,
			404: 1156,
			477: 1459,
			573: 1459,
			718: 1460,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			573: 4586,
			718: 4600,
			393: 4082,
			404: 4083,
			477: 4586,
		},
	},
	{
		"minecraft:jungle_button[face=wall,facing=west,powered=true]",
		nil,
		NewMapping{
			393: 5387,
			404: 5388,
			477: 5894,
			573: 5894,
			718: 6430,
		},
	},
	{
		"minecraft:red_sandstone_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			718: 8400,
			393: 7345,
			404: 7346,
			477: 7864,
			573: 7864,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=2,powered=false]",
		nil,
		NewMapping{
			404: 703,
			477: 703,
			573: 703,
			718: 704,
			393: 703,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10486,
			573: 10486,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=8,powered=false]",
		nil,
		NewMapping{
			393: 415,
			404: 415,
			477: 415,
			573: 415,
			718: 416,
		},
	},
	{
		"minecraft:loom[facing=south]",
		nil,
		NewMapping{
			477: 11132,
			573: 11132,
			718: 14788,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=10,south=none,west=up]",
		nil,
		NewMapping{
			573: 3016,
			718: 3018,
			393: 2712,
			404: 2713,
			477: 3016,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=low,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14160,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=none,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12708,
		},
	},
	{
		"minecraft:warped_fence_gate[facing=west,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			718: 15310,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=tall,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12568,
		},
	},
	{
		"minecraft:brown_glazed_terracotta[facing=east]",
		nil,
		NewMapping{
			573: 8889,
			718: 9425,
			4:   3955,
			393: 8364,
			404: 8365,
			477: 8889,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=west,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			573: 7962,
			718: 8498,
			393: 7437,
			404: 7438,
			477: 7962,
		},
	},
	{
		"minecraft:beehive[facing=north,honey_level=4]",
		nil,
		NewMapping{
			573: 11315,
			718: 15804,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11080,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=1,south=none,west=none]",
		nil,
		NewMapping{
			393: 2345,
			404: 2346,
			477: 2649,
			573: 2649,
			718: 2651,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 5754,
			393: 5250,
			404: 5251,
			477: 5754,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16932,
		},
	},
	{
		"minecraft:stripped_birch_wood[axis=y]",
		nil,
		NewMapping{
			393: 133,
			404: 133,
			477: 133,
			573: 133,
			718: 134,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=22,powered=true]",
		nil,
		NewMapping{
			477: 842,
			573: 842,
			718: 843,
		},
	},
	{
		"minecraft:warped_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 15408,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=tall,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 17097,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=1,powered=false]",
		nil,
		NewMapping{
			477: 851,
			573: 851,
			718: 852,
		},
	},
	{
		"minecraft:bell[attachment=single_wall,facing=west,powered=false]",
		nil,
		NewMapping{
			573: 11219,
			718: 14875,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=south,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3685,
			404: 3686,
			477: 4189,
			573: 4189,
			718: 4203,
		},
	},
	{
		"minecraft:sign[rotation=8,waterlogged=false]",
		nil,
		NewMapping{
			393: 3092,
			4:   1016,
		},
	},
	{
		"minecraft:birch_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4993,
			404: 4994,
			477: 5497,
			573: 5497,
			718: 5513,
		},
	},
	{
		"minecraft:birch_sign[rotation=0,waterlogged=false]",
		nil,
		NewMapping{
			477: 3444,
			573: 3444,
			718: 3446,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=low,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14047,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13412,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 6332,
			477: 6838,
			573: 6838,
			718: 7374,
			4:   2575,
			393: 6331,
		},
	},
	{
		"minecraft:oak_sign[rotation=8,waterlogged=true]",
		nil,
		NewMapping{
			573: 3395,
			718: 3397,
			404: 3092,
			477: 3395,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9280,
			573: 9280,
			718: 9816,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=low,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11983,
		},
	},
	{
		"minecraft:birch_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4973,
			404: 4974,
			477: 5477,
			573: 5477,
			718: 5493,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9818,
			573: 9818,
			718: 10354,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 6865,
			718: 7401,
			393: 6358,
			404: 6359,
			477: 6865,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9589,
			573: 9589,
			718: 10125,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=17,powered=true]",
		nil,
		NewMapping{
			477: 532,
			573: 532,
			718: 533,
			393: 532,
			404: 532,
		},
	},
	{
		"minecraft:acacia_fence[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7623,
			404: 7624,
			477: 8148,
			573: 8148,
			718: 8684,
		},
	},
	{
		"minecraft:white_banner[rotation=4]",
		nil,
		NewMapping{
			4:   2820,
			393: 6858,
			404: 6859,
			477: 7365,
			573: 7365,
			718: 7901,
		},
	},
	{
		"minecraft:warped_door[facing=west,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			718: 15623,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5952,
			404: 5953,
			477: 6459,
			573: 6459,
			718: 6995,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9468,
			573: 9468,
			718: 10004,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11482,
		},
	},
	{
		"minecraft:birch_door[facing=south,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7766,
			404: 7767,
			477: 8291,
			573: 8291,
			718: 8827,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=south,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3872,
			404: 3873,
			477: 4376,
			573: 4376,
			718: 4390,
		},
	},
	{
		"minecraft:purpur_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 9147,
			393: 8086,
			404: 8087,
			477: 8611,
			573: 8611,
		},
	},
	{
		"minecraft:repeater[delay=3,facing=north,locked=false,powered=false]",
		nil,
		NewMapping{
			404: 3549,
			477: 4052,
			573: 4052,
			718: 4066,
			4:   1498,
			393: 3548,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=none,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13141,
		},
	},
	{
		"minecraft:cake[bites=6]",
		nil,
		NewMapping{
			404: 3513,
			477: 4016,
			573: 4016,
			718: 4030,
			4:   1478,
			393: 3512,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9652,
			573: 9652,
			718: 10188,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=low,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16074,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5735,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=18,powered=true]",
		nil,
		NewMapping{
			477: 884,
			573: 884,
			718: 885,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=none,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11327,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=low,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6239,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5717,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=west,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			404: 7411,
			477: 7935,
			573: 7935,
			718: 8471,
			4:   2949,
			393: 7410,
		},
	},
	{
		"minecraft:purple_bed[facing=south,occupied=false,part=foot]",
		nil,
		NewMapping{
			477: 1215,
			573: 1215,
			718: 1216,
			393: 915,
			404: 915,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=8,south=up,west=none]",
		nil,
		NewMapping{
			573: 2994,
			718: 2996,
			393: 2690,
			404: 2691,
			477: 2994,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=14,powered=false]",
		nil,
		NewMapping{
			393: 327,
			404: 327,
			477: 327,
			573: 327,
			718: 328,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11173,
		},
	},
	{
		"minecraft:potted_orange_tulip",
		nil,
		NewMapping{
			393: 5279,
			404: 5280,
			477: 5783,
			573: 5783,
			718: 6319,
		},
	},
	{
		"minecraft:bee_nest[facing=south,honey_level=3]",
		nil,
		NewMapping{
			573: 11296,
			718: 15785,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=none,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16349,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=2,south=none,west=side]",
		nil,
		NewMapping{
			393: 2065,
			404: 2066,
			477: 2369,
			573: 2369,
			718: 2371,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=14,waterlogged=true]",
		nil,
		NewMapping{
			477: 3567,
			573: 3567,
			718: 3569,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13661,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=low,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14577,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11641,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=low,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 15965,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			718: 6863,
			393: 5820,
			404: 5821,
			477: 6327,
			573: 6327,
		},
	},
	{
		"minecraft:birch_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 4995,
			404: 4996,
			477: 5499,
			573: 5499,
			718: 5515,
			4:   2162,
		},
	},
	{
		"minecraft:oak_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			718: 8302,
			393: 7259,
			404: 7260,
			477: 7766,
			573: 7766,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10856,
			573: 10856,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=3,south=side,west=side]",
		nil,
		NewMapping{
			573: 3239,
			718: 3241,
			393: 2935,
			404: 2936,
			477: 3239,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13298,
		},
	},
	{
		"minecraft:crimson_door[facing=north,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			718: 15538,
		},
	},
	{
		"minecraft:repeater[delay=4,facing=north,locked=false,powered=false]",
		nil,
		NewMapping{
			404: 3565,
			477: 4068,
			573: 4068,
			718: 4082,
			4:   1502,
			393: 3564,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=0,south=none,west=up]",
		nil,
		NewMapping{
			393: 2766,
			404: 2767,
			477: 3070,
			573: 3070,
			718: 3072,
		},
	},
	{
		"minecraft:cracked_stone_bricks",
		nil,
		NewMapping{
			393: 3985,
			404: 3986,
			477: 4483,
			573: 4483,
			718: 4497,
			4:   1570,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10020,
			573: 10020,
			718: 10556,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=south,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			404: 7372,
			477: 7896,
			573: 7896,
			718: 8432,
			4:   2936,
			393: 7371,
		},
	},
	{
		"minecraft:jungle_door[facing=west,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			404: 7841,
			477: 8365,
			573: 8365,
			718: 8901,
			393: 7840,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=tall,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16872,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 16295,
		},
	},
	{
		"minecraft:crimson_fence[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			718: 15065,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11208,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=21,powered=true]",
		nil,
		NewMapping{
			393: 740,
			404: 740,
			477: 740,
			573: 740,
			718: 741,
		},
	},
	{
		"minecraft:oak_door[facing=south,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			393: 3128,
			404: 3129,
			477: 3592,
			573: 3592,
			718: 3594,
		},
	},
	{
		"minecraft:dark_oak_wall_sign[facing=north,waterlogged=false]",
		nil,
		NewMapping{
			477: 3774,
			573: 3774,
			718: 3776,
		},
	},
	{
		"minecraft:end_stone_brick_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			477: 10285,
			573: 10285,
			718: 10821,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=none,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16994,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=false,north=false,powered=false,south=false,west=true]",
		nil,
		NewMapping{
			393: 4849,
			404: 4850,
			477: 5353,
			573: 5353,
			718: 5369,
		},
	},
	{
		"minecraft:birch_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 5510,
			718: 5526,
			393: 5006,
			404: 5007,
			477: 5510,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=false,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			718: 4564,
			393: 4046,
			404: 4047,
			477: 4550,
			573: 4550,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=none,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16059,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=low,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11462,
		},
	},
	{
		"minecraft:fire[age=5,east=true,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			477: 1605,
			573: 1605,
			718: 1606,
			393: 1301,
			404: 1302,
		},
	},
	{
		"minecraft:granite_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			477: 10306,
			573: 10306,
			718: 10842,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=17,powered=false]",
		nil,
		NewMapping{
			477: 1033,
			573: 1033,
			718: 1034,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=none,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14223,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=none,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12836,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=low,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13524,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14286,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=12,south=up,west=up]",
		nil,
		NewMapping{
			393: 2868,
			404: 2869,
			477: 3172,
			573: 3172,
			718: 3174,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=15,south=none,west=up]",
		nil,
		NewMapping{
			573: 3061,
			718: 3063,
			393: 2757,
			404: 2758,
			477: 3061,
		},
	},
	{
		"minecraft:spruce_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 5404,
			393: 4884,
			404: 4885,
			477: 5388,
			573: 5388,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10089,
			573: 10089,
			718: 10625,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=none,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12388,
		},
	},
	{
		"minecraft:oak_leaves[distance=1,persistent=false]",
		nil,
		NewMapping{
			4:   288,
			393: 145,
			404: 145,
			477: 145,
			573: 145,
			718: 146,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=north,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3851,
			404: 3852,
			477: 4355,
			573: 4355,
			718: 4369,
		},
	},
	{
		"minecraft:beehive[facing=east,honey_level=1]",
		nil,
		NewMapping{
			573: 11330,
			718: 15819,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13644,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 3685,
			573: 3685,
			718: 3687,
			393: 3221,
			404: 3222,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=23,powered=false]",
		nil,
		NewMapping{
			477: 895,
			573: 895,
			718: 896,
		},
	},
	{
		"minecraft:oak_button[face=wall,facing=west,powered=true]",
		nil,
		NewMapping{
			573: 5822,
			718: 6358,
			4:   2298,
			393: 5315,
			404: 5316,
			477: 5822,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			718: 9082,
			393: 8021,
			404: 8022,
			477: 8546,
			573: 8546,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=2,south=side,west=none]",
		nil,
		NewMapping{
			573: 2655,
			718: 2657,
			393: 2351,
			404: 2352,
			477: 2655,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=4,south=side,west=up]",
		nil,
		NewMapping{
			393: 2223,
			404: 2224,
			477: 2527,
			573: 2527,
			718: 2529,
		},
	},
	{
		"minecraft:purpur_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 8093,
			404: 8094,
			477: 8618,
			573: 8618,
			718: 9154,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=21,powered=true]",
		nil,
		NewMapping{
			573: 490,
			718: 491,
			393: 490,
			404: 490,
			477: 490,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=tall,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14099,
		},
	},
	{
		"minecraft:polished_blackstone_brick_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			718: 16256,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5887,
		},
	},
	{
		"minecraft:yellow_shulker_box[facing=north]",
		nil,
		NewMapping{
			393: 8241,
			404: 8242,
			477: 8766,
			573: 8766,
			718: 9302,
			4:   3570,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=12,south=side,west=none]",
		nil,
		NewMapping{
			718: 3179,
			393: 2873,
			404: 2874,
			477: 3177,
			573: 3177,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 9437,
			718: 9973,
			477: 9437,
		},
	},
	{
		"minecraft:smooth_red_sandstone_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			477: 10262,
			573: 10262,
			718: 10798,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10721,
			477: 10721,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=4,south=none,west=side]",
		nil,
		NewMapping{
			573: 2963,
			718: 2965,
			393: 2659,
			404: 2660,
			477: 2963,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=14,powered=true]",
		nil,
		NewMapping{
			393: 626,
			404: 626,
			477: 626,
			573: 626,
			718: 627,
		},
	},
	{
		"minecraft:spruce_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 5448,
			718: 5464,
			393: 4944,
			404: 4945,
			477: 5448,
		},
	},
	{
		"minecraft:fire[age=5,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			404: 1310,
			477: 1613,
			573: 1613,
			718: 1614,
			393: 1309,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10664,
			573: 10664,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13435,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=low,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5713,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=tall,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 5744,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=13,south=up,west=side]",
		nil,
		NewMapping{
			718: 3040,
			393: 2734,
			404: 2735,
			477: 3038,
			573: 3038,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=5,powered=true]",
		nil,
		NewMapping{
			393: 308,
			404: 308,
			477: 308,
			573: 308,
			718: 309,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4581,
			404: 4582,
			477: 5085,
			573: 5085,
			718: 5101,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5905,
			404: 5906,
			477: 6412,
			573: 6412,
			718: 6948,
		},
	},
	{
		"minecraft:iron_door[facing=east,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			718: 3857,
			393: 3351,
			404: 3352,
			477: 3855,
			573: 3855,
		},
	},
	{
		"minecraft:iron_door[facing=north,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			573: 3810,
			718: 3812,
			393: 3306,
			404: 3307,
			477: 3810,
		},
	},
	{
		"minecraft:warped_door[facing=south,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15613,
		},
	},
	{
		"minecraft:fire[age=7,east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1378,
			404: 1379,
			477: 1682,
			573: 1682,
			718: 1683,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10528,
			573: 10528,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=low,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11776,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12689,
		},
	},
	{
		"minecraft:pink_banner[rotation=1]",
		nil,
		NewMapping{
			404: 6952,
			477: 7458,
			573: 7458,
			718: 7994,
			393: 6951,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=4,waterlogged=false]",
		nil,
		NewMapping{
			477: 3548,
			573: 3548,
			718: 3550,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=tall,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 6189,
		},
	},
	{
		"minecraft:player_wall_head[facing=north]",
		nil,
		NewMapping{
			393: 5507,
			404: 5508,
			477: 6030,
			573: 6030,
			718: 6566,
		},
	},
	{
		"minecraft:spruce_leaves[distance=6,persistent=true]",
		nil,
		NewMapping{
			477: 168,
			573: 168,
			718: 169,
			393: 168,
			404: 168,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 10241,
			718: 10777,
			477: 10241,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=tall,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 17069,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=tall,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13230,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16505,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=tall,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16544,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13624,
		},
	},
	{
		"minecraft:dandelion",
		nil,
		NewMapping{
			573: 1411,
			718: 1412,
			4:   592,
			393: 1111,
			404: 1111,
			477: 1411,
		},
	},
	{
		"minecraft:purpur_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			393: 7352,
			404: 7353,
			477: 7877,
			573: 7877,
			718: 8413,
			4:   3264,
		},
	},
	{
		"minecraft:stone_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9624,
			573: 9624,
			718: 10160,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=11,powered=false]",
		nil,
		NewMapping{
			393: 621,
			404: 621,
			477: 621,
			573: 621,
			718: 622,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 3722,
			393: 3256,
			404: 3257,
			477: 3720,
			573: 3720,
		},
	},
	{
		"minecraft:jungle_sign[rotation=9,waterlogged=true]",
		nil,
		NewMapping{
			477: 3525,
			573: 3525,
			718: 3527,
		},
	},
	{
		"minecraft:acacia_door[facing=south,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7896,
			404: 7897,
			477: 8421,
			573: 8421,
			718: 8957,
		},
	},
	{
		"minecraft:brain_coral_fan[waterlogged=true]",
		nil,
		NewMapping{
			393: 8556,
			404: 8572,
			477: 9016,
			573: 9016,
			718: 9552,
		},
	},
	{
		"minecraft:moving_piston[facing=west,type=normal]",
		nil,
		NewMapping{
			393: 1105,
			404: 1105,
			477: 1405,
			573: 1405,
			718: 1406,
			4:   580,
		},
	},
	{
		"minecraft:acacia_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 6903,
			573: 6903,
			718: 7439,
			393: 6396,
			404: 6397,
		},
	},
	{
		"minecraft:crimson_sign[rotation=0,waterlogged=true]",
		nil,
		NewMapping{
			718: 15655,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=11,south=none,west=none]",
		nil,
		NewMapping{
			477: 2883,
			573: 2883,
			718: 2885,
			393: 2579,
			404: 2580,
		},
	},
	{
		"minecraft:gray_banner[rotation=9]",
		nil,
		NewMapping{
			393: 6975,
			404: 6976,
			477: 7482,
			573: 7482,
			718: 8018,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=13,powered=false]",
		nil,
		NewMapping{
			477: 825,
			573: 825,
			718: 826,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11912,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=low,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13199,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16121,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=19,powered=false]",
		nil,
		NewMapping{
			393: 637,
			404: 637,
			477: 637,
			573: 637,
			718: 638,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11277,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 17031,
		},
	},
	{
		"minecraft:crimson_sign[rotation=8,waterlogged=true]",
		nil,
		NewMapping{
			718: 15671,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=tall,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13207,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4441,
			404: 4442,
			477: 4945,
			573: 4945,
			718: 4961,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10246,
			573: 10246,
			718: 10782,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=low,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12650,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=tall,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 5738,
		},
	},
	{
		"minecraft:command_block[conditional=false,facing=up]",
		nil,
		NewMapping{
			477: 5638,
			573: 5638,
			718: 5654,
			4:   2193,
			393: 5134,
			404: 5135,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=4,south=none,west=none]",
		nil,
		NewMapping{
			477: 3108,
			573: 3108,
			718: 3110,
			393: 2804,
			404: 2805,
		},
	},
	{
		"minecraft:coarse_dirt",
		nil,
		NewMapping{
			393: 11,
			404: 11,
			477: 11,
			573: 11,
			718: 11,
			4:   49,
		},
	},
	{
		"minecraft:water[level=14]",
		nil,
		NewMapping{
			4:   158,
			393: 48,
			404: 48,
			477: 48,
			573: 48,
			718: 48,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=low,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11352,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 6210,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 10404,
			477: 9868,
			573: 9868,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=low,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12002,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=none,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11210,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13643,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=none,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12838,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=tall,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11386,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=tall,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14643,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=none,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13490,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=south,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7464,
			404: 7465,
			477: 7989,
			573: 7989,
			718: 8525,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=13,south=side,west=up]",
		nil,
		NewMapping{
			393: 2736,
			404: 2737,
			477: 3040,
			573: 3040,
			718: 3042,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=6,south=up,west=side]",
		nil,
		NewMapping{
			404: 2672,
			477: 2975,
			573: 2975,
			718: 2977,
			393: 2671,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16569,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14235,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6726,
			404: 6727,
			477: 7233,
			573: 7233,
			718: 7769,
		},
	},
	{
		"minecraft:fire[age=4,east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1280,
			404: 1281,
			477: 1584,
			573: 1584,
			718: 1585,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10781,
			573: 10781,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=none,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 10873,
		},
	},
	{
		"minecraft:fire[age=10,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1458,
			404: 1459,
			477: 1762,
			573: 1762,
			718: 1763,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 16717,
		},
	},
	{
		"minecraft:jigsaw[facing=west]",
		nil,
		NewMapping{
			477: 11259,
			573: 11275,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=none,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16050,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=none,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16807,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=none,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14465,
		},
	},
	{
		"minecraft:iron_bars[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 4209,
			477: 4712,
			573: 4712,
			718: 4726,
			393: 4208,
		},
	},
	{
		"minecraft:birch_button[face=floor,facing=west,powered=true]",
		nil,
		NewMapping{
			393: 5355,
			404: 5356,
			477: 5862,
			573: 5862,
			718: 6398,
		},
	},
	{
		"minecraft:jungle_log[axis=y]",
		nil,
		NewMapping{
			4:   275,
			393: 82,
			404: 82,
			477: 82,
			573: 82,
			718: 83,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 7219,
			573: 7219,
			718: 7755,
			393: 6712,
			404: 6713,
		},
	},
	{
		"minecraft:birch_door[facing=north,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			404: 7757,
			477: 8281,
			573: 8281,
			718: 8817,
			4:   3107,
			393: 7756,
		},
	},
	{
		"minecraft:birch_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 5554,
			393: 5034,
			404: 5035,
			477: 5538,
			573: 5538,
		},
	},
	{
		"minecraft:scaffolding[bottom=true,distance=7,waterlogged=true]",
		nil,
		NewMapping{
			477: 11113,
			573: 11113,
			718: 14769,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=tall,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12588,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 6611,
			404: 6612,
			477: 7118,
			573: 7118,
			718: 7654,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 6487,
			573: 6487,
			718: 7023,
			393: 5980,
			404: 5981,
		},
	},
	{
		"minecraft:diorite_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10207,
			573: 10207,
			718: 10743,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16661,
		},
	},
	{
		"minecraft:stone_button[face=wall,facing=north,powered=false]",
		nil,
		NewMapping{
			393: 3400,
			404: 3401,
			477: 3904,
			573: 3904,
			718: 3906,
			4:   1236,
		},
	},
	{
		"minecraft:birch_button[face=wall,facing=north,powered=false]",
		nil,
		NewMapping{
			393: 5360,
			404: 5361,
			477: 5867,
			573: 5867,
			718: 6403,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=24,powered=true]",
		nil,
		NewMapping{
			477: 846,
			573: 846,
			718: 847,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5737,
		},
	},
	{
		"minecraft:red_banner[rotation=3]",
		nil,
		NewMapping{
			393: 7081,
			404: 7082,
			477: 7588,
			573: 7588,
			718: 8124,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=west,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3691,
			404: 3692,
			477: 4195,
			573: 4195,
			718: 4209,
		},
	},
	{
		"minecraft:fire[age=15,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1626,
			404: 1627,
			477: 1930,
			573: 1930,
			718: 1931,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4690,
			404: 4691,
			477: 5194,
			573: 5194,
			718: 5210,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=tall,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 6273,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13314,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11101,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=14,south=none,west=up]",
		nil,
		NewMapping{
			573: 2332,
			718: 2334,
			393: 2028,
			404: 2029,
			477: 2332,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9827,
			573: 9827,
			718: 10363,
		},
	},
	{
		"minecraft:crimson_fence_gate[facing=south,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			718: 15267,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=none,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12506,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 9945,
			477: 9409,
			573: 9409,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11854,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=none,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5772,
		},
	},
	{
		"minecraft:fire[age=4,east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1287,
			404: 1288,
			477: 1591,
			573: 1591,
			718: 1592,
		},
	},
	{
		"minecraft:acacia_door[facing=south,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7890,
			404: 7891,
			477: 8415,
			573: 8415,
			718: 8951,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9838,
			573: 9838,
			718: 10374,
		},
	},
	{
		"minecraft:diorite_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 10717,
			477: 10181,
			573: 10181,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=2,south=none,west=none]",
		nil,
		NewMapping{
			393: 2354,
			404: 2355,
			477: 2658,
			573: 2658,
			718: 2660,
		},
	},
	{
		"minecraft:acacia_wall_sign[facing=north,waterlogged=true]",
		nil,
		NewMapping{
			477: 3757,
			573: 3757,
			718: 3759,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=low,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5941,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=low,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11138,
		},
	},
	{
		"minecraft:frosted_ice[age=1]",
		nil,
		NewMapping{
			573: 8714,
			718: 9250,
			4:   3393,
			393: 8189,
			404: 8190,
			477: 8714,
		},
	},
	{
		"minecraft:kelp[age=6]",
		nil,
		NewMapping{
			718: 9476,
			393: 8415,
			404: 8416,
			477: 8940,
			573: 8940,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 5890,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=low,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11461,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=12,south=up,west=side]",
		nil,
		NewMapping{
			404: 1862,
			477: 2165,
			573: 2165,
			718: 2167,
			393: 1861,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6102,
			404: 6103,
			477: 6609,
			573: 6609,
			718: 7145,
		},
	},
	{
		"minecraft:vine[east=true,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 4276,
			404: 4277,
			477: 4780,
			573: 4780,
			718: 4796,
		},
	},
	{
		"minecraft:spruce_sign[rotation=0,waterlogged=false]",
		nil,
		NewMapping{
			477: 3412,
			573: 3412,
			718: 3414,
		},
	},
	{
		"minecraft:oak_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			393: 7261,
			404: 7262,
			477: 7768,
			573: 7768,
			718: 8304,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9246,
			573: 9246,
			718: 9782,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=low,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12337,
		},
	},
	{
		"minecraft:fire[age=15,east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1639,
			404: 1640,
			477: 1943,
			573: 1943,
			718: 1944,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=12,south=none,west=none]",
		nil,
		NewMapping{
			393: 2012,
			404: 2013,
			477: 2316,
			573: 2316,
			718: 2318,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=11,south=up,west=up]",
		nil,
		NewMapping{
			393: 1995,
			404: 1996,
			477: 2299,
			573: 2299,
			718: 2301,
		},
	},
	{
		"minecraft:stone_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9684,
			573: 9684,
			718: 10220,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=1,powered=true]",
		nil,
		NewMapping{
			718: 601,
			393: 600,
			404: 600,
			477: 600,
			573: 600,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=low,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11141,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 3231,
			404: 3232,
			477: 3695,
			573: 3695,
			718: 3697,
		},
	},
	{
		"minecraft:oak_fence[east=true,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 3461,
			404: 3462,
			477: 3965,
			573: 3965,
			718: 3967,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9487,
			573: 9487,
			718: 10023,
		},
	},
	{
		"minecraft:skeleton_skull[rotation=11]",
		nil,
		NewMapping{
			573: 5965,
			718: 6501,
			393: 5462,
			404: 5463,
			477: 5965,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 6087,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=west,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3628,
			404: 3629,
			477: 4132,
			573: 4132,
			718: 4146,
			4:   1550,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 6232,
			477: 6738,
			573: 6738,
			718: 7274,
			393: 6231,
		},
	},
	{
		"minecraft:jungle_sign[rotation=7,waterlogged=false]",
		nil,
		NewMapping{
			573: 3522,
			718: 3524,
			477: 3522,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=tall,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13025,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 16735,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=0,south=side,west=none]",
		nil,
		NewMapping{
			393: 2045,
			404: 2046,
			477: 2349,
			573: 2349,
			718: 2351,
		},
	},
	{
		"minecraft:jungle_door[facing=east,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7853,
			404: 7854,
			477: 8378,
			573: 8378,
			718: 8914,
		},
	},
	{
		"minecraft:lever[face=wall,facing=north,powered=true]",
		nil,
		NewMapping{
			573: 3789,
			718: 3791,
			4:   1116,
			393: 3285,
			404: 3286,
			477: 3789,
		},
	},
	{
		"minecraft:birch_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 5516,
			573: 5516,
			718: 5532,
			393: 5012,
			404: 5013,
		},
	},
	{
		"minecraft:jungle_door[facing=west,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7843,
			404: 7844,
			477: 8368,
			573: 8368,
			718: 8904,
		},
	},
	{
		"minecraft:brick_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4358,
			404: 4359,
			477: 4862,
			573: 4862,
			718: 4878,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16013,
		},
	},
	{
		"minecraft:jungle_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 5090,
			477: 5593,
			573: 5593,
			718: 5609,
			393: 5089,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13646,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=low,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12961,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10855,
			573: 10855,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=none,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14541,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5839,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12310,
		},
	},
	{
		"minecraft:fire[age=8,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			477: 1696,
			573: 1696,
			718: 1697,
			393: 1392,
			404: 1393,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=low,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11470,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=tall,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16867,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11670,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=7,south=side,west=side]",
		nil,
		NewMapping{
			718: 3133,
			393: 2827,
			404: 2828,
			477: 3131,
			573: 3131,
		},
	},
	{
		"minecraft:oak_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 1696,
			404: 1697,
			477: 2000,
			573: 2000,
			718: 2002,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=false,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10970,
			573: 10970,
		},
	},
	{
		"minecraft:lectern[facing=south,has_book=false,powered=false]",
		nil,
		NewMapping{
			477: 11184,
			573: 11184,
			718: 14840,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13857,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=none,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13792,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 10906,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 7708,
			718: 8244,
			393: 7201,
			404: 7202,
			477: 7708,
		},
	},
	{
		"minecraft:fire[age=6,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1357,
			404: 1358,
			477: 1661,
			573: 1661,
			718: 1662,
		},
	},
	{
		"minecraft:gray_bed[facing=south,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 866,
			404: 866,
			477: 1166,
			573: 1166,
			718: 1167,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=20,powered=false]",
		nil,
		NewMapping{
			477: 989,
			573: 989,
			718: 990,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9431,
			573: 9431,
			718: 9967,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14737,
		},
	},
	{
		"minecraft:polished_basalt[axis=z]",
		nil,
		NewMapping{
			718: 4007,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11514,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=3,waterlogged=true]",
		nil,
		NewMapping{
			477: 3545,
			573: 3545,
			718: 3547,
		},
	},
	{
		"minecraft:fire[age=7,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			4:   823,
			393: 1390,
			404: 1391,
			477: 1694,
			573: 1694,
			718: 1695,
		},
	},
	{
		"minecraft:oak_button[face=floor,facing=east,powered=false]",
		nil,
		NewMapping{
			393: 5310,
			404: 5311,
			477: 5817,
			573: 5817,
			718: 6353,
		},
	},
	{
		"minecraft:oak_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 1992,
			393: 1686,
			404: 1687,
			477: 1990,
			573: 1990,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10790,
			573: 10790,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=none,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13062,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=low,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6047,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12329,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=south,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			477: 7923,
			573: 7923,
			718: 8459,
			393: 7398,
			404: 7399,
		},
	},
	{
		"minecraft:acacia_log[axis=z]",
		nil,
		NewMapping{
			573: 86,
			718: 87,
			4:   2600,
			393: 86,
			404: 86,
			477: 86,
		},
	},
	{
		"minecraft:frosted_ice[age=3]",
		nil,
		NewMapping{
			393: 8191,
			404: 8192,
			477: 8716,
			573: 8716,
			718: 9252,
			4:   3395,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9793,
			573: 9793,
			718: 10329,
		},
	},
	{
		"minecraft:cracked_nether_bricks",
		nil,
		NewMapping{
			718: 17102,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12018,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5251,
			404: 5252,
			477: 5755,
			573: 5755,
		},
	},
	{
		"minecraft:iron_bars[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 4189,
			404: 4190,
			477: 4693,
			573: 4693,
			718: 4707,
		},
	},
	{
		"minecraft:fire[age=5,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1297,
			404: 1298,
			477: 1601,
			573: 1601,
			718: 1602,
		},
	},
	{
		"minecraft:fire[age=8,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1395,
			404: 1396,
			477: 1699,
			573: 1699,
			718: 1700,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=6,south=up,west=none]",
		nil,
		NewMapping{
			404: 2817,
			477: 3120,
			573: 3120,
			718: 3122,
			393: 2816,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6060,
			404: 6061,
			477: 6567,
			573: 6567,
			718: 7103,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=none,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13030,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=tall,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13325,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12615,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=14,south=none,west=up]",
		nil,
		NewMapping{
			404: 2317,
			477: 2620,
			573: 2620,
			718: 2622,
			393: 2316,
		},
	},
	{
		"minecraft:fire[age=0,east=false,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1157,
			404: 1158,
			477: 1461,
			573: 1461,
			718: 1462,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=5,south=side,west=side]",
		nil,
		NewMapping{
			393: 2377,
			404: 2378,
			477: 2681,
			573: 2681,
			718: 2683,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9576,
			573: 9576,
			718: 10112,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=tall,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12911,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=east,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3907,
			404: 3908,
			477: 4411,
			573: 4411,
			718: 4425,
		},
	},
	{
		"minecraft:redstone_wall_torch[facing=south,lit=true]",
		nil,
		NewMapping{
			4:   1219,
			393: 3385,
			404: 3386,
			477: 3889,
			573: 3889,
			718: 3891,
		},
	},
	{
		"minecraft:brown_banner[rotation=9]",
		nil,
		NewMapping{
			573: 7562,
			718: 8098,
			393: 7055,
			404: 7056,
			477: 7562,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6666,
			404: 6667,
			477: 7173,
			573: 7173,
			718: 7709,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6783,
			404: 6784,
			477: 7290,
			573: 7290,
			718: 7826,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 5185,
			573: 5185,
			718: 5201,
			4:   2050,
			393: 4681,
			404: 4682,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 3209,
			404: 3210,
			477: 3673,
			573: 3673,
			718: 3675,
		},
	},
	{
		"minecraft:fire[age=1,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			477: 1486,
			573: 1486,
			718: 1487,
			393: 1182,
			404: 1183,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			573: 8394,
			718: 8930,
			393: 7869,
			404: 7870,
			477: 8394,
		},
	},
	{
		"minecraft:polished_blackstone_button[face=ceiling,facing=south,powered=false]",
		nil,
		NewMapping{
			718: 16772,
		},
	},
	{
		"minecraft:fire[age=10,east=true,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			573: 1767,
			718: 1768,
			393: 1463,
			404: 1464,
			477: 1767,
		},
	},
	{
		"minecraft:purpur_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 8113,
			404: 8114,
			477: 8638,
			573: 8638,
			718: 9174,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 10909,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=low,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12638,
		},
	},
	{
		"minecraft:fire[age=11,east=false,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1509,
			404: 1510,
			477: 1813,
			573: 1813,
			718: 1814,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 7690,
			718: 8226,
			393: 7183,
			404: 7184,
			477: 7690,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=west,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			404: 3953,
			477: 4456,
			573: 4456,
			718: 4470,
			393: 3952,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 3988,
			404: 3989,
			477: 4492,
			573: 4492,
			718: 4506,
		},
	},
	{
		"minecraft:oak_sign[rotation=0,waterlogged=false]",
		nil,
		NewMapping{
			404: 3077,
			477: 3380,
			573: 3380,
			718: 3382,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=west,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3891,
			404: 3892,
			477: 4395,
			573: 4395,
			718: 4409,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=1,south=none,west=side]",
		nil,
		NewMapping{
			393: 2056,
			404: 2057,
			477: 2360,
			573: 2360,
			718: 2362,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=tall,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16880,
		},
	},
	{
		"minecraft:oak_button[face=floor,facing=west,powered=true]",
		nil,
		NewMapping{
			393: 5307,
			404: 5308,
			477: 5814,
			573: 5814,
			718: 6350,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10417,
			573: 10417,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16149,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=tall,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5759,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13871,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=low,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12123,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=low,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14372,
		},
	},
	{
		"minecraft:chain_command_block[conditional=true,facing=west]",
		nil,
		NewMapping{
			393: 8179,
			404: 8180,
			477: 8704,
			573: 8704,
			718: 9240,
			4:   3388,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=6,south=side,west=side]",
		nil,
		NewMapping{
			718: 2116,
			393: 1810,
			404: 1811,
			477: 2114,
			573: 2114,
		},
	},
	{
		"minecraft:red_mushroom",
		nil,
		NewMapping{
			718: 1426,
			4:   640,
			393: 1122,
			404: 1122,
			477: 1425,
			573: 1425,
		},
	},
	{
		"minecraft:spruce_door[facing=south,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			718: 8758,
			393: 7697,
			404: 7698,
			477: 8222,
			573: 8222,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11048,
			573: 11048,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=west,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 6528,
			404: 6529,
			477: 7035,
			573: 7035,
			718: 7571,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11716,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=low,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13741,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=north,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			718: 15134,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=false,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 4012,
			404: 4013,
			477: 4516,
			573: 4516,
			718: 4530,
		},
	},
	{
		"minecraft:nether_brick_fence[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 4512,
			477: 5015,
			573: 5015,
			718: 5031,
			393: 4511,
		},
	},
	{
		"minecraft:fire_coral_block",
		nil,
		NewMapping{
			393: 8457,
			404: 8458,
			477: 8982,
			573: 8982,
			718: 9518,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13951,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9848,
			573: 9848,
			718: 10384,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15888,
		},
	},
	{
		"minecraft:fire[age=8,east=true,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1401,
			404: 1402,
			477: 1705,
			573: 1705,
			718: 1706,
		},
	},
	{
		"minecraft:yellow_bed[facing=east,occupied=true,part=foot]",
		nil,
		NewMapping{
			477: 1125,
			573: 1125,
			718: 1126,
			393: 825,
			404: 825,
		},
	},
	{
		"minecraft:fire[age=9,east=false,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			404: 1445,
			477: 1748,
			573: 1748,
			718: 1749,
			393: 1444,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=11,waterlogged=false]",
		nil,
		NewMapping{
			477: 3562,
			573: 3562,
			718: 3564,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 10941,
		},
	},
	{
		"minecraft:weeping_vines[age=13]",
		nil,
		NewMapping{
			718: 15003,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=low,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12553,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=15,south=side,west=up]",
		nil,
		NewMapping{
			404: 2323,
			477: 2626,
			573: 2626,
			718: 2628,
			393: 2322,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=12,south=none,west=side]",
		nil,
		NewMapping{
			573: 3035,
			718: 3037,
			393: 2731,
			404: 2732,
			477: 3035,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16402,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=none,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16036,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=13,south=up,west=none]",
		nil,
		NewMapping{
			393: 3023,
			404: 3024,
			477: 3327,
			573: 3327,
			718: 3329,
		},
	},
	{
		"minecraft:brick_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 4904,
			718: 4920,
			393: 4400,
			404: 4401,
			477: 4904,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=none,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12935,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10654,
			573: 10654,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=24,powered=false]",
		nil,
		NewMapping{
			477: 897,
			573: 897,
			718: 898,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=none,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11089,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16978,
		},
	},
	{
		"minecraft:crimson_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15342,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=low,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12523,
		},
	},
	{
		"minecraft:potted_oxeye_daisy",
		nil,
		NewMapping{
			393: 5282,
			404: 5283,
			477: 5786,
			573: 5786,
			718: 6322,
		},
	},
	{
		"minecraft:terracotta",
		nil,
		NewMapping{
			393: 6839,
			404: 6840,
			477: 7346,
			573: 7346,
			718: 7882,
			4:   2752,
		},
	},
	{
		"minecraft:iron_door[facing=east,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 3355,
			404: 3356,
			477: 3859,
			573: 3859,
			718: 3861,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=15,powered=false]",
		nil,
		NewMapping{
			477: 979,
			573: 979,
			718: 980,
		},
	},
	{
		"minecraft:zombie_wall_head[facing=west]",
		nil,
		NewMapping{
			718: 6548,
			393: 5489,
			404: 5490,
			477: 6012,
			573: 6012,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14559,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=west,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7469,
			404: 7470,
			477: 7994,
			573: 7994,
			718: 8530,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11029,
			573: 11029,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 7406,
			4:   2610,
			393: 6363,
			404: 6364,
			477: 6870,
			573: 6870,
		},
	},
	{
		"minecraft:repeater[delay=3,facing=east,locked=true,powered=true]",
		nil,
		NewMapping{
			718: 4075,
			393: 3557,
			404: 3558,
			477: 4061,
			573: 4061,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 6709,
			718: 7245,
			393: 6202,
			404: 6203,
			477: 6709,
		},
	},
	{
		"minecraft:fire[age=0,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			477: 1440,
			573: 1440,
			718: 1441,
			393: 1136,
			404: 1137,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=tall,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11614,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=7,south=up,west=up]",
		nil,
		NewMapping{
			718: 2265,
			393: 1959,
			404: 1960,
			477: 2263,
			573: 2263,
		},
	},
	{
		"minecraft:moving_piston[facing=down,type=normal]",
		nil,
		NewMapping{
			393: 1109,
			404: 1109,
			477: 1409,
			573: 1409,
			718: 1410,
			4:   576,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=east,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3651,
			404: 3652,
			477: 4155,
			573: 4155,
			718: 4169,
		},
	},
	{
		"minecraft:fire[age=4,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			718: 1583,
			393: 1278,
			404: 1279,
			477: 1582,
			573: 1582,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9158,
			573: 9158,
			718: 9694,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10498,
			573: 10498,
		},
	},
	{
		"minecraft:lava[level=10]",
		nil,
		NewMapping{
			4:   170,
			393: 60,
			404: 60,
			477: 60,
			573: 60,
			718: 60,
		},
	},
	{
		"minecraft:oak_fence[east=true,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 3472,
			404: 3473,
			477: 3976,
			573: 3976,
			718: 3978,
		},
	},
	{
		"minecraft:dead_horn_coral_wall_fan[facing=south,waterlogged=false]",
		nil,
		NewMapping{
			393: 8499,
			404: 8515,
			477: 9059,
			573: 9059,
			718: 9595,
		},
	},
	{
		"minecraft:yellow_banner[rotation=6]",
		nil,
		NewMapping{
			393: 6924,
			404: 6925,
			477: 7431,
			573: 7431,
			718: 7967,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11763,
		},
	},
	{
		"minecraft:gilded_blackstone",
		nil,
		NewMapping{
			718: 16664,
		},
	},
	{
		"minecraft:fire[age=15,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1643,
			404: 1644,
			477: 1947,
			573: 1947,
			718: 1948,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 6862,
			573: 6862,
			718: 7398,
			393: 6355,
			404: 6356,
		},
	},
	{
		"minecraft:jungle_wall_sign[facing=east,waterlogged=true]",
		nil,
		NewMapping{
			477: 3771,
			573: 3771,
			718: 3773,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=none,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11950,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13172,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=12,south=none,west=side]",
		nil,
		NewMapping{
			477: 3179,
			573: 3179,
			718: 3181,
			393: 2875,
			404: 2876,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6366,
			404: 6367,
			477: 6873,
			573: 6873,
			718: 7409,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 1653,
			404: 1654,
			477: 1957,
			573: 1957,
			718: 1959,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=low,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 6044,
		},
	},
	{
		"minecraft:smooth_red_sandstone_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			718: 10800,
			477: 10264,
			573: 10264,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=low,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13619,
		},
	},
	{
		"minecraft:crimson_sign[rotation=1,waterlogged=false]",
		nil,
		NewMapping{
			718: 15658,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5192,
			404: 5193,
			477: 5696,
			573: 5696,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11007,
			573: 11007,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 10166,
			718: 10702,
			477: 10166,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=low,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11011,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 6685,
			477: 7191,
			573: 7191,
			718: 7727,
			393: 6684,
		},
	},
	{
		"minecraft:dead_brain_coral[waterlogged=false]",
		nil,
		NewMapping{
			404: 8463,
			477: 8987,
			573: 8987,
			718: 9523,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9785,
			573: 9785,
			718: 10321,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16014,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=tall,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13130,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13257,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13421,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7874,
			404: 7875,
			477: 8399,
			573: 8399,
			718: 8935,
		},
	},
	{
		"minecraft:piston_head[facing=down,short=true,type=sticky]",
		nil,
		NewMapping{
			404: 1080,
			477: 1380,
			573: 1380,
			718: 1381,
			393: 1080,
		},
	},
	{
		"minecraft:light_blue_bed[facing=east,occupied=true,part=head]",
		nil,
		NewMapping{
			404: 808,
			477: 1108,
			573: 1108,
			718: 1109,
			393: 808,
		},
	},
	{
		"minecraft:diorite_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10195,
			573: 10195,
			718: 10731,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=4,south=up,west=up]",
		nil,
		NewMapping{
			393: 2076,
			404: 2077,
			477: 2380,
			573: 2380,
			718: 2382,
		},
	},
	{
		"minecraft:fire[age=0,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			477: 1443,
			573: 1443,
			718: 1444,
			393: 1139,
			404: 1140,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=tall,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12902,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12283,
		},
	},
	{
		"minecraft:stripped_spruce_wood[axis=y]",
		nil,
		NewMapping{
			393: 130,
			404: 130,
			477: 130,
			573: 130,
			718: 131,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12642,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11745,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11264,
		},
	},
	{
		"minecraft:acacia_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6381,
			404: 6382,
			477: 6888,
			573: 6888,
			718: 7424,
		},
	},
	{
		"minecraft:birch_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 5015,
			404: 5016,
			477: 5519,
			573: 5519,
			718: 5535,
			4:   2161,
		},
	},
	{
		"minecraft:sign[rotation=0,waterlogged=false]",
		nil,
		NewMapping{
			393: 3076,
			4:   1008,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=none,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11764,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=none,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11196,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=3,south=side,west=none]",
		nil,
		NewMapping{
			393: 2648,
			404: 2649,
			477: 2952,
			573: 2952,
			718: 2954,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=13,south=side,west=up]",
		nil,
		NewMapping{
			404: 2881,
			477: 3184,
			573: 3184,
			718: 3186,
			393: 2880,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=15,powered=true]",
		nil,
		NewMapping{
			393: 578,
			404: 578,
			477: 578,
			573: 578,
			718: 579,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=tall,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16125,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=17,powered=false]",
		nil,
		NewMapping{
			718: 734,
			393: 733,
			404: 733,
			477: 733,
			573: 733,
		},
	},
	{
		"minecraft:lectern[facing=east,has_book=false,powered=true]",
		nil,
		NewMapping{
			477: 11191,
			573: 11191,
			718: 14847,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10622,
			573: 10622,
		},
	},
	{
		"minecraft:soul_campfire[facing=south,lit=false,signal_fire=false,waterlogged=true]",
		nil,
		NewMapping{
			718: 14936,
		},
	},
	{
		"minecraft:spruce_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4917,
			404: 4918,
			477: 5421,
			573: 5421,
			718: 5437,
		},
	},
	{
		"minecraft:brick_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 4856,
			393: 4336,
			404: 4337,
			477: 4840,
			573: 4840,
		},
	},
	{
		"minecraft:dark_oak_pressure_plate[powered=false]",
		nil,
		NewMapping{
			393: 3378,
			404: 3379,
			477: 3882,
			573: 3882,
			718: 3884,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10816,
			573: 10816,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4713,
			404: 4714,
			477: 5217,
			573: 5217,
			718: 5233,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=4,south=none,west=side]",
		nil,
		NewMapping{
			393: 2947,
			404: 2948,
			477: 3251,
			573: 3251,
			718: 3253,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5117,
			404: 5118,
			477: 5621,
			573: 5621,
			718: 5637,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=1,south=side,west=none]",
		nil,
		NewMapping{
			393: 2054,
			404: 2055,
			477: 2358,
			573: 2358,
			718: 2360,
		},
	},
	{
		"minecraft:crimson_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15392,
		},
	},
	{
		"minecraft:chest[facing=west,type=right,waterlogged=false]",
		nil,
		NewMapping{
			393: 1745,
			404: 1746,
			477: 2049,
			573: 2049,
			718: 2051,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=4,south=none,west=up]",
		nil,
		NewMapping{
			393: 2802,
			404: 2803,
			477: 3106,
			573: 3106,
			718: 3108,
		},
	},
	{
		"minecraft:fire[age=7,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			718: 1666,
			393: 1361,
			404: 1362,
			477: 1665,
			573: 1665,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11370,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=0,south=none,west=side]",
		nil,
		NewMapping{
			393: 2767,
			404: 2768,
			477: 3071,
			573: 3071,
			718: 3073,
		},
	},
	{
		"minecraft:oak_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			404: 7259,
			477: 7765,
			573: 7765,
			718: 8301,
			4:   2024,
			393: 7258,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10047,
			573: 10047,
			718: 10583,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=9,powered=true]",
		nil,
		NewMapping{
			573: 616,
			718: 617,
			393: 616,
			404: 616,
			477: 616,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12627,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=low,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16500,
		},
	},
	{
		"minecraft:jungle_fence[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 8113,
			573: 8113,
			718: 8649,
			393: 7588,
			404: 7589,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			718: 9092,
			393: 8031,
			404: 8032,
			477: 8556,
			573: 8556,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4428,
			404: 4429,
			477: 4932,
			573: 4932,
			718: 4948,
		},
	},
	{
		"minecraft:fire[age=3,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1259,
			404: 1260,
			477: 1563,
			573: 1563,
			718: 1564,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=none,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14034,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=4,powered=false]",
		nil,
		NewMapping{
			404: 457,
			477: 457,
			573: 457,
			718: 458,
			393: 457,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10462,
			573: 10462,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10707,
			573: 10707,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=low,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13508,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=none,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11973,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=low,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5923,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 17085,
		},
	},
	{
		"minecraft:sugar_cane[age=2]",
		nil,
		NewMapping{
			393: 3444,
			404: 3445,
			477: 3948,
			573: 3948,
			718: 3950,
			4:   1330,
		},
	},
	{
		"minecraft:dark_oak_door[facing=north,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7944,
			404: 7945,
			477: 8469,
			573: 8469,
			718: 9005,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5837,
			404: 5838,
			477: 6344,
			573: 6344,
			718: 6880,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=true,north=true,powered=true,south=true,west=true]",
		nil,
		NewMapping{
			393: 4851,
			404: 4852,
			477: 5355,
			573: 5355,
			718: 5371,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=3,powered=true]",
		nil,
		NewMapping{
			573: 254,
			718: 255,
			393: 254,
			404: 254,
			477: 254,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=low,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14698,
		},
	},
	{
		"minecraft:prismarine_brick_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			573: 7315,
			718: 7851,
			393: 6808,
			404: 6809,
			477: 7315,
		},
	},
	{
		"minecraft:andesite_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9934,
			573: 9934,
			718: 10470,
		},
	},
	{
		"minecraft:cut_red_sandstone_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			718: 8407,
			477: 7871,
			573: 7871,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 6276,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=true,powered=false,south=true,west=true]",
		nil,
		NewMapping{
			404: 4808,
			477: 5311,
			573: 5311,
			718: 5327,
			393: 4807,
		},
	},
	{
		"minecraft:sign[rotation=14,waterlogged=false]",
		nil,
		NewMapping{
			4:   1022,
			393: 3104,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=low,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11133,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=none,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16887,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14361,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=tall,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12360,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=none,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16449,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=3,south=up,west=none]",
		nil,
		NewMapping{
			393: 2789,
			404: 2790,
			477: 3093,
			573: 3093,
			718: 3095,
		},
	},
	{
		"minecraft:brewing_stand[has_bottle_0=true,has_bottle_1=false,has_bottle_2=false]",
		nil,
		NewMapping{
			718: 5136,
			4:   1873,
			393: 4616,
			404: 4617,
			477: 5120,
			573: 5120,
		},
	},
	{
		"minecraft:brewing_stand[has_bottle_0=true,has_bottle_1=true,has_bottle_2=true]",
		nil,
		NewMapping{
			4:   1879,
			393: 4613,
			404: 4614,
			477: 5117,
			573: 5117,
			718: 5133,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			718: 4637,
			393: 4119,
			404: 4120,
			477: 4623,
			573: 4623,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6700,
			404: 6701,
			477: 7207,
			573: 7207,
			718: 7743,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 4463,
			404: 4464,
			477: 4967,
			573: 4967,
			718: 4983,
			4:   1745,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			477: 8405,
			573: 8405,
			718: 8941,
			393: 7880,
			404: 7881,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=12,south=side,west=none]",
		nil,
		NewMapping{
			393: 2297,
			404: 2298,
			477: 2601,
			573: 2601,
			718: 2603,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13300,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 6767,
			477: 7273,
			573: 7273,
			718: 7809,
			393: 6766,
		},
	},
	{
		"minecraft:magenta_shulker_box[facing=east]",
		nil,
		NewMapping{
			718: 9291,
			4:   3541,
			393: 8230,
			404: 8231,
			477: 8755,
			573: 8755,
		},
	},
	{
		"minecraft:heavy_weighted_pressure_plate[power=2]",
		nil,
		NewMapping{
			404: 5622,
			477: 6128,
			573: 6128,
			718: 6664,
			4:   2370,
			393: 5621,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=20,powered=false]",
		nil,
		NewMapping{
			718: 1040,
			477: 1039,
			573: 1039,
		},
	},
	{
		"minecraft:andesite_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10000,
			573: 10000,
			718: 10536,
		},
	},
	{
		"minecraft:jungle_sign[rotation=13,waterlogged=true]",
		nil,
		NewMapping{
			573: 3533,
			718: 3535,
			477: 3533,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 4954,
			718: 4970,
			393: 4450,
			404: 4451,
			477: 4954,
		},
	},
	{
		"minecraft:acacia_leaves[distance=3,persistent=false]",
		nil,
		NewMapping{
			393: 205,
			404: 205,
			477: 205,
			573: 205,
			718: 206,
		},
	},
	{
		"minecraft:glass_pane[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			718: 4741,
			393: 4221,
			404: 4222,
			477: 4725,
			573: 4725,
		},
	},
	{
		"minecraft:iron_door[facing=east,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			404: 3365,
			477: 3868,
			573: 3868,
			718: 3870,
			4:   1140,
			393: 3364,
		},
	},
	{
		"minecraft:creeper_head[rotation=6]",
		nil,
		NewMapping{
			393: 5537,
			404: 5538,
			477: 6040,
			573: 6040,
			718: 6576,
		},
	},
	{
		"minecraft:scaffolding[bottom=true,distance=1,waterlogged=false]",
		nil,
		NewMapping{
			477: 11102,
			573: 11102,
			718: 14758,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=tall,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16019,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=12,powered=false]",
		nil,
		NewMapping{
			477: 773,
			573: 773,
			718: 774,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=low,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12967,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=low,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11257,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6370,
			404: 6371,
			477: 6877,
			573: 6877,
			718: 7413,
		},
	},
	{
		"minecraft:fire[age=9,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1446,
			404: 1447,
			477: 1750,
			573: 1750,
			718: 1751,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13965,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14146,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=false,north=false,powered=true,south=true,west=false]",
		nil,
		NewMapping{
			393: 4844,
			404: 4845,
			477: 5348,
			573: 5348,
			718: 5364,
		},
	},
	{
		"minecraft:chain_command_block[conditional=true,facing=east]",
		nil,
		NewMapping{
			393: 8177,
			404: 8178,
			477: 8702,
			573: 8702,
			718: 9238,
			4:   3389,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16841,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 5647,
			393: 5143,
			404: 5144,
			477: 5647,
		},
	},
	{
		"minecraft:acacia_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 6386,
			477: 6892,
			573: 6892,
			718: 7428,
			393: 6385,
		},
	},
	{
		"minecraft:black_banner[rotation=10]",
		nil,
		NewMapping{
			477: 7611,
			573: 7611,
			718: 8147,
			393: 7104,
			404: 7105,
		},
	},
	{
		"minecraft:repeater[delay=1,facing=south,locked=false,powered=true]",
		nil,
		NewMapping{
			718: 4037,
			4:   1504,
			393: 3519,
			404: 3520,
			477: 4023,
			573: 4023,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10488,
			573: 10488,
		},
	},
	{
		"minecraft:warped_door[facing=north,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15597,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15884,
		},
	},
	{
		"minecraft:oak_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 1670,
			404: 1671,
			477: 1974,
			573: 1974,
			718: 1976,
		},
	},
	{
		"minecraft:magenta_shulker_box[facing=west]",
		nil,
		NewMapping{
			4:   3540,
			393: 8232,
			404: 8233,
			477: 8757,
			573: 8757,
			718: 9293,
		},
	},
	{
		"minecraft:brick_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 4910,
			573: 4910,
			718: 4926,
			393: 4406,
			404: 4407,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=east,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3654,
			404: 3655,
			477: 4158,
			573: 4158,
			718: 4172,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=15,south=none,west=side]",
		nil,
		NewMapping{
			393: 2902,
			404: 2903,
			477: 3206,
			573: 3206,
			718: 3208,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5913,
			404: 5914,
			477: 6420,
			573: 6420,
			718: 6956,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10677,
			573: 10677,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10905,
			573: 10905,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10955,
			573: 10955,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9779,
			573: 9779,
			718: 10315,
		},
	},
	{
		"minecraft:warped_fence[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			718: 15108,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=3,powered=false]",
		nil,
		NewMapping{
			477: 605,
			573: 605,
			718: 606,
			393: 605,
			404: 605,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16447,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=none,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14007,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=south,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			718: 15157,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=west,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			573: 7936,
			718: 8472,
			4:   2953,
			393: 7411,
			404: 7412,
			477: 7936,
		},
	},
	{
		"minecraft:fire[age=8,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1422,
			404: 1423,
			477: 1726,
			573: 1726,
			718: 1727,
			4:   824,
		},
	},
	{
		"minecraft:stone_button[face=ceiling,facing=south,powered=false]",
		nil,
		NewMapping{
			393: 3410,
			404: 3411,
			477: 3914,
			573: 3914,
			718: 3916,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=5,south=up,west=side]",
		nil,
		NewMapping{
			393: 2662,
			404: 2663,
			477: 2966,
			573: 2966,
			718: 2968,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10902,
			573: 10902,
		},
	},
	{
		"minecraft:creeper_head[rotation=14]",
		nil,
		NewMapping{
			393: 5545,
			404: 5546,
			477: 6048,
			573: 6048,
			718: 6584,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=east,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 4465,
			573: 4465,
			718: 4479,
			393: 3961,
			404: 3962,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=3,south=none,west=none]",
		nil,
		NewMapping{
			477: 3243,
			573: 3243,
			718: 3245,
			4:   883,
			393: 2939,
			404: 2940,
		},
	},
	{
		"minecraft:brown_banner[rotation=15]",
		nil,
		NewMapping{
			477: 7568,
			573: 7568,
			718: 8104,
			393: 7061,
			404: 7062,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=south,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3688,
			404: 3689,
			477: 4192,
			573: 4192,
			718: 4206,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9326,
			573: 9326,
			718: 9862,
		},
	},
	{
		"minecraft:warped_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15402,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=tall,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12354,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=tall,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12803,
		},
	},
	{
		"minecraft:acacia_door[facing=east,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7919,
			404: 7920,
			477: 8444,
			573: 8444,
			718: 8980,
			4:   3146,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=low,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16186,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=low,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11553,
		},
	},
	{
		"minecraft:warped_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15462,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=tall,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14410,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=none,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12278,
		},
	},
	{
		"minecraft:zombie_head[rotation=15]",
		nil,
		NewMapping{
			404: 5507,
			477: 6009,
			573: 6009,
			718: 6545,
			393: 5506,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 6120,
			477: 6626,
			573: 6626,
			718: 7162,
			393: 6119,
		},
	},
	{
		"minecraft:cactus[age=5]",
		nil,
		NewMapping{
			4:   1301,
			393: 3430,
			404: 3431,
			477: 3934,
			573: 3934,
			718: 3936,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=low,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5832,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13197,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 7614,
			393: 6571,
			404: 6572,
			477: 7078,
			573: 7078,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 5672,
			573: 5672,
			393: 5168,
			404: 5169,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=5,powered=false]",
		nil,
		NewMapping{
			393: 409,
			404: 409,
			477: 409,
			573: 409,
			718: 410,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=west,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3892,
			404: 3893,
			477: 4396,
			573: 4396,
			718: 4410,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 5036,
			718: 5052,
			393: 4532,
			404: 4533,
			477: 5036,
		},
	},
	{
		"minecraft:quartz_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 5761,
			404: 5762,
			477: 6268,
			573: 6268,
			718: 6804,
			4:   2500,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=tall,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14645,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=low,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13516,
		},
	},
	{
		"minecraft:fire[age=8,east=false,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			718: 1717,
			393: 1412,
			404: 1413,
			477: 1716,
			573: 1716,
		},
	},
	{
		"minecraft:vine[east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			477: 4800,
			573: 4800,
			718: 4816,
			393: 4296,
			404: 4297,
		},
	},
	{
		"minecraft:spruce_button[face=ceiling,facing=south,powered=true]",
		nil,
		NewMapping{
			393: 5345,
			404: 5346,
			477: 5852,
			573: 5852,
			718: 6388,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=none,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11537,
		},
	},
	{
		"minecraft:warped_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15473,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=14,south=side,west=up]",
		nil,
		NewMapping{
			393: 1881,
			404: 1882,
			477: 2185,
			573: 2185,
			718: 2187,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6720,
			404: 6721,
			477: 7227,
			573: 7227,
			718: 7763,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 6702,
			718: 7238,
			393: 6195,
			404: 6196,
			477: 6702,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=14,south=up,west=none]",
		nil,
		NewMapping{
			573: 3336,
			718: 3338,
			393: 3032,
			404: 3033,
			477: 3336,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10977,
			477: 10977,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13338,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=none,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13049,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11743,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6243,
			404: 6244,
			477: 6750,
			573: 6750,
			718: 7286,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=6,powered=false]",
		nil,
		NewMapping{
			573: 761,
			718: 762,
			477: 761,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10042,
			573: 10042,
			718: 10578,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=tall,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14518,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9602,
			573: 9602,
			718: 10138,
		},
	},
	{
		"minecraft:diorite_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 10214,
			573: 10214,
			718: 10750,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=none,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16361,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9805,
			573: 9805,
			718: 10341,
		},
	},
	{
		"minecraft:crimson_door[facing=east,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15577,
		},
	},
	{
		"minecraft:jungle_button[face=ceiling,facing=north,powered=true]",
		nil,
		NewMapping{
			477: 5898,
			573: 5898,
			718: 6434,
			393: 5391,
			404: 5392,
		},
	},
	{
		"minecraft:beetroots[age=0]",
		nil,
		NewMapping{
			393: 8158,
			404: 8159,
			477: 8683,
			573: 8683,
			718: 9219,
			4:   3312,
		},
	},
	{
		"minecraft:orange_bed[facing=west,occupied=false,part=head]",
		nil,
		NewMapping{
			573: 1074,
			718: 1075,
			393: 774,
			404: 774,
			477: 1074,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=19,powered=true]",
		nil,
		NewMapping{
			393: 636,
			404: 636,
			477: 636,
			573: 636,
			718: 637,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=1,south=side,west=side]",
		nil,
		NewMapping{
			573: 2213,
			718: 2215,
			393: 1909,
			404: 1910,
			477: 2213,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11768,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=low,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11672,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=true,north=true,powered=true,south=true,west=true]",
		nil,
		NewMapping{
			393: 4787,
			404: 4788,
			477: 5291,
			573: 5291,
			718: 5307,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6311,
			404: 6312,
			477: 6818,
			573: 6818,
			718: 7354,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10990,
			573: 10990,
		},
	},
	{
		"minecraft:beehive[facing=south,honey_level=5]",
		nil,
		NewMapping{
			573: 11322,
			718: 15811,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=south,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3619,
			404: 3620,
			477: 4123,
			573: 4123,
			718: 4137,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 6729,
			718: 7265,
			393: 6222,
			404: 6223,
			477: 6729,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=16,powered=false]",
		nil,
		NewMapping{
			477: 531,
			573: 531,
			718: 532,
			393: 531,
			404: 531,
		},
	},
	{
		"minecraft:fire[age=1,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			573: 1482,
			718: 1483,
			393: 1178,
			404: 1179,
			477: 1482,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11358,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16067,
		},
	},
	{
		"minecraft:birch_door[facing=south,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7761,
			404: 7762,
			477: 8286,
			573: 8286,
			718: 8822,
		},
	},
	{
		"minecraft:brick_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 4925,
			393: 4405,
			404: 4406,
			477: 4909,
			573: 4909,
		},
	},
	{
		"minecraft:lever[face=floor,facing=east,powered=false]",
		nil,
		NewMapping{
			404: 3285,
			477: 3788,
			573: 3788,
			718: 3790,
			393: 3284,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16797,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13737,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12092,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11883,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=none,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13916,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4698,
			404: 4699,
			477: 5202,
			573: 5202,
			718: 5218,
		},
	},
	{
		"minecraft:green_wall_banner[facing=east]",
		nil,
		NewMapping{
			393: 7165,
			404: 7166,
			477: 7672,
			573: 7672,
			718: 8208,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11681,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16820,
		},
	},
	{
		"minecraft:fire[age=3,east=true,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1244,
			404: 1245,
			477: 1548,
			573: 1548,
			718: 1549,
		},
	},
	{
		"minecraft:carved_pumpkin[facing=north]",
		nil,
		NewMapping{
			393: 3498,
			404: 3499,
			477: 4002,
			573: 4002,
			718: 4016,
			4:   1378,
		},
	},
	{
		"minecraft:oak_pressure_plate[powered=true]",
		nil,
		NewMapping{
			404: 3368,
			477: 3871,
			573: 3871,
			718: 3873,
			4:   1153,
			393: 3367,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=12,powered=false]",
		nil,
		NewMapping{
			718: 274,
			393: 273,
			404: 273,
			477: 273,
			573: 273,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 3666,
			573: 3666,
			718: 3668,
			393: 3202,
			404: 3203,
		},
	},
	{
		"minecraft:pink_bed[facing=south,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 850,
			404: 850,
			477: 1150,
			573: 1150,
			718: 1151,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=14,waterlogged=false]",
		nil,
		NewMapping{
			477: 3568,
			573: 3568,
			718: 3570,
		},
	},
	{
		"minecraft:lime_banner[rotation=13]",
		nil,
		NewMapping{
			573: 7454,
			718: 7990,
			393: 6947,
			404: 6948,
			477: 7454,
		},
	},
	{
		"minecraft:cyan_banner[rotation=4]",
		nil,
		NewMapping{
			477: 7509,
			573: 7509,
			718: 8045,
			393: 7002,
			404: 7003,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6149,
			404: 6150,
			477: 6656,
			573: 6656,
			718: 7192,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=22,powered=false]",
		nil,
		NewMapping{
			404: 443,
			477: 443,
			573: 443,
			718: 444,
			393: 443,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=8,south=up,west=up]",
		nil,
		NewMapping{
			404: 2113,
			477: 2416,
			573: 2416,
			718: 2418,
			393: 2112,
		},
	},
	{
		"minecraft:fire[age=0,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			477: 1467,
			573: 1467,
			718: 1468,
			393: 1163,
			404: 1164,
		},
	},
	{
		"minecraft:light_blue_banner[rotation=3]",
		nil,
		NewMapping{
			477: 7412,
			573: 7412,
			718: 7948,
			393: 6905,
			404: 6906,
		},
	},
	{
		"minecraft:polished_diorite_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			477: 10272,
			573: 10272,
			718: 10808,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=3,south=up,west=none]",
		nil,
		NewMapping{
			393: 1925,
			404: 1926,
			477: 2229,
			573: 2229,
			718: 2231,
		},
	},
	{
		"minecraft:iron_bars[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 4707,
			718: 4721,
			393: 4203,
			404: 4204,
			477: 4707,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=23,powered=false]",
		nil,
		NewMapping{
			393: 495,
			404: 495,
			477: 495,
			573: 495,
			718: 496,
		},
	},
	{
		"minecraft:oak_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 1682,
			404: 1683,
			477: 1986,
			573: 1986,
			718: 1988,
		},
	},
	{
		"minecraft:warped_fence[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			718: 15122,
		},
	},
	{
		"minecraft:crimson_sign[rotation=12,waterlogged=false]",
		nil,
		NewMapping{
			718: 15680,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=low,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11227,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9828,
			573: 9828,
			718: 10364,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9564,
			573: 9564,
			718: 10100,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=none,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13362,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=tall,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13001,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=4,powered=true]",
		nil,
		NewMapping{
			393: 306,
			404: 306,
			477: 306,
			573: 306,
			718: 307,
		},
	},
	{
		"minecraft:turtle_egg[eggs=2,hatch=1]",
		nil,
		NewMapping{
			404: 8442,
			477: 8966,
			573: 8966,
			718: 9502,
			393: 8441,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6449,
			404: 6450,
			477: 6956,
			573: 6956,
			718: 7492,
		},
	},
	{
		"minecraft:fire[age=15,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1627,
			404: 1628,
			477: 1931,
			573: 1931,
			718: 1932,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=none,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5799,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=low,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12107,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10894,
			573: 10894,
		},
	},
	{
		"minecraft:warped_trapdoor[facing=north,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			718: 15205,
		},
	},
	{
		"minecraft:polished_blackstone_brick_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			718: 16259,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=none,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 6116,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2624,
			393: 6483,
			404: 6484,
			477: 6990,
			573: 6990,
			718: 7526,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 6136,
			477: 6642,
			573: 6642,
			718: 7178,
			393: 6135,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=east,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			393: 4325,
			404: 4326,
			477: 4829,
			573: 4829,
			718: 4845,
		},
	},
	{
		"minecraft:acacia_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			718: 8328,
			393: 7285,
			404: 7286,
			477: 7792,
			573: 7792,
		},
	},
	{
		"minecraft:jungle_door[facing=south,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			477: 8361,
			573: 8361,
			718: 8897,
			4:   3121,
			393: 7836,
			404: 7837,
		},
	},
	{
		"minecraft:spruce_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4953,
			404: 4954,
			477: 5457,
			573: 5457,
			718: 5473,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16198,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 16321,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=10,south=up,west=up]",
		nil,
		NewMapping{
			393: 1986,
			404: 1987,
			477: 2290,
			573: 2290,
			718: 2292,
		},
	},
	{
		"minecraft:oak_button[face=wall,facing=east,powered=true]",
		nil,
		NewMapping{
			393: 5317,
			404: 5318,
			477: 5824,
			573: 5824,
			718: 6360,
			4:   2297,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 10414,
			477: 10414,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 17065,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=19,powered=false]",
		nil,
		NewMapping{
			477: 737,
			573: 737,
			718: 738,
			393: 737,
			404: 737,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9550,
			718: 10086,
			477: 9550,
		},
	},
	{
		"minecraft:andesite_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9993,
			573: 9993,
			718: 10529,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16088,
		},
	},
	{
		"minecraft:warped_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15422,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=east,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 6551,
			404: 6552,
			477: 7058,
			573: 7058,
			718: 7594,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5220,
			404: 5221,
			477: 5724,
			573: 5724,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4538,
			404: 4539,
			477: 5042,
			573: 5042,
			718: 5058,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=4,south=up,west=none]",
		nil,
		NewMapping{
			404: 2799,
			477: 3102,
			573: 3102,
			718: 3104,
			393: 2798,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5984,
			404: 5985,
			477: 6491,
			573: 6491,
			718: 7027,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9486,
			573: 9486,
			718: 10022,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=low,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12216,
		},
	},
	{
		"minecraft:warped_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 15419,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			718: 7125,
			393: 6082,
			404: 6083,
			477: 6589,
			573: 6589,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=5,powered=false]",
		nil,
		NewMapping{
			477: 309,
			573: 309,
			718: 310,
			393: 309,
			404: 309,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=tall,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16632,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16837,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6200,
			404: 6201,
			477: 6707,
			573: 6707,
			718: 7243,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=east,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			4:   2675,
			393: 6557,
			404: 6558,
			477: 7064,
			573: 7064,
			718: 7600,
		},
	},
	{
		"minecraft:command_block[conditional=true,facing=south]",
		nil,
		NewMapping{
			393: 5126,
			404: 5127,
			477: 5630,
			573: 5630,
			718: 5646,
			4:   2203,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=none,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16364,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5971,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12780,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=tall,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14421,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=east,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			393: 4324,
			404: 4325,
			477: 4828,
			573: 4828,
			718: 4844,
		},
	},
	{
		"minecraft:fire[age=2,east=true,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1207,
			404: 1208,
			477: 1511,
			573: 1511,
			718: 1512,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=east,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3847,
			404: 3848,
			477: 4351,
			573: 4351,
			718: 4365,
		},
	},
	{
		"minecraft:crimson_fence[east=true,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			718: 15071,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5891,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=tall,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11288,
		},
	},
	{
		"minecraft:iron_door[facing=north,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			573: 3818,
			718: 3820,
			393: 3314,
			404: 3315,
			477: 3818,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=10,south=side,west=up]",
		nil,
		NewMapping{
			393: 2565,
			404: 2566,
			477: 2869,
			573: 2869,
			718: 2871,
		},
	},
	{
		"minecraft:orange_bed[facing=north,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 764,
			404: 764,
			477: 1064,
			573: 1064,
			718: 1065,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10402,
			573: 10402,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11296,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5872,
			404: 5873,
			477: 6379,
			573: 6379,
			718: 6915,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=12,south=up,west=side]",
		nil,
		NewMapping{
			477: 3173,
			573: 3173,
			718: 3175,
			393: 2869,
			404: 2870,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10616,
			573: 10616,
		},
	},
	{
		"minecraft:warped_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 15448,
		},
	},
	{
		"minecraft:basalt[axis=z]",
		nil,
		NewMapping{
			718: 4004,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=18,powered=false]",
		nil,
		NewMapping{
			718: 686,
			393: 685,
			404: 685,
			477: 685,
			573: 685,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10657,
			573: 10657,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12458,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=tall,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5761,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=tall,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12252,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=low,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 10908,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=true,powered=true,south=true,west=false]",
		nil,
		NewMapping{
			393: 4804,
			404: 4805,
			477: 5308,
			573: 5308,
			718: 5324,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=13,south=up,west=up]",
		nil,
		NewMapping{
			477: 2893,
			573: 2893,
			718: 2895,
			393: 2589,
			404: 2590,
		},
	},
	{
		"minecraft:lever[face=floor,facing=west,powered=false]",
		nil,
		NewMapping{
			393: 3282,
			404: 3283,
			477: 3786,
			573: 3786,
			718: 3788,
			4:   1110,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=tall,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16653,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=none,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14141,
		},
	},
	{
		"minecraft:spruce_button[face=wall,facing=south,powered=true]",
		nil,
		NewMapping{
			477: 5844,
			573: 5844,
			718: 6380,
			393: 5337,
			404: 5338,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=8,powered=true]",
		nil,
		NewMapping{
			477: 1014,
			573: 1014,
			718: 1015,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 15999,
		},
	},
	{
		"minecraft:sign[rotation=2,waterlogged=true]",
		nil,
		NewMapping{
			393: 3079,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16206,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=tall,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11071,
		},
	},
	{
		"minecraft:crimson_button[face=floor,facing=south,powered=false]",
		nil,
		NewMapping{
			718: 15482,
		},
	},
	{
		"minecraft:yellow_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			477: 1121,
			573: 1121,
			718: 1122,
			393: 821,
			404: 821,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=15,south=side,west=up]",
		nil,
		NewMapping{
			393: 2034,
			404: 2035,
			477: 2338,
			573: 2338,
			718: 2340,
		},
	},
	{
		"minecraft:chain_command_block[conditional=false,facing=south]",
		nil,
		NewMapping{
			404: 8185,
			477: 8709,
			573: 8709,
			718: 9245,
			4:   3379,
			393: 8184,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=none,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16472,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=none,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12492,
		},
	},
	{
		"minecraft:black_glazed_terracotta[facing=north]",
		nil,
		NewMapping{
			477: 8898,
			573: 8898,
			718: 9434,
			4:   4002,
			393: 8373,
			404: 8374,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=11,south=up,west=none]",
		nil,
		NewMapping{
			393: 3005,
			404: 3006,
			477: 3309,
			573: 3309,
			718: 3311,
		},
	},
	{
		"minecraft:spruce_fence[east=true,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 7518,
			404: 7519,
			477: 8043,
			573: 8043,
			718: 8579,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10065,
			573: 10065,
			718: 10601,
		},
	},
	{
		"minecraft:lava[level=7]",
		nil,
		NewMapping{
			573: 57,
			718: 57,
			4:   167,
			393: 57,
			404: 57,
			477: 57,
		},
	},
	{
		"minecraft:granite_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9913,
			573: 9913,
			718: 10449,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 10390,
			477: 9854,
			573: 9854,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=low,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11999,
		},
	},
	{
		"minecraft:void_air",
		nil,
		NewMapping{
			718: 9665,
			393: 8574,
			404: 8591,
			477: 9129,
			573: 9129,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9309,
			573: 9309,
			718: 9845,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10237,
			573: 10237,
			718: 10773,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=low,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12110,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=west,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			718: 8534,
			4:   3005,
			393: 7473,
			404: 7474,
			477: 7998,
			573: 7998,
		},
	},
	{
		"minecraft:skeleton_skull[rotation=1]",
		nil,
		NewMapping{
			393: 5452,
			404: 5453,
			477: 5955,
			573: 5955,
			718: 6491,
		},
	},
	{
		"minecraft:dark_oak_door[facing=east,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			404: 7993,
			477: 8517,
			573: 8517,
			718: 9053,
			393: 7992,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 5644,
			573: 5644,
			393: 5140,
			404: 5141,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13566,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=none,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14328,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=low,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11551,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=3,south=side,west=side]",
		nil,
		NewMapping{
			573: 2087,
			718: 2089,
			393: 1783,
			404: 1784,
			477: 2087,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=north,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			573: 4230,
			718: 4244,
			393: 3726,
			404: 3727,
			477: 4230,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=14,powered=false]",
		nil,
		NewMapping{
			477: 1027,
			573: 1027,
			718: 1028,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=south,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 16695,
		},
	},
	{
		"minecraft:brewing_stand[has_bottle_0=false,has_bottle_1=false,has_bottle_2=true]",
		nil,
		NewMapping{
			718: 5139,
			4:   1876,
			393: 4619,
			404: 4620,
			477: 5123,
			573: 5123,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=4,south=up,west=side]",
		nil,
		NewMapping{
			393: 1789,
			404: 1790,
			477: 2093,
			573: 2093,
			718: 2095,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=north,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			404: 3730,
			477: 4233,
			573: 4233,
			718: 4247,
			393: 3729,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=east,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7478,
			404: 7479,
			477: 8003,
			573: 8003,
			718: 8539,
		},
	},
	{
		"minecraft:polished_blackstone_brick_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			718: 16257,
		},
	},
	{
		"minecraft:fire[age=1,east=true,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			718: 1472,
			393: 1167,
			404: 1168,
			477: 1471,
			573: 1471,
		},
	},
	{
		"minecraft:composter[level=4]",
		nil,
		NewMapping{
			477: 11266,
			573: 11282,
			718: 15755,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 10017,
			477: 9481,
			573: 9481,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 5870,
			477: 6376,
			573: 6376,
			718: 6912,
			393: 5869,
		},
	},
	{
		"minecraft:red_wall_banner[facing=south]",
		nil,
		NewMapping{
			573: 7674,
			718: 8210,
			393: 7167,
			404: 7168,
			477: 7674,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9143,
			573: 9143,
			718: 9679,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10750,
			573: 10750,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=north,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			573: 4421,
			718: 4435,
			393: 3917,
			404: 3918,
			477: 4421,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=7,powered=true]",
		nil,
		NewMapping{
			573: 1012,
			718: 1013,
			477: 1012,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9223,
			573: 9223,
			718: 9759,
		},
	},
	{
		"minecraft:barrel[facing=south,open=true]",
		nil,
		NewMapping{
			718: 14795,
			477: 11139,
			573: 11139,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 10945,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=2,south=side,west=up]",
		nil,
		NewMapping{
			393: 1773,
			404: 1774,
			477: 2077,
			573: 2077,
			718: 2079,
		},
	},
	{
		"minecraft:light_blue_bed[facing=west,occupied=false,part=foot]",
		nil,
		NewMapping{
			573: 1107,
			718: 1108,
			393: 807,
			404: 807,
			477: 1107,
		},
	},
	{
		"minecraft:oak_sign[rotation=11,waterlogged=false]",
		nil,
		NewMapping{
			477: 3402,
			573: 3402,
			718: 3404,
			404: 3099,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=15,powered=true]",
		nil,
		NewMapping{
			477: 978,
			573: 978,
			718: 979,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			718: 6993,
			393: 5950,
			404: 5951,
			477: 6457,
			573: 6457,
		},
	},
	{
		"minecraft:andesite_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9971,
			573: 9971,
			718: 10507,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10600,
			573: 10600,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 7181,
			404: 7182,
			477: 7688,
			573: 7688,
			718: 8224,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			573: 4652,
			718: 4666,
			393: 4148,
			404: 4149,
			477: 4652,
		},
	},
	{
		"minecraft:dragon_head[rotation=0]",
		nil,
		NewMapping{
			404: 5552,
			477: 6054,
			573: 6054,
			718: 6590,
			393: 5551,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14214,
		},
	},
	{
		"minecraft:crimson_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15382,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=3,south=up,west=none]",
		nil,
		NewMapping{
			477: 3237,
			573: 3237,
			718: 3239,
			393: 2933,
			404: 2934,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10708,
			573: 10708,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10868,
			477: 10868,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9594,
			573: 9594,
			718: 10130,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=tall,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16972,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=none,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11652,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 7154,
			718: 7690,
			393: 6647,
			404: 6648,
			477: 7154,
		},
	},
	{
		"minecraft:oak_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 1694,
			404: 1695,
			477: 1998,
			573: 1998,
			718: 2000,
		},
	},
	{
		"minecraft:diorite_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10215,
			573: 10215,
			718: 10751,
		},
	},
	{
		"minecraft:stone_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9620,
			573: 9620,
			718: 10156,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=south,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			573: 7022,
			718: 7558,
			393: 6515,
			404: 6516,
			477: 7022,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 3252,
			404: 3253,
			477: 3716,
			573: 3716,
			718: 3718,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6435,
			404: 6436,
			477: 6942,
			573: 6942,
			718: 7478,
		},
	},
	{
		"minecraft:brick_stairs[facing=south,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 4866,
			718: 4882,
			393: 4362,
			404: 4363,
			477: 4866,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4927,
			404: 4928,
			477: 5431,
			573: 5431,
			718: 5447,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10140,
			573: 10140,
			718: 10676,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=none,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11654,
		},
	},
	{
		"minecraft:cyan_glazed_terracotta[facing=east]",
		nil,
		NewMapping{
			477: 8877,
			573: 8877,
			718: 9413,
			4:   3907,
			393: 8352,
			404: 8353,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 6757,
			718: 7293,
			393: 6250,
			404: 6251,
			477: 6757,
		},
	},
	{
		"minecraft:moving_piston[facing=down,type=sticky]",
		nil,
		NewMapping{
			477: 1410,
			573: 1410,
			718: 1411,
			4:   584,
			393: 1110,
			404: 1110,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 4067,
			404: 4068,
			477: 4571,
			573: 4571,
			718: 4585,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12448,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=tall,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16531,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=7,south=none,west=none]",
		nil,
		NewMapping{
			573: 2271,
			718: 2273,
			393: 1967,
			404: 1968,
			477: 2271,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12772,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9717,
			573: 9717,
			718: 10253,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6623,
			404: 6624,
			477: 7130,
			573: 7130,
			718: 7666,
		},
	},
	{
		"minecraft:acacia_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 6881,
			573: 6881,
			718: 7417,
			393: 6374,
			404: 6375,
		},
	},
	{
		"minecraft:crafting_table",
		nil,
		NewMapping{
			393: 3050,
			404: 3051,
			477: 3354,
			573: 3354,
			718: 3356,
			4:   928,
		},
	},
	{
		"minecraft:purpur_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 8131,
			404: 8132,
			477: 8656,
			573: 8656,
			718: 9192,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=6,powered=true]",
		nil,
		NewMapping{
			393: 410,
			404: 410,
			477: 410,
			573: 410,
			718: 411,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9566,
			718: 10102,
			477: 9566,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=low,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12431,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11160,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 7765,
			393: 6722,
			404: 6723,
			477: 7229,
			573: 7229,
		},
	},
	{
		"minecraft:creeper_head[rotation=1]",
		nil,
		NewMapping{
			404: 5533,
			477: 6035,
			573: 6035,
			718: 6571,
			393: 5532,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=none,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5874,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 10927,
		},
	},
	{
		"minecraft:iron_bars[east=false,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 4711,
			718: 4725,
			393: 4207,
			404: 4208,
			477: 4711,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=15,south=side,west=side]",
		nil,
		NewMapping{
			393: 2611,
			404: 2612,
			477: 2915,
			573: 2915,
			718: 2917,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13411,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=none,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13899,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5188,
			404: 5189,
			477: 5692,
			573: 5692,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=21,powered=true]",
		nil,
		NewMapping{
			393: 590,
			404: 590,
			477: 590,
			573: 590,
			718: 591,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=7,south=none,west=none]",
		nil,
		NewMapping{
			404: 2112,
			477: 2415,
			573: 2415,
			718: 2417,
			393: 2111,
		},
	},
	{
		"minecraft:dark_prismarine",
		nil,
		NewMapping{
			718: 7603,
			4:   2690,
			393: 6560,
			404: 6561,
			477: 7067,
			573: 7067,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=low,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16185,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16871,
		},
	},
	{
		"minecraft:lava[level=6]",
		nil,
		NewMapping{
			404: 56,
			477: 56,
			573: 56,
			718: 56,
			4:   166,
			393: 56,
		},
	},
	{
		"minecraft:acacia_button[face=ceiling,facing=east,powered=true]",
		nil,
		NewMapping{
			718: 6464,
			393: 5421,
			404: 5422,
			477: 5928,
			573: 5928,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10748,
			573: 10748,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11230,
		},
	},
	{
		"minecraft:quartz_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 5710,
			404: 5711,
			477: 6217,
			573: 6217,
			718: 6753,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=18,powered=false]",
		nil,
		NewMapping{
			477: 435,
			573: 435,
			718: 436,
			393: 435,
			404: 435,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14556,
		},
	},
	{
		"minecraft:crimson_sign[rotation=4,waterlogged=true]",
		nil,
		NewMapping{
			718: 15663,
		},
	},
	{
		"minecraft:vine[east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 4280,
			404: 4281,
			477: 4784,
			573: 4784,
			718: 4800,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9347,
			573: 9347,
			718: 9883,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10725,
			573: 10725,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13822,
		},
	},
	{
		"minecraft:crimson_door[facing=south,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15549,
		},
	},
	{
		"minecraft:fire[age=8,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1420,
			404: 1421,
			477: 1724,
			573: 1724,
			718: 1725,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 3980,
			573: 3980,
			718: 3982,
			393: 3476,
			404: 3477,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=low,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5819,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=low,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16621,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=false,north=true,powered=true,south=true,west=true]",
		nil,
		NewMapping{
			393: 4835,
			404: 4836,
			477: 5339,
			573: 5339,
			718: 5355,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=22,powered=false]",
		nil,
		NewMapping{
			718: 294,
			393: 293,
			404: 293,
			477: 293,
			573: 293,
		},
	},
	{
		"minecraft:bell[attachment=double_wall,facing=west,powered=true]",
		nil,
		NewMapping{
			573: 11226,
			718: 14882,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=low,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13203,
		},
	},
	{
		"minecraft:player_head[rotation=7]",
		nil,
		NewMapping{
			393: 5518,
			404: 5519,
			477: 6021,
			573: 6021,
			718: 6557,
		},
	},
	{
		"minecraft:potatoes[age=7]",
		nil,
		NewMapping{
			718: 6345,
			4:   2279,
			393: 5302,
			404: 5303,
			477: 5809,
			573: 5809,
		},
	},
	{
		"minecraft:jungle_leaves[distance=7,persistent=false]",
		nil,
		NewMapping{
			718: 200,
			393: 199,
			404: 199,
			477: 199,
			573: 199,
		},
	},
	{
		"minecraft:iron_bars[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			718: 4720,
			393: 4202,
			404: 4203,
			477: 4706,
			573: 4706,
		},
	},
	{
		"minecraft:crimson_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15343,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=low,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12641,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16480,
		},
	},
	{
		"minecraft:jigsaw[orientation=up_east]",
		nil,
		NewMapping{
			718: 15743,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=5,south=side,west=side]",
		nil,
		NewMapping{
			573: 2393,
			718: 2395,
			393: 2089,
			404: 2090,
			477: 2393,
		},
	},
	{
		"minecraft:oak_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 1727,
			404: 1728,
			477: 2031,
			573: 2031,
			718: 2033,
		},
	},
	{
		"minecraft:bell[attachment=double_wall,facing=west]",
		nil,
		NewMapping{
			477: 11212,
		},
	},
	{
		"minecraft:diorite_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10182,
			573: 10182,
			718: 10718,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=13,south=side,west=side]",
		nil,
		NewMapping{
			404: 3026,
			477: 3329,
			573: 3329,
			718: 3331,
			393: 3025,
		},
	},
	{
		"minecraft:acacia_button[face=floor,facing=west,powered=false]",
		nil,
		NewMapping{
			393: 5404,
			404: 5405,
			477: 5911,
			573: 5911,
			718: 6447,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=19,powered=false]",
		nil,
		NewMapping{
			477: 1037,
			573: 1037,
			718: 1038,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=tall,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12361,
		},
	},
	{
		"minecraft:crimson_fence[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			718: 15086,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 3251,
			404: 3252,
			477: 3715,
			573: 3715,
			718: 3717,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 11070,
			477: 11070,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=tall,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 6062,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=low,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12984,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=15,powered=false]",
		nil,
		NewMapping{
			573: 629,
			718: 630,
			393: 629,
			404: 629,
			477: 629,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6275,
			404: 6276,
			477: 6782,
			573: 6782,
			718: 7318,
		},
	},
	{
		"minecraft:lectern[facing=east,has_book=true,powered=false]",
		nil,
		NewMapping{
			477: 11190,
			573: 11190,
			718: 14846,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10993,
			573: 10993,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6356,
			718: 6892,
			393: 5849,
			404: 5850,
			477: 6356,
		},
	},
	{
		"minecraft:player_wall_head[facing=east]",
		nil,
		NewMapping{
			393: 5510,
			404: 5511,
			477: 6033,
			573: 6033,
			718: 6569,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=10,south=none,west=side]",
		nil,
		NewMapping{
			393: 1849,
			404: 1850,
			477: 2153,
			573: 2153,
			718: 2155,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=3,powered=false]",
		nil,
		NewMapping{
			404: 405,
			477: 405,
			573: 405,
			718: 406,
			393: 405,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=low,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14382,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=none,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12403,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10793,
			573: 10793,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12775,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=none,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16811,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=none,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12945,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 9941,
			477: 9405,
			573: 9405,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=south,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			718: 15152,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16424,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12147,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 6682,
			404: 6683,
			477: 7189,
			573: 7189,
			718: 7725,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 6842,
			573: 6842,
			718: 7378,
			393: 6335,
			404: 6336,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=12,powered=true]",
		nil,
		NewMapping{
			573: 322,
			718: 323,
			393: 322,
			404: 322,
			477: 322,
		},
	},
	{
		"minecraft:spruce_fence[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 7540,
			404: 7541,
			477: 8065,
			573: 8065,
			718: 8601,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=none,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14655,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10335,
			573: 10335,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10878,
			573: 10878,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=low,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11041,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11298,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12149,
		},
	},
	{
		"minecraft:lava[level=13]",
		nil,
		NewMapping{
			4:   189,
			393: 63,
			404: 63,
			477: 63,
			573: 63,
			718: 63,
		},
	},
	{
		"minecraft:birch_door[facing=west,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			573: 8303,
			718: 8839,
			393: 7778,
			404: 7779,
			477: 8303,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 10307,
			477: 9771,
			573: 9771,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16168,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12778,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=2,powered=true]",
		nil,
		NewMapping{
			718: 303,
			393: 302,
			404: 302,
			477: 302,
			573: 302,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=18,powered=false]",
		nil,
		NewMapping{
			393: 535,
			404: 535,
			477: 535,
			573: 535,
			718: 536,
		},
	},
	{
		"minecraft:purple_terracotta",
		nil,
		NewMapping{
			573: 6321,
			718: 6857,
			4:   2554,
			393: 5814,
			404: 5815,
			477: 6321,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=low,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12444,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 4170,
			404: 4171,
			477: 4674,
			573: 4674,
			718: 4688,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 3269,
			477: 3732,
			573: 3732,
			718: 3734,
			393: 3268,
		},
	},
	{
		"minecraft:oak_planks",
		nil,
		NewMapping{
			477: 15,
			573: 15,
			718: 15,
			4:   80,
			393: 15,
			404: 15,
		},
	},
	{
		"minecraft:crimson_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			718: 15051,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10236,
			573: 10236,
			718: 10772,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 4130,
			404: 4131,
			477: 4634,
			573: 4634,
			718: 4648,
		},
	},
	{
		"minecraft:fire[age=4,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1292,
			404: 1293,
			477: 1596,
			573: 1596,
			718: 1597,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=13,south=none,west=none]",
		nil,
		NewMapping{
			718: 3335,
			4:   893,
			393: 3029,
			404: 3030,
			477: 3333,
			573: 3333,
		},
	},
	{
		"minecraft:piston_head[facing=south,short=true,type=sticky]",
		nil,
		NewMapping{
			393: 1068,
			404: 1068,
			477: 1368,
			573: 1368,
			718: 1369,
		},
	},
	{
		"minecraft:stone_brick_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10667,
			573: 10667,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13856,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=low,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16399,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=5,south=up,west=side]",
		nil,
		NewMapping{
			393: 2374,
			404: 2375,
			477: 2678,
			573: 2678,
			718: 2680,
		},
	},
	{
		"minecraft:fire[age=3,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1238,
			404: 1239,
			477: 1542,
			573: 1542,
			718: 1543,
		},
	},
	{
		"minecraft:dead_bubble_coral_wall_fan[facing=south,waterlogged=false]",
		nil,
		NewMapping{
			393: 8483,
			404: 8499,
			477: 9043,
			573: 9043,
			718: 9579,
		},
	},
	{
		"minecraft:birch_button[face=wall,facing=south,powered=true]",
		nil,
		NewMapping{
			393: 5361,
			404: 5362,
			477: 5868,
			573: 5868,
			718: 6404,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=south,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			477: 4440,
			573: 4440,
			718: 4454,
			393: 3936,
			404: 3937,
		},
	},
	{
		"minecraft:iron_door[facing=south,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			393: 3334,
			404: 3335,
			477: 3838,
			573: 3838,
			718: 3840,
			4:   1137,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=true,north=true,powered=true,south=false,west=false]",
		nil,
		NewMapping{
			393: 4854,
			404: 4855,
			477: 5358,
			573: 5358,
			718: 5374,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5193,
			404: 5194,
			477: 5697,
			573: 5697,
		},
	},
	{
		"minecraft:oak_door[facing=south,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			404: 3126,
			477: 3589,
			573: 3589,
			718: 3591,
			393: 3125,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9545,
			573: 9545,
			718: 10081,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13861,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=south,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 4319,
			393: 3801,
			404: 3802,
			477: 4305,
			573: 4305,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=14,south=side,west=none]",
		nil,
		NewMapping{
			393: 2027,
			404: 2028,
			477: 2331,
			573: 2331,
			718: 2333,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13688,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 6061,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=false,north=false,powered=true,south=true,west=true]",
		nil,
		NewMapping{
			573: 5379,
			718: 5395,
			393: 4875,
			404: 4876,
			477: 5379,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=none,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13247,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 5860,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=low,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16376,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			404: 4117,
			477: 4620,
			573: 4620,
			718: 4634,
			393: 4116,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			404: 3992,
			477: 4495,
			573: 4495,
			718: 4509,
			393: 3991,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=22,powered=false]",
		nil,
		NewMapping{
			573: 943,
			718: 944,
			477: 943,
		},
	},
	{
		"minecraft:bell[attachment=ceiling,facing=east,powered=false]",
		nil,
		NewMapping{
			573: 11213,
			718: 14869,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=west,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			718: 4281,
			393: 3763,
			404: 3764,
			477: 4267,
			573: 4267,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=west,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			718: 15164,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11769,
		},
	},
	{
		"minecraft:crimson_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15383,
		},
	},
	{
		"minecraft:andesite_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 9933,
			718: 10469,
			477: 9933,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9310,
			573: 9310,
			718: 9846,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9198,
			573: 9198,
			718: 9734,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13425,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=north,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 7005,
			573: 7005,
			718: 7541,
			393: 6498,
			404: 6499,
		},
	},
	{
		"minecraft:fire[age=5,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1313,
			404: 1314,
			477: 1617,
			573: 1617,
			718: 1618,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=tall,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16233,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=tall,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16236,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=none,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 17028,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=3,powered=true]",
		nil,
		NewMapping{
			393: 354,
			404: 354,
			477: 354,
			573: 354,
			718: 355,
		},
	},
	{
		"minecraft:white_terracotta",
		nil,
		NewMapping{
			477: 6311,
			573: 6311,
			718: 6847,
			4:   2544,
			393: 5804,
			404: 5805,
		},
	},
	{
		"minecraft:creeper_wall_head[facing=north]",
		nil,
		NewMapping{
			404: 5528,
			477: 6050,
			573: 6050,
			718: 6586,
			393: 5527,
		},
	},
	{
		"minecraft:oak_wall_sign[facing=west,waterlogged=true]",
		nil,
		NewMapping{
			477: 3737,
			573: 3737,
			718: 3739,
			404: 3274,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 15937,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14274,
		},
	},
	{
		"minecraft:pink_banner[rotation=0]",
		nil,
		NewMapping{
			393: 6950,
			404: 6951,
			477: 7457,
			573: 7457,
			718: 7993,
		},
	},
	{
		"minecraft:brown_bed[facing=south,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 945,
			404: 945,
			477: 1245,
			573: 1245,
			718: 1246,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=14,south=none,west=side]",
		nil,
		NewMapping{
			718: 2911,
			393: 2605,
			404: 2606,
			477: 2909,
			573: 2909,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			477: 8589,
			573: 8589,
			718: 9125,
			393: 8064,
			404: 8065,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9826,
			718: 10362,
			477: 9826,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9332,
			573: 9332,
			718: 9868,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=9,powered=true]",
		nil,
		NewMapping{
			718: 867,
			477: 866,
			573: 866,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16897,
		},
	},
	{
		"minecraft:cobblestone_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			393: 7312,
			404: 7313,
			477: 7831,
			573: 7831,
			718: 8367,
			4:   715,
		},
	},
	{
		"minecraft:fire[age=15,east=true,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1620,
			404: 1621,
			477: 1924,
			573: 1924,
			718: 1925,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 7829,
			393: 6786,
			404: 6787,
			477: 7293,
			573: 7293,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13287,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=3,powered=true]",
		nil,
		NewMapping{
			573: 554,
			718: 555,
			393: 554,
			404: 554,
			477: 554,
		},
	},
	{
		"minecraft:blue_glazed_terracotta[facing=north]",
		nil,
		NewMapping{
			4:   3938,
			393: 8357,
			404: 8358,
			477: 8882,
			573: 8882,
			718: 9418,
		},
	},
	{
		"minecraft:birch_door[facing=north,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			477: 8275,
			573: 8275,
			718: 8811,
			393: 7750,
			404: 7751,
		},
	},
	{
		"minecraft:campfire[facing=east,lit=false,signal_fire=false,waterlogged=true]",
		nil,
		NewMapping{
			477: 11246,
			573: 11262,
			718: 14920,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5742,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=none,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16798,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 3713,
			573: 3713,
			718: 3715,
			393: 3249,
			404: 3250,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5163,
			404: 5164,
			477: 5667,
			573: 5667,
		},
	},
	{
		"minecraft:crimson_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15332,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=none,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 6117,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=15,south=side,west=none]",
		nil,
		NewMapping{
			477: 2916,
			573: 2916,
			718: 2918,
			393: 2612,
			404: 2613,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=false,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10522,
			573: 10522,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10357,
			573: 10357,
		},
	},
	{
		"minecraft:target[power=4]",
		nil,
		NewMapping{
			718: 15764,
		},
	},
	{
		"minecraft:oak_door[facing=south,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			718: 3592,
			393: 3126,
			404: 3127,
			477: 3590,
			573: 3590,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12397,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=none,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13578,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=south,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			477: 7986,
			573: 7986,
			718: 8522,
			393: 7461,
			404: 7462,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=0,south=up,west=none]",
		nil,
		NewMapping{
			393: 1754,
			404: 1755,
			477: 2058,
			573: 2058,
			718: 2060,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=low,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12227,
		},
	},
	{
		"minecraft:horn_coral_wall_fan[facing=south,waterlogged=false]",
		nil,
		NewMapping{
			393: 8539,
			404: 8555,
			477: 9099,
			573: 9099,
			718: 9635,
		},
	},
	{
		"minecraft:brewing_stand[has_bottle_0=false,has_bottle_1=true,has_bottle_2=true]",
		nil,
		NewMapping{
			718: 5137,
			4:   1878,
			393: 4617,
			404: 4618,
			477: 5121,
			573: 5121,
		},
	},
	{
		"minecraft:orange_bed[facing=south,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 768,
			404: 768,
			477: 1068,
			573: 1068,
			718: 1069,
		},
	},
	{
		"minecraft:spruce_log[axis=z]",
		nil,
		NewMapping{
			393: 77,
			404: 77,
			477: 77,
			573: 77,
			718: 78,
			4:   281,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10342,
			573: 10342,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=low,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13603,
		},
	},
	{
		"minecraft:cyan_wool",
		nil,
		NewMapping{
			477: 1392,
			573: 1392,
			718: 1393,
			4:   569,
			393: 1092,
			404: 1092,
		},
	},
	{
		"minecraft:stone_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9686,
			573: 9686,
			718: 10222,
		},
	},
	{
		"minecraft:beehive[facing=west,honey_level=2]",
		nil,
		NewMapping{
			573: 11325,
			718: 15814,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 5674,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=0,powered=true]",
		nil,
		NewMapping{
			718: 299,
			393: 298,
			404: 298,
			477: 298,
			573: 298,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=2,south=side,west=side]",
		nil,
		NewMapping{
			718: 2080,
			393: 1774,
			404: 1775,
			477: 2078,
			573: 2078,
		},
	},
	{
		"minecraft:purple_banner[rotation=15]",
		nil,
		NewMapping{
			573: 7536,
			718: 8072,
			393: 7029,
			404: 7030,
			477: 7536,
		},
	},
	{
		"minecraft:brick_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 4353,
			404: 4354,
			477: 4857,
			573: 4857,
			718: 4873,
			4:   1734,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=16,powered=false]",
		nil,
		NewMapping{
			393: 681,
			404: 681,
			477: 681,
			573: 681,
			718: 682,
		},
	},
	{
		"minecraft:red_glazed_terracotta[facing=south]",
		nil,
		NewMapping{
			404: 8371,
			477: 8895,
			573: 8895,
			718: 9431,
			4:   3984,
			393: 8370,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=low,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12321,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=false,powered=false,south=false,west=true]",
		nil,
		NewMapping{
			393: 4817,
			404: 4818,
			477: 5321,
			573: 5321,
			718: 5337,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=14,south=up,west=up]",
		nil,
		NewMapping{
			393: 2022,
			404: 2023,
			477: 2326,
			573: 2326,
			718: 2328,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=none,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16469,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=low,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14710,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13719,
		},
	},
	{
		"minecraft:sign[rotation=7,waterlogged=false]",
		nil,
		NewMapping{
			393: 3090,
			4:   1015,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6336,
			404: 6337,
			477: 6843,
			573: 6843,
			718: 7379,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=13,south=side,west=none]",
		nil,
		NewMapping{
			393: 3026,
			404: 3027,
			477: 3330,
			573: 3330,
			718: 3332,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=low,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14713,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9648,
			573: 9648,
			718: 10184,
		},
	},
	{
		"minecraft:acacia_wall_sign[facing=east,waterlogged=true]",
		nil,
		NewMapping{
			573: 3763,
			718: 3765,
			477: 3763,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=low,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11036,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 16714,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6714,
			404: 6715,
			477: 7221,
			573: 7221,
			718: 7757,
		},
	},
	{
		"minecraft:fire[age=13,east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			718: 1880,
			393: 1575,
			404: 1576,
			477: 1879,
			573: 1879,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 4017,
			404: 4018,
			477: 4521,
			573: 4521,
			718: 4535,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14038,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12530,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=none,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14140,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=none,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12839,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=2,south=up,west=none]",
		nil,
		NewMapping{
			393: 2492,
			404: 2493,
			477: 2796,
			573: 2796,
			718: 2798,
		},
	},
	{
		"minecraft:smooth_red_sandstone_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			477: 10263,
			573: 10263,
			718: 10799,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10829,
			477: 10829,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=none,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11305,
		},
	},
	{
		"minecraft:acacia_button[face=floor,facing=east,powered=false]",
		nil,
		NewMapping{
			573: 5913,
			718: 6449,
			393: 5406,
			404: 5407,
			477: 5913,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=low,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13417,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=tall,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 10971,
		},
	},
	{
		"minecraft:blue_banner[rotation=9]",
		nil,
		NewMapping{
			393: 7039,
			404: 7040,
			477: 7546,
			573: 7546,
			718: 8082,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 3259,
			404: 3260,
			477: 3723,
			573: 3723,
			718: 3725,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=tall,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12136,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=tall,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14209,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 4963,
			4:   1746,
			393: 4443,
			404: 4444,
			477: 4947,
			573: 4947,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9331,
			573: 9331,
			718: 9867,
		},
	},
	{
		"minecraft:crimson_fence[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			718: 15089,
		},
	},
	{
		"minecraft:observer[facing=south,powered=true]",
		nil,
		NewMapping{
			404: 8204,
			477: 8728,
			573: 8728,
			718: 9264,
			4:   3499,
			393: 8203,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 4678,
			477: 5181,
			573: 5181,
			718: 5197,
			393: 4677,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10072,
			573: 10072,
			718: 10608,
		},
	},
	{
		"minecraft:crimson_door[facing=west,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15573,
		},
	},
	{
		"minecraft:purple_shulker_box[facing=south]",
		nil,
		NewMapping{
			404: 8280,
			477: 8804,
			573: 8804,
			718: 9340,
			4:   3667,
			393: 8279,
		},
	},
	{
		"minecraft:warped_button[face=floor,facing=north,powered=true]",
		nil,
		NewMapping{
			718: 15503,
		},
	},
	{
		"minecraft:soul_campfire[facing=west,lit=false,signal_fire=true,waterlogged=false]",
		nil,
		NewMapping{
			718: 14943,
		},
	},
	{
		"minecraft:fire[age=2,east=false,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1220,
			404: 1221,
			477: 1524,
			573: 1524,
			718: 1525,
		},
	},
	{
		"minecraft:campfire[facing=west,lit=false,signal_fire=false,waterlogged=false]",
		nil,
		NewMapping{
			477: 11239,
			573: 11255,
			718: 14913,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 11050,
			477: 11050,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9506,
			573: 9506,
			718: 10042,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=east,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			404: 3649,
			477: 4152,
			573: 4152,
			718: 4166,
			4:   1547,
			393: 3648,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12094,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=none,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14131,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12774,
		},
	},
	{
		"minecraft:twisting_vines[age=1]",
		nil,
		NewMapping{
			718: 15018,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=none,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14025,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=tall,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13542,
		},
	},
	{
		"minecraft:cactus[age=4]",
		nil,
		NewMapping{
			573: 3933,
			718: 3935,
			4:   1300,
			393: 3429,
			404: 3430,
			477: 3933,
		},
	},
	{
		"minecraft:fire[age=2,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1210,
			404: 1211,
			477: 1514,
			573: 1514,
			718: 1515,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 4935,
			404: 4936,
			477: 5439,
			573: 5439,
			718: 5455,
			4:   2145,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10730,
			573: 10730,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=tall,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14189,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14370,
		},
	},
	{
		"minecraft:oak_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 1681,
			404: 1682,
			477: 1985,
			573: 1985,
			718: 1987,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=west,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			477: 7930,
			573: 7930,
			718: 8466,
			393: 7405,
			404: 7406,
		},
	},
	{
		"minecraft:oak_door[facing=west,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			477: 3610,
			573: 3610,
			718: 3612,
			393: 3146,
			404: 3147,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=tall,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13335,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=true,north=false,powered=false,south=true,west=true]",
		nil,
		NewMapping{
			718: 5351,
			393: 4831,
			404: 4832,
			477: 5335,
			573: 5335,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10962,
			573: 10962,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 9672,
			477: 9136,
			573: 9136,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=tall,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12903,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 16715,
		},
	},
	{
		"minecraft:note_block[instrument=iron_xylophone,note=20,powered=true]",
		nil,
		NewMapping{
			477: 788,
			573: 788,
			718: 789,
		},
	},
	{
		"minecraft:gray_shulker_box[facing=east]",
		nil,
		NewMapping{
			393: 8260,
			404: 8261,
			477: 8785,
			573: 8785,
			718: 9321,
			4:   3621,
		},
	},
	{
		"minecraft:cyan_banner[rotation=11]",
		nil,
		NewMapping{
			718: 8052,
			393: 7009,
			404: 7010,
			477: 7516,
			573: 7516,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 7718,
			393: 6675,
			404: 6676,
			477: 7182,
			573: 7182,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10820,
			573: 10820,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=tall,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11813,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=none,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5907,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13426,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=none,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16585,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			573: 8576,
			718: 9112,
			393: 8051,
			404: 8052,
			477: 8576,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=15,south=side,west=up]",
		nil,
		NewMapping{
			393: 2466,
			404: 2467,
			477: 2770,
			573: 2770,
			718: 2772,
		},
	},
	{
		"minecraft:glass_pane[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 4224,
			404: 4225,
			477: 4728,
			573: 4728,
			718: 4744,
		},
	},
	{
		"minecraft:dragon_head[rotation=3]",
		nil,
		NewMapping{
			477: 6057,
			573: 6057,
			718: 6593,
			393: 5554,
			404: 5555,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 7464,
			393: 6421,
			404: 6422,
			477: 6928,
			573: 6928,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=east,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7451,
			404: 7452,
			477: 7976,
			573: 7976,
			718: 8512,
			4:   2971,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9187,
			573: 9187,
			718: 9723,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9705,
			573: 9705,
			718: 10241,
		},
	},
	{
		"minecraft:warped_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15423,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6076,
			404: 6077,
			477: 6583,
			573: 6583,
			718: 7119,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9586,
			573: 9586,
			718: 10122,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=11,powered=false]",
		nil,
		NewMapping{
			477: 971,
			573: 971,
			718: 972,
		},
	},
	{
		"minecraft:warped_fence_gate[facing=west,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15305,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=north,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			573: 4353,
			718: 4367,
			393: 3849,
			404: 3850,
			477: 4353,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10745,
			573: 10745,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 10092,
			718: 10628,
			477: 10092,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=11,powered=true]",
		nil,
		NewMapping{
			393: 270,
			404: 270,
			477: 270,
			573: 270,
			718: 271,
		},
	},
	{
		"minecraft:pink_banner[rotation=11]",
		nil,
		NewMapping{
			393: 6961,
			404: 6962,
			477: 7468,
			573: 7468,
			718: 8004,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 10156,
			718: 10692,
			477: 10156,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=low,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12422,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			404: 6444,
			477: 6950,
			573: 6950,
			718: 7486,
			4:   2626,
			393: 6443,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=true,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 4083,
			404: 4084,
			477: 4587,
			573: 4587,
			718: 4601,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11012,
			573: 11012,
		},
	},
	{
		"minecraft:hopper[enabled=false,facing=north]",
		nil,
		NewMapping{
			477: 6198,
			573: 6198,
			718: 6734,
			4:   2474,
			393: 5691,
			404: 5692,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=11,powered=true]",
		nil,
		NewMapping{
			477: 1020,
			573: 1020,
			718: 1021,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=tall,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11391,
		},
	},
	{
		"minecraft:oak_door[facing=north,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 3121,
			404: 3122,
			477: 3585,
			573: 3585,
			718: 3587,
		},
	},
	{
		"minecraft:dark_oak_door[facing=north,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			718: 9002,
			393: 7941,
			404: 7942,
			477: 8466,
			573: 8466,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=0,south=side,west=none]",
		nil,
		NewMapping{
			393: 2477,
			404: 2478,
			477: 2781,
			573: 2781,
			718: 2783,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=north,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 4234,
			573: 4234,
			718: 4248,
			393: 3730,
			404: 3731,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=low,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16385,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 17077,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=east,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			477: 7941,
			573: 7941,
			718: 8477,
			393: 7416,
			404: 7417,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9856,
			573: 9856,
			718: 10392,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13904,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=none,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12490,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=22,powered=false]",
		nil,
		NewMapping{
			477: 1043,
			573: 1043,
			718: 1044,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=low,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11885,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 16313,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 7643,
			393: 6600,
			404: 6601,
			477: 7107,
			573: 7107,
		},
	},
	{
		"minecraft:purpur_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 8141,
			404: 8142,
			477: 8666,
			573: 8666,
			718: 9202,
		},
	},
	{
		"minecraft:fire[age=3,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1254,
			404: 1255,
			477: 1558,
			573: 1558,
			718: 1559,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10558,
			573: 10558,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=none,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13902,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14392,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=13,powered=true]",
		nil,
		NewMapping{
			393: 474,
			404: 474,
			477: 474,
			573: 474,
			718: 475,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6139,
			404: 6140,
			477: 6646,
			573: 6646,
			718: 7182,
			4:   2569,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 10169,
			718: 10705,
			477: 10169,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=none,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16035,
		},
	},
	{
		"minecraft:crimson_door[facing=south,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			718: 15555,
		},
	},
	{
		"minecraft:warped_sign[rotation=9,waterlogged=true]",
		nil,
		NewMapping{
			718: 15705,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16830,
		},
	},
	{
		"minecraft:birch_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 5558,
			393: 5038,
			404: 5039,
			477: 5542,
			573: 5542,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=15,south=side,west=none]",
		nil,
		NewMapping{
			477: 2484,
			573: 2484,
			718: 2486,
			393: 2180,
			404: 2181,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=15,south=up,west=none]",
		nil,
		NewMapping{
			404: 2610,
			477: 2913,
			573: 2913,
			718: 2915,
			393: 2609,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 5824,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=8,south=side,west=side]",
		nil,
		NewMapping{
			404: 2981,
			477: 3284,
			573: 3284,
			718: 3286,
			393: 2980,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9436,
			573: 9436,
			718: 9972,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 17064,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13506,
		},
	},
	{
		"minecraft:warped_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			718: 15055,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 7607,
			393: 6564,
			404: 6565,
			477: 7071,
			573: 7071,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=7,south=up,west=side]",
		nil,
		NewMapping{
			573: 2552,
			718: 2554,
			393: 2248,
			404: 2249,
			477: 2552,
		},
	},
	{
		"minecraft:jungle_fence[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			718: 8658,
			393: 7597,
			404: 7598,
			477: 8122,
			573: 8122,
		},
	},
	{
		"minecraft:smoker[facing=south,lit=true]",
		nil,
		NewMapping{
			477: 11149,
			573: 11149,
			718: 14805,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6485,
			404: 6486,
			477: 6992,
			573: 6992,
			718: 7528,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16794,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10899,
			477: 10899,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13537,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13390,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13867,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10367,
			573: 10367,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=tall,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12375,
		},
	},
	{
		"minecraft:pumpkin_stem[age=2]",
		nil,
		NewMapping{
			393: 4254,
			404: 4255,
			477: 4758,
			573: 4758,
			718: 4774,
			4:   1666,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=2,south=none,west=up]",
		nil,
		NewMapping{
			393: 2496,
			404: 2497,
			477: 2800,
			573: 2800,
			718: 2802,
		},
	},
	{
		"minecraft:fire[age=15,east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			404: 1635,
			477: 1938,
			573: 1938,
			718: 1939,
			393: 1634,
		},
	},
	{
		"minecraft:jungle_fence[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7605,
			404: 7606,
			477: 8130,
			573: 8130,
			718: 8666,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=low,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5695,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=6,south=up,west=up]",
		nil,
		NewMapping{
			393: 2814,
			404: 2815,
			477: 3118,
			573: 3118,
			718: 3120,
		},
	},
	{
		"minecraft:jungle_door[facing=east,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7865,
			404: 7866,
			477: 8390,
			573: 8390,
			718: 8926,
		},
	},
	{
		"minecraft:oak_sign[rotation=3,waterlogged=true]",
		nil,
		NewMapping{
			718: 3387,
			404: 3082,
			477: 3385,
			573: 3385,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 9213,
			718: 9749,
			477: 9213,
		},
	},
	{
		"minecraft:crimson_fence_gate[facing=south,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15265,
		},
	},
	{
		"minecraft:dark_oak_leaves[distance=6,persistent=false]",
		nil,
		NewMapping{
			477: 225,
			573: 225,
			718: 226,
			393: 225,
			404: 225,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=21,powered=false]",
		nil,
		NewMapping{
			393: 491,
			404: 491,
			477: 491,
			573: 491,
			718: 492,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9590,
			573: 9590,
			718: 10126,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=5,powered=false]",
		nil,
		NewMapping{
			477: 859,
			573: 859,
			718: 860,
		},
	},
	{
		"minecraft:tube_coral",
		nil,
		NewMapping{
			393: 8459,
		},
	},
	{
		"minecraft:jungle_leaves[distance=3,persistent=true]",
		nil,
		NewMapping{
			393: 190,
			404: 190,
			477: 190,
			573: 190,
			718: 191,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11115,
		},
	},
	{
		"minecraft:repeater[delay=3,facing=south,locked=false,powered=false]",
		nil,
		NewMapping{
			718: 4070,
			4:   1496,
			393: 3552,
			404: 3553,
			477: 4056,
			573: 4056,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=12,south=none,west=none]",
		nil,
		NewMapping{
			393: 2300,
			404: 2301,
			477: 2604,
			573: 2604,
			718: 2606,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10943,
			573: 10943,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=low,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 5828,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=low,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13526,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=15,powered=true]",
		nil,
		NewMapping{
			393: 678,
			404: 678,
			477: 678,
			573: 678,
			718: 679,
		},
	},
	{
		"minecraft:spruce_button[face=floor,facing=north,powered=false]",
		nil,
		NewMapping{
			477: 5835,
			573: 5835,
			718: 6371,
			393: 5328,
			404: 5329,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10996,
			573: 10996,
		},
	},
	{
		"minecraft:warped_fence[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			718: 15119,
		},
	},
	{
		"minecraft:spruce_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 5481,
			393: 4961,
			404: 4962,
			477: 5465,
			573: 5465,
		},
	},
	{
		"minecraft:oak_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 2008,
			718: 2010,
			393: 1704,
			404: 1705,
			477: 2008,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16626,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=14,south=side,west=up]",
		nil,
		NewMapping{
			573: 3337,
			718: 3339,
			393: 3033,
			404: 3034,
			477: 3337,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13383,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12070,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=none,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16564,
		},
	},
	{
		"minecraft:blue_banner[rotation=12]",
		nil,
		NewMapping{
			393: 7042,
			404: 7043,
			477: 7549,
			573: 7549,
			718: 8085,
		},
	},
	{
		"minecraft:spruce_button[face=wall,facing=south,powered=false]",
		nil,
		NewMapping{
			393: 5338,
			404: 5339,
			477: 5845,
			573: 5845,
			718: 6381,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14754,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16209,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=low,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13400,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 15891,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=none,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11304,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=12,powered=true]",
		nil,
		NewMapping{
			393: 572,
			404: 572,
			477: 572,
			573: 572,
			718: 573,
		},
	},
	{
		"minecraft:dropper[facing=south,triggered=false]",
		nil,
		NewMapping{
			404: 5798,
			477: 6304,
			573: 6304,
			718: 6840,
			4:   2531,
			393: 5797,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 9857,
			718: 10393,
			477: 9857,
		},
	},
	{
		"minecraft:beehive[facing=east,honey_level=5]",
		nil,
		NewMapping{
			573: 11334,
			718: 15823,
		},
	},
	{
		"minecraft:acacia_sign[rotation=11,waterlogged=true]",
		nil,
		NewMapping{
			477: 3497,
			573: 3497,
			718: 3499,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9491,
			573: 9491,
			718: 10027,
		},
	},
	{
		"minecraft:emerald_ore",
		nil,
		NewMapping{
			404: 4731,
			477: 5234,
			573: 5234,
			718: 5250,
			4:   2064,
			393: 4730,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 6382,
			573: 6382,
			718: 6918,
			393: 5875,
			404: 5876,
		},
	},
	{
		"minecraft:oak_door[facing=east,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 3165,
			404: 3166,
			477: 3629,
			573: 3629,
			718: 3631,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10904,
			477: 10904,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 4547,
			404: 4548,
			477: 5051,
			573: 5051,
			718: 5067,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=2,south=none,west=up]",
		nil,
		NewMapping{
			404: 2209,
			477: 2512,
			573: 2512,
			718: 2514,
			393: 2208,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=false,north=false,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			393: 4813,
			404: 4814,
			477: 5317,
			573: 5317,
			718: 5333,
		},
	},
	{
		"minecraft:birch_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5037,
			404: 5038,
			477: 5541,
			573: 5541,
			718: 5557,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=south,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			718: 4140,
			393: 3622,
			404: 3623,
			477: 4126,
			573: 4126,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=low,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12869,
		},
	},
	{
		"minecraft:andesite_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 10540,
			477: 10004,
			573: 10004,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11767,
		},
	},
	{
		"minecraft:pink_glazed_terracotta[facing=east]",
		nil,
		NewMapping{
			393: 8340,
			404: 8341,
			477: 8865,
			573: 8865,
			718: 9401,
			4:   3859,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6759,
			404: 6760,
			477: 7266,
			573: 7266,
			718: 7802,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=none,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16586,
		},
	},
	{
		"minecraft:warped_sign[rotation=14,waterlogged=false]",
		nil,
		NewMapping{
			718: 15716,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 7178,
			404: 7179,
			477: 7685,
			573: 7685,
			718: 8221,
			4:   2887,
		},
	},
	{
		"minecraft:tripwire_hook[attached=true,facing=south,powered=true]",
		nil,
		NewMapping{
			573: 5245,
			718: 5261,
			4:   2108,
			393: 4741,
			404: 4742,
			477: 5245,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=0,south=up,west=none]",
		nil,
		NewMapping{
			393: 2762,
			404: 2763,
			477: 3066,
			573: 3066,
			718: 3068,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 5994,
			477: 6500,
			573: 6500,
			718: 7036,
			393: 5993,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10910,
			573: 10910,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11251,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=none,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14134,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=none,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 17016,
		},
	},
	{
		"minecraft:cut_red_sandstone",
		nil,
		NewMapping{
			393: 7176,
			404: 7177,
			477: 7683,
			573: 7683,
			718: 8219,
			4:   2866,
		},
	},
	{
		"minecraft:end_rod[facing=east]",
		nil,
		NewMapping{
			4:   3173,
			393: 7998,
			404: 7999,
			477: 8523,
			573: 8523,
			718: 9059,
		},
	},
	{
		"minecraft:brain_coral_wall_fan[facing=north,waterlogged=true]",
		nil,
		NewMapping{
			404: 8528,
			477: 9072,
			573: 9072,
			718: 9608,
			393: 8512,
		},
	},
	{
		"minecraft:stone_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			477: 7800,
			573: 7800,
			718: 8336,
			393: 7293,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=5,south=none,west=side]",
		nil,
		NewMapping{
			393: 2236,
			404: 2237,
			477: 2540,
			573: 2540,
			718: 2542,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6072,
			404: 6073,
			477: 6579,
			573: 6579,
			718: 7115,
		},
	},
	{
		"minecraft:acacia_door[facing=west,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7902,
			404: 7903,
			477: 8427,
			573: 8427,
			718: 8963,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=8,south=side,west=none]",
		nil,
		NewMapping{
			393: 2693,
			404: 2694,
			477: 2997,
			573: 2997,
			718: 2999,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=tall,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16022,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11120,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=tall,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16546,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14343,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 4957,
			393: 4437,
			404: 4438,
			477: 4941,
			573: 4941,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=13,south=up,west=up]",
		nil,
		NewMapping{
			393: 2013,
			404: 2014,
			477: 2317,
			573: 2317,
			718: 2319,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10476,
			573: 10476,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=low,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11015,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16174,
		},
	},
	{
		"minecraft:warped_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15403,
		},
	},
	{
		"minecraft:purpur_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 8103,
			404: 8104,
			477: 8628,
			573: 8628,
			718: 9164,
		},
	},
	{
		"minecraft:fire[age=7,east=true,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			404: 1365,
			477: 1668,
			573: 1668,
			718: 1669,
			393: 1364,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13366,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=none,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14240,
		},
	},
	{
		"minecraft:birch_door[facing=north,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			573: 8279,
			718: 8815,
			4:   3111,
			393: 7754,
			404: 7755,
			477: 8279,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=true,north=false,powered=true,south=true,west=false]",
		nil,
		NewMapping{
			393: 4828,
			404: 4829,
			477: 5332,
			573: 5332,
			718: 5348,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=south,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			404: 3749,
			477: 4252,
			573: 4252,
			718: 4266,
			393: 3748,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 4074,
			404: 4075,
			477: 4578,
			573: 4578,
			718: 4592,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=0,south=up,west=up]",
		nil,
		NewMapping{
			404: 2329,
			477: 2632,
			573: 2632,
			718: 2634,
			393: 2328,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13290,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=none,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11521,
		},
	},
	{
		"minecraft:soul_campfire[facing=east,lit=false,signal_fire=false,waterlogged=false]",
		nil,
		NewMapping{
			718: 14953,
		},
	},
	{
		"minecraft:oak_door[facing=north,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			477: 3580,
			573: 3580,
			718: 3582,
			393: 3116,
			404: 3117,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=9,south=none,west=side]",
		nil,
		NewMapping{
			573: 2576,
			718: 2578,
			393: 2272,
			404: 2273,
			477: 2576,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=east,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3707,
			404: 3708,
			477: 4211,
			573: 4211,
			718: 4225,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10986,
			573: 10986,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=east,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7385,
			404: 7386,
			477: 7910,
			573: 7910,
			718: 8446,
			4:   2943,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=low,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 6246,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5843,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=west,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7374,
			404: 7375,
			477: 7899,
			573: 7899,
			718: 8435,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 6652,
			404: 6653,
			477: 7159,
			573: 7159,
			718: 7695,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 17081,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=none,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5667,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=8,powered=true]",
		nil,
		NewMapping{
			477: 914,
			573: 914,
			718: 915,
		},
	},
	{
		"minecraft:warped_sign[rotation=12,waterlogged=false]",
		nil,
		NewMapping{
			718: 15712,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=7,powered=false]",
		nil,
		NewMapping{
			718: 314,
			393: 313,
			404: 313,
			477: 313,
			573: 313,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 6780,
			477: 7286,
			573: 7286,
			718: 7822,
			393: 6779,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=0,south=none,west=none]",
		nil,
		NewMapping{
			404: 1905,
			477: 2208,
			573: 2208,
			718: 2210,
			393: 1904,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=15,powered=true]",
		nil,
		NewMapping{
			477: 928,
			573: 928,
			718: 929,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 9371,
			718: 9907,
			477: 9371,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13971,
		},
	},
	{
		"minecraft:potted_red_tulip",
		nil,
		NewMapping{
			718: 6318,
			393: 5278,
			404: 5279,
			477: 5782,
			573: 5782,
		},
	},
	{
		"minecraft:jungle_door[facing=east,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7867,
			404: 7868,
			477: 8392,
			573: 8392,
			718: 8928,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=10,south=up,west=none]",
		nil,
		NewMapping{
			477: 2868,
			573: 2868,
			718: 2870,
			393: 2564,
			404: 2565,
		},
	},
	{
		"minecraft:potted_azure_bluet",
		nil,
		NewMapping{
			393: 5277,
			404: 5278,
			477: 5781,
			573: 5781,
			718: 6317,
		},
	},
	{
		"minecraft:hopper[enabled=true,facing=north]",
		nil,
		NewMapping{
			718: 6729,
			4:   2466,
			393: 5686,
			404: 5687,
			477: 6193,
			573: 6193,
		},
	},
	{
		"minecraft:turtle_egg[eggs=1,hatch=0]",
		nil,
		NewMapping{
			393: 8437,
			404: 8438,
			477: 8962,
			573: 8962,
			718: 9498,
		},
	},
	{
		"minecraft:spruce_sign[rotation=0,waterlogged=true]",
		nil,
		NewMapping{
			477: 3411,
			573: 3411,
			718: 3413,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			573: 4682,
			718: 4696,
			393: 4178,
			404: 4179,
			477: 4682,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 3489,
			404: 3490,
			477: 3993,
			573: 3993,
			718: 3995,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16415,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13316,
		},
	},
	{
		"minecraft:oak_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   852,
			393: 1709,
			404: 1710,
			477: 2013,
			573: 2013,
			718: 2015,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6595,
			404: 6596,
			477: 7102,
			573: 7102,
			718: 7638,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16425,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=tall,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11604,
		},
	},
	{
		"minecraft:jungle_sign[rotation=12,waterlogged=true]",
		nil,
		NewMapping{
			573: 3531,
			718: 3533,
			477: 3531,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=none,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12184,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=tall,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13664,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14609,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=4,south=none,west=side]",
		nil,
		NewMapping{
			393: 1939,
			404: 1940,
			477: 2243,
			573: 2243,
			718: 2245,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=2,south=none,west=up]",
		nil,
		NewMapping{
			393: 1920,
			404: 1921,
			477: 2224,
			573: 2224,
			718: 2226,
		},
	},
	{
		"minecraft:light_blue_bed[facing=east,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 811,
			404: 811,
			477: 1111,
			573: 1111,
			718: 1112,
		},
	},
	{
		"minecraft:warped_wall_sign[facing=north,waterlogged=false]",
		nil,
		NewMapping{
			718: 15728,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=low,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14388,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=low,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11772,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12469,
		},
	},
	{
		"minecraft:weeping_vines[age=14]",
		nil,
		NewMapping{
			718: 15004,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 7708,
			393: 6665,
			404: 6666,
			477: 7172,
			573: 7172,
		},
	},
	{
		"minecraft:red_banner[rotation=5]",
		nil,
		NewMapping{
			573: 7590,
			718: 8126,
			393: 7083,
			404: 7084,
			477: 7590,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10723,
			477: 10723,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9470,
			573: 9470,
			718: 10006,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5955,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=true,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10363,
			573: 10363,
		},
	},
	{
		"minecraft:warped_pressure_plate[powered=false]",
		nil,
		NewMapping{
			718: 15062,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=low,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14480,
		},
	},
	{
		"minecraft:spruce_button[face=floor,facing=south,powered=false]",
		nil,
		NewMapping{
			573: 5837,
			718: 6373,
			393: 5330,
			404: 5331,
			477: 5837,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=0,south=none,west=side]",
		nil,
		NewMapping{
			393: 2911,
			404: 2912,
			477: 3215,
			573: 3215,
			718: 3217,
		},
	},
	{
		"minecraft:purpur_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 8087,
			404: 8088,
			477: 8612,
			573: 8612,
			718: 9148,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=2,powered=true]",
		nil,
		NewMapping{
			477: 452,
			573: 452,
			718: 453,
			393: 452,
			404: 452,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9776,
			573: 9776,
			718: 10312,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=low,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13416,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13860,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			404: 4000,
			477: 4503,
			573: 4503,
			718: 4517,
			393: 3999,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=false,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			477: 4548,
			573: 4548,
			718: 4562,
			4:   1592,
			393: 4044,
			404: 4045,
		},
	},
	{
		"minecraft:barrel[facing=west,open=true]",
		nil,
		NewMapping{
			477: 11141,
			573: 11141,
			718: 14797,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=low,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16092,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=1,south=up,west=side]",
		nil,
		NewMapping{
			404: 1763,
			477: 2066,
			573: 2066,
			718: 2068,
			393: 1762,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=east,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7510,
			404: 7511,
			477: 8035,
			573: 8035,
			718: 8571,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=none,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12614,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=low,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13523,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14536,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13477,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=none,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16459,
		},
	},
	{
		"minecraft:kelp[age=10]",
		nil,
		NewMapping{
			477: 8944,
			573: 8944,
			718: 9480,
			393: 8419,
			404: 8420,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=east,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			573: 7942,
			718: 8478,
			4:   2959,
			393: 7417,
			404: 7418,
			477: 7942,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12672,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=low,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6029,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=8,south=up,west=up]",
		nil,
		NewMapping{
			477: 2272,
			573: 2272,
			718: 2274,
			393: 1968,
			404: 1969,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=tall,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13234,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 9872,
			718: 10408,
			477: 9872,
		},
	},
	{
		"minecraft:birch_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 5486,
			573: 5486,
			718: 5502,
			393: 4982,
			404: 4983,
		},
	},
	{
		"minecraft:jungle_door[facing=east,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7864,
			404: 7865,
			477: 8389,
			573: 8389,
			718: 8925,
		},
	},
	{
		"minecraft:fire[age=8,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1407,
			404: 1408,
			477: 1711,
			573: 1711,
			718: 1712,
		},
	},
	{
		"minecraft:stone_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9687,
			573: 9687,
			718: 10223,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=low,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12435,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=tall,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12901,
		},
	},
	{
		"minecraft:brick_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4370,
			404: 4371,
			477: 4874,
			573: 4874,
			718: 4890,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=7,south=none,west=side]",
		nil,
		NewMapping{
			393: 2686,
			404: 2687,
			477: 2990,
			573: 2990,
			718: 2992,
		},
	},
	{
		"minecraft:diorite_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10180,
			573: 10180,
			718: 10716,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=tall,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13009,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14660,
		},
	},
	{
		"minecraft:fire[age=12,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 1548,
			404: 1549,
			477: 1852,
			573: 1852,
			718: 1853,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=6,powered=false]",
		nil,
		NewMapping{
			393: 411,
			404: 411,
			477: 411,
			573: 411,
			718: 412,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=0,south=none,west=side]",
		nil,
		NewMapping{
			393: 2335,
			404: 2336,
			477: 2639,
			573: 2639,
			718: 2641,
		},
	},
	{
		"minecraft:andesite_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9956,
			573: 9956,
			718: 10492,
		},
	},
	{
		"minecraft:crimson_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 15341,
		},
	},
	{
		"minecraft:fire[age=1,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1171,
			404: 1172,
			477: 1475,
			573: 1475,
			718: 1476,
		},
	},
	{
		"minecraft:birch_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5031,
			404: 5032,
			477: 5535,
			573: 5535,
			718: 5551,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=tall,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 15992,
		},
	},
	{
		"minecraft:crimson_door[facing=east,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15585,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=7,south=up,west=side]",
		nil,
		NewMapping{
			573: 3128,
			718: 3130,
			393: 2824,
			404: 2825,
			477: 3128,
		},
	},
	{
		"minecraft:fire[age=15,east=true,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1630,
			404: 1631,
			477: 1934,
			573: 1934,
			718: 1935,
		},
	},
	{
		"minecraft:birch_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4970,
			404: 4971,
			477: 5474,
			573: 5474,
			718: 5490,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=tall,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12258,
		},
	},
	{
		"minecraft:blast_furnace[facing=south,lit=true]",
		nil,
		NewMapping{
			477: 11157,
			573: 11157,
			718: 14813,
		},
	},
	{
		"minecraft:zombie_head[rotation=6]",
		nil,
		NewMapping{
			404: 5498,
			477: 6000,
			573: 6000,
			718: 6536,
			393: 5497,
		},
	},
	{
		"minecraft:repeater[delay=4,facing=east,locked=true,powered=false]",
		nil,
		NewMapping{
			393: 3574,
			404: 3575,
			477: 4078,
			573: 4078,
			718: 4092,
		},
	},
	{
		"minecraft:dark_oak_fence[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 8199,
			718: 8735,
			393: 7674,
			404: 7675,
			477: 8199,
		},
	},
	{
		"minecraft:oak_sign[rotation=14,waterlogged=true]",
		nil,
		NewMapping{
			477: 3407,
			573: 3407,
			718: 3409,
			404: 3104,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 6863,
			573: 6863,
			718: 7399,
			393: 6356,
			404: 6357,
		},
	},
	{
		"minecraft:smoker[facing=east,lit=false]",
		nil,
		NewMapping{
			477: 11154,
			573: 11154,
			718: 14810,
		},
	},
	{
		"minecraft:andesite_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9973,
			573: 9973,
			718: 10509,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16045,
		},
	},
	{
		"minecraft:birch_button[face=ceiling,facing=west,powered=true]",
		nil,
		NewMapping{
			718: 6414,
			393: 5371,
			404: 5372,
			477: 5878,
			573: 5878,
		},
	},
	{
		"minecraft:light_gray_wool",
		nil,
		NewMapping{
			573: 1391,
			718: 1392,
			4:   568,
			393: 1091,
			404: 1091,
			477: 1391,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 7714,
			718: 8250,
			393: 7207,
			404: 7208,
			477: 7714,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 10115,
			477: 9579,
			573: 9579,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=north,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3861,
			404: 3862,
			477: 4365,
			573: 4365,
			718: 4379,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=low,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12550,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=low,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11998,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14144,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9582,
			573: 9582,
			718: 10118,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12125,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12782,
		},
	},
	{
		"minecraft:fire[age=1,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			477: 1472,
			573: 1472,
			718: 1473,
			393: 1168,
			404: 1169,
		},
	},
	{
		"minecraft:spruce_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 5410,
			718: 5426,
			393: 4906,
			404: 4907,
			477: 5410,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=19,powered=true]",
		nil,
		NewMapping{
			573: 586,
			718: 587,
			393: 586,
			404: 586,
			477: 586,
		},
	},
	{
		"minecraft:barrel[facing=east,open=false]",
		nil,
		NewMapping{
			477: 11138,
			573: 11138,
			718: 14794,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=south,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			718: 4264,
			393: 3746,
			404: 3747,
			477: 4250,
			573: 4250,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=10,powered=true]",
		nil,
		NewMapping{
			718: 919,
			477: 918,
			573: 918,
		},
	},
	{
		"minecraft:soul_campfire[facing=south,lit=true,signal_fire=false,waterlogged=true]",
		nil,
		NewMapping{
			718: 14932,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=none,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11002,
		},
	},
	{
		"minecraft:acacia_leaves[distance=5,persistent=false]",
		nil,
		NewMapping{
			393: 209,
			404: 209,
			477: 209,
			573: 209,
			718: 210,
		},
	},
	{
		"minecraft:jungle_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 5550,
			573: 5550,
			718: 5566,
			393: 5046,
			404: 5047,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6001,
			404: 6002,
			477: 6508,
			573: 6508,
			718: 7044,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=16,powered=false]",
		nil,
		NewMapping{
			393: 331,
			404: 331,
			477: 331,
			573: 331,
			718: 332,
		},
	},
	{
		"minecraft:detector_rail[powered=true,shape=ascending_west]",
		nil,
		NewMapping{
			393: 1019,
			404: 1019,
			477: 1319,
			573: 1319,
			718: 1320,
			4:   459,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12459,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=tall,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16879,
		},
	},
	{
		"minecraft:fire[age=6,east=true,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			477: 1636,
			573: 1636,
			718: 1637,
			393: 1332,
			404: 1333,
		},
	},
	{
		"minecraft:warped_fence[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			718: 15114,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 1654,
			404: 1655,
			477: 1958,
			573: 1958,
			718: 1960,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=18,powered=true]",
		nil,
		NewMapping{
			404: 584,
			477: 584,
			573: 584,
			718: 585,
			393: 584,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 8059,
			404: 8060,
			477: 8584,
			573: 8584,
			718: 9120,
		},
	},
	{
		"minecraft:diorite_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10222,
			573: 10222,
			718: 10758,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=3,south=none,west=none]",
		nil,
		NewMapping{
			718: 2525,
			393: 2219,
			404: 2220,
			477: 2523,
			573: 2523,
		},
	},
	{
		"minecraft:soul_sand",
		nil,
		NewMapping{
			477: 3998,
			573: 3998,
			718: 4000,
			4:   1408,
			393: 3494,
			404: 3495,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=none,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11540,
		},
	},
	{
		"minecraft:polished_andesite",
		nil,
		NewMapping{
			404: 7,
			477: 7,
			573: 7,
			718: 7,
			4:   22,
			393: 7,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=true,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			393: 4000,
			404: 4001,
			477: 4504,
			573: 4504,
			718: 4518,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9572,
			573: 9572,
			718: 10108,
		},
	},
	{
		"minecraft:crimson_door[facing=north,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			718: 15536,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=none,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13146,
		},
	},
	{
		"minecraft:crimson_fence_gate[facing=west,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			718: 15274,
		},
	},
	{
		"minecraft:oak_door[facing=west,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			404: 3142,
			477: 3605,
			573: 3605,
			718: 3607,
			393: 3141,
		},
	},
	{
		"minecraft:bell[attachment=ceiling,facing=south,powered=false]",
		nil,
		NewMapping{
			573: 11209,
			718: 14865,
		},
	},
	{
		"minecraft:bell[attachment=double_wall,facing=north,powered=false]",
		nil,
		NewMapping{
			718: 14879,
			573: 11223,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10347,
			573: 10347,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=true,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10747,
			573: 10747,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5778,
		},
	},
	{
		"minecraft:yellow_concrete_powder",
		nil,
		NewMapping{
			393: 8397,
			404: 8398,
			477: 8922,
			573: 8922,
			718: 9458,
			4:   4036,
		},
	},
	{
		"minecraft:brick_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4398,
			404: 4399,
			477: 4902,
			573: 4902,
			718: 4918,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 5127,
			393: 4607,
			404: 4608,
			477: 5111,
			573: 5111,
		},
	},
	{
		"minecraft:dark_oak_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			393: 7292,
			404: 7293,
			477: 7799,
			573: 7799,
			718: 8335,
			4:   2005,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14407,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6532,
			718: 7068,
			393: 6025,
			404: 6026,
			477: 6532,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=12,south=side,west=none]",
		nil,
		NewMapping{
			393: 2009,
			404: 2010,
			477: 2313,
			573: 2313,
			718: 2315,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=low,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12541,
		},
	},
	{
		"minecraft:weeping_vines[age=3]",
		nil,
		NewMapping{
			718: 14993,
		},
	},
	{
		"minecraft:warped_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15427,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=tall,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14530,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=west,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3957,
			404: 3958,
			477: 4461,
			573: 4461,
			718: 4475,
		},
	},
	{
		"minecraft:jungle_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 5063,
			404: 5064,
			477: 5567,
			573: 5567,
			718: 5583,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5968,
			404: 5969,
			477: 6475,
			573: 6475,
			718: 7011,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=13,powered=true]",
		nil,
		NewMapping{
			718: 575,
			393: 574,
			404: 574,
			477: 574,
			573: 574,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=none,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6095,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=none,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13682,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12715,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14169,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=8,south=none,west=side]",
		nil,
		NewMapping{
			393: 2551,
			404: 2552,
			477: 2855,
			573: 2855,
			718: 2857,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=9,south=side,west=up]",
		nil,
		NewMapping{
			404: 2125,
			477: 2428,
			573: 2428,
			718: 2430,
			393: 2124,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 7075,
			718: 7611,
			393: 6568,
			404: 6569,
			477: 7075,
		},
	},
	{
		"minecraft:granite_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9919,
			573: 9919,
			718: 10455,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12378,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 10955,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=none,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16909,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=south,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			404: 3612,
			477: 4115,
			573: 4115,
			718: 4129,
			393: 3611,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 4957,
			573: 4957,
			718: 4973,
			4:   1749,
			393: 4453,
			404: 4454,
		},
	},
	{
		"minecraft:spruce_wall_sign[facing=north,waterlogged=false]",
		nil,
		NewMapping{
			718: 3744,
			477: 3742,
			573: 3742,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13752,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14230,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=north,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			393: 4300,
			404: 4301,
			477: 4804,
			573: 4804,
			718: 4820,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=north,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			404: 3598,
			477: 4101,
			573: 4101,
			718: 4115,
			393: 3597,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 7625,
			393: 6582,
			404: 6583,
			477: 7089,
			573: 7089,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=4,south=none,west=none]",
		nil,
		NewMapping{
			393: 2372,
			404: 2373,
			477: 2676,
			573: 2676,
			718: 2678,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10435,
			573: 10435,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=tall,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11179,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16400,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=5,south=none,west=side]",
		nil,
		NewMapping{
			393: 2956,
			404: 2957,
			477: 3260,
			573: 3260,
			718: 3262,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=4,powered=false]",
		nil,
		NewMapping{
			393: 507,
			404: 507,
			477: 507,
			573: 507,
			718: 508,
		},
	},
	{
		"minecraft:prismarine_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10431,
			573: 10431,
		},
	},
	{
		"minecraft:spruce_wall_sign[facing=south,waterlogged=true]",
		nil,
		NewMapping{
			477: 3743,
			573: 3743,
			718: 3745,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 7468,
			393: 6425,
			404: 6426,
			477: 6932,
			573: 6932,
		},
	},
	{
		"minecraft:crimson_door[facing=north,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15533,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11270,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11802,
		},
	},
	{
		"minecraft:nether_wart[age=0]",
		nil,
		NewMapping{
			393: 4608,
			404: 4609,
			477: 5112,
			573: 5112,
			718: 5128,
			4:   1840,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11533,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=11,south=none,west=none]",
		nil,
		NewMapping{
			718: 2453,
			393: 2147,
			404: 2148,
			477: 2451,
			573: 2451,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=9,south=side,west=side]",
		nil,
		NewMapping{
			393: 1981,
			404: 1982,
			477: 2285,
			573: 2285,
			718: 2287,
		},
	},
	{
		"minecraft:brick_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 4380,
			477: 4883,
			573: 4883,
			718: 4899,
			393: 4379,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=none,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11107,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14413,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16222,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=low,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12538,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14186,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=none,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5883,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14340,
		},
	},
	{
		"minecraft:fire[age=0,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			718: 1443,
			393: 1138,
			404: 1139,
			477: 1442,
			573: 1442,
		},
	},
	{
		"minecraft:dark_oak_pressure_plate[powered=true]",
		nil,
		NewMapping{
			718: 3883,
			393: 3377,
			404: 3378,
			477: 3881,
			573: 3881,
		},
	},
	{
		"minecraft:jungle_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 5583,
			573: 5583,
			718: 5599,
			393: 5079,
			404: 5080,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10469,
			573: 10469,
		},
	},
	{
		"minecraft:spruce_button[face=ceiling,facing=west,powered=false]",
		nil,
		NewMapping{
			393: 5348,
			404: 5349,
			477: 5855,
			573: 5855,
			718: 6391,
		},
	},
	{
		"minecraft:red_bed[facing=south,occupied=true,part=foot]",
		nil,
		NewMapping{
			477: 1277,
			573: 1277,
			718: 1278,
			393: 977,
			404: 977,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 5634,
			393: 5114,
			404: 5115,
			477: 5618,
			573: 5618,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=6,south=up,west=up]",
		nil,
		NewMapping{
			393: 1950,
			404: 1951,
			477: 2254,
			573: 2254,
			718: 2256,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=15,south=up,west=none]",
		nil,
		NewMapping{
			393: 2897,
			404: 2898,
			477: 3201,
			573: 3201,
			718: 3203,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10877,
			573: 10877,
		},
	},
	{
		"minecraft:stonecutter[facing=south]",
		nil,
		NewMapping{
			477: 11195,
			573: 11195,
			718: 14851,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=8,south=up,west=side]",
		nil,
		NewMapping{
			393: 2545,
			404: 2546,
			477: 2849,
			573: 2849,
			718: 2851,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6153,
			404: 6154,
			477: 6660,
			573: 6660,
			718: 7196,
		},
	},
	{
		"minecraft:dispenser[facing=south,triggered=false]",
		nil,
		NewMapping{
			393: 238,
			404: 238,
			477: 238,
			573: 238,
			718: 239,
			4:   371,
		},
	},
	{
		"minecraft:fire[age=2,east=false,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			718: 1527,
			393: 1222,
			404: 1223,
			477: 1526,
			573: 1526,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 10899,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=none,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14460,
		},
	},
	{
		"minecraft:white_wall_banner[facing=east]",
		nil,
		NewMapping{
			393: 7113,
			404: 7114,
			477: 7620,
			573: 7620,
			718: 8156,
			4:   2837,
		},
	},
	{
		"minecraft:fire[age=9,east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1447,
			404: 1448,
			477: 1751,
			573: 1751,
			718: 1752,
		},
	},
	{
		"minecraft:lime_bed[facing=south,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 832,
			404: 832,
			477: 1132,
			573: 1132,
			718: 1133,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=true,north=false,powered=true,south=false,west=false]",
		nil,
		NewMapping{
			404: 4831,
			477: 5334,
			573: 5334,
			718: 5350,
			393: 4830,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 5642,
			393: 5138,
			404: 5139,
			477: 5642,
		},
	},
	{
		"minecraft:white_wool",
		nil,
		NewMapping{
			4:   560,
			393: 1083,
			404: 1083,
			477: 1383,
			573: 1383,
			718: 1384,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=low,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 17051,
		},
	},
	{
		"minecraft:acacia_leaves[distance=2,persistent=true]",
		nil,
		NewMapping{
			477: 202,
			573: 202,
			718: 203,
			4:   2588,
			393: 202,
			404: 202,
		},
	},
	{
		"minecraft:acacia_door[facing=south,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7895,
			404: 7896,
			477: 8420,
			573: 8420,
			718: 8956,
		},
	},
	{
		"minecraft:bamboo[age=1,leaves=none,stage=1]",
		nil,
		NewMapping{
			477: 9123,
			573: 9123,
			718: 9659,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=north,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			718: 15138,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11313,
		},
	},
	{
		"minecraft:fire[age=15,east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			573: 1923,
			718: 1924,
			393: 1619,
			404: 1620,
			477: 1923,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=15,south=up,west=up]",
		nil,
		NewMapping{
			393: 2463,
			404: 2464,
			477: 2767,
			573: 2767,
			718: 2769,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5232,
			404: 5233,
			477: 5736,
			573: 5736,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10112,
			573: 10112,
			718: 10648,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=low,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11890,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12795,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 5105,
			404: 5106,
			477: 5609,
			573: 5609,
			718: 5625,
			4:   2180,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			718: 4677,
			393: 4159,
			404: 4160,
			477: 4663,
			573: 4663,
		},
	},
	{
		"minecraft:birch_sign[rotation=5,waterlogged=true]",
		nil,
		NewMapping{
			477: 3453,
			573: 3453,
			718: 3455,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10995,
			573: 10995,
		},
	},
	{
		"minecraft:birch_door[facing=west,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			477: 8311,
			573: 8311,
			718: 8847,
			4:   3110,
			393: 7786,
			404: 7787,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6281,
		},
	},
	{
		"minecraft:daylight_detector[inverted=true,power=6]",
		nil,
		NewMapping{
			393: 5657,
			404: 5658,
			477: 6164,
			573: 6164,
			718: 6700,
			4:   2854,
		},
	},
	{
		"minecraft:warped_trapdoor[facing=north,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			718: 15204,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11838,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=true,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6312,
			404: 6313,
			477: 6819,
			573: 6819,
			718: 7355,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=4,south=up,west=up]",
		nil,
		NewMapping{
			718: 2670,
			393: 2364,
			404: 2365,
			477: 2668,
			573: 2668,
		},
	},
	{
		"minecraft:grindstone[face=wall,facing=west]",
		nil,
		NewMapping{
			477: 11171,
			573: 11171,
			718: 14827,
		},
	},
	{
		"minecraft:warped_fence[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			718: 15106,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13293,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 5043,
			573: 5043,
			718: 5059,
			4:   1827,
			393: 4539,
			404: 4540,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=1,south=side,west=up]",
		nil,
		NewMapping{
			477: 2788,
			573: 2788,
			718: 2790,
			393: 2484,
			404: 2485,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=2,south=side,west=none]",
		nil,
		NewMapping{
			573: 3087,
			718: 3089,
			393: 2783,
			404: 2784,
			477: 3087,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9274,
			573: 9274,
			718: 9810,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 6271,
		},
	},
	{
		"minecraft:jungle_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 5058,
			404: 5059,
			477: 5562,
			573: 5562,
			718: 5578,
		},
	},
	{
		"minecraft:powered_rail[powered=true,shape=north_south]",
		nil,
		NewMapping{
			393: 1004,
			404: 1004,
			477: 1304,
			573: 1304,
			718: 1305,
			4:   440,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=5,south=none,west=none]",
		nil,
		NewMapping{
			393: 2093,
			404: 2094,
			477: 2397,
			573: 2397,
			718: 2399,
		},
	},
	{
		"minecraft:andesite_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 10003,
			573: 10003,
			718: 10539,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13766,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 5720,
			393: 5216,
			404: 5217,
			477: 5720,
		},
	},
	{
		"minecraft:stripped_acacia_log[axis=y]",
		nil,
		NewMapping{
			393: 100,
			404: 100,
			477: 100,
			573: 100,
			718: 101,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10392,
			573: 10392,
		},
	},
	{
		"minecraft:stone_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9616,
			573: 9616,
			718: 10152,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9754,
			573: 9754,
			718: 10290,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16958,
		},
	},
	{
		"minecraft:warped_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15413,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=none,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11415,
		},
	},
	{
		"minecraft:black_banner[rotation=5]",
		nil,
		NewMapping{
			393: 7099,
			404: 7100,
			477: 7606,
			573: 7606,
			718: 8142,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=south,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			404: 3811,
			477: 4314,
			573: 4314,
			718: 4328,
			393: 3810,
		},
	},
	{
		"minecraft:chorus_flower[age=4]",
		nil,
		NewMapping{
			718: 9132,
			4:   3204,
			393: 8071,
			404: 8072,
			477: 8596,
			573: 8596,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9657,
			573: 9657,
			718: 10193,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=low,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11884,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 4932,
			477: 5435,
			573: 5435,
			718: 5451,
			393: 4931,
		},
	},
	{
		"minecraft:fire[age=10,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1469,
			404: 1470,
			477: 1773,
			573: 1773,
			718: 1774,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9215,
			573: 9215,
			718: 9751,
		},
	},
	{
		"minecraft:campfire[facing=south,lit=true,signal_fire=false,waterlogged=true]",
		nil,
		NewMapping{
			573: 11242,
			718: 14900,
			477: 11226,
		},
	},
	{
		"minecraft:iron_door[facing=west,half=lower,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			573: 3853,
			718: 3855,
			393: 3349,
			404: 3350,
			477: 3853,
		},
	},
	{
		"minecraft:quartz_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 5745,
			477: 6251,
			573: 6251,
			718: 6787,
			393: 5744,
		},
	},
	{
		"minecraft:prismarine_brick_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			404: 6812,
			477: 7318,
			573: 7318,
			718: 7854,
			393: 6811,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=10,powered=false]",
		nil,
		NewMapping{
			718: 620,
			393: 619,
			404: 619,
			477: 619,
			573: 619,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=true,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			477: 4564,
			573: 4564,
			718: 4578,
			393: 4060,
			404: 4061,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=11,south=up,west=side]",
		nil,
		NewMapping{
			393: 1852,
			404: 1853,
			477: 2156,
			573: 2156,
			718: 2158,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12315,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=none,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13036,
		},
	},
	{
		"minecraft:purpur_block",
		nil,
		NewMapping{
			393: 8073,
			404: 8074,
			477: 8598,
			573: 8598,
			718: 9134,
			4:   3216,
		},
	},
	{
		"minecraft:jungle_button[face=floor,facing=west,powered=true]",
		nil,
		NewMapping{
			477: 5886,
			573: 5886,
			718: 6422,
			393: 5379,
			404: 5380,
		},
	},
	{
		"minecraft:pink_shulker_box[facing=east]",
		nil,
		NewMapping{
			4:   3605,
			393: 8254,
			404: 8255,
			477: 8779,
			573: 8779,
			718: 9315,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=12,powered=false]",
		nil,
		NewMapping{
			477: 923,
			573: 923,
			718: 924,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 15959,
		},
	},
	{
		"minecraft:pumpkin_stem[age=6]",
		nil,
		NewMapping{
			393: 4258,
			404: 4259,
			477: 4762,
			573: 4762,
			718: 4778,
			4:   1670,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=13,south=up,west=none]",
		nil,
		NewMapping{
			477: 3183,
			573: 3183,
			718: 3185,
			393: 2879,
			404: 2880,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9606,
			718: 10142,
			477: 9606,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10873,
			573: 10873,
		},
	},
	{
		"minecraft:jungle_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 5566,
			718: 5582,
			393: 5062,
			404: 5063,
			477: 5566,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10944,
			573: 10944,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 10701,
			477: 10165,
			573: 10165,
		},
	},
	{
		"minecraft:bell[attachment=ceiling,facing=north,powered=true]",
		nil,
		NewMapping{
			573: 11206,
			718: 14862,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=1,south=up,west=side]",
		nil,
		NewMapping{
			477: 2210,
			573: 2210,
			718: 2212,
			393: 1906,
			404: 1907,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=true,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			404: 4065,
			477: 4568,
			573: 4568,
			718: 4582,
			393: 4064,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=11,powered=true]",
		nil,
		NewMapping{
			573: 720,
			718: 721,
			393: 720,
			404: 720,
			477: 720,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5826,
			404: 5827,
			477: 6333,
			573: 6333,
			718: 6869,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=tall,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12377,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14061,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11512,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=15,powered=true]",
		nil,
		NewMapping{
			718: 329,
			393: 328,
			404: 328,
			477: 328,
			573: 328,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=east,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			573: 4216,
			718: 4230,
			393: 3712,
			404: 3713,
			477: 4216,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=2,south=up,west=side]",
		nil,
		NewMapping{
			718: 2077,
			393: 1771,
			404: 1772,
			477: 2075,
			573: 2075,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6580,
			573: 6580,
			718: 7116,
			393: 6073,
			404: 6074,
		},
	},
	{
		"minecraft:warped_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15417,
		},
	},
	{
		"minecraft:hopper[enabled=false,facing=east]",
		nil,
		NewMapping{
			477: 6201,
			573: 6201,
			718: 6737,
			4:   2477,
			393: 5694,
			404: 5695,
		},
	},
	{
		"minecraft:nether_portal[axis=z]",
		nil,
		NewMapping{
			393: 3497,
			404: 3498,
			477: 4001,
			573: 4001,
			718: 4015,
			4:   1442,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=7,south=side,west=side]",
		nil,
		NewMapping{
			393: 2395,
			404: 2396,
			477: 2699,
			573: 2699,
			718: 2701,
		},
	},
	{
		"minecraft:bamboo[age=0,leaves=large,stage=0]",
		nil,
		NewMapping{
			477: 9120,
			573: 9120,
			718: 9656,
		},
	},
	{
		"minecraft:oak_stairs[facing=east,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 1711,
			404: 1712,
			477: 2015,
			573: 2015,
			718: 2017,
		},
	},
	{
		"minecraft:iron_bars[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			718: 4706,
			393: 4188,
			404: 4189,
			477: 4692,
			573: 4692,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=south,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3745,
			404: 3746,
			477: 4249,
			573: 4249,
			718: 4263,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13579,
		},
	},
	{
		"minecraft:crimson_fence_gate[facing=south,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			718: 15270,
		},
	},
	{
		"minecraft:birch_fence_gate[facing=north,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			404: 7397,
			477: 7921,
			573: 7921,
			718: 8457,
			4:   2946,
			393: 7396,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=5,south=up,west=side]",
		nil,
		NewMapping{
			393: 2518,
			404: 2519,
			477: 2822,
			573: 2822,
			718: 2824,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=14,south=up,west=none]",
		nil,
		NewMapping{
			393: 2024,
			404: 2025,
			477: 2328,
			573: 2328,
			718: 2330,
		},
	},
	{
		"minecraft:warped_fence_gate[facing=east,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			718: 15311,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			404: 4151,
			477: 4654,
			573: 4654,
			718: 4668,
			393: 4150,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 6058,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14183,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=low,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13615,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=12,south=side,west=side]",
		nil,
		NewMapping{
			718: 3178,
			393: 2872,
			404: 2873,
			477: 3176,
			573: 3176,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=none,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16812,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5802,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13065,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 4964,
			573: 4964,
			718: 4980,
			393: 4460,
			404: 4461,
		},
	},
	{
		"minecraft:grindstone[face=wall,facing=east]",
		nil,
		NewMapping{
			477: 11172,
			573: 11172,
			718: 14828,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9709,
			573: 9709,
			718: 10245,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11427,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=2,powered=true]",
		nil,
		NewMapping{
			404: 552,
			477: 552,
			573: 552,
			718: 553,
			393: 552,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=9,powered=true]",
		nil,
		NewMapping{
			718: 517,
			393: 516,
			404: 516,
			477: 516,
			573: 516,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6340,
			404: 6341,
			477: 6847,
			573: 6847,
			718: 7383,
		},
	},
	{
		"minecraft:spruce_sapling[stage=0]",
		nil,
		NewMapping{
			393: 23,
			404: 23,
			477: 23,
			573: 23,
			718: 23,
			4:   97,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=south,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			718: 15153,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=low,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12544,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14094,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=none,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16470,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13302,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13310,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=none,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14249,
		},
	},
	{
		"minecraft:fire[age=9,east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			573: 1744,
			718: 1745,
			393: 1440,
			404: 1441,
			477: 1744,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6033,
			404: 6034,
			477: 6540,
			573: 6540,
			718: 7076,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10739,
			573: 10739,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11486,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5675,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=true,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 4124,
			404: 4125,
			477: 4628,
			573: 4628,
			718: 4642,
		},
	},
	{
		"minecraft:fire_coral_wall_fan[facing=west,waterlogged=false]",
		nil,
		NewMapping{
			393: 8533,
			404: 8549,
			477: 9093,
			573: 9093,
			718: 9629,
		},
	},
	{
		"minecraft:glass_pane[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 4722,
			573: 4722,
			718: 4738,
			393: 4218,
			404: 4219,
		},
	},
	{
		"minecraft:granite_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 10446,
			477: 9910,
			573: 9910,
		},
	},
	{
		"minecraft:comparator[facing=south,mode=compare,powered=true]",
		nil,
		NewMapping{
			718: 6682,
			4:   2392,
			393: 5639,
			404: 5640,
			477: 6146,
			573: 6146,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=east,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7450,
			404: 7451,
			477: 7975,
			573: 7975,
			718: 8511,
			4:   2967,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=tall,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12806,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10424,
			573: 10424,
		},
	},
	{
		"minecraft:crimson_sign[rotation=10,waterlogged=false]",
		nil,
		NewMapping{
			718: 15676,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 5978,
		},
	},
	{
		"minecraft:jungle_pressure_plate[powered=false]",
		nil,
		NewMapping{
			573: 3878,
			718: 3880,
			393: 3374,
			404: 3375,
			477: 3878,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=5,south=side,west=side]",
		nil,
		NewMapping{
			393: 2809,
			404: 2810,
			477: 3113,
			573: 3113,
			718: 3115,
		},
	},
	{
		"minecraft:spruce_door[facing=north,half=upper,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			404: 7679,
			477: 8203,
			573: 8203,
			718: 8739,
			393: 7678,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 5663,
			573: 5663,
			393: 5159,
			404: 5160,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=east,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			718: 15184,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=none,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14129,
		},
	},
	{
		"minecraft:bell[attachment=ceiling,facing=south,powered=true]",
		nil,
		NewMapping{
			573: 11208,
			718: 14864,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 17068,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=east,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			718: 4426,
			393: 3908,
			404: 3909,
			477: 4412,
			573: 4412,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=4,south=side,west=up]",
		nil,
		NewMapping{
			393: 1791,
			404: 1792,
			477: 2095,
			573: 2095,
			718: 2097,
		},
	},
	{
		"minecraft:jungle_door[facing=north,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			718: 8876,
			393: 7815,
			404: 7816,
			477: 8340,
			573: 8340,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10121,
			573: 10121,
			718: 10657,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=none,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16580,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5783,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=west,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3638,
			404: 3639,
			477: 4142,
			573: 4142,
			718: 4156,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=9,powered=true]",
		nil,
		NewMapping{
			477: 916,
			573: 916,
			718: 917,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=tall,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12245,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=none,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5793,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=13,south=side,west=up]",
		nil,
		NewMapping{
			718: 3330,
			393: 3024,
			404: 3025,
			477: 3328,
			573: 3328,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9454,
			573: 9454,
			718: 9990,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=none,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13788,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=tall,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14307,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4426,
			404: 4427,
			477: 4930,
			573: 4930,
			718: 4946,
		},
	},
	{
		"minecraft:redstone_wall_torch[facing=west,lit=true]",
		nil,
		NewMapping{
			477: 3891,
			573: 3891,
			718: 3893,
			4:   1218,
			393: 3387,
			404: 3388,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=11,south=side,west=up]",
		nil,
		NewMapping{
			573: 3022,
			718: 3024,
			393: 2718,
			404: 2719,
			477: 3022,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13763,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=tall,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11273,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=3,powered=true]",
		nil,
		NewMapping{
			393: 654,
			404: 654,
			477: 654,
			573: 654,
			718: 655,
		},
	},
	{
		"minecraft:gray_glazed_terracotta[facing=east]",
		nil,
		NewMapping{
			393: 8344,
			404: 8345,
			477: 8869,
			573: 8869,
			718: 9405,
			4:   3875,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9144,
			573: 9144,
			718: 9680,
		},
	},
	{
		"minecraft:acacia_sign[rotation=0,waterlogged=true]",
		nil,
		NewMapping{
			477: 3475,
			573: 3475,
			718: 3477,
		},
	},
	{
		"minecraft:skeleton_wall_skull[facing=east]",
		nil,
		NewMapping{
			718: 6509,
			4:   2309,
			393: 5450,
			404: 5451,
			477: 5973,
			573: 5973,
		},
	},
	{
		"minecraft:light_blue_bed[facing=south,occupied=true,part=head]",
		nil,
		NewMapping{
			718: 1101,
			393: 800,
			404: 800,
			477: 1100,
			573: 1100,
		},
	},
	{
		"minecraft:heavy_weighted_pressure_plate[power=8]",
		nil,
		NewMapping{
			718: 6670,
			4:   2376,
			393: 5627,
			404: 5628,
			477: 6134,
			573: 6134,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13150,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11048,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=tall,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12586,
		},
	},
	{
		"minecraft:light_blue_shulker_box[facing=up]",
		nil,
		NewMapping{
			718: 9300,
			4:   3553,
			393: 8239,
			404: 8240,
			477: 8764,
			573: 8764,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5965,
			404: 5966,
			477: 6472,
			573: 6472,
			718: 7008,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 6165,
		},
	},
	{
		"minecraft:warped_button[face=wall,facing=west,powered=true]",
		nil,
		NewMapping{
			718: 15515,
		},
	},
	{
		"minecraft:warped_trapdoor[facing=north,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			718: 15202,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=1,south=none,west=side]",
		nil,
		NewMapping{
			477: 2792,
			573: 2792,
			718: 2794,
			393: 2488,
			404: 2489,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9646,
			573: 9646,
			718: 10182,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11959,
		},
	},
	{
		"minecraft:crimson_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15387,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9787,
			573: 9787,
			718: 10323,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=none,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11324,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12667,
		},
	},
	{
		"minecraft:snow[layers=8]",
		nil,
		NewMapping{
			477: 3926,
			573: 3926,
			718: 3928,
			4:   1255,
			393: 3422,
			404: 3423,
		},
	},
	{
		"minecraft:observer[facing=down,powered=false]",
		nil,
		NewMapping{
			573: 8735,
			718: 9271,
			4:   3488,
			393: 8210,
			404: 8211,
			477: 8735,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=2,south=none,west=none]",
		nil,
		NewMapping{
			573: 3090,
			718: 3092,
			393: 2786,
			404: 2787,
			477: 3090,
		},
	},
	{
		"minecraft:andesite_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 10001,
			573: 10001,
			718: 10537,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12575,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13909,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=east,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			573: 7062,
			718: 7598,
			393: 6555,
			404: 6556,
			477: 7062,
		},
	},
	{
		"minecraft:purpur_slab[type=bottom,waterlogged=true]",
		nil,
		NewMapping{
			718: 8410,
			393: 7349,
			404: 7350,
			477: 7874,
			573: 7874,
		},
	},
	{
		"minecraft:birch_button[face=wall,facing=west,powered=true]",
		nil,
		NewMapping{
			393: 5363,
			404: 5364,
			477: 5870,
			573: 5870,
			718: 6406,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9265,
			573: 9265,
			718: 9801,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=14,south=up,west=none]",
		nil,
		NewMapping{
			573: 2616,
			718: 2618,
			393: 2312,
			404: 2313,
			477: 2616,
		},
	},
	{
		"minecraft:wall_torch[facing=west]",
		nil,
		NewMapping{
			393: 1133,
			404: 1134,
			477: 1437,
			573: 1437,
			718: 1438,
			4:   802,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5180,
			404: 5181,
			477: 5684,
			573: 5684,
		},
	},
	{
		"minecraft:dark_oak_leaves[distance=6,persistent=true]",
		nil,
		NewMapping{
			573: 224,
			718: 225,
			393: 224,
			404: 224,
			477: 224,
		},
	},
	{
		"minecraft:lectern[facing=south,has_book=true,powered=false]",
		nil,
		NewMapping{
			718: 14838,
			477: 11182,
			573: 11182,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12888,
		},
	},
	{
		"minecraft:polished_blackstone_button[face=wall,facing=north,powered=false]",
		nil,
		NewMapping{
			718: 16762,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=tall,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11402,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 6599,
			477: 7105,
			573: 7105,
			718: 7641,
			393: 6598,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=true,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			718: 4505,
			4:   1598,
			393: 3987,
			404: 3988,
			477: 4491,
			573: 4491,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 6698,
			404: 6699,
			477: 7205,
			573: 7205,
			718: 7741,
		},
	},
	{
		"minecraft:tube_coral[waterlogged=false]",
		nil,
		NewMapping{
			404: 8471,
			477: 8995,
			573: 8995,
			718: 9531,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=tall,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16128,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 15895,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=none,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 15945,
		},
	},
	{
		"minecraft:acacia_fence_gate[facing=north,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			573: 7985,
			718: 8521,
			4:   2994,
			393: 7460,
			404: 7461,
			477: 7985,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6716,
			718: 7252,
			393: 6209,
			404: 6210,
			477: 6716,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=1,powered=true]",
		nil,
		NewMapping{
			718: 901,
			477: 900,
			573: 900,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13096,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9171,
			573: 9171,
			718: 9707,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=tall,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14520,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=west,in_wall=true,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7502,
			404: 7503,
			477: 8027,
			573: 8027,
			718: 8563,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10372,
			573: 10372,
		},
	},
	{
		"minecraft:yellow_bed[facing=south,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 818,
			404: 818,
			477: 1118,
			573: 1118,
			718: 1119,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=18,powered=true]",
		nil,
		NewMapping{
			393: 484,
			404: 484,
			477: 484,
			573: 484,
			718: 485,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6360,
			404: 6361,
			477: 6867,
			573: 6867,
			718: 7403,
		},
	},
	{
		"minecraft:blue_banner[rotation=10]",
		nil,
		NewMapping{
			393: 7040,
			404: 7041,
			477: 7547,
			573: 7547,
			718: 8083,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=6,powered=true]",
		nil,
		NewMapping{
			404: 610,
			477: 610,
			573: 610,
			718: 611,
			393: 610,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13321,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14623,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=none,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14434,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14201,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 1667,
			404: 1668,
			477: 1971,
			573: 1971,
			718: 1973,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=tall,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16119,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 15933,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14726,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13220,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=low,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11244,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 16330,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=east,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3836,
			404: 3837,
			477: 4340,
			573: 4340,
			718: 4354,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=1,south=none,west=side]",
		nil,
		NewMapping{
			393: 1768,
			404: 1769,
			477: 2072,
			573: 2072,
			718: 2074,
		},
	},
	{
		"minecraft:dark_oak_wood[axis=x]",
		nil,
		NewMapping{
			404: 123,
			477: 123,
			573: 123,
			718: 124,
			393: 123,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=19,powered=true]",
		nil,
		NewMapping{
			393: 486,
			404: 486,
			477: 486,
			573: 486,
			718: 487,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11145,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12808,
		},
	},
	{
		"minecraft:dead_tube_coral_wall_fan[facing=east,waterlogged=true]",
		nil,
		NewMapping{
			393: 8470,
			404: 8486,
			477: 9030,
			573: 9030,
			718: 9566,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=false,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			573: 4676,
			718: 4690,
			393: 4172,
			404: 4173,
			477: 4676,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10953,
			573: 10953,
		},
	},
	{
		"minecraft:granite_stairs[facing=west,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9904,
			573: 9904,
			718: 10440,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=south,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			404: 7432,
			477: 7956,
			573: 7956,
			718: 8492,
			393: 7431,
		},
	},
	{
		"minecraft:warped_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15466,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11404,
		},
	},
	{
		"minecraft:light_weighted_pressure_plate[power=3]",
		nil,
		NewMapping{
			393: 5606,
			404: 5607,
			477: 6113,
			573: 6113,
			718: 6649,
			4:   2355,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9392,
			573: 9392,
			718: 9928,
		},
	},
	{
		"minecraft:bee_nest[facing=north,honey_level=4]",
		nil,
		NewMapping{
			573: 11291,
			718: 15780,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=north,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			718: 15132,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=none,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5765,
		},
	},
	{
		"minecraft:purpur_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 8077,
			404: 8078,
			477: 8602,
			573: 8602,
			718: 9138,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9555,
			573: 9555,
			718: 10091,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10532,
			573: 10532,
		},
	},
	{
		"minecraft:campfire[facing=east,lit=true,signal_fire=false,waterlogged=false]",
		nil,
		NewMapping{
			477: 11243,
			573: 11259,
			718: 14917,
		},
	},
	{
		"minecraft:acacia_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 6395,
			477: 6901,
			573: 6901,
			718: 7437,
			393: 6394,
		},
	},
	{
		"minecraft:soul_campfire[facing=north,lit=false,signal_fire=true,waterlogged=false]",
		nil,
		NewMapping{
			718: 14927,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=none,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16566,
		},
	},
	{
		"minecraft:warped_sign[rotation=11,waterlogged=false]",
		nil,
		NewMapping{
			718: 15710,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=7,south=up,west=none]",
		nil,
		NewMapping{
			393: 2969,
			404: 2970,
			477: 3273,
			573: 3273,
			718: 3275,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6736,
			573: 6736,
			718: 7272,
			393: 6229,
			404: 6230,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=3,powered=false]",
		nil,
		NewMapping{
			393: 305,
			404: 305,
			477: 305,
			573: 305,
			718: 306,
		},
	},
	{
		"minecraft:birch_door[facing=south,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			404: 7764,
			477: 8288,
			573: 8288,
			718: 8824,
			393: 7763,
		},
	},
	{
		"minecraft:brown_banner[rotation=8]",
		nil,
		NewMapping{
			393: 7054,
			404: 7055,
			477: 7561,
			573: 7561,
			718: 8097,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=west,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			477: 4335,
			573: 4335,
			718: 4349,
			393: 3831,
			404: 3832,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10757,
			573: 10757,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=tall,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16982,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=tall,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11405,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=low,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13630,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11285,
		},
	},
	{
		"minecraft:iron_ore",
		nil,
		NewMapping{
			477: 70,
			573: 70,
			718: 70,
			4:   240,
			393: 70,
			404: 70,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=14,south=side,west=none]",
		nil,
		NewMapping{
			393: 2315,
			404: 2316,
			477: 2619,
			573: 2619,
			718: 2621,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 7306,
			718: 7842,
			393: 6799,
			404: 6800,
			477: 7306,
		},
	},
	{
		"minecraft:spruce_sign[rotation=14,waterlogged=true]",
		nil,
		NewMapping{
			477: 3439,
			573: 3439,
			718: 3441,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=none,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11416,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=low,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 5696,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=none,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13600,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13337,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=west,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			718: 7578,
			393: 6535,
			404: 6536,
			477: 7042,
			573: 7042,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=11,south=side,west=side]",
		nil,
		NewMapping{
			393: 2719,
			404: 2720,
			477: 3023,
			573: 3023,
			718: 3025,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6753,
			404: 6754,
			477: 7260,
			573: 7260,
			718: 7796,
		},
	},
	{
		"minecraft:birch_sign[rotation=14,waterlogged=true]",
		nil,
		NewMapping{
			477: 3471,
			573: 3471,
			718: 3473,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 15936,
		},
	},
	{
		"minecraft:polished_blackstone_brick_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			718: 16258,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=tall,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13433,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 16692,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=low,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11556,
		},
	},
	{
		"minecraft:quartz_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 6817,
			393: 5774,
			404: 5775,
			477: 6281,
			573: 6281,
		},
	},
	{
		"minecraft:orange_wall_banner[facing=south]",
		nil,
		NewMapping{
			393: 7115,
			404: 7116,
			477: 7622,
			573: 7622,
			718: 8158,
		},
	},
	{
		"minecraft:hopper[enabled=false,facing=south]",
		nil,
		NewMapping{
			4:   2475,
			393: 5692,
			404: 5693,
			477: 6199,
			573: 6199,
			718: 6735,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9175,
			573: 9175,
			718: 9711,
		},
	},
	{
		"minecraft:repeater[delay=4,facing=east,locked=false,powered=true]",
		nil,
		NewMapping{
			477: 4079,
			573: 4079,
			718: 4093,
			4:   1519,
			393: 3575,
			404: 3576,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=low,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12424,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9610,
			573: 9610,
			718: 10146,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=tall,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 5848,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=tall,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16970,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=3,powered=true]",
		nil,
		NewMapping{
			393: 454,
			404: 454,
			477: 454,
			573: 454,
			718: 455,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14508,
		},
	},
	{
		"minecraft:warped_door[facing=south,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15617,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=none,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11409,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 5099,
			718: 5115,
			393: 4595,
			404: 4596,
			477: 5099,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=23,powered=false]",
		nil,
		NewMapping{
			393: 545,
			404: 545,
			477: 545,
			573: 545,
			718: 546,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9232,
			573: 9232,
			718: 9768,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13765,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=tall,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 5866,
		},
	},
	{
		"minecraft:dragon_head[rotation=2]",
		nil,
		NewMapping{
			404: 5554,
			477: 6056,
			573: 6056,
			718: 6592,
			393: 5553,
		},
	},
	{
		"minecraft:lever[face=wall,facing=west,powered=true]",
		nil,
		NewMapping{
			393: 3289,
			404: 3290,
			477: 3793,
			573: 3793,
			718: 3795,
			4:   1114,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9592,
			573: 9592,
			718: 10128,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16041,
		},
	},
	{
		"minecraft:fire[age=8,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1409,
			404: 1410,
			477: 1713,
			573: 1713,
			718: 1714,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=12,south=up,west=side]",
		nil,
		NewMapping{
			393: 3013,
			404: 3014,
			477: 3317,
			573: 3317,
			718: 3319,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=3,south=side,west=side]",
		nil,
		NewMapping{
			393: 2215,
			404: 2216,
			477: 2519,
			573: 2519,
			718: 2521,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 3229,
			404: 3230,
			477: 3693,
			573: 3693,
			718: 3695,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10635,
			573: 10635,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16213,
		},
	},
	{
		"minecraft:diorite_stairs[facing=south,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 10203,
			573: 10203,
			718: 10739,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=tall,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14098,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=east,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7447,
			404: 7448,
			477: 7972,
			573: 7972,
			718: 8508,
		},
	},
	{
		"minecraft:lime_bed[facing=south,occupied=false,part=foot]",
		nil,
		NewMapping{
			477: 1135,
			573: 1135,
			718: 1136,
			393: 835,
			404: 835,
		},
	},
	{
		"minecraft:fire[age=12,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			718: 1831,
			393: 1526,
			404: 1527,
			477: 1830,
			573: 1830,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 10208,
			477: 9672,
			573: 9672,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=22,powered=true]",
		nil,
		NewMapping{
			477: 442,
			573: 442,
			718: 443,
			393: 442,
			404: 442,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=none,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14461,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=low,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11570,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5979,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 6491,
			477: 6997,
			573: 6997,
			718: 7533,
			393: 6490,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10242,
			573: 10242,
			718: 10778,
		},
	},
	{
		"minecraft:polished_blackstone_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			718: 16749,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=2,south=up,west=up]",
		nil,
		NewMapping{
			718: 2652,
			393: 2346,
			404: 2347,
			477: 2650,
			573: 2650,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 7752,
			718: 8288,
			393: 7245,
			404: 7246,
			477: 7752,
		},
	},
	{
		"minecraft:warped_fence[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			718: 15113,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12284,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9231,
			573: 9231,
			718: 9767,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 15880,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6103,
			404: 6104,
			477: 6610,
			573: 6610,
			718: 7146,
		},
	},
	{
		"minecraft:jungle_fence[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7599,
			404: 7600,
			477: 8124,
			573: 8124,
			718: 8660,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=8,powered=false]",
		nil,
		NewMapping{
			477: 865,
			573: 865,
			718: 866,
		},
	},
	{
		"minecraft:end_stone_brick_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			477: 10286,
			573: 10286,
			718: 10822,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=none,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13789,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=low,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14045,
		},
	},
	{
		"minecraft:skeleton_skull[rotation=15]",
		nil,
		NewMapping{
			393: 5466,
			404: 5467,
			477: 5969,
			573: 5969,
			718: 6505,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=none,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 17001,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=tall,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12250,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=none,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11004,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=false,north=true,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			718: 5293,
			393: 4773,
			404: 4774,
			477: 5277,
			573: 5277,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5195,
			404: 5196,
			477: 5699,
			573: 5699,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=0,south=up,west=side]",
		nil,
		NewMapping{
			718: 2923,
			393: 2617,
			404: 2618,
			477: 2921,
			573: 2921,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12031,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=east,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 4274,
			573: 4274,
			718: 4288,
			393: 3770,
			404: 3771,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=10,south=up,west=side]",
		nil,
		NewMapping{
			404: 2708,
			477: 3011,
			573: 3011,
			718: 3013,
			393: 2707,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10810,
			573: 10810,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=none,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14671,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=none,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13710,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=tall,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6293,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=11,powered=true]",
		nil,
		NewMapping{
			393: 370,
			404: 370,
			477: 370,
			573: 370,
			718: 371,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=1,south=none,west=none]",
		nil,
		NewMapping{
			573: 3081,
			718: 3083,
			393: 2777,
			404: 2778,
			477: 3081,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=west,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3832,
			404: 3833,
			477: 4336,
			573: 4336,
			718: 4350,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=none,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5691,
		},
	},
	{
		"minecraft:crimson_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			718: 15052,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=tall,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12045,
		},
	},
	{
		"minecraft:dark_oak_door[facing=south,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7956,
			404: 7957,
			477: 8481,
			573: 8481,
			718: 9017,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13211,
		},
	},
	{
		"minecraft:warped_roots",
		nil,
		NewMapping{
			718: 14973,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=none,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16913,
		},
	},
	{
		"minecraft:stripped_jungle_wood[axis=z]",
		nil,
		NewMapping{
			404: 137,
			477: 137,
			573: 137,
			718: 138,
			393: 137,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9593,
			573: 9593,
			718: 10129,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=none,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13698,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16883,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=8,south=up,west=up]",
		nil,
		NewMapping{
			393: 2832,
			404: 2833,
			477: 3136,
			573: 3136,
			718: 3138,
		},
	},
	{
		"minecraft:stripped_oak_log[axis=x]",
		nil,
		NewMapping{
			477: 105,
			573: 105,
			718: 106,
			393: 105,
			404: 105,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10798,
			573: 10798,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 15915,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 5962,
		},
	},
	{
		"minecraft:bricks",
		nil,
		NewMapping{
			393: 1125,
			404: 1125,
			477: 1428,
			573: 1428,
			718: 1429,
			4:   720,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9301,
			573: 9301,
			718: 9837,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9462,
			573: 9462,
			718: 9998,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=tall,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11617,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=none,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12198,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=none,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16474,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=tall,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16537,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=tall,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13434,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=north,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			404: 7488,
			477: 8012,
			573: 8012,
			718: 8548,
			393: 7487,
		},
	},
	{
		"minecraft:kelp[age=0]",
		nil,
		NewMapping{
			393: 8409,
			404: 8410,
			477: 8934,
			573: 8934,
			718: 9470,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=none,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11947,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 16337,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11336,
		},
	},
	{
		"minecraft:polished_blackstone_button[face=ceiling,facing=north,powered=true]",
		nil,
		NewMapping{
			718: 16769,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=low,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11896,
		},
	},
	{
		"minecraft:gray_wall_banner[facing=east]",
		nil,
		NewMapping{
			393: 7141,
			404: 7142,
			477: 7648,
			573: 7648,
			718: 8184,
		},
	},
	{
		"minecraft:fire[age=12,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1520,
			404: 1521,
			477: 1824,
			573: 1824,
			718: 1825,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=tall,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5850,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16834,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=22,powered=false]",
		nil,
		NewMapping{
			477: 993,
			573: 993,
			718: 994,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=none,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12706,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=3,south=side,west=none]",
		nil,
		NewMapping{
			404: 1929,
			477: 2232,
			573: 2232,
			718: 2234,
			393: 1928,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=8,south=side,west=none]",
		nil,
		NewMapping{
			393: 1973,
			404: 1974,
			477: 2277,
			573: 2277,
			718: 2279,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=true,north=false,powered=false,south=false,west=false]",
		nil,
		NewMapping{
			393: 4866,
			404: 4867,
			477: 5370,
			573: 5370,
			718: 5386,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=west,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			393: 4322,
			404: 4323,
			477: 4826,
			573: 4826,
			718: 4842,
			4:   1721,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=none,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16918,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=none,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13488,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=none,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13245,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14395,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=15,powered=false]",
		nil,
		NewMapping{
			718: 480,
			393: 479,
			404: 479,
			477: 479,
			573: 479,
		},
	},
	{
		"minecraft:orange_bed[facing=west,occupied=false,part=foot]",
		nil,
		NewMapping{
			404: 775,
			477: 1075,
			573: 1075,
			718: 1076,
			393: 775,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10872,
			573: 10872,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11501,
		},
	},
	{
		"minecraft:birch_door[facing=north,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7745,
			404: 7746,
			477: 8270,
			573: 8270,
			718: 8806,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			404: 4424,
			477: 4927,
			573: 4927,
			718: 4943,
			4:   1747,
			393: 4423,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=1,powered=true]",
		nil,
		NewMapping{
			573: 950,
			718: 951,
			477: 950,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 9841,
			718: 10377,
			477: 9841,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14601,
		},
	},
	{
		"minecraft:warped_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15407,
		},
	},
	{
		"minecraft:fire[age=13,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			404: 1566,
			477: 1869,
			573: 1869,
			718: 1870,
			393: 1565,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 6686,
			477: 7192,
			573: 7192,
			718: 7728,
			393: 6685,
		},
	},
	{
		"minecraft:jungle_button[face=floor,facing=south,powered=true]",
		nil,
		NewMapping{
			393: 5377,
			404: 5378,
			477: 5884,
			573: 5884,
			718: 6420,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=south,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 4310,
			573: 4310,
			718: 4324,
			393: 3806,
			404: 3807,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10129,
			573: 10129,
			718: 10665,
		},
	},
	{
		"minecraft:light_gray_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 885,
			404: 885,
			477: 1185,
			573: 1185,
			718: 1186,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=9,south=none,west=side]",
		nil,
		NewMapping{
			404: 2129,
			477: 2432,
			573: 2432,
			718: 2434,
			393: 2128,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=2,south=none,west=side]",
		nil,
		NewMapping{
			573: 2513,
			718: 2515,
			393: 2209,
			404: 2210,
			477: 2513,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 5151,
			477: 5654,
			573: 5654,
			393: 5150,
		},
	},
	{
		"minecraft:andesite_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10840,
			573: 10840,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=15,powered=false]",
		nil,
		NewMapping{
			477: 929,
			573: 929,
			718: 930,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=low,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 6264,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13289,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5831,
			404: 5832,
			477: 6338,
			573: 6338,
			718: 6874,
		},
	},
	{
		"minecraft:spruce_leaves[distance=5,persistent=false]",
		nil,
		NewMapping{
			393: 167,
			404: 167,
			477: 167,
			573: 167,
			718: 168,
		},
	},
	{
		"minecraft:iron_door[facing=east,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 3362,
			404: 3363,
			477: 3866,
			573: 3866,
			718: 3868,
		},
	},
	{
		"minecraft:bamboo[age=0,leaves=none,stage=1]",
		nil,
		NewMapping{
			477: 9117,
			573: 9117,
			718: 9653,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=low,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13399,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=tall,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16111,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=low,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 17047,
		},
	},
	{
		"minecraft:crimson_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15337,
		},
	},
	{
		"minecraft:iron_door[facing=south,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 3321,
			404: 3322,
			477: 3825,
			573: 3825,
			718: 3827,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=12,powered=true]",
		nil,
		NewMapping{
			573: 472,
			718: 473,
			393: 472,
			404: 472,
			477: 472,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13040,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=none,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13809,
		},
	},
	{
		"minecraft:moving_piston[facing=north,type=sticky]",
		nil,
		NewMapping{
			4:   586,
			393: 1100,
			404: 1100,
			477: 1400,
			573: 1400,
			718: 1401,
		},
	},
	{
		"minecraft:yellow_banner[rotation=15]",
		nil,
		NewMapping{
			477: 7440,
			573: 7440,
			718: 7976,
			393: 6933,
			404: 6934,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12549,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13695,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=none,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14118,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=east,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 3966,
			404: 3967,
			477: 4470,
			573: 4470,
			718: 4484,
		},
	},
	{
		"minecraft:piston_head[facing=east,short=true,type=normal]",
		nil,
		NewMapping{
			393: 1063,
			404: 1063,
			477: 1363,
			573: 1363,
			718: 1364,
		},
	},
	{
		"minecraft:structure_block[mode=load]",
		nil,
		NewMapping{
			404: 8596,
			477: 11253,
			573: 11269,
			718: 15736,
			4:   4081,
			393: 8579,
		},
	},
	{
		"minecraft:jigsaw[facing=up]",
		nil,
		NewMapping{
			477: 11260,
			573: 11276,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9605,
			573: 9605,
			718: 10141,
		},
	},
	{
		"minecraft:lime_stained_glass",
		nil,
		NewMapping{
			393: 3582,
			404: 3583,
			477: 4086,
			573: 4086,
			718: 4100,
			4:   1525,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=2,south=up,west=none]",
		nil,
		NewMapping{
			393: 2636,
			404: 2637,
			477: 2940,
			573: 2940,
			718: 2942,
		},
	},
	{
		"minecraft:jungle_wood[axis=x]",
		nil,
		NewMapping{
			573: 117,
			718: 118,
			393: 117,
			404: 117,
			477: 117,
		},
	},
	{
		"minecraft:black_concrete",
		nil,
		NewMapping{
			4:   4031,
			393: 8392,
			404: 8393,
			477: 8917,
			573: 8917,
			718: 9453,
		},
	},
	{
		"minecraft:glass_pane[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			718: 4737,
			393: 4217,
			404: 4218,
			477: 4721,
			573: 4721,
		},
	},
	{
		"minecraft:barrel[facing=north,open=false]",
		nil,
		NewMapping{
			477: 11136,
			573: 11136,
			718: 14792,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=tall,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12142,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 15885,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9388,
			573: 9388,
			718: 9924,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 10300,
			477: 9764,
			573: 9764,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14633,
		},
	},
	{
		"minecraft:oak_button[face=floor,facing=east,powered=true]",
		nil,
		NewMapping{
			393: 5309,
			404: 5310,
			477: 5816,
			573: 5816,
			718: 6352,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=15,south=up,west=up]",
		nil,
		NewMapping{
			393: 3039,
			404: 3040,
			477: 3343,
			573: 3343,
			718: 3345,
		},
	},
	{
		"minecraft:oak_button[face=floor,facing=west,powered=false]",
		nil,
		NewMapping{
			718: 6351,
			393: 5308,
			404: 5309,
			477: 5815,
			573: 5815,
		},
	},
	{
		"minecraft:mossy_cobblestone_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9521,
			573: 9521,
			718: 10057,
		},
	},
	{
		"minecraft:weeping_vines[age=16]",
		nil,
		NewMapping{
			718: 15006,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5967,
		},
	},
	{
		"minecraft:cactus[age=13]",
		nil,
		NewMapping{
			718: 3944,
			4:   1309,
			393: 3438,
			404: 3439,
			477: 3942,
			573: 3942,
		},
	},
	{
		"minecraft:granite_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9893,
			573: 9893,
			718: 10429,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12024,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 10952,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11584,
		},
	},
	{
		"minecraft:birch_door[facing=east,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7797,
			404: 7798,
			477: 8322,
			573: 8322,
			718: 8858,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 3483,
			404: 3484,
			477: 3987,
			573: 3987,
			718: 3989,
		},
	},
	{
		"minecraft:light_gray_banner[rotation=2]",
		nil,
		NewMapping{
			718: 8027,
			393: 6984,
			404: 6985,
			477: 7491,
			573: 7491,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=tall,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12800,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6772,
			573: 6772,
			718: 7308,
			393: 6265,
			404: 6266,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9796,
			573: 9796,
			718: 10332,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16426,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14122,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9862,
			573: 9862,
			718: 10398,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=low,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 15986,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14685,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=low,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11362,
		},
	},
	{
		"minecraft:dark_oak_door[facing=south,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7953,
			404: 7954,
			477: 8478,
			573: 8478,
			718: 9014,
		},
	},
	{
		"minecraft:spruce_door[facing=west,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7713,
			404: 7714,
			477: 8238,
			573: 8238,
			718: 8774,
		},
	},
	{
		"minecraft:birch_door[facing=south,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7760,
			404: 7761,
			477: 8285,
			573: 8285,
			718: 8821,
		},
	},
	{
		"minecraft:cyan_wall_banner[facing=north]",
		nil,
		NewMapping{
			393: 7146,
			404: 7147,
			477: 7653,
			573: 7653,
			718: 8189,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16644,
		},
	},
	{
		"minecraft:oak_sapling[stage=0]",
		nil,
		NewMapping{
			718: 21,
			4:   96,
			393: 21,
			404: 21,
			477: 21,
			573: 21,
		},
	},
	{
		"minecraft:oak_leaves[distance=4,persistent=false]",
		nil,
		NewMapping{
			718: 152,
			393: 151,
			404: 151,
			477: 151,
			573: 151,
		},
	},
	{
		"minecraft:brick_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4342,
			404: 4343,
			477: 4846,
			573: 4846,
			718: 4862,
		},
	},
	{
		"minecraft:fire_coral_wall_fan[facing=north,waterlogged=true]",
		nil,
		NewMapping{
			718: 9624,
			393: 8528,
			404: 8544,
			477: 9088,
			573: 9088,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 3212,
			404: 3213,
			477: 3676,
			573: 3676,
			718: 3678,
		},
	},
	{
		"minecraft:iron_bars[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 4697,
			718: 4711,
			393: 4193,
			404: 4194,
			477: 4697,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=tall,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16552,
		},
	},
	{
		"minecraft:green_banner[rotation=3]",
		nil,
		NewMapping{
			718: 8108,
			393: 7065,
			404: 7066,
			477: 7572,
			573: 7572,
		},
	},
	{
		"minecraft:brain_coral_block",
		nil,
		NewMapping{
			573: 8980,
			718: 9516,
			393: 8455,
			404: 8456,
			477: 8980,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=south,in_wall=false,open=true,powered=true]",
		nil,
		NewMapping{
			477: 7958,
			573: 7958,
			718: 8494,
			4:   2972,
			393: 7433,
			404: 7434,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=south,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			393: 4315,
			404: 4316,
			477: 4819,
			573: 4819,
			718: 4835,
			4:   1712,
		},
	},
	{
		"minecraft:green_bed[facing=south,occupied=true,part=head]",
		nil,
		NewMapping{
			573: 1260,
			718: 1261,
			393: 960,
			404: 960,
			477: 1260,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9313,
			573: 9313,
			718: 9849,
		},
	},
	{
		"minecraft:warped_door[facing=west,half=lower,hinge=left,open=true,powered=false]",
		nil,
		NewMapping{
			718: 15632,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6059,
			404: 6060,
			477: 6566,
			573: 6566,
			718: 7102,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=east,in_wall=false,open=false,powered=false]",
		nil,
		NewMapping{
			477: 7913,
			573: 7913,
			718: 8449,
			4:   2931,
			393: 7388,
			404: 7389,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11096,
			573: 11096,
		},
	},
	{
		"minecraft:bell[attachment=floor,facing=south]",
		nil,
		NewMapping{
			477: 11199,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5835,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 15995,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 16329,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			718: 6967,
			393: 5924,
			404: 5925,
			477: 6431,
			573: 6431,
		},
	},
	{
		"minecraft:stripped_acacia_wood[axis=y]",
		nil,
		NewMapping{
			393: 139,
			404: 139,
			477: 139,
			573: 139,
			718: 140,
		},
	},
	{
		"minecraft:birch_sign[rotation=4,waterlogged=true]",
		nil,
		NewMapping{
			477: 3451,
			573: 3451,
			718: 3453,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=none,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16138,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=low,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13733,
		},
	},
	{
		"minecraft:acacia_trapdoor[facing=west,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			404: 3887,
			477: 4390,
			573: 4390,
			718: 4404,
			393: 3886,
		},
	},
	{
		"minecraft:granite_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 10636,
			477: 10636,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10521,
			573: 10521,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=low,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12447,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=2,south=none,west=none]",
		nil,
		NewMapping{
			573: 2226,
			718: 2228,
			393: 1922,
			404: 1923,
			477: 2226,
		},
	},
	{
		"minecraft:dark_oak_door[facing=west,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			718: 9038,
			393: 7977,
			404: 7978,
			477: 8502,
			573: 8502,
		},
	},
	{
		"minecraft:light_gray_bed[facing=north,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 879,
			404: 879,
			477: 1179,
			573: 1179,
			718: 1180,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=13,powered=false]",
		nil,
		NewMapping{
			404: 575,
			477: 575,
			573: 575,
			718: 576,
			393: 575,
		},
	},
	{
		"minecraft:trapped_chest[facing=west,type=right,waterlogged=true]",
		nil,
		NewMapping{
			573: 6102,
			718: 6638,
			393: 5595,
			404: 5596,
			477: 6102,
		},
	},
	{
		"minecraft:yellow_banner[rotation=13]",
		nil,
		NewMapping{
			393: 6931,
			404: 6932,
			477: 7438,
			573: 7438,
			718: 7974,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 6530,
			573: 6530,
			718: 7066,
			393: 6023,
			404: 6024,
		},
	},
	{
		"minecraft:jungle_sign[rotation=3,waterlogged=false]",
		nil,
		NewMapping{
			477: 3514,
			573: 3514,
			718: 3516,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=none,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5671,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 10402,
			477: 9866,
			573: 9866,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9444,
			573: 9444,
			718: 9980,
		},
	},
	{
		"minecraft:granite_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10602,
			573: 10602,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5951,
			404: 5952,
			477: 6458,
			573: 6458,
			718: 6994,
		},
	},
	{
		"minecraft:chorus_flower[age=0]",
		nil,
		NewMapping{
			573: 8592,
			718: 9128,
			4:   3200,
			393: 8067,
			404: 8068,
			477: 8592,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=4,south=up,west=none]",
		nil,
		NewMapping{
			393: 2366,
			404: 2367,
			477: 2670,
			573: 2670,
			718: 2672,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=false,north=false,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			393: 4877,
			404: 4878,
			477: 5381,
			573: 5381,
			718: 5397,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=none,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12170,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16625,
		},
	},
	{
		"minecraft:purpur_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 9184,
			393: 8123,
			404: 8124,
			477: 8648,
			573: 8648,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10379,
			573: 10379,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=low,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14157,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14667,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=7,south=none,west=side]",
		nil,
		NewMapping{
			393: 1822,
			404: 1823,
			477: 2126,
			573: 2126,
			718: 2128,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=1,powered=false]",
		nil,
		NewMapping{
			404: 651,
			477: 651,
			573: 651,
			718: 652,
			393: 651,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 16333,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11775,
		},
	},
	{
		"minecraft:dropper[facing=up,triggered=false]",
		nil,
		NewMapping{
			4:   2529,
			393: 5801,
			404: 5802,
			477: 6308,
			573: 6308,
			718: 6844,
		},
	},
	{
		"minecraft:jungle_door[facing=south,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7825,
			404: 7826,
			477: 8350,
			573: 8350,
			718: 8886,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 9211,
			718: 9747,
			477: 9211,
		},
	},
	{
		"minecraft:crimson_door[facing=west,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15565,
		},
	},
	{
		"minecraft:warped_fence_gate[facing=west,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			718: 15308,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=low,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6143,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=east,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			718: 15188,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 6080,
			404: 6081,
			477: 6587,
			573: 6587,
			718: 7123,
		},
	},
	{
		"minecraft:cyan_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 6625,
			573: 6625,
			718: 7161,
			393: 6118,
			404: 6119,
		},
	},
	{
		"minecraft:stone_button[face=wall,facing=west,powered=false]",
		nil,
		NewMapping{
			718: 3910,
			4:   1234,
			393: 3404,
			404: 3405,
			477: 3908,
			573: 3908,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 5956,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 6543,
			573: 6543,
			718: 7079,
			393: 6036,
			404: 6037,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=none,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11213,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=west,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3753,
			404: 3754,
			477: 4257,
			573: 4257,
			718: 4271,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=none,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11215,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16898,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=2,south=side,west=side]",
		nil,
		NewMapping{
			393: 2782,
			404: 2783,
			477: 3086,
			573: 3086,
			718: 3088,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=west,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			4:   2933,
			393: 7378,
			404: 7379,
			477: 7903,
			573: 7903,
			718: 8439,
		},
	},
	{
		"minecraft:oak_fence[east=true,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 3977,
			573: 3977,
			718: 3979,
			393: 3473,
			404: 3474,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10055,
			573: 10055,
			718: 10591,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=tall,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14534,
		},
	},
	{
		"minecraft:stone_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9680,
			573: 9680,
			718: 10216,
		},
	},
	{
		"minecraft:jungle_sign[rotation=1,waterlogged=false]",
		nil,
		NewMapping{
			718: 3512,
			477: 3510,
			573: 3510,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16097,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=11,south=side,west=up]",
		nil,
		NewMapping{
			393: 1854,
			404: 1855,
			477: 2158,
			573: 2158,
			718: 2160,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=east,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			4:   1551,
			393: 3644,
			404: 3645,
			477: 4148,
			573: 4148,
			718: 4162,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 6270,
			404: 6271,
			477: 6777,
			573: 6777,
			718: 7313,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 4020,
			404: 4021,
			477: 4524,
			573: 4524,
			718: 4538,
		},
	},
	{
		"minecraft:stone_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9619,
			573: 9619,
			718: 10155,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14078,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12999,
		},
	},
	{
		"minecraft:crimson_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 15350,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12918,
		},
	},
	{
		"minecraft:yellow_bed[facing=west,occupied=false,part=foot]",
		nil,
		NewMapping{
			573: 1123,
			718: 1124,
			393: 823,
			404: 823,
			477: 1123,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=4,powered=true]",
		nil,
		NewMapping{
			718: 507,
			393: 506,
			404: 506,
			477: 506,
			573: 506,
		},
	},
	{
		"minecraft:birch_fence[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 7564,
			404: 7565,
			477: 8089,
			573: 8089,
			718: 8625,
		},
	},
	{
		"minecraft:stone_slab[type=top,waterlogged=false]",
		nil,
		NewMapping{
			393: 7294,
			477: 7801,
			573: 7801,
			718: 8337,
			4:   712,
		},
	},
	{
		"minecraft:warped_wall_sign[facing=east,waterlogged=true]",
		nil,
		NewMapping{
			718: 15733,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=south,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 3678,
			718: 3680,
			393: 3214,
			404: 3215,
			477: 3678,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=north,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3923,
			404: 3924,
			477: 4427,
			573: 4427,
			718: 4441,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=22,powered=true]",
		nil,
		NewMapping{
			393: 742,
			404: 742,
			477: 742,
			573: 742,
			718: 743,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=true,north=false,powered=true,south=false,west=false]",
		nil,
		NewMapping{
			393: 4862,
			404: 4863,
			477: 5366,
			573: 5366,
			718: 5382,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 17030,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			404: 8037,
			477: 8561,
			573: 8561,
			718: 9097,
			393: 8036,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=9,south=side,west=none]",
		nil,
		NewMapping{
			404: 1983,
			477: 2286,
			573: 2286,
			718: 2288,
			393: 1982,
		},
	},
	{
		"minecraft:birch_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 5524,
			718: 5540,
			393: 5020,
			404: 5021,
			477: 5524,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 10131,
			718: 10667,
			477: 10131,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			477: 4638,
			573: 4638,
			718: 4652,
			393: 4134,
			404: 4135,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=8,powered=true]",
		nil,
		NewMapping{
			718: 315,
			393: 314,
			404: 314,
			477: 314,
			573: 314,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5943,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13221,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=tall,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11726,
		},
	},
	{
		"minecraft:redstone_wall_torch[facing=north,lit=true]",
		nil,
		NewMapping{
			718: 3889,
			4:   1220,
			393: 3383,
			404: 3384,
			477: 3887,
			573: 3887,
		},
	},
	{
		"minecraft:stone_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9636,
			573: 9636,
			718: 10172,
		},
	},
	{
		"minecraft:crimson_fence_gate[facing=west,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			718: 15276,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11176,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=false,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10544,
			573: 10544,
		},
	},
	{
		"minecraft:beehive[facing=north,honey_level=5]",
		nil,
		NewMapping{
			573: 11316,
			718: 15805,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=none,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13892,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=low,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13415,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=4,south=up,west=none]",
		nil,
		NewMapping{
			573: 2238,
			718: 2240,
			393: 1934,
			404: 1935,
			477: 2238,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5908,
			404: 5909,
			477: 6415,
			573: 6415,
			718: 6951,
		},
	},
	{
		"minecraft:nether_brick_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10766,
			573: 10766,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=none,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14567,
		},
	},
	{
		"minecraft:fire[age=9,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			4:   825,
			393: 1454,
			404: 1455,
			477: 1758,
			573: 1758,
			718: 1759,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11331,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13283,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=14,south=none,west=side]",
		nil,
		NewMapping{
			393: 2893,
			404: 2894,
			477: 3197,
			573: 3197,
			718: 3199,
		},
	},
	{
		"minecraft:fire[age=14,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			393: 1590,
			404: 1591,
			477: 1894,
			573: 1894,
			718: 1895,
		},
	},
	{
		"minecraft:fire[age=6,east=true,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			573: 1637,
			718: 1638,
			393: 1333,
			404: 1334,
			477: 1637,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=low,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11134,
		},
	},
	{
		"minecraft:oak_wall_sign[facing=north,waterlogged=false]",
		nil,
		NewMapping{
			573: 3734,
			718: 3736,
			404: 3271,
			477: 3734,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5763,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=13,south=side,west=none]",
		nil,
		NewMapping{
			393: 2738,
			404: 2739,
			477: 3042,
			573: 3042,
			718: 3044,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16155,
		},
	},
	{
		"minecraft:light_gray_shulker_box[facing=up]",
		nil,
		NewMapping{
			4:   3633,
			393: 8269,
			404: 8270,
			477: 8794,
			573: 8794,
			718: 9330,
		},
	},
	{
		"minecraft:sticky_piston[extended=false,facing=down]",
		nil,
		NewMapping{
			4:   464,
			393: 1039,
			404: 1039,
			477: 1339,
			573: 1339,
			718: 1340,
		},
	},
	{
		"minecraft:magenta_wall_banner[facing=north]",
		nil,
		NewMapping{
			573: 7625,
			718: 8161,
			393: 7118,
			404: 7119,
			477: 7625,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10503,
			573: 10503,
		},
	},
	{
		"minecraft:jungle_door[facing=north,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7805,
			404: 7806,
			477: 8330,
			573: 8330,
			718: 8866,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=none,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13478,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=none,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12407,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13966,
		},
	},
	{
		"minecraft:stone_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9689,
			573: 9689,
			718: 10225,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 16703,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=4,south=none,west=none]",
		nil,
		NewMapping{
			573: 2100,
			718: 2102,
			393: 1796,
			404: 1797,
			477: 2100,
		},
	},
	{
		"minecraft:andesite_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10795,
			573: 10795,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=tall,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 6186,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11152,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=true,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6552,
			573: 6552,
			718: 7088,
			393: 6045,
			404: 6046,
		},
	},
	{
		"minecraft:wall_torch[facing=south]",
		nil,
		NewMapping{
			4:   803,
			393: 1132,
			404: 1133,
			477: 1436,
			573: 1436,
			718: 1437,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=5,south=up,west=up]",
		nil,
		NewMapping{
			393: 2229,
			404: 2230,
			477: 2533,
			573: 2533,
			718: 2535,
		},
	},
	{
		"minecraft:brick_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4396,
			404: 4397,
			477: 4900,
			573: 4900,
			718: 4916,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=low,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12428,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=15,south=side,west=none]",
		nil,
		NewMapping{
			393: 2900,
			404: 2901,
			477: 3204,
			573: 3204,
			718: 3206,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=true,north=true,powered=true,south=false,west=false]",
		nil,
		NewMapping{
			393: 4790,
			404: 4791,
			477: 5294,
			573: 5294,
			718: 5310,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16971,
		},
	},
	{
		"minecraft:spruce_button[face=wall,facing=east,powered=true]",
		nil,
		NewMapping{
			718: 6384,
			393: 5341,
			404: 5342,
			477: 5848,
			573: 5848,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=false,north=false,powered=false,south=false,west=false]",
		nil,
		NewMapping{
			393: 4786,
			404: 4787,
			477: 5290,
			573: 5290,
			718: 5306,
			4:   2126,
		},
	},
	{
		"minecraft:birch_door[facing=east,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7801,
			404: 7802,
			477: 8326,
			573: 8326,
			718: 8862,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13045,
		},
	},
	{
		"minecraft:crimson_door[facing=north,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			718: 15531,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=tall,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16645,
		},
	},
	{
		"minecraft:iron_bars[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 4186,
			404: 4187,
			477: 4690,
			573: 4690,
			718: 4704,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=tall,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14626,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 6164,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14252,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11361,
		},
	},
	{
		"minecraft:tube_coral_wall_fan[facing=south,waterlogged=true]",
		nil,
		NewMapping{
			393: 8506,
			404: 8522,
			477: 9066,
			573: 9066,
			718: 9602,
		},
	},
	{
		"minecraft:iron_door[facing=north,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			404: 3312,
			477: 3815,
			573: 3815,
			718: 3817,
			393: 3311,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=south,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3937,
			404: 3938,
			477: 4441,
			573: 4441,
			718: 4455,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 6290,
		},
	},
	{
		"minecraft:detector_rail[powered=true,shape=east_west]",
		nil,
		NewMapping{
			404: 1017,
			477: 1317,
			573: 1317,
			718: 1318,
			4:   457,
			393: 1017,
		},
	},
	{
		"minecraft:chorus_plant[down=true,east=true,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 8003,
			404: 8004,
			477: 8528,
			573: 8528,
			718: 9064,
		},
	},
	{
		"minecraft:spruce_sign[rotation=11,waterlogged=true]",
		nil,
		NewMapping{
			477: 3433,
			573: 3433,
			718: 3435,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=tall,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13129,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=tall,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13345,
		},
	},
	{
		"minecraft:jungle_door[facing=east,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			477: 8380,
			573: 8380,
			718: 8916,
			4:   3130,
			393: 7855,
			404: 7856,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=6,south=side,west=up]",
		nil,
		NewMapping{
			477: 2977,
			573: 2977,
			718: 2979,
			393: 2673,
			404: 2674,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=tall,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13015,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16571,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=low,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 15974,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11102,
		},
	},
	{
		"minecraft:white_banner[rotation=2]",
		nil,
		NewMapping{
			4:   2818,
			393: 6856,
			404: 6857,
			477: 7363,
			573: 7363,
			718: 7899,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=3,south=side,west=up]",
		nil,
		NewMapping{
			393: 2070,
			404: 2071,
			477: 2374,
			573: 2374,
			718: 2376,
		},
	},
	{
		"minecraft:campfire[facing=east,lit=false,signal_fire=false,waterlogged=false]",
		nil,
		NewMapping{
			477: 11247,
			573: 11263,
			718: 14921,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=3,powered=false]",
		nil,
		NewMapping{
			573: 955,
			718: 956,
			477: 955,
		},
	},
	{
		"minecraft:repeating_command_block[conditional=false,facing=up]",
		nil,
		NewMapping{
			393: 8174,
			404: 8175,
			477: 8699,
			573: 8699,
			718: 9235,
			4:   3361,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10885,
			573: 10885,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=low,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12656,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 4131,
			404: 4132,
			477: 4635,
			573: 4635,
			718: 4649,
		},
	},
	{
		"minecraft:fire[age=5,east=false,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 1315,
			404: 1316,
			477: 1619,
			573: 1619,
			718: 1620,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=none,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16362,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12702,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=15,powered=true]",
		nil,
		NewMapping{
			404: 728,
			477: 728,
			573: 728,
			718: 729,
			393: 728,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 17025,
		},
	},
	{
		"minecraft:warped_trapdoor[facing=south,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			718: 15217,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12722,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4722,
			404: 4723,
			477: 5226,
			573: 5226,
			718: 5242,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=south,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3929,
			404: 3930,
			477: 4433,
			573: 4433,
			718: 4447,
		},
	},
	{
		"minecraft:dark_oak_wood[axis=z]",
		nil,
		NewMapping{
			393: 125,
			404: 125,
			477: 125,
			573: 125,
			718: 126,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=3,south=none,west=side]",
		nil,
		NewMapping{
			718: 3244,
			393: 2938,
			404: 2939,
			477: 3242,
			573: 3242,
		},
	},
	{
		"minecraft:twisting_vines[age=15]",
		nil,
		NewMapping{
			718: 15032,
		},
	},
	{
		"minecraft:red_bed[facing=south,occupied=false,part=head]",
		nil,
		NewMapping{
			573: 1278,
			718: 1279,
			4:   424,
			393: 978,
			404: 978,
			477: 1278,
		},
	},
	{
		"minecraft:jungle_leaves[distance=7,persistent=true]",
		nil,
		NewMapping{
			718: 199,
			393: 198,
			404: 198,
			477: 198,
			573: 198,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9443,
			573: 9443,
			718: 9979,
		},
	},
	{
		"minecraft:granite_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 10415,
			477: 9879,
			573: 9879,
		},
	},
	{
		"minecraft:sign[rotation=13,waterlogged=true]",
		nil,
		NewMapping{
			393: 3101,
		},
	},
	{
		"minecraft:jungle_fence[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 7586,
			404: 7587,
			477: 8111,
			573: 8111,
			718: 8647,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=east,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			718: 15186,
		},
	},
	{
		"minecraft:andesite_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9980,
			573: 9980,
			718: 10516,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=tall,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 6286,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=low,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12432,
		},
	},
	{
		"minecraft:repeater[delay=4,facing=south,locked=true,powered=false]",
		nil,
		NewMapping{
			573: 4070,
			718: 4084,
			393: 3566,
			404: 3567,
			477: 4070,
		},
	},
	{
		"minecraft:cyan_bed[facing=east,occupied=true,part=head]",
		nil,
		NewMapping{
			477: 1204,
			573: 1204,
			718: 1205,
			393: 904,
			404: 904,
		},
	},
	{
		"minecraft:vine[east=false,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			718: 4808,
			393: 4288,
			404: 4289,
			477: 4792,
			573: 4792,
		},
	},
	{
		"minecraft:andesite_stairs[facing=south,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9962,
			573: 9962,
			718: 10498,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12988,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=tall,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14306,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11683,
		},
	},
	{
		"minecraft:anvil[facing=west]",
		nil,
		NewMapping{
			393: 5569,
			404: 5570,
			477: 6076,
			573: 6076,
			718: 6612,
			4:   2321,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=7,south=up,west=side]",
		nil,
		NewMapping{
			393: 1960,
			404: 1961,
			477: 2264,
			573: 2264,
			718: 2266,
		},
	},
	{
		"minecraft:warped_fence[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			718: 15100,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=west,half=top,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 15159,
		},
	},
	{
		"minecraft:note_block[instrument=snare,note=2,powered=true]",
		nil,
		NewMapping{
			477: 352,
			573: 352,
			718: 353,
			393: 352,
			404: 352,
		},
	},
	{
		"minecraft:rail[shape=ascending_east]",
		nil,
		NewMapping{
			404: 3182,
			477: 3645,
			573: 3645,
			718: 3647,
			4:   1058,
			393: 3181,
		},
	},
	{
		"minecraft:crimson_fence[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			718: 15070,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5932,
			404: 5933,
			477: 6439,
			573: 6439,
			718: 6975,
		},
	},
	{
		"minecraft:jungle_button[face=floor,facing=east,powered=true]",
		nil,
		NewMapping{
			393: 5381,
			404: 5382,
			477: 5888,
			573: 5888,
			718: 6424,
		},
	},
	{
		"minecraft:scaffolding[bottom=true,distance=7,waterlogged=false]",
		nil,
		NewMapping{
			573: 11114,
			718: 14770,
			477: 11114,
		},
	},
	{
		"minecraft:andesite_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			477: 10311,
			573: 10311,
			718: 10847,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16224,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=tall,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13887,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=7,south=side,west=none]",
		nil,
		NewMapping{
			477: 2700,
			573: 2700,
			718: 2702,
			393: 2396,
			404: 2397,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=1,powered=false]",
		nil,
		NewMapping{
			393: 601,
			404: 601,
			477: 601,
			573: 601,
			718: 602,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 6649,
			573: 6649,
			718: 7185,
			393: 6142,
			404: 6143,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9407,
			573: 9407,
			718: 9943,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=none,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11220,
		},
	},
	{
		"minecraft:crimson_button[face=ceiling,facing=north,powered=true]",
		nil,
		NewMapping{
			718: 15495,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=low,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 6040,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=12,south=up,west=none]",
		nil,
		NewMapping{
			718: 2600,
			393: 2294,
			404: 2295,
			477: 2598,
			573: 2598,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			573: 8398,
			718: 8934,
			393: 7873,
			404: 7874,
			477: 8398,
		},
	},
	{
		"minecraft:stone_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			393: 7298,
			477: 7805,
			573: 7805,
			718: 8341,
			4:   688,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=6,south=up,west=up]",
		nil,
		NewMapping{
			477: 2830,
			573: 2830,
			718: 2832,
			393: 2526,
			404: 2527,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11636,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=none,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12290,
		},
	},
	{
		"minecraft:crimson_sign[rotation=2,waterlogged=false]",
		nil,
		NewMapping{
			718: 15660,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=low,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14049,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6580,
			404: 6581,
			477: 7087,
			573: 7087,
			718: 7623,
		},
	},
	{
		"minecraft:purpur_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 8154,
			404: 8155,
			477: 8679,
			573: 8679,
			718: 9215,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16384,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=low,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 6154,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=none,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16915,
		},
	},
	{
		"minecraft:bell[attachment=floor,facing=north,powered=false]",
		nil,
		NewMapping{
			573: 11199,
			718: 14855,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13475,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=none,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12919,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=tall,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16535,
		},
	},
	{
		"minecraft:iron_door[facing=south,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			404: 3320,
			477: 3823,
			573: 3823,
			718: 3825,
			393: 3319,
		},
	},
	{
		"minecraft:brain_coral_wall_fan[facing=north,waterlogged=false]",
		nil,
		NewMapping{
			393: 8513,
			404: 8529,
			477: 9073,
			573: 9073,
			718: 9609,
		},
	},
	{
		"minecraft:birch_sign[rotation=8,waterlogged=false]",
		nil,
		NewMapping{
			477: 3460,
			573: 3460,
			718: 3462,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10125,
			573: 10125,
			718: 10661,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9138,
			573: 9138,
			718: 9674,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 16693,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9747,
			573: 9747,
			718: 10283,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			573: 9145,
			718: 9681,
			477: 9145,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10110,
			573: 10110,
			718: 10646,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=none,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14137,
		},
	},
	{
		"minecraft:light_weighted_pressure_plate[power=1]",
		nil,
		NewMapping{
			477: 6111,
			573: 6111,
			718: 6647,
			4:   2353,
			393: 5604,
			404: 5605,
		},
	},
	{
		"minecraft:fire[age=2,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1201,
			404: 1202,
			477: 1505,
			573: 1505,
			718: 1506,
		},
	},
	{
		"minecraft:pink_shulker_box[facing=south]",
		nil,
		NewMapping{
			404: 8256,
			477: 8780,
			573: 8780,
			718: 9316,
			4:   3603,
			393: 8255,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4672,
			404: 4673,
			477: 5176,
			573: 5176,
			718: 5192,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13318,
		},
	},
	{
		"minecraft:blackstone_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			718: 16247,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=none,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14437,
		},
	},
	{
		"minecraft:coal_ore",
		nil,
		NewMapping{
			718: 71,
			4:   256,
			393: 71,
			404: 71,
			477: 71,
			573: 71,
		},
	},
	{
		"minecraft:magenta_banner[rotation=4]",
		nil,
		NewMapping{
			393: 6890,
			404: 6891,
			477: 7397,
			573: 7397,
			718: 7933,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5176,
			404: 5177,
			477: 5680,
			573: 5680,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=tall,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11921,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=tall,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13557,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 16274,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=13,south=up,west=up]",
		nil,
		NewMapping{
			718: 3039,
			393: 2733,
			404: 2734,
			477: 3037,
			573: 3037,
		},
	},
	{
		"minecraft:oak_fence[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 3973,
			718: 3975,
			393: 3469,
			404: 3470,
			477: 3973,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=north,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 4187,
			393: 3669,
			404: 3670,
			477: 4173,
			573: 4173,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9170,
			573: 9170,
			718: 9706,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=none,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16778,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=1,south=up,west=up]",
		nil,
		NewMapping{
			404: 2626,
			477: 2929,
			573: 2929,
			718: 2931,
			393: 2625,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=3,south=up,west=none]",
		nil,
		NewMapping{
			393: 2501,
			404: 2502,
			477: 2805,
			573: 2805,
			718: 2807,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=5,south=side,west=none]",
		nil,
		NewMapping{
			393: 2810,
			404: 2811,
			477: 3114,
			573: 3114,
			718: 3116,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11468,
		},
	},
	{
		"minecraft:light_weighted_pressure_plate[power=5]",
		nil,
		NewMapping{
			718: 6651,
			4:   2357,
			393: 5608,
			404: 5609,
			477: 6115,
			573: 6115,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=0,south=none,west=none]",
		nil,
		NewMapping{
			477: 3072,
			573: 3072,
			718: 3074,
			393: 2768,
			404: 2769,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9734,
			573: 9734,
			718: 10270,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9220,
			573: 9220,
			718: 9756,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=11,south=up,west=up]",
		nil,
		NewMapping{
			393: 2283,
			404: 2284,
			477: 2587,
			573: 2587,
			718: 2589,
		},
	},
	{
		"minecraft:fire[age=6,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			477: 1660,
			573: 1660,
			718: 1661,
			393: 1356,
			404: 1357,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=true,east=false,north=false,powered=false,south=false,west=true]",
		nil,
		NewMapping{
			718: 5305,
			393: 4785,
			404: 4786,
			477: 5289,
			573: 5289,
		},
	},
	{
		"minecraft:fire[age=7,east=true,north=true,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1365,
			404: 1366,
			477: 1669,
			573: 1669,
			718: 1670,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14428,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11589,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=tall,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12915,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=none,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14656,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14516,
		},
	},
	{
		"minecraft:fire[age=7,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			404: 1376,
			477: 1679,
			573: 1679,
			718: 1680,
			393: 1375,
		},
	},
	{
		"minecraft:oak_door[facing=west,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 3149,
			404: 3150,
			477: 3613,
			573: 3613,
			718: 3615,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=3,south=up,west=up]",
		nil,
		NewMapping{
			393: 2643,
			404: 2644,
			477: 2947,
			573: 2947,
			718: 2949,
		},
	},
	{
		"minecraft:cobblestone",
		nil,
		NewMapping{
			477: 14,
			573: 14,
			718: 14,
			4:   64,
			393: 14,
			404: 14,
		},
	},
	{
		"minecraft:repeater[delay=4,facing=north,locked=true,powered=false]",
		nil,
		NewMapping{
			404: 3563,
			477: 4066,
			573: 4066,
			718: 4080,
			393: 3562,
		},
	},
	{
		"minecraft:dark_oak_door[facing=north,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 7945,
			404: 7946,
			477: 8470,
			573: 8470,
			718: 9006,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=4,south=none,west=none]",
		nil,
		NewMapping{
			477: 3252,
			573: 3252,
			718: 3254,
			4:   884,
			393: 2948,
			404: 2949,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16230,
		},
	},
	{
		"minecraft:daylight_detector[inverted=false,power=14]",
		nil,
		NewMapping{
			573: 6188,
			718: 6724,
			4:   2430,
			393: 5681,
			404: 5682,
			477: 6188,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16818,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=3,south=up,west=up]",
		nil,
		NewMapping{
			393: 2211,
			404: 2212,
			477: 2515,
			573: 2515,
			718: 2517,
		},
	},
	{
		"minecraft:fire[age=7,east=true,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			573: 1671,
			718: 1672,
			393: 1367,
			404: 1368,
			477: 1671,
		},
	},
	{
		"minecraft:nether_brick_fence[east=true,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			404: 4504,
			477: 5007,
			573: 5007,
			718: 5023,
			393: 4503,
		},
	},
	{
		"minecraft:jungle_sign[rotation=9,waterlogged=false]",
		nil,
		NewMapping{
			477: 3526,
			573: 3526,
			718: 3528,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16822,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=tall,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11062,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=none,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14126,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=tall,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12372,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=none,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12514,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=tall,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11492,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=3,south=up,west=side]",
		nil,
		NewMapping{
			393: 2068,
			404: 2069,
			477: 2372,
			573: 2372,
			718: 2374,
		},
	},
	{
		"minecraft:lime_banner[rotation=4]",
		nil,
		NewMapping{
			393: 6938,
			404: 6939,
			477: 7445,
			573: 7445,
			718: 7981,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9654,
			573: 9654,
			718: 10190,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9218,
			573: 9218,
			718: 9754,
		},
	},
	{
		"minecraft:piston_head[facing=south,short=false,type=normal]",
		nil,
		NewMapping{
			393: 1069,
			404: 1069,
			477: 1369,
			573: 1369,
			718: 1370,
			4:   547,
		},
	},
	{
		"minecraft:furnace[facing=north,lit=false]",
		nil,
		NewMapping{
			718: 3374,
			4:   978,
			393: 3068,
			404: 3069,
			477: 3372,
			573: 3372,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10028,
			573: 10028,
			718: 10564,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=13,powered=true]",
		nil,
		NewMapping{
			477: 874,
			573: 874,
			718: 875,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=11,south=none,west=up]",
		nil,
		NewMapping{
			718: 3171,
			393: 2865,
			404: 2866,
			477: 3169,
			573: 3169,
		},
	},
	{
		"minecraft:scaffolding[bottom=true,distance=6,waterlogged=true]",
		nil,
		NewMapping{
			477: 11111,
			573: 11111,
			718: 14767,
		},
	},
	{
		"minecraft:fire[age=2,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1225,
			404: 1226,
			477: 1529,
			573: 1529,
			718: 1530,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=true,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			573: 8570,
			718: 9106,
			393: 8045,
			404: 8046,
			477: 8570,
		},
	},
	{
		"minecraft:iron_door[facing=south,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			393: 3331,
			404: 3332,
			477: 3835,
			573: 3835,
			718: 3837,
		},
	},
	{
		"minecraft:acacia_button[face=ceiling,facing=west,powered=true]",
		nil,
		NewMapping{
			393: 5419,
			404: 5420,
			477: 5926,
			573: 5926,
			718: 6462,
		},
	},
	{
		"minecraft:lime_glazed_terracotta[facing=east]",
		nil,
		NewMapping{
			393: 8336,
			404: 8337,
			477: 8861,
			573: 8861,
			718: 9397,
			4:   3843,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=low,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12983,
		},
	},
	{
		"minecraft:dark_oak_wall_sign[facing=north,waterlogged=true]",
		nil,
		NewMapping{
			477: 3773,
			573: 3773,
			718: 3775,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9219,
			573: 9219,
			718: 9755,
		},
	},
	{
		"minecraft:bee_nest[facing=west,honey_level=5]",
		nil,
		NewMapping{
			718: 15793,
			573: 11304,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=none,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13593,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=1,south=up,west=up]",
		nil,
		NewMapping{
			393: 2193,
			404: 2194,
			477: 2497,
			573: 2497,
			718: 2499,
		},
	},
	{
		"minecraft:zombie_head[rotation=0]",
		nil,
		NewMapping{
			393: 5491,
			404: 5492,
			477: 5994,
			573: 5994,
			718: 6530,
		},
	},
	{
		"minecraft:magenta_glazed_terracotta[facing=west]",
		nil,
		NewMapping{
			477: 8848,
			573: 8848,
			718: 9384,
			4:   3793,
			393: 8323,
			404: 8324,
		},
	},
	{
		"minecraft:note_block[instrument=basedrum,note=10,powered=true]",
		nil,
		NewMapping{
			573: 318,
			718: 319,
			393: 318,
			404: 318,
			477: 318,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=none,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14114,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=tall,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14206,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=15,south=none,west=up]",
		nil,
		NewMapping{
			393: 2181,
			404: 2182,
			477: 2485,
			573: 2485,
			718: 2487,
		},
	},
	{
		"minecraft:quartz_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 6285,
			573: 6285,
			718: 6821,
			393: 5778,
			404: 5779,
		},
	},
	{
		"minecraft:green_shulker_box[facing=down]",
		nil,
		NewMapping{
			718: 9361,
			4:   3712,
			393: 8300,
			404: 8301,
			477: 8825,
			573: 8825,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9814,
			573: 9814,
			718: 10350,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 6393,
			573: 6393,
			718: 6929,
			393: 5886,
			404: 5887,
		},
	},
	{
		"minecraft:fire[age=13,east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			404: 1569,
			477: 1872,
			573: 1872,
			718: 1873,
			393: 1568,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=3,south=side,west=none]",
		nil,
		NewMapping{
			404: 2217,
			477: 2520,
			573: 2520,
			718: 2522,
			393: 2216,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2048,
			393: 4721,
			404: 4722,
			477: 5225,
			573: 5225,
			718: 5241,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=east,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			477: 4471,
			573: 4471,
			718: 4485,
			393: 3967,
			404: 3968,
		},
	},
	{
		"minecraft:stripped_jungle_wood[axis=x]",
		nil,
		NewMapping{
			477: 135,
			573: 135,
			718: 136,
			393: 135,
			404: 135,
		},
	},
	{
		"minecraft:weeping_vines[age=19]",
		nil,
		NewMapping{
			718: 15009,
		},
	},
	{
		"minecraft:light_weighted_pressure_plate[power=15]",
		nil,
		NewMapping{
			718: 6661,
			4:   2367,
			393: 5618,
			404: 5619,
			477: 6125,
			573: 6125,
		},
	},
	{
		"minecraft:diorite_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10210,
			573: 10210,
			718: 10746,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11500,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13744,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=8,powered=false]",
		nil,
		NewMapping{
			393: 665,
			404: 665,
			477: 665,
			573: 665,
			718: 666,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=north,half=top,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3664,
			404: 3665,
			477: 4168,
			573: 4168,
			718: 4182,
		},
	},
	{
		"minecraft:piston_head[facing=up,short=false,type=normal]",
		nil,
		NewMapping{
			477: 1377,
			573: 1377,
			718: 1378,
			4:   545,
			393: 1077,
			404: 1077,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=9,south=side,west=up]",
		nil,
		NewMapping{
			393: 1980,
			404: 1981,
			477: 2284,
			573: 2284,
			718: 2286,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=none,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16563,
		},
	},
	{
		"minecraft:dark_oak_fence[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 8181,
			573: 8181,
			718: 8717,
			393: 7656,
			404: 7657,
		},
	},
	{
		"minecraft:petrified_oak_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			393: 7305,
			404: 7306,
			477: 7824,
			573: 7824,
			718: 8360,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 10788,
			477: 10252,
			573: 10252,
		},
	},
	{
		"minecraft:polished_diorite_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			477: 10276,
			573: 10276,
			718: 10812,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13437,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=low,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11583,
		},
	},
	{
		"minecraft:crimson_button[face=wall,facing=south,powered=true]",
		nil,
		NewMapping{
			718: 15489,
		},
	},
	{
		"minecraft:birch_button[face=ceiling,facing=east,powered=false]",
		nil,
		NewMapping{
			393: 5374,
			404: 5375,
			477: 5881,
			573: 5881,
			718: 6417,
		},
	},
	{
		"minecraft:stone_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 9674,
			718: 10210,
			477: 9674,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=none,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13811,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=low,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14251,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9658,
			573: 9658,
			718: 10194,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 5617,
			718: 5633,
			393: 5113,
			404: 5114,
			477: 5617,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=south,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			573: 7021,
			718: 7557,
			393: 6514,
			404: 6515,
			477: 7021,
		},
	},
	{
		"minecraft:vine[east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 4298,
			404: 4299,
			477: 4802,
			573: 4802,
			718: 4818,
			4:   1698,
		},
	},
	{
		"minecraft:carrots[age=7]",
		nil,
		NewMapping{
			393: 5294,
			404: 5295,
			477: 5801,
			573: 5801,
			718: 6337,
			4:   2263,
		},
	},
	{
		"minecraft:sweet_berry_bush[age=1]",
		nil,
		NewMapping{
			573: 11265,
			718: 14955,
			477: 11249,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=low,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12314,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=low,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12960,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12269,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			718: 7014,
			393: 5971,
			404: 5972,
			477: 6478,
			573: 6478,
		},
	},
	{
		"minecraft:oak_fence[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 3482,
			477: 3985,
			573: 3985,
			718: 3987,
			393: 3481,
		},
	},
	{
		"minecraft:diorite",
		nil,
		NewMapping{
			393: 4,
			404: 4,
			477: 4,
			573: 4,
			718: 4,
			4:   19,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10968,
			573: 10968,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=none,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11761,
		},
	},
	{
		"minecraft:vine[east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 4287,
			404: 4288,
			477: 4791,
			573: 4791,
			718: 4807,
			4:   1701,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 9653,
			718: 10189,
			477: 9653,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16570,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16990,
		},
	},
	{
		"minecraft:fire[age=6,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1341,
			404: 1342,
			477: 1645,
			573: 1645,
			718: 1646,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=5,south=side,west=side]",
		nil,
		NewMapping{
			393: 1801,
			404: 1802,
			477: 2105,
			573: 2105,
			718: 2107,
		},
	},
	{
		"minecraft:dark_oak_wall_sign[facing=west,waterlogged=false]",
		nil,
		NewMapping{
			477: 3778,
			573: 3778,
			718: 3780,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12997,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10475,
			477: 10475,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5958,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 6742,
			404: 6743,
			477: 7249,
			573: 7249,
			718: 7785,
		},
	},
	{
		"minecraft:brewing_stand[has_bottle_0=true,has_bottle_1=true,has_bottle_2=false]",
		nil,
		NewMapping{
			4:   1875,
			393: 4614,
			404: 4615,
			477: 5118,
			573: 5118,
			718: 5134,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=9,south=none,west=side]",
		nil,
		NewMapping{
			393: 1840,
			404: 1841,
			477: 2144,
			573: 2144,
			718: 2146,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=12,south=up,west=up]",
		nil,
		NewMapping{
			393: 3012,
			404: 3013,
			477: 3316,
			573: 3316,
			718: 3318,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=low,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13079,
		},
	},
	{
		"minecraft:end_portal_frame[eye=false,facing=south]",
		nil,
		NewMapping{
			718: 5151,
			4:   1920,
			393: 4631,
			404: 4632,
			477: 5135,
			573: 5135,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9759,
			573: 9759,
			718: 10295,
		},
	},
	{
		"minecraft:birch_sign[rotation=13,waterlogged=false]",
		nil,
		NewMapping{
			477: 3470,
			573: 3470,
			718: 3472,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=tall,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13227,
		},
	},
	{
		"minecraft:chain_command_block[conditional=true,facing=down]",
		nil,
		NewMapping{
			573: 8706,
			718: 9242,
			4:   3384,
			393: 8181,
			404: 8182,
			477: 8706,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=none,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13057,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=tall,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16652,
		},
	},
	{
		"minecraft:orange_bed[facing=north,occupied=false,part=head]",
		nil,
		NewMapping{
			404: 766,
			477: 1066,
			573: 1066,
			718: 1067,
			393: 766,
		},
	},
	{
		"minecraft:fire[age=2,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			573: 1506,
			718: 1507,
			393: 1202,
			404: 1203,
			477: 1506,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10545,
			573: 10545,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 9668,
			718: 10204,
			477: 9668,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 9714,
			718: 10250,
			477: 9714,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=low,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14482,
		},
	},
	{
		"minecraft:magenta_bed[facing=east,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 793,
			404: 793,
			477: 1093,
			573: 1093,
			718: 1094,
		},
	},
	{
		"minecraft:comparator[facing=east,mode=compare,powered=true]",
		nil,
		NewMapping{
			393: 5647,
			404: 5648,
			477: 6154,
			573: 6154,
			718: 6690,
			4:   2411,
		},
	},
	{
		"minecraft:lilac[half=lower]",
		nil,
		NewMapping{
			393: 6845,
			404: 6846,
			477: 7352,
			573: 7352,
			718: 7888,
			4:   2801,
		},
	},
	{
		"minecraft:jungle_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 5603,
			718: 5619,
			393: 5099,
			404: 5100,
			477: 5603,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=south,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 16282,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=none,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16063,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11696,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=tall,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12574,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14228,
		},
	},
	{
		"minecraft:jungle_door[facing=south,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7826,
			404: 7827,
			477: 8351,
			573: 8351,
			718: 8887,
		},
	},
	{
		"minecraft:glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 4228,
			404: 4229,
			477: 4732,
			573: 4732,
			718: 4748,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=15,south=side,west=side]",
		nil,
		NewMapping{
			393: 2323,
			404: 2324,
			477: 2627,
			573: 2627,
			718: 2629,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10461,
			573: 10461,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=none,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 15954,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9268,
			573: 9268,
			718: 9804,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=south,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			404: 7436,
			477: 7960,
			573: 7960,
			718: 8496,
			4:   2968,
			393: 7435,
		},
	},
	{
		"minecraft:daylight_detector[inverted=false,power=2]",
		nil,
		NewMapping{
			404: 5670,
			477: 6176,
			573: 6176,
			718: 6712,
			4:   2418,
			393: 5669,
		},
	},
	{
		"minecraft:nether_wart_block",
		nil,
		NewMapping{
			573: 8718,
			718: 9254,
			4:   3424,
			393: 8193,
			404: 8194,
			477: 8718,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=east,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 4282,
			573: 4282,
			718: 4296,
			393: 3778,
			404: 3779,
		},
	},
	{
		"minecraft:yellow_glazed_terracotta[facing=south]",
		nil,
		NewMapping{
			477: 8855,
			573: 8855,
			718: 9391,
			4:   3824,
			393: 8330,
			404: 8331,
		},
	},
	{
		"minecraft:black_wall_banner[facing=east]",
		nil,
		NewMapping{
			404: 7174,
			477: 7680,
			573: 7680,
			718: 8216,
			393: 7173,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10938,
			573: 10938,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11917,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=8,powered=true]",
		nil,
		NewMapping{
			718: 865,
			477: 864,
			573: 864,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10744,
			573: 10744,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=6,powered=true]",
		nil,
		NewMapping{
			477: 810,
			573: 810,
			718: 811,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16592,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=tall,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14641,
		},
	},
	{
		"minecraft:warped_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 15454,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12343,
		},
	},
	{
		"minecraft:dark_oak_trapdoor[facing=south,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			477: 4438,
			573: 4438,
			718: 4452,
			393: 3934,
			404: 3935,
		},
	},
	{
		"minecraft:observer[facing=north,powered=false]",
		nil,
		NewMapping{
			4:   3490,
			393: 8200,
			404: 8201,
			477: 8725,
			573: 8725,
			718: 9261,
		},
	},
	{
		"minecraft:oak_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 2022,
			573: 2022,
			718: 2024,
			393: 1718,
			404: 1719,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=7,powered=false]",
		nil,
		NewMapping{
			718: 964,
			477: 963,
			573: 963,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=tall,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14638,
		},
	},
	{
		"minecraft:purple_wall_banner[facing=north]",
		nil,
		NewMapping{
			573: 7657,
			718: 8193,
			393: 7150,
			404: 7151,
			477: 7657,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=15,powered=true]",
		nil,
		NewMapping{
			477: 878,
			573: 878,
			718: 879,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11747,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11908,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=none,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11308,
		},
	},
	{
		"minecraft:cocoa[age=0,facing=south]",
		nil,
		NewMapping{
			404: 4640,
			477: 5143,
			573: 5143,
			718: 5159,
			4:   2032,
			393: 4639,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 10263,
			477: 9727,
			573: 9727,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=none,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13783,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=none,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16590,
		},
	},
	{
		"minecraft:light_gray_bed[facing=west,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 884,
			404: 884,
			477: 1184,
			573: 1184,
			718: 1185,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			404: 4085,
			477: 4588,
			573: 4588,
			718: 4602,
			393: 4084,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13754,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12995,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14043,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13906,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 5997,
			477: 6503,
			573: 6503,
			718: 7039,
			393: 5996,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   1074,
			393: 3220,
			404: 3221,
			477: 3684,
			573: 3684,
			718: 3686,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=5,south=up,west=up]",
		nil,
		NewMapping{
			573: 2965,
			718: 2967,
			393: 2661,
			404: 2662,
			477: 2965,
		},
	},
	{
		"minecraft:purpur_stairs[facing=east,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 8147,
			404: 8148,
			477: 8672,
			573: 8672,
			718: 9208,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=13,south=none,west=up]",
		nil,
		NewMapping{
			393: 3027,
			404: 3028,
			477: 3331,
			573: 3331,
			718: 3333,
		},
	},
	{
		"minecraft:smooth_red_sandstone_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9253,
			573: 9253,
			718: 9789,
		},
	},
	{
		"minecraft:diorite_wall[east=false,north=false,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 11084,
			573: 11084,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 10558,
			477: 10022,
			573: 10022,
		},
	},
	{
		"minecraft:fire[age=15,east=false,north=false,south=true,up=false,west=true]",
		nil,
		NewMapping{
			718: 1946,
			393: 1641,
			404: 1642,
			477: 1945,
			573: 1945,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 6325,
			477: 6831,
			573: 6831,
			718: 7367,
			393: 6324,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6233,
			404: 6234,
			477: 6740,
			573: 6740,
			718: 7276,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=6,powered=true]",
		nil,
		NewMapping{
			404: 510,
			477: 510,
			573: 510,
			718: 511,
			393: 510,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11823,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=tall,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12786,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=low,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 15987,
		},
	},
	{
		"minecraft:potted_birch_sapling",
		nil,
		NewMapping{
			477: 5772,
			573: 5772,
			718: 6308,
			393: 5268,
			404: 5269,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=tall,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16858,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13074,
		},
	},
	{
		"minecraft:diorite_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10218,
			573: 10218,
			718: 10754,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=low,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12864,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 4700,
			477: 5203,
			573: 5203,
			718: 5219,
			393: 4699,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10923,
			573: 10923,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=low,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 6130,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13153,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=low,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13955,
		},
	},
	{
		"minecraft:fire[age=14,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			404: 1613,
			477: 1916,
			573: 1916,
			718: 1917,
			393: 1612,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=low,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11365,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=low,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12332,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=none,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16565,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 10128,
			718: 10664,
			477: 10128,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14291,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=10,south=side,west=side]",
		nil,
		NewMapping{
			393: 2998,
			404: 2999,
			477: 3302,
			573: 3302,
			718: 3304,
		},
	},
	{
		"minecraft:lime_concrete",
		nil,
		NewMapping{
			573: 8907,
			718: 9443,
			4:   4021,
			393: 8382,
			404: 8383,
			477: 8907,
		},
	},
	{
		"minecraft:stripped_jungle_log[axis=y]",
		nil,
		NewMapping{
			404: 97,
			477: 97,
			573: 97,
			718: 98,
			393: 97,
		},
	},
	{
		"minecraft:jungle_door[facing=south,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			718: 8893,
			393: 7832,
			404: 7833,
			477: 8357,
			573: 8357,
		},
	},
	{
		"minecraft:magenta_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5904,
			404: 5905,
			477: 6411,
			573: 6411,
			718: 6947,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=0,south=none,west=side]",
		nil,
		NewMapping{
			393: 1759,
			404: 1760,
			477: 2063,
			573: 2063,
			718: 2065,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=none,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13462,
		},
	},
	{
		"minecraft:quartz_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 5733,
			477: 6239,
			573: 6239,
			718: 6775,
			393: 5732,
		},
	},
	{
		"minecraft:spruce_door[facing=east,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			477: 8250,
			573: 8250,
			718: 8786,
			393: 7725,
			404: 7726,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=11,south=up,west=none]",
		nil,
		NewMapping{
			477: 2733,
			573: 2733,
			718: 2735,
			393: 2429,
			404: 2430,
		},
	},
	{
		"minecraft:campfire[facing=west,lit=false,signal_fire=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 11236,
			573: 11252,
			718: 14910,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=west,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 4151,
			393: 3633,
			404: 3634,
			477: 4137,
			573: 4137,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			573: 6512,
			718: 7048,
			393: 6005,
			404: 6006,
			477: 6512,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14359,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5178,
			404: 5179,
			477: 5682,
			573: 5682,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 5454,
			393: 4934,
			404: 4935,
			477: 5438,
			573: 5438,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10502,
			573: 10502,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 10936,
		},
	},
	{
		"minecraft:gray_glazed_terracotta[facing=north]",
		nil,
		NewMapping{
			4:   3874,
			393: 8341,
			404: 8342,
			477: 8866,
			573: 8866,
			718: 9402,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=none,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14111,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=8,powered=true]",
		nil,
		NewMapping{
			393: 564,
			404: 564,
			477: 564,
			573: 564,
			718: 565,
		},
	},
	{
		"minecraft:cocoa[age=2,facing=north]",
		nil,
		NewMapping{
			393: 4646,
			404: 4647,
			477: 5150,
			573: 5150,
			718: 5166,
			4:   2042,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=tall,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13344,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13745,
		},
	},
	{
		"minecraft:andesite_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9991,
			573: 9991,
			718: 10527,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=none,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11195,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=none,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16342,
		},
	},
	{
		"minecraft:note_block[instrument=hat,note=13,powered=true]",
		nil,
		NewMapping{
			393: 424,
			404: 424,
			477: 424,
			573: 424,
			718: 425,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6277,
			404: 6278,
			477: 6784,
			573: 6784,
			718: 7320,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=true,north=true,powered=false,south=true,west=true]",
		nil,
		NewMapping{
			393: 4791,
			404: 4792,
			477: 5295,
			573: 5295,
			718: 5311,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10942,
			573: 10942,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=east,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			573: 4275,
			718: 4289,
			393: 3771,
			404: 3772,
			477: 4275,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11156,
		},
	},
	{
		"minecraft:fire[age=14,east=true,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			477: 1899,
			573: 1899,
			718: 1900,
			393: 1595,
			404: 1596,
		},
	},
	{
		"minecraft:fire[age=2,east=false,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			718: 1529,
			393: 1224,
			404: 1225,
			477: 1528,
			573: 1528,
		},
	},
	{
		"minecraft:black_banner[rotation=11]",
		nil,
		NewMapping{
			393: 7105,
			404: 7106,
			477: 7612,
			573: 7612,
			718: 8148,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 6595,
			477: 7101,
			573: 7101,
			718: 7637,
			393: 6594,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5857,
			404: 5858,
			477: 6364,
			573: 6364,
			718: 6900,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=none,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 15934,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=low,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11900,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=6,powered=true]",
		nil,
		NewMapping{
			477: 910,
			573: 910,
			718: 911,
		},
	},
	{
		"minecraft:note_block[instrument=bit,note=19,powered=false]",
		nil,
		NewMapping{
			477: 937,
			573: 937,
			718: 938,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13431,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10867,
			573: 10867,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10740,
			573: 10740,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=north,half=top,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			718: 15133,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13284,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=false,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 6289,
			477: 6795,
			573: 6795,
			718: 7331,
			393: 6288,
		},
	},
	{
		"minecraft:dark_oak_button[face=floor,facing=south,powered=false]",
		nil,
		NewMapping{
			573: 5933,
			718: 6469,
			393: 5426,
			404: 5427,
			477: 5933,
		},
	},
	{
		"minecraft:birch_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 5011,
			477: 5514,
			573: 5514,
			718: 5530,
			393: 5010,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9858,
			573: 9858,
			718: 10394,
		},
	},
	{
		"minecraft:warped_door[facing=north,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			718: 15595,
		},
	},
	{
		"minecraft:pink_glazed_terracotta[facing=north]",
		nil,
		NewMapping{
			393: 8337,
			404: 8338,
			477: 8862,
			573: 8862,
			718: 9398,
			4:   3858,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=8,south=side,west=side]",
		nil,
		NewMapping{
			393: 1972,
			404: 1973,
			477: 2276,
			573: 2276,
			718: 2278,
		},
	},
	{
		"minecraft:sign[rotation=1,waterlogged=false]",
		nil,
		NewMapping{
			393: 3078,
			4:   1009,
		},
	},
	{
		"minecraft:chiseled_sandstone",
		nil,
		NewMapping{
			393: 246,
			404: 246,
			477: 246,
			573: 246,
			718: 247,
			4:   385,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=3,south=side,west=side]",
		nil,
		NewMapping{
			393: 1927,
			404: 1928,
			477: 2231,
			573: 2231,
			718: 2233,
		},
	},
	{
		"minecraft:black_banner[rotation=6]",
		nil,
		NewMapping{
			404: 7101,
			477: 7607,
			573: 7607,
			718: 8143,
			393: 7100,
		},
	},
	{
		"minecraft:dark_oak_door[facing=east,half=upper,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			573: 8510,
			718: 9046,
			393: 7985,
			404: 7986,
			477: 8510,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=4,south=none,west=up]",
		nil,
		NewMapping{
			718: 2676,
			393: 2370,
			404: 2371,
			477: 2674,
			573: 2674,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=tall,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16011,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=tall,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13110,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10338,
			573: 10338,
		},
	},
	{
		"minecraft:crimson_sign[rotation=1,waterlogged=true]",
		nil,
		NewMapping{
			718: 15657,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5959,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=south,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 7175,
			573: 7175,
			718: 7711,
			393: 6668,
			404: 6669,
		},
	},
	{
		"minecraft:shulker_box[facing=west]",
		nil,
		NewMapping{
			573: 8739,
			718: 9275,
			393: 8214,
			404: 8215,
			477: 8739,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 3257,
			404: 3258,
			477: 3721,
			573: 3721,
			718: 3723,
		},
	},
	{
		"minecraft:black_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 6326,
			477: 6832,
			573: 6832,
			718: 7368,
			393: 6325,
		},
	},
	{
		"minecraft:warped_trapdoor[facing=west,half=bottom,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			718: 15236,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14739,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=tall,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13235,
		},
	},
	{
		"minecraft:green_glazed_terracotta[facing=south]",
		nil,
		NewMapping{
			393: 8366,
			404: 8367,
			477: 8891,
			573: 8891,
			718: 9427,
			4:   3968,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16215,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=tall,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12802,
		},
	},
	{
		"minecraft:crimson_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 15359,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6686,
			404: 6687,
			477: 7193,
			573: 7193,
			718: 7729,
		},
	},
	{
		"minecraft:yellow_wall_banner[facing=east]",
		nil,
		NewMapping{
			718: 8172,
			393: 7129,
			404: 7130,
			477: 7636,
			573: 7636,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10032,
			573: 10032,
			718: 10568,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16156,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=low,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11555,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=west,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 16302,
		},
	},
	{
		"minecraft:note_block[instrument=bell,note=17,powered=false]",
		nil,
		NewMapping{
			393: 583,
			404: 583,
			477: 583,
			573: 583,
			718: 584,
		},
	},
	{
		"minecraft:warped_sign[rotation=2,waterlogged=true]",
		nil,
		NewMapping{
			718: 15691,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=tall,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16003,
		},
	},
	{
		"minecraft:warped_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15456,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9777,
			573: 9777,
			718: 10313,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=tall,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16126,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=low,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 17043,
		},
	},
	{
		"minecraft:oak_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 2028,
			393: 1722,
			404: 1723,
			477: 2026,
			573: 2026,
		},
	},
	{
		"minecraft:dark_oak_log[axis=y]",
		nil,
		NewMapping{
			393: 88,
			404: 88,
			477: 88,
			573: 88,
			718: 89,
			4:   2593,
		},
	},
	{
		"minecraft:stripped_dark_oak_wood[axis=y]",
		nil,
		NewMapping{
			393: 142,
			404: 142,
			477: 142,
			573: 142,
			718: 143,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=13,south=none,west=side]",
		nil,
		NewMapping{
			393: 2740,
			404: 2741,
			477: 3044,
			573: 3044,
			718: 3046,
		},
	},
	{
		"minecraft:yellow_banner[rotation=1]",
		nil,
		NewMapping{
			477: 7426,
			573: 7426,
			718: 7962,
			393: 6919,
			404: 6920,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9542,
			573: 9542,
			718: 10078,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12720,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13047,
		},
	},
	{
		"minecraft:smooth_sandstone",
		nil,
		NewMapping{
			477: 7879,
			573: 7879,
			718: 8415,
			4:   697,
			393: 7354,
			404: 7355,
		},
	},
	{
		"minecraft:spruce_door[facing=north,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7684,
			404: 7685,
			477: 8209,
			573: 8209,
			718: 8745,
		},
	},
	{
		"minecraft:fire[age=13,east=false,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			404: 1570,
			477: 1873,
			573: 1873,
			718: 1874,
			393: 1569,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=false,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			718: 4551,
			393: 4033,
			404: 4034,
			477: 4537,
			573: 4537,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=low,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12771,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12286,
		},
	},
	{
		"minecraft:fire[age=0,east=true,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			477: 1452,
			573: 1452,
			718: 1453,
			393: 1148,
			404: 1149,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=3,south=up,west=none]",
		nil,
		NewMapping{
			404: 2358,
			477: 2661,
			573: 2661,
			718: 2663,
			393: 2357,
		},
	},
	{
		"minecraft:light_gray_shulker_box[facing=south]",
		nil,
		NewMapping{
			477: 8792,
			573: 8792,
			718: 9328,
			4:   3635,
			393: 8267,
			404: 8268,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=none,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14346,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=tall,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16439,
		},
	},
	{
		"minecraft:fire[age=11,east=true,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			477: 1796,
			573: 1796,
			718: 1797,
			393: 1492,
			404: 1493,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6255,
			404: 6256,
			477: 6762,
			573: 6762,
			718: 7298,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=none,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12940,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12992,
		},
	},
	{
		"minecraft:oak_door[facing=west,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			393: 3154,
			404: 3155,
			477: 3618,
			573: 3618,
			718: 3620,
			4:   1026,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 4065,
			404: 4066,
			477: 4569,
			573: 4569,
			718: 4583,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 16727,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=east,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 15179,
		},
	},
	{
		"minecraft:sticky_piston[extended=false,facing=north]",
		nil,
		NewMapping{
			393: 1034,
			404: 1034,
			477: 1334,
			573: 1334,
			718: 1335,
			4:   466,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 5226,
			393: 4706,
			404: 4707,
			477: 5210,
			573: 5210,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=none,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12947,
		},
	},
	{
		"minecraft:pink_bed[facing=east,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 856,
			404: 856,
			477: 1156,
			573: 1156,
			718: 1157,
		},
	},
	{
		"minecraft:light_blue_banner[rotation=2]",
		nil,
		NewMapping{
			404: 6905,
			477: 7411,
			573: 7411,
			718: 7947,
			393: 6904,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=10,south=side,west=up]",
		nil,
		NewMapping{
			718: 3015,
			393: 2709,
			404: 2710,
			477: 3013,
			573: 3013,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6770,
			404: 6771,
			477: 7277,
			573: 7277,
			718: 7813,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=north,half=bottom,open=false,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3800,
			404: 3801,
			477: 4304,
			573: 4304,
			718: 4318,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=19,powered=true]",
		nil,
		NewMapping{
			477: 836,
			573: 836,
			718: 837,
		},
	},
	{
		"minecraft:dark_oak_sign[rotation=6,waterlogged=false]",
		nil,
		NewMapping{
			477: 3552,
			573: 3552,
			718: 3554,
		},
	},
	{
		"minecraft:bell[attachment=floor,facing=south,powered=true]",
		nil,
		NewMapping{
			573: 11200,
			718: 14856,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 7830,
			393: 6787,
			404: 6788,
			477: 7294,
			573: 7294,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=14,south=up,west=side]",
		nil,
		NewMapping{
			393: 2455,
			404: 2456,
			477: 2759,
			573: 2759,
			718: 2761,
		},
	},
	{
		"minecraft:stone_button[face=wall,facing=east,powered=false]",
		nil,
		NewMapping{
			4:   1233,
			393: 3406,
			404: 3407,
			477: 3910,
			573: 3910,
			718: 3912,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13369,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=tall,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11398,
		},
	},
	{
		"minecraft:dark_oak_button[face=floor,facing=north,powered=true]",
		nil,
		NewMapping{
			393: 5423,
			404: 5424,
			477: 5930,
			573: 5930,
			718: 6466,
		},
	},
	{
		"minecraft:purpur_stairs[facing=north,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 8079,
			404: 8080,
			477: 8604,
			573: 8604,
			718: 9140,
		},
	},
	{
		"minecraft:fire[age=9,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1437,
			404: 1438,
			477: 1741,
			573: 1741,
			718: 1742,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 5812,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=north,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			718: 4249,
			393: 3731,
			404: 3732,
			477: 4235,
			573: 4235,
		},
	},
	{
		"minecraft:note_block[instrument=banjo,note=12,powered=true]",
		nil,
		NewMapping{
			477: 972,
			573: 972,
			718: 973,
		},
	},
	{
		"minecraft:jungle_sign[rotation=2,waterlogged=false]",
		nil,
		NewMapping{
			477: 3512,
			573: 3512,
			718: 3514,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13953,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5225,
			404: 5226,
			477: 5729,
			573: 5729,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=4,south=none,west=up]",
		nil,
		NewMapping{
			393: 2082,
			404: 2083,
			477: 2386,
			573: 2386,
			718: 2388,
		},
	},
	{
		"minecraft:oak_stairs[facing=south,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 1685,
			404: 1686,
			477: 1989,
			573: 1989,
			718: 1991,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=true,east=true,north=false,powered=true,south=false,west=true]",
		nil,
		NewMapping{
			393: 4829,
			404: 4830,
			477: 5333,
			573: 5333,
			718: 5349,
		},
	},
	{
		"minecraft:kelp[age=17]",
		nil,
		NewMapping{
			477: 8951,
			573: 8951,
			718: 9487,
			393: 8426,
			404: 8427,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 6524,
			573: 6524,
			718: 7060,
			393: 6017,
			404: 6018,
		},
	},
	{
		"minecraft:campfire[facing=south,lit=true,signal_fire=true,waterlogged=true]",
		nil,
		NewMapping{
			573: 11240,
			718: 14898,
			477: 11224,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11393,
		},
	},
	{
		"minecraft:magenta_stained_glass",
		nil,
		NewMapping{
			393: 3579,
			404: 3580,
			477: 4083,
			573: 4083,
			718: 4097,
			4:   1522,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=tall,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 6193,
		},
	},
	{
		"minecraft:gray_shulker_box[facing=north]",
		nil,
		NewMapping{
			477: 8784,
			573: 8784,
			718: 9320,
			4:   3618,
			393: 8259,
			404: 8260,
		},
	},
	{
		"minecraft:iron_bars[east=false,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 4205,
			404: 4206,
			477: 4709,
			573: 4709,
			718: 4723,
		},
	},
	{
		"minecraft:repeater[delay=3,facing=west,locked=false,powered=true]",
		nil,
		NewMapping{
			393: 3555,
			404: 3556,
			477: 4059,
			573: 4059,
			718: 4073,
			4:   1513,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=11,south=up,west=none]",
		nil,
		NewMapping{
			477: 2589,
			573: 2589,
			718: 2591,
			393: 2285,
			404: 2286,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 6433,
			573: 6433,
			718: 6969,
			393: 5926,
			404: 5927,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 6009,
			404: 6010,
			477: 6516,
			573: 6516,
			718: 7052,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14149,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=14,south=none,west=none]",
		nil,
		NewMapping{
			393: 2462,
			404: 2463,
			477: 2766,
			573: 2766,
			718: 2768,
		},
	},
	{
		"minecraft:yellow_terracotta",
		nil,
		NewMapping{
			4:   2548,
			393: 5808,
			404: 5809,
			477: 6315,
			573: 6315,
			718: 6851,
		},
	},
	{
		"minecraft:melon_stem[age=2]",
		nil,
		NewMapping{
			4:   1682,
			393: 4262,
			404: 4263,
			477: 4766,
			573: 4766,
			718: 4782,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=10,powered=true]",
		nil,
		NewMapping{
			573: 268,
			718: 269,
			393: 268,
			404: 268,
			477: 268,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 5215,
			477: 5718,
			573: 5718,
			393: 5214,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=none,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 10999,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=none,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13058,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=low,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5815,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=23,powered=true]",
		nil,
		NewMapping{
			718: 545,
			393: 544,
			404: 544,
			477: 544,
			573: 544,
		},
	},
	{
		"minecraft:light_blue_bed[facing=south,occupied=true,part=foot]",
		nil,
		NewMapping{
			404: 801,
			477: 1101,
			573: 1101,
			718: 1102,
			393: 801,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 7694,
			573: 7694,
			718: 8230,
			393: 7187,
			404: 7188,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13322,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=none,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11850,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=none,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12730,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   1079,
			393: 3190,
			404: 3191,
			477: 3654,
			573: 3654,
			718: 3656,
		},
	},
	{
		"minecraft:spruce_fence_gate[facing=west,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7379,
			404: 7380,
			477: 7904,
			573: 7904,
			718: 8440,
			4:   2937,
		},
	},
	{
		"minecraft:dead_bubble_coral_fan[waterlogged=true]",
		nil,
		NewMapping{
			404: 8564,
			477: 9008,
			573: 9008,
			718: 9544,
			393: 8548,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=false,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10987,
			573: 10987,
		},
	},
	{
		"minecraft:dead_horn_coral_wall_fan[facing=east,waterlogged=false]",
		nil,
		NewMapping{
			718: 9599,
			393: 8503,
			404: 8519,
			477: 9063,
			573: 9063,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 6214,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=tall,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14105,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=14,powered=false]",
		nil,
		NewMapping{
			718: 478,
			393: 477,
			404: 477,
			477: 477,
			573: 477,
		},
	},
	{
		"minecraft:dark_oak_fence[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 7655,
			477: 8179,
			573: 8179,
			718: 8715,
			393: 7654,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=true,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5144,
			404: 5145,
			477: 5648,
			573: 5648,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 5112,
			393: 4592,
			404: 4593,
			477: 5096,
			573: 5096,
		},
	},
	{
		"minecraft:blue_banner[rotation=1]",
		nil,
		NewMapping{
			393: 7031,
			404: 7032,
			477: 7538,
			573: 7538,
			718: 8074,
		},
	},
	{
		"minecraft:brain_coral[waterlogged=false]",
		nil,
		NewMapping{
			404: 8473,
			477: 8997,
			573: 8997,
			718: 9533,
		},
	},
	{
		"minecraft:sandstone_wall[east=false,north=true,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10950,
			573: 10950,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 12330,
		},
	},
	{
		"minecraft:warped_fence_gate[facing=west,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15309,
		},
	},
	{
		"minecraft:purple_wall_banner[facing=west]",
		nil,
		NewMapping{
			477: 7659,
			573: 7659,
			718: 8195,
			393: 7152,
			404: 7153,
		},
	},
	{
		"minecraft:sign[rotation=4,waterlogged=true]",
		nil,
		NewMapping{
			393: 3083,
		},
	},
	{
		"minecraft:jungle_door[facing=north,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			477: 8341,
			573: 8341,
			718: 8877,
			393: 7816,
			404: 7817,
		},
	},
	{
		"minecraft:jungle_sign[rotation=15,waterlogged=false]",
		nil,
		NewMapping{
			477: 3538,
			573: 3538,
			718: 3540,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=0,south=none,west=up]",
		nil,
		NewMapping{
			718: 2640,
			393: 2334,
			404: 2335,
			477: 2638,
			573: 2638,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16419,
		},
	},
	{
		"minecraft:weeping_vines[age=21]",
		nil,
		NewMapping{
			718: 15011,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=tall,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11506,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=none,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14236,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=tall,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13449,
		},
	},
	{
		"minecraft:dark_oak_stairs[facing=north,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 6926,
			718: 7462,
			393: 6419,
			404: 6420,
			477: 6926,
		},
	},
	{
		"minecraft:blue_bed[facing=west,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 934,
			404: 934,
			477: 1234,
			573: 1234,
			718: 1235,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=8,powered=true]",
		nil,
		NewMapping{
			393: 614,
			404: 614,
			477: 614,
			573: 614,
			718: 615,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=none,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14325,
		},
	},
	{
		"minecraft:andesite_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9986,
			573: 9986,
			718: 10522,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=none,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14027,
		},
	},
	{
		"minecraft:fire[age=1,east=true,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			573: 1480,
			718: 1481,
			393: 1176,
			404: 1177,
			477: 1480,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=6,powered=false]",
		nil,
		NewMapping{
			393: 711,
			404: 711,
			477: 711,
			573: 711,
			718: 712,
		},
	},
	{
		"minecraft:lapis_block",
		nil,
		NewMapping{
			393: 232,
			404: 232,
			477: 232,
			573: 232,
			718: 233,
			4:   352,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=2,powered=true]",
		nil,
		NewMapping{
			718: 653,
			393: 652,
			404: 652,
			477: 652,
			573: 652,
		},
	},
	{
		"minecraft:crimson_stem[axis=x]",
		nil,
		NewMapping{
			718: 14975,
		},
	},
	{
		"minecraft:carrots[age=6]",
		nil,
		NewMapping{
			393: 5293,
			404: 5294,
			477: 5800,
			573: 5800,
			718: 6336,
			4:   2262,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=north,half=top,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3661,
			404: 3662,
			477: 4165,
			573: 4165,
			718: 4179,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=false,south=true,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 10350,
			573: 10350,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=low,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5723,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14180,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=20,powered=true]",
		nil,
		NewMapping{
			393: 538,
			404: 538,
			477: 538,
			573: 538,
			718: 539,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=12,south=none,west=side]",
		nil,
		NewMapping{
			393: 2155,
			404: 2156,
			477: 2459,
			573: 2459,
			718: 2461,
		},
	},
	{
		"minecraft:dark_oak_door[facing=east,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			393: 7986,
			404: 7987,
			477: 8511,
			573: 8511,
			718: 9047,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=tall,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11827,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=tall,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12483,
		},
	},
	{
		"minecraft:warped_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 15412,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14599,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14233,
		},
	},
	{
		"minecraft:diorite_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 10184,
			573: 10184,
			718: 10720,
		},
	},
	{
		"minecraft:diorite_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 10188,
			718: 10724,
			477: 10188,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=tall,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11936,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=low,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13499,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=6,south=up,west=up]",
		nil,
		NewMapping{
			718: 2112,
			393: 1806,
			404: 1807,
			477: 2110,
			573: 2110,
		},
	},
	{
		"minecraft:brain_coral_wall_fan[facing=east,waterlogged=false]",
		nil,
		NewMapping{
			393: 8519,
			404: 8535,
			477: 9079,
			573: 9079,
			718: 9615,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=4,south=none,west=side]",
		nil,
		NewMapping{
			477: 2531,
			573: 2531,
			718: 2533,
			393: 2227,
			404: 2228,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=true,north=true,south=false,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10980,
			573: 10980,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14741,
		},
	},
	{
		"minecraft:jungle_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5057,
			404: 5058,
			477: 5561,
			573: 5561,
			718: 5577,
		},
	},
	{
		"minecraft:fire[age=5,east=true,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1303,
			404: 1304,
			477: 1607,
			573: 1607,
			718: 1608,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16827,
		},
	},
	{
		"minecraft:crimson_fence_gate[facing=north,in_wall=true,open=false,powered=false]",
		nil,
		NewMapping{
			718: 15258,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10116,
			573: 10116,
			718: 10652,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11063,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 13660,
		},
	},
	{
		"minecraft:dead_horn_coral_wall_fan[facing=north,waterlogged=true]",
		nil,
		NewMapping{
			393: 8496,
			404: 8512,
			477: 9056,
			573: 9056,
			718: 9592,
		},
	},
	{
		"minecraft:acacia_fence[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 7638,
			404: 7639,
			477: 8163,
			573: 8163,
			718: 8699,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=8,south=side,west=none]",
		nil,
		NewMapping{
			393: 1829,
			404: 1830,
			477: 2133,
			573: 2133,
			718: 2135,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10933,
			573: 10933,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10487,
			573: 10487,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=tall,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14620,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=none,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13596,
		},
	},
	{
		"minecraft:spruce_door[facing=west,half=lower,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7719,
			404: 7720,
			477: 8244,
			573: 8244,
			718: 8780,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=none,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13380,
		},
	},
	{
		"minecraft:crimson_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 15395,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=north,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9149,
			573: 9149,
			718: 9685,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=false,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10381,
			573: 10381,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11559,
		},
	},
	{
		"minecraft:red_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 981,
			404: 981,
			477: 1281,
			573: 1281,
			718: 1282,
		},
	},
	{
		"minecraft:quartz_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 5712,
			404: 5713,
			477: 6219,
			573: 6219,
			718: 6755,
		},
	},
	{
		"minecraft:acacia_fence[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7635,
			404: 7636,
			477: 8160,
			573: 8160,
			718: 8696,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=12,south=side,west=side]",
		nil,
		NewMapping{
			404: 2585,
			477: 2888,
			573: 2888,
			718: 2890,
			393: 2584,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=none,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16341,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=tall,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13981,
		},
	},
	{
		"minecraft:mushroom_stem[down=false,east=false,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 4166,
			404: 4167,
			477: 4670,
			573: 4670,
			718: 4684,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=0,south=side,west=side]",
		nil,
		NewMapping{
			477: 2780,
			573: 2780,
			718: 2782,
			393: 2476,
			404: 2477,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 6605,
			573: 6605,
			718: 7141,
			393: 6098,
			404: 6099,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=low,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16939,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 16310,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=7,south=up,west=up]",
		nil,
		NewMapping{
			393: 2247,
			404: 2248,
			477: 2551,
			573: 2551,
			718: 2553,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 7729,
			718: 8265,
			393: 7222,
			404: 7223,
			477: 7729,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=12,powered=false]",
		nil,
		NewMapping{
			393: 623,
			404: 623,
			477: 623,
			573: 623,
			718: 624,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 10240,
			573: 10240,
			718: 10776,
		},
	},
	{
		"minecraft:tripwire[attached=true,disarmed=false,east=true,north=false,powered=false,south=false,west=true]",
		nil,
		NewMapping{
			393: 4801,
			404: 4802,
			477: 5305,
			573: 5305,
			718: 5321,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=west,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			573: 7268,
			718: 7804,
			393: 6761,
			404: 6762,
			477: 7268,
		},
	},
	{
		"minecraft:jungle_sign[rotation=10,waterlogged=true]",
		nil,
		NewMapping{
			573: 3527,
			718: 3529,
			477: 3527,
		},
	},
	{
		"minecraft:jigsaw[facing=north]",
		nil,
		NewMapping{
			477: 11256,
			573: 11272,
		},
	},
	{
		"minecraft:jungle_door[facing=north,half=upper,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			573: 8335,
			718: 8871,
			393: 7810,
			404: 7811,
			477: 8335,
		},
	},
	{
		"minecraft:birch_button[face=ceiling,facing=west,powered=false]",
		nil,
		NewMapping{
			393: 5372,
			404: 5373,
			477: 5879,
			573: 5879,
			718: 6415,
		},
	},
	{
		"minecraft:dark_oak_fence[east=false,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 8198,
			573: 8198,
			718: 8734,
			393: 7673,
			404: 7674,
		},
	},
	{
		"minecraft:acacia_sign[rotation=14,waterlogged=true]",
		nil,
		NewMapping{
			477: 3503,
			573: 3503,
			718: 3505,
		},
	},
	{
		"minecraft:acacia_button[face=wall,facing=south,powered=true]",
		nil,
		NewMapping{
			573: 5916,
			718: 6452,
			393: 5409,
			404: 5410,
			477: 5916,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 6088,
		},
	},
	{
		"minecraft:netherite_block",
		nil,
		NewMapping{
			718: 15826,
		},
	},
	{
		"minecraft:warped_trapdoor[facing=south,half=bottom,open=true,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			718: 15215,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=tall,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5739,
		},
	},
	{
		"minecraft:acacia_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6397,
			404: 6398,
			477: 6904,
			573: 6904,
			718: 7440,
		},
	},
	{
		"minecraft:red_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 6286,
			477: 6792,
			573: 6792,
			718: 7328,
			393: 6285,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 15896,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=low,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 15972,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9370,
			718: 9906,
			477: 9370,
		},
	},
	{
		"minecraft:spruce_stairs[facing=north,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 4891,
			477: 5394,
			573: 5394,
			718: 5410,
			393: 4890,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=10,powered=false]",
		nil,
		NewMapping{
			573: 869,
			718: 870,
			477: 869,
		},
	},
	{
		"minecraft:brick_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 10910,
		},
	},
	{
		"minecraft:acacia_door[facing=north,half=lower,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			393: 7884,
			404: 7885,
			477: 8409,
			573: 8409,
			718: 8945,
			4:   3139,
		},
	},
	{
		"minecraft:piston[extended=true,facing=east]",
		nil,
		NewMapping{
			718: 1349,
			4:   541,
			393: 1048,
			404: 1048,
			477: 1348,
			573: 1348,
		},
	},
	{
		"minecraft:warped_hyphae[axis=z]",
		nil,
		NewMapping{
			718: 14966,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=6,south=up,west=none]",
		nil,
		NewMapping{
			718: 2690,
			393: 2384,
			404: 2385,
			477: 2688,
			573: 2688,
		},
	},
	{
		"minecraft:jungle_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			404: 5071,
			477: 5574,
			573: 5574,
			718: 5590,
			393: 5070,
		},
	},
	{
		"minecraft:tnt[unstable=true]",
		nil,
		NewMapping{
			404: 1126,
			477: 1429,
			573: 1429,
			718: 1430,
		},
	},
	{
		"minecraft:cyan_shulker_box[facing=south]",
		nil,
		NewMapping{
			718: 9334,
			4:   3651,
			393: 8273,
			404: 8274,
			477: 8798,
			573: 8798,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=8,south=side,west=none]",
		nil,
		NewMapping{
			393: 2981,
			404: 2982,
			477: 3285,
			573: 3285,
			718: 3287,
		},
	},
	{
		"minecraft:acacia_button[face=ceiling,facing=north,powered=false]",
		nil,
		NewMapping{
			573: 5923,
			718: 6459,
			393: 5416,
			404: 5417,
			477: 5923,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=north,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			393: 3200,
			404: 3201,
			477: 3664,
			573: 3664,
			718: 3666,
			4:   1075,
		},
	},
	{
		"minecraft:oak_fence_gate[facing=north,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			404: 4306,
			477: 4809,
			573: 4809,
			718: 4825,
			4:   1718,
			393: 4305,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11371,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=east,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 16740,
		},
	},
	{
		"minecraft:brown_mushroom_block[down=true,east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 4011,
			404: 4012,
			477: 4515,
			573: 4515,
			718: 4529,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 16298,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=tall,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11911,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=none,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12055,
		},
	},
	{
		"minecraft:birch_leaves[distance=1,persistent=false]",
		nil,
		NewMapping{
			477: 173,
			573: 173,
			718: 174,
			4:   290,
			393: 173,
			404: 173,
		},
	},
	{
		"minecraft:jungle_sign[rotation=3,waterlogged=true]",
		nil,
		NewMapping{
			477: 3513,
			573: 3513,
			718: 3515,
		},
	},
	{
		"minecraft:scaffolding[bottom=false,distance=6,waterlogged=true]",
		nil,
		NewMapping{
			477: 11127,
			573: 11127,
			718: 14783,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5973,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=tall,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11930,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6215,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=8,powered=true]",
		nil,
		NewMapping{
			393: 464,
			404: 464,
			477: 464,
			573: 464,
			718: 465,
		},
	},
	{
		"minecraft:gray_banner[rotation=0]",
		nil,
		NewMapping{
			393: 6966,
			404: 6967,
			477: 7473,
			573: 7473,
			718: 8009,
		},
	},
	{
		"minecraft:oak_door[facing=south,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			393: 3131,
			404: 3132,
			477: 3595,
			573: 3595,
			718: 3597,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9419,
			573: 9419,
			718: 9955,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13439,
		},
	},
	{
		"minecraft:jungle_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 5579,
			573: 5579,
			718: 5595,
			4:   2178,
			393: 5075,
			404: 5076,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=1,south=none,west=side]",
		nil,
		NewMapping{
			404: 2345,
			477: 2648,
			573: 2648,
			718: 2650,
			393: 2344,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 6094,
			477: 6600,
			573: 6600,
			718: 7136,
			393: 6093,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=north,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 10097,
			718: 10633,
			477: 10097,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 9146,
			718: 9682,
			477: 9146,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=tall,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13023,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=none,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 6091,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14625,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=10,south=none,west=side]",
		nil,
		NewMapping{
			718: 3019,
			393: 2713,
			404: 2714,
			477: 3017,
			573: 3017,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 6586,
			573: 6586,
			718: 7122,
			393: 6079,
			404: 6080,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10568,
			573: 10568,
		},
	},
	{
		"minecraft:blue_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			718: 7233,
			393: 6190,
			404: 6191,
			477: 6697,
			573: 6697,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=0,south=none,west=up]",
		nil,
		NewMapping{
			477: 2206,
			573: 2206,
			718: 2208,
			393: 1902,
			404: 1903,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=low,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11906,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13502,
		},
	},
	{
		"minecraft:fire[age=15,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 1617,
			404: 1618,
			477: 1921,
			573: 1921,
			718: 1922,
		},
	},
	{
		"minecraft:white_bed[facing=west,occupied=true,part=foot]",
		nil,
		NewMapping{
			393: 757,
			404: 757,
			477: 1057,
			573: 1057,
			718: 1058,
		},
	},
	{
		"minecraft:yellow_bed[facing=west,occupied=true,part=head]",
		nil,
		NewMapping{
			573: 1120,
			718: 1121,
			393: 820,
			404: 820,
			477: 1120,
		},
	},
	{
		"minecraft:conduit",
		nil,
		NewMapping{
			393: 8573,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=none,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 14248,
		},
	},
	{
		"minecraft:vine[east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 4270,
			404: 4271,
			477: 4774,
			573: 4774,
			718: 4790,
			4:   1711,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5830,
			404: 5831,
			477: 6337,
			573: 6337,
			718: 6873,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10359,
			573: 10359,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=tall,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6185,
		},
	},
	{
		"minecraft:fire[age=9,east=true,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			477: 1736,
			573: 1736,
			718: 1737,
			393: 1432,
			404: 1433,
		},
	},
	{
		"minecraft:oak_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 2005,
			573: 2005,
			718: 2007,
			393: 1701,
			404: 1702,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=none,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 10995,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=3,south=side,west=up]",
		nil,
		NewMapping{
			393: 1782,
			404: 1783,
			477: 2086,
			573: 2086,
			718: 2088,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=13,south=side,west=side]",
		nil,
		NewMapping{
			404: 2162,
			477: 2465,
			573: 2465,
			718: 2467,
			393: 2161,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=none,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16919,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 5797,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=none,south=low,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5670,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=13,south=none,west=up]",
		nil,
		NewMapping{
			393: 2019,
			404: 2020,
			477: 2323,
			573: 2323,
			718: 2325,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=false,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 6453,
			573: 6453,
			718: 6989,
			393: 5946,
			404: 5947,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=none,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13033,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=6,powered=false]",
		nil,
		NewMapping{
			393: 661,
			404: 661,
			477: 661,
			573: 661,
			718: 662,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=11,south=side,west=none]",
		nil,
		NewMapping{
			404: 2145,
			477: 2448,
			573: 2448,
			718: 2450,
			393: 2144,
		},
	},
	{
		"minecraft:trapped_chest[facing=south,type=right,waterlogged=false]",
		nil,
		NewMapping{
			393: 5590,
			404: 5591,
			477: 6097,
			573: 6097,
			718: 6633,
		},
	},
	{
		"minecraft:lime_bed[facing=south,occupied=false,part=head]",
		nil,
		NewMapping{
			404: 834,
			477: 1134,
			573: 1134,
			718: 1135,
			393: 834,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=south,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 3227,
			404: 3228,
			477: 3691,
			573: 3691,
			718: 3693,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=low,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5931,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16517,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=none,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13999,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5166,
			404: 5167,
			477: 5670,
			573: 5670,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=none,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 11197,
		},
	},
	{
		"minecraft:warped_door[facing=east,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			718: 15652,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13445,
		},
	},
	{
		"minecraft:vine[east=true,north=false,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 4277,
			404: 4278,
			477: 4781,
			573: 4781,
			718: 4797,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 10879,
			477: 10879,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13195,
		},
	},
	{
		"minecraft:warped_door[facing=east,half=lower,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			718: 15647,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=none,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5658,
		},
	},
	{
		"minecraft:fire[age=8,east=false,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1421,
			404: 1422,
			477: 1725,
			573: 1725,
			718: 1726,
		},
	},
	{
		"minecraft:pink_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 6038,
			477: 6544,
			573: 6544,
			718: 7080,
			393: 6037,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=3,south=up,west=none]",
		nil,
		NewMapping{
			573: 2949,
			718: 2951,
			393: 2645,
			404: 2646,
			477: 2949,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9302,
			573: 9302,
			718: 9838,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=low,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11679,
		},
	},
	{
		"minecraft:dark_oak_door[facing=west,half=lower,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			573: 8501,
			718: 9037,
			393: 7976,
			404: 7977,
			477: 8501,
		},
	},
	{
		"minecraft:iron_door[facing=north,half=upper,hinge=right,open=false,powered=false]",
		nil,
		NewMapping{
			393: 3310,
			404: 3311,
			477: 3814,
			573: 3814,
			718: 3816,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=tall,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12909,
		},
	},
	{
		"minecraft:warped_trapdoor[facing=west,half=top,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			718: 15225,
		},
	},
	{
		"minecraft:dark_oak_button[face=wall,facing=west,powered=true]",
		nil,
		NewMapping{
			393: 5435,
			404: 5436,
			477: 5942,
			573: 5942,
			718: 6478,
		},
	},
	{
		"minecraft:note_block[instrument=cow_bell,note=21,powered=false]",
		nil,
		NewMapping{
			477: 841,
			573: 841,
			718: 842,
		},
	},
	{
		"minecraft:fire[age=4,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			573: 1578,
			718: 1579,
			393: 1274,
			404: 1275,
			477: 1578,
		},
	},
	{
		"minecraft:birch_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 5549,
			393: 5029,
			404: 5030,
			477: 5533,
			573: 5533,
		},
	},
	{
		"minecraft:stone_brick_wall[east=false,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			573: 10705,
			477: 10705,
		},
	},
	{
		"minecraft:campfire[facing=south,lit=false,signal_fire=true,waterlogged=true]",
		nil,
		NewMapping{
			477: 11228,
			573: 11244,
			718: 14902,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13964,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=4,south=up,west=side]",
		nil,
		NewMapping{
			393: 2221,
			404: 2222,
			477: 2525,
			573: 2525,
			718: 2527,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11051,
		},
	},
	{
		"minecraft:crimson_fence[east=true,north=false,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			718: 15077,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12141,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=east,half=top,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 6543,
			404: 6544,
			477: 7050,
			573: 7050,
			718: 7586,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=east,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			404: 6710,
			477: 7216,
			573: 7216,
			718: 7752,
			393: 6709,
		},
	},
	{
		"minecraft:fire[age=4,east=false,north=false,south=false,up=false,west=false]",
		nil,
		NewMapping{
			477: 1598,
			573: 1598,
			718: 1599,
			4:   820,
			393: 1294,
			404: 1295,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=none,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11742,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9293,
			573: 9293,
			718: 9829,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 10238,
			573: 10238,
			718: 10774,
		},
	},
	{
		"minecraft:crimson_fence_gate[facing=south,in_wall=false,open=false,powered=true]",
		nil,
		NewMapping{
			718: 15269,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=low,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13393,
		},
	},
	{
		"minecraft:acacia_sapling[stage=0]",
		nil,
		NewMapping{
			4:   100,
			393: 29,
			404: 29,
			477: 29,
			573: 29,
			718: 29,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=true,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5215,
			404: 5216,
			477: 5719,
			573: 5719,
		},
	},
	{
		"minecraft:acacia_stairs[facing=south,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 6860,
			718: 7396,
			4:   2614,
			393: 6353,
			404: 6354,
			477: 6860,
		},
	},
	{
		"minecraft:birch_stairs[facing=east,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2164,
			393: 5025,
			404: 5026,
			477: 5529,
			573: 5529,
			718: 5545,
		},
	},
	{
		"minecraft:spruce_door[facing=west,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			718: 8773,
			393: 7712,
			404: 7713,
			477: 8237,
			573: 8237,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=tall,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12801,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12562,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			718: 4662,
			393: 4144,
			404: 4145,
			477: 4648,
			573: 4648,
		},
	},
	{
		"minecraft:fire[age=15,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			573: 1935,
			718: 1936,
			393: 1631,
			404: 1632,
			477: 1935,
		},
	},
	{
		"minecraft:diorite_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14448,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=4,south=none,west=side]",
		nil,
		NewMapping{
			477: 3107,
			573: 3107,
			718: 3109,
			393: 2803,
			404: 2804,
		},
	},
	{
		"minecraft:chipped_anvil[facing=west]",
		nil,
		NewMapping{
			4:   2325,
			393: 5573,
			404: 5574,
			477: 6080,
			573: 6080,
			718: 6616,
		},
	},
	{
		"minecraft:command_block[conditional=false,facing=south]",
		nil,
		NewMapping{
			4:   2195,
			393: 5132,
			404: 5133,
			477: 5636,
			573: 5636,
			718: 5652,
		},
	},
	{
		"minecraft:iron_trapdoor[facing=north,half=bottom,open=true,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			393: 6503,
			404: 6504,
			477: 7010,
			573: 7010,
			718: 7546,
		},
	},
	{
		"minecraft:note_block[instrument=chime,note=9,powered=false]",
		nil,
		NewMapping{
			393: 667,
			404: 667,
			477: 667,
			573: 667,
			718: 668,
		},
	},
	{
		"minecraft:fire[age=11,east=false,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			573: 1812,
			718: 1813,
			393: 1508,
			404: 1509,
			477: 1812,
		},
	},
	{
		"minecraft:white_bed[facing=east,occupied=true,part=head]",
		nil,
		NewMapping{
			393: 760,
			404: 760,
			477: 1060,
			573: 1060,
			718: 1061,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=south,half=bottom,open=true,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			393: 3747,
			404: 3748,
			477: 4251,
			573: 4251,
			718: 4265,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16205,
		},
	},
	{
		"minecraft:dark_oak_fence[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 7651,
			404: 7652,
			477: 8176,
			573: 8176,
			718: 8712,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=none,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14013,
		},
	},
	{
		"minecraft:crimson_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 15385,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14718,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16960,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=west,half=top,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			393: 3756,
			404: 3757,
			477: 4260,
			573: 4260,
			718: 4274,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=tall,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13872,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13105,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=none,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16896,
		},
	},
	{
		"minecraft:anvil[facing=south]",
		nil,
		NewMapping{
			393: 5568,
			404: 5569,
			477: 6075,
			573: 6075,
			718: 6611,
			4:   2320,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=8,south=side,west=side]",
		nil,
		NewMapping{
			718: 2998,
			393: 2692,
			404: 2693,
			477: 2996,
			573: 2996,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=west,half=top,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9420,
			573: 9420,
			718: 9956,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=none,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11642,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=4,south=side,west=up]",
		nil,
		NewMapping{
			393: 2367,
			404: 2368,
			477: 2671,
			573: 2671,
			718: 2673,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=15,south=none,west=none]",
		nil,
		NewMapping{
			393: 2327,
			404: 2328,
			477: 2631,
			573: 2631,
			718: 2633,
		},
	},
	{
		"minecraft:jungle_stairs[facing=east,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 5108,
			404: 5109,
			477: 5612,
			573: 5612,
			718: 5628,
		},
	},
	{
		"minecraft:sandstone_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10920,
			573: 10920,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=none,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 6005,
		},
	},
	{
		"minecraft:gray_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 6555,
			573: 6555,
			718: 7091,
			393: 6048,
			404: 6049,
		},
	},
	{
		"minecraft:fire[age=7,east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1383,
			404: 1384,
			477: 1687,
			573: 1687,
			718: 1688,
		},
	},
	{
		"minecraft:cave_air",
		nil,
		NewMapping{
			393: 8575,
			404: 8592,
			477: 9130,
			573: 9130,
			718: 9666,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=10,south=none,west=none]",
		nil,
		NewMapping{
			477: 3306,
			573: 3306,
			718: 3308,
			4:   890,
			393: 3002,
			404: 3003,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 6759,
			718: 7295,
			393: 6252,
			404: 6253,
			477: 6759,
		},
	},
	{
		"minecraft:fire[age=2,east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1216,
			404: 1217,
			477: 1520,
			573: 1520,
			718: 1521,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=low,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14163,
		},
	},
	{
		"minecraft:stone_brick_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12721,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			718: 16280,
		},
	},
	{
		"minecraft:blackstone_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 15964,
		},
	},
	{
		"minecraft:nether_brick_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			404: 7330,
			477: 7848,
			573: 7848,
			718: 8384,
			393: 7329,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=14,south=side,west=up]",
		nil,
		NewMapping{
			477: 2329,
			573: 2329,
			718: 2331,
			393: 2025,
			404: 2026,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=low,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 6146,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=tall,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14639,
		},
	},
	{
		"minecraft:birch_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			404: 5006,
			477: 5509,
			573: 5509,
			718: 5525,
			4:   2165,
			393: 5005,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=true,north=true,south=false,up=false,west=false]",
		nil,
		NewMapping{
			404: 8043,
			477: 8567,
			573: 8567,
			718: 9103,
			393: 8042,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16522,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 10418,
			477: 10418,
		},
	},
	{
		"minecraft:red_nether_brick_stairs[facing=west,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 10598,
			477: 10062,
			573: 10062,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=tall,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16882,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 7255,
			404: 7256,
			477: 7762,
			573: 7762,
			718: 8298,
		},
	},
	{
		"minecraft:potatoes[age=0]",
		nil,
		NewMapping{
			718: 6338,
			4:   2272,
			393: 5295,
			404: 5296,
			477: 5802,
			573: 5802,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=west,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4462,
			404: 4463,
			477: 4966,
			573: 4966,
			718: 4982,
		},
	},
	{
		"minecraft:green_banner[rotation=15]",
		nil,
		NewMapping{
			393: 7077,
			404: 7078,
			477: 7584,
			573: 7584,
			718: 8120,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12891,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=true,north=false,south=false,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 6502,
			718: 7038,
			393: 5995,
			404: 5996,
			477: 6502,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=side,power=7,south=none,west=up]",
		nil,
		NewMapping{
			718: 3135,
			393: 2829,
			404: 2830,
			477: 3133,
			573: 3133,
		},
	},
	{
		"minecraft:cobblestone_wall[east=true,north=false,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 5165,
			404: 5166,
			477: 5669,
			573: 5669,
		},
	},
	{
		"minecraft:prismarine_wall[east=true,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10408,
			573: 10408,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=none,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16561,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=false,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 6381,
			573: 6381,
			718: 6917,
			393: 5874,
			404: 5875,
		},
	},
	{
		"minecraft:light_gray_stained_glass_pane[east=false,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 6095,
			404: 6096,
			477: 6602,
			573: 6602,
			718: 7138,
		},
	},
	{
		"minecraft:light_weighted_pressure_plate[power=0]",
		nil,
		NewMapping{
			393: 5603,
			404: 5604,
			477: 6110,
			573: 6110,
			718: 6646,
			4:   2352,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9821,
			573: 9821,
			718: 10357,
		},
	},
	{
		"minecraft:quartz_stairs[facing=east,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 6286,
			718: 6822,
			393: 5779,
			404: 5780,
			477: 6286,
		},
	},
	{
		"minecraft:dragon_wall_head[facing=east]",
		nil,
		NewMapping{
			393: 5550,
			404: 5551,
			477: 6073,
			573: 6073,
			718: 6609,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=false,north=false,south=false,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 11027,
			573: 11027,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=low,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 5920,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16485,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=low,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 6248,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			477: 6498,
			573: 6498,
			718: 7034,
			393: 5991,
			404: 5992,
		},
	},
	{
		"minecraft:mycelium[snowy=true]",
		nil,
		NewMapping{
			393: 4492,
			404: 4493,
			477: 4996,
			573: 4996,
			718: 5012,
		},
	},
	{
		"minecraft:glass_pane[east=true,north=false,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 4223,
			404: 4224,
			477: 4727,
			573: 4727,
			718: 4743,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9296,
			573: 9296,
			718: 9832,
		},
	},
	{
		"minecraft:dark_oak_door[facing=north,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7939,
			404: 7940,
			477: 8464,
			573: 8464,
			718: 9000,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			573: 7242,
			718: 7778,
			393: 6735,
			404: 6736,
			477: 7242,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=tall,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11937,
		},
	},
	{
		"minecraft:cobblestone_wall[east=false,north=true,south=true,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			393: 5174,
			404: 5175,
			477: 5678,
			573: 5678,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=14,south=side,west=up]",
		nil,
		NewMapping{
			393: 2313,
			404: 2314,
			477: 2617,
			573: 2617,
			718: 2619,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9693,
			573: 9693,
			718: 10229,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=low,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14051,
		},
	},
	{
		"minecraft:brown_stained_glass_pane[east=true,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 6214,
			477: 6720,
			573: 6720,
			718: 7256,
			393: 6213,
		},
	},
	{
		"minecraft:stone_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9665,
			573: 9665,
			718: 10201,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14294,
		},
	},
	{
		"minecraft:birch_stairs[facing=east,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 5556,
			393: 5036,
			404: 5037,
			477: 5540,
			573: 5540,
		},
	},
	{
		"minecraft:andesite_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 10482,
			477: 9946,
			573: 9946,
		},
	},
	{
		"minecraft:granite_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12178,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=none,south=tall,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13274,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=none,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12629,
		},
	},
	{
		"minecraft:yellow_bed[facing=east,occupied=true,part=head]",
		nil,
		NewMapping{
			404: 824,
			477: 1124,
			573: 1124,
			718: 1125,
			393: 824,
		},
	},
	{
		"minecraft:green_shulker_box[facing=west]",
		nil,
		NewMapping{
			393: 8298,
			404: 8299,
			477: 8823,
			573: 8823,
			718: 9359,
			4:   3716,
		},
	},
	{
		"minecraft:green_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 6745,
			573: 6745,
			718: 7281,
			393: 6238,
			404: 6239,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=west,half=bottom,shape=outer_left,waterlogged=false]",
		nil,
		NewMapping{
			718: 16317,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=none,north=none,south=tall,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 6014,
		},
	},
	{
		"minecraft:brewing_stand[has_bottle_0=true,has_bottle_1=false,has_bottle_2=true]",
		nil,
		NewMapping{
			4:   1877,
			393: 4615,
			404: 4616,
			477: 5119,
			573: 5119,
			718: 5135,
		},
	},
	{
		"minecraft:chorus_plant[down=false,east=true,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			477: 8560,
			573: 8560,
			718: 9096,
			393: 8035,
			404: 8036,
		},
	},
	{
		"minecraft:weeping_vines[age=9]",
		nil,
		NewMapping{
			718: 14999,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=low,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13613,
		},
	},
	{
		"minecraft:orange_stained_glass_pane[east=false,north=false,south=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			718: 6920,
			393: 5877,
			404: 5878,
			477: 6384,
			573: 6384,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11153,
		},
	},
	{
		"minecraft:soul_wall_torch[facing=east]",
		nil,
		NewMapping{
			718: 4012,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 4425,
			404: 4426,
			477: 4929,
			573: 4929,
			718: 4945,
		},
	},
	{
		"minecraft:prismarine_brick_stairs[facing=north,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 6646,
			404: 6647,
			477: 7153,
			573: 7153,
			718: 7689,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=true,north=true,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 6330,
			718: 6866,
			393: 5823,
			404: 5824,
			477: 6330,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=1,south=none,west=up]",
		nil,
		NewMapping{
			404: 2056,
			477: 2359,
			573: 2359,
			718: 2361,
			393: 2055,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=none,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16044,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=side,power=3,south=none,west=none]",
		nil,
		NewMapping{
			393: 1931,
			404: 1932,
			477: 2235,
			573: 2235,
			718: 2237,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10493,
			573: 10493,
		},
	},
	{
		"minecraft:polished_granite_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 9697,
			477: 9161,
			573: 9161,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=tall,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11507,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16616,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 1663,
			404: 1664,
			477: 1967,
			573: 1967,
			718: 1969,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=7,south=up,west=up]",
		nil,
		NewMapping{
			477: 2983,
			573: 2983,
			718: 2985,
			393: 2679,
			404: 2680,
		},
	},
	{
		"minecraft:stripped_acacia_wood[axis=z]",
		nil,
		NewMapping{
			393: 140,
			404: 140,
			477: 140,
			573: 140,
			718: 141,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 16709,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=none,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11645,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=north,half=bottom,open=false,powered=false,waterlogged=true]",
		nil,
		NewMapping{
			573: 4303,
			718: 4317,
			393: 3799,
			404: 3800,
			477: 4303,
		},
	},
	{
		"minecraft:purple_stained_glass_pane[east=true,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 6145,
			477: 6651,
			573: 6651,
			718: 7187,
			393: 6144,
		},
	},
	{
		"minecraft:beehive[facing=south,honey_level=4]",
		nil,
		NewMapping{
			573: 11321,
			718: 15810,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=tall,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12907,
		},
	},
	{
		"minecraft:birch_stairs[facing=west,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 5528,
			393: 5008,
			404: 5009,
			477: 5512,
			573: 5512,
		},
	},
	{
		"minecraft:warped_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 15464,
		},
	},
	{
		"minecraft:granite_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 10407,
			477: 9871,
			573: 9871,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16588,
		},
	},
	{
		"minecraft:cobblestone_stairs[facing=north,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 3189,
			404: 3190,
			477: 3653,
			573: 3653,
			718: 3655,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=east,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			573: 4220,
			718: 4234,
			393: 3716,
			404: 3717,
			477: 4220,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			718: 4572,
			393: 4054,
			404: 4055,
			477: 4558,
			573: 4558,
		},
	},
	{
		"minecraft:cyan_bed[facing=east,occupied=false,part=head]",
		nil,
		NewMapping{
			393: 906,
			404: 906,
			477: 1206,
			573: 1206,
			718: 1207,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=low,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12449,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=west,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 5107,
			393: 4587,
			404: 4588,
			477: 5091,
			573: 5091,
		},
	},
	{
		"minecraft:birch_door[facing=west,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7775,
			404: 7776,
			477: 8300,
			573: 8300,
			718: 8836,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=tall,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13973,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=low,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13847,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 6747,
			404: 6748,
			477: 7254,
			573: 7254,
			718: 7790,
		},
	},
	{
		"minecraft:jungle_slab[type=double,waterlogged=false]",
		nil,
		NewMapping{
			393: 7280,
			404: 7281,
			477: 7787,
			573: 7787,
			718: 8323,
			4:   2003,
		},
	},
	{
		"minecraft:crimson_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 15390,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=none,south=low,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14669,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=low,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11799,
		},
	},
	{
		"minecraft:stone_button[face=ceiling,facing=north,powered=false]",
		nil,
		NewMapping{
			393: 3408,
			404: 3409,
			477: 3912,
			573: 3912,
			718: 3914,
			4:   1232,
		},
	},
	{
		"minecraft:yellow_bed[facing=south,occupied=true,part=head]",
		nil,
		NewMapping{
			573: 1116,
			718: 1117,
			393: 816,
			404: 816,
			477: 1116,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 7732,
			573: 7732,
			718: 8268,
			393: 7225,
			404: 7226,
		},
	},
	{
		"minecraft:acacia_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			4:   2615,
			393: 6333,
			404: 6334,
			477: 6840,
			573: 6840,
			718: 7376,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=tall,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13982,
		},
	},
	{
		"minecraft:crimson_sign[rotation=4,waterlogged=false]",
		nil,
		NewMapping{
			718: 15664,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=none,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14024,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=10,south=none,west=up]",
		nil,
		NewMapping{
			573: 2728,
			718: 2730,
			393: 2424,
			404: 2425,
			477: 2728,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=18,powered=true]",
		nil,
		NewMapping{
			393: 734,
			404: 734,
			477: 734,
			573: 734,
			718: 735,
		},
	},
	{
		"minecraft:rail[shape=north_south]",
		nil,
		NewMapping{
			393: 3179,
			404: 3180,
			477: 3643,
			573: 3643,
			718: 3645,
			4:   1056,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9317,
			573: 9317,
			718: 9853,
		},
	},
	{
		"minecraft:birch_leaves[distance=1,persistent=true]",
		nil,
		NewMapping{
			573: 172,
			718: 173,
			4:   294,
			393: 172,
			404: 172,
			477: 172,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10716,
			573: 10716,
		},
	},
	{
		"minecraft:birch_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 5482,
			573: 5482,
			718: 5498,
			393: 4978,
			404: 4979,
		},
	},
	{
		"minecraft:spruce_button[face=wall,facing=north,powered=false]",
		nil,
		NewMapping{
			573: 5843,
			718: 6379,
			393: 5336,
			404: 5337,
			477: 5843,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=none,south=none,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14117,
		},
	},
	{
		"minecraft:sandstone_wall[east=none,north=none,south=tall,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13815,
		},
	},
	{
		"minecraft:repeater[delay=3,facing=east,locked=false,powered=false]",
		nil,
		NewMapping{
			718: 4078,
			4:   1499,
			393: 3560,
			404: 3561,
			477: 4064,
			573: 4064,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			573: 7717,
			718: 8253,
			393: 7210,
			404: 7211,
			477: 7717,
		},
	},
	{
		"minecraft:dark_oak_fence_gate[facing=south,in_wall=true,open=false,powered=true]",
		nil,
		NewMapping{
			573: 8020,
			718: 8556,
			393: 7495,
			404: 7496,
			477: 8020,
		},
	},
	{
		"minecraft:pink_bed[facing=north,occupied=true,part=head]",
		nil,
		NewMapping{
			477: 1144,
			573: 1144,
			718: 1145,
			393: 844,
			404: 844,
		},
	},
	{
		"minecraft:sandstone_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			393: 4670,
			404: 4671,
			477: 5174,
			573: 5174,
			718: 5190,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=tall,south=none,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 5945,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13214,
		},
	},
	{
		"minecraft:blackstone_wall[east=tall,north=low,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16176,
		},
	},
	{
		"minecraft:purple_banner[rotation=13]",
		nil,
		NewMapping{
			393: 7027,
			404: 7028,
			477: 7534,
			573: 7534,
			718: 8070,
		},
	},
	{
		"minecraft:jungle_trapdoor[facing=north,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			573: 4294,
			718: 4308,
			393: 3790,
			404: 3791,
			477: 4294,
		},
	},
	{
		"minecraft:fire[age=13,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 1554,
			404: 1555,
			477: 1858,
			573: 1858,
			718: 1859,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=1,powered=true]",
		nil,
		NewMapping{
			477: 1000,
			573: 1000,
			718: 1001,
		},
	},
	{
		"minecraft:diorite_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 14707,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=none,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11214,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=up,power=1,south=side,west=up]",
		nil,
		NewMapping{
			573: 2932,
			718: 2934,
			393: 2628,
			404: 2629,
			477: 2932,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9308,
			573: 9308,
			718: 9844,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=tall,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13021,
		},
	},
	{
		"minecraft:nether_brick_wall[east=low,north=none,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12944,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=low,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13291,
		},
	},
	{
		"minecraft:fire[age=3,east=false,north=true,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1247,
			404: 1248,
			477: 1551,
			573: 1551,
			718: 1552,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			477: 9598,
			573: 9598,
			718: 10134,
		},
	},
	{
		"minecraft:bee_nest[facing=north,honey_level=3]",
		nil,
		NewMapping{
			573: 11290,
			718: 15779,
		},
	},
	{
		"minecraft:crimson_button[face=floor,facing=north,powered=false]",
		nil,
		NewMapping{
			718: 15480,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=20,powered=true]",
		nil,
		NewMapping{
			477: 888,
			573: 888,
			718: 889,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=tall,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 16854,
		},
	},
	{
		"minecraft:white_stained_glass_pane[east=false,north=false,south=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			393: 5847,
			404: 5848,
			477: 6354,
			573: 6354,
			718: 6890,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=none,power=11,south=up,west=up]",
		nil,
		NewMapping{
			393: 2571,
			404: 2572,
			477: 2875,
			573: 2875,
			718: 2877,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=south,half=top,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9797,
			573: 9797,
			718: 10333,
		},
	},
	{
		"minecraft:brick_wall[east=true,north=true,south=true,up=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10331,
			573: 10331,
		},
	},
	{
		"minecraft:brick_wall[east=low,north=none,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 10982,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=low,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 6238,
		},
	},
	{
		"minecraft:note_block[instrument=harp,note=24,powered=true]",
		nil,
		NewMapping{
			573: 296,
			718: 297,
			393: 296,
			404: 296,
			477: 296,
		},
	},
	{
		"minecraft:red_sandstone_stairs[facing=east,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			404: 7243,
			477: 7749,
			573: 7749,
			718: 8285,
			393: 7242,
		},
	},
	{
		"minecraft:prismarine_stairs[facing=west,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			404: 6603,
			477: 7109,
			573: 7109,
			718: 7645,
			393: 6602,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=false,south=false,up=true,waterlogged=false,west=false]",
		nil,
		NewMapping{
			573: 5732,
			393: 5228,
			404: 5229,
			477: 5732,
		},
	},
	{
		"minecraft:brown_concrete",
		nil,
		NewMapping{
			393: 8389,
			404: 8390,
			477: 8914,
			573: 8914,
			718: 9450,
			4:   4028,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13223,
		},
	},
	{
		"minecraft:bubble_coral[waterlogged=false]",
		nil,
		NewMapping{
			404: 8475,
			477: 8999,
			573: 8999,
			718: 9535,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			477: 10504,
			573: 10504,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16814,
		},
	},
	{
		"minecraft:repeater[delay=1,facing=north,locked=false,powered=false]",
		nil,
		NewMapping{
			404: 3517,
			477: 4020,
			573: 4020,
			718: 4034,
			4:   1490,
			393: 3516,
		},
	},
	{
		"minecraft:dark_oak_fence[east=false,north=true,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			393: 7661,
			404: 7662,
			477: 8186,
			573: 8186,
			718: 8722,
		},
	},
	{
		"minecraft:note_block[instrument=xylophone,note=10,powered=true]",
		nil,
		NewMapping{
			393: 718,
			404: 718,
			477: 718,
			573: 718,
			718: 719,
		},
	},
	{
		"minecraft:oak_door[facing=east,half=upper,hinge=left,open=false,powered=true]",
		nil,
		NewMapping{
			4:   1034,
			393: 3157,
			404: 3158,
			477: 3621,
			573: 3621,
			718: 3623,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=true,north=false,south=false,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5231,
			404: 5232,
			477: 5735,
			573: 5735,
		},
	},
	{
		"minecraft:lime_stained_glass_pane[east=true,north=false,south=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5990,
			404: 5991,
			477: 6497,
			573: 6497,
			718: 7033,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=tall,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16984,
		},
	},
	{
		"minecraft:fire[age=8,east=false,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			404: 1412,
			477: 1715,
			573: 1715,
			718: 1716,
			393: 1411,
		},
	},
	{
		"minecraft:spruce_door[facing=west,half=upper,hinge=right,open=false,powered=true]",
		nil,
		NewMapping{
			393: 7715,
			404: 7716,
			477: 8240,
			573: 8240,
			718: 8776,
		},
	},
	{
		"minecraft:dark_prismarine_stairs[facing=west,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 6769,
			404: 6770,
			477: 7276,
			573: 7276,
			718: 7812,
		},
	},
	{
		"minecraft:tripwire[attached=false,disarmed=false,east=false,north=true,powered=true,south=true,west=true]",
		nil,
		NewMapping{
			393: 4867,
			404: 4868,
			477: 5371,
			573: 5371,
			718: 5387,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=none,south=tall,up=false,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13385,
		},
	},
	{
		"minecraft:fire[age=7,east=false,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			393: 1376,
			404: 1377,
			477: 1680,
			573: 1680,
			718: 1681,
		},
	},
	{
		"minecraft:spruce_stairs[facing=west,half=top,shape=inner_right,waterlogged=false]",
		nil,
		NewMapping{
			573: 5433,
			718: 5449,
			393: 4929,
			404: 4930,
			477: 5433,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13413,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=none,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13370,
		},
	},
	{
		"minecraft:fire[age=2,east=false,north=false,south=true,up=true,west=true]",
		nil,
		NewMapping{
			393: 1223,
			404: 1224,
			477: 1527,
			573: 1527,
			718: 1528,
		},
	},
	{
		"minecraft:light_blue_stained_glass_pane[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			404: 5946,
			477: 6452,
			573: 6452,
			718: 6988,
			393: 5945,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=west,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			718: 15894,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16523,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=low,north=low,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 6142,
		},
	},
	{
		"minecraft:stripped_oak_wood[axis=x]",
		nil,
		NewMapping{
			477: 126,
			573: 126,
			718: 127,
			393: 126,
			404: 126,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=2,south=none,west=side]",
		nil,
		NewMapping{
			477: 2657,
			573: 2657,
			718: 2659,
			393: 2353,
			404: 2354,
		},
	},
	{
		"minecraft:spruce_fence[east=true,north=false,south=true,waterlogged=true,west=true]",
		nil,
		NewMapping{
			404: 7526,
			477: 8050,
			573: 8050,
			718: 8586,
			393: 7525,
		},
	},
	{
		"minecraft:prismarine_wall[east=low,north=tall,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11394,
		},
	},
	{
		"minecraft:light_blue_wall_banner[facing=south]",
		nil,
		NewMapping{
			393: 7123,
			404: 7124,
			477: 7630,
			573: 7630,
			718: 8166,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=east,half=top,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 9835,
			573: 9835,
			718: 10371,
		},
	},
	{
		"minecraft:brick_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4366,
			404: 4367,
			477: 4870,
			573: 4870,
			718: 4886,
		},
	},
	{
		"minecraft:blue_shulker_box[facing=up]",
		nil,
		NewMapping{
			4:   3681,
			393: 8287,
			404: 8288,
			477: 8812,
			573: 8812,
			718: 9348,
		},
	},
	{
		"minecraft:fire[age=14,east=true,north=true,south=true,up=true,west=false]",
		nil,
		NewMapping{
			404: 1585,
			477: 1888,
			573: 1888,
			718: 1889,
			393: 1584,
		},
	},
	{
		"minecraft:note_block[instrument=flute,note=5,powered=false]",
		nil,
		NewMapping{
			404: 509,
			477: 509,
			573: 509,
			718: 510,
			393: 509,
		},
	},
	{
		"minecraft:polished_blackstone_slab[type=bottom,waterlogged=false]",
		nil,
		NewMapping{
			718: 16748,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=none,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16562,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14612,
		},
	},
	{
		"minecraft:vine[east=false,north=true,south=false,up=true,west=false]",
		nil,
		NewMapping{
			718: 4809,
			393: 4289,
			404: 4290,
			477: 4793,
			573: 4793,
		},
	},
	{
		"minecraft:cauldron[level=0]",
		nil,
		NewMapping{
			393: 4621,
			404: 4622,
			477: 5125,
			573: 5125,
			718: 5141,
			4:   1888,
		},
	},
	{
		"minecraft:dead_bubble_coral[waterlogged=true]",
		nil,
		NewMapping{
			404: 8464,
			477: 8988,
			573: 8988,
			718: 9524,
		},
	},
	{
		"minecraft:polished_blackstone_brick_stairs[facing=west,half=bottom,shape=inner_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 16312,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=low,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 13943,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=none,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11109,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=low,north=low,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13628,
		},
	},
	{
		"minecraft:blackstone_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			718: 15911,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=tall,south=none,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11706,
		},
	},
	{
		"minecraft:pink_shulker_box[facing=up]",
		nil,
		NewMapping{
			718: 9318,
			4:   3601,
			393: 8257,
			404: 8258,
			477: 8782,
			573: 8782,
		},
	},
	{
		"minecraft:red_shulker_box[facing=west]",
		nil,
		NewMapping{
			393: 8304,
			404: 8305,
			477: 8829,
			573: 8829,
			718: 9365,
			4:   3732,
		},
	},
	{
		"minecraft:crimson_fence_gate[facing=north,in_wall=true,open=true,powered=true]",
		nil,
		NewMapping{
			718: 15255,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=tall,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11938,
		},
	},
	{
		"minecraft:dropper[facing=down,triggered=false]",
		nil,
		NewMapping{
			393: 5803,
			404: 5804,
			477: 6310,
			573: 6310,
			718: 6846,
			4:   2528,
		},
	},
	{
		"minecraft:mossy_stone_brick_stairs[facing=north,half=bottom,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9303,
			573: 9303,
			718: 9839,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 17010,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=none,north=none,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 14123,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=tall,south=tall,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 11291,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=none,north=none,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11858,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=tall,north=low,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 16604,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=tall,north=low,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14367,
		},
	},
	{
		"minecraft:andesite_wall[east=none,north=tall,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13237,
		},
	},
	{
		"minecraft:spruce_leaves[distance=2,persistent=true]",
		nil,
		NewMapping{
			718: 161,
			4:   301,
			393: 160,
			404: 160,
			477: 160,
			573: 160,
		},
	},
	{
		"minecraft:acacia_door[facing=east,half=lower,hinge=right,open=true,powered=false]",
		nil,
		NewMapping{
			4:   3140,
			393: 7930,
			404: 7931,
			477: 8455,
			573: 8455,
			718: 8991,
		},
	},
	{
		"minecraft:moving_piston[facing=up,type=sticky]",
		nil,
		NewMapping{
			4:   585,
			393: 1108,
			404: 1108,
			477: 1408,
			573: 1408,
			718: 1409,
		},
	},
	{
		"minecraft:brick_wall[east=tall,north=tall,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11181,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=none,power=12,south=up,west=side]",
		nil,
		NewMapping{
			404: 2150,
			477: 2453,
			573: 2453,
			718: 2455,
			393: 2149,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=5,powered=true]",
		nil,
		NewMapping{
			477: 858,
			573: 858,
			718: 859,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=none,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 12405,
		},
	},
	{
		"minecraft:diorite_wall[east=low,north=low,south=none,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14580,
		},
	},
	{
		"minecraft:oak_door[facing=south,half=upper,hinge=left,open=true,powered=true]",
		nil,
		NewMapping{
			404: 3124,
			477: 3587,
			573: 3587,
			718: 3589,
			393: 3123,
		},
	},
	{
		"minecraft:nether_brick_stairs[facing=south,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			573: 5063,
			718: 5079,
			4:   1826,
			393: 4559,
			404: 4560,
			477: 5063,
		},
	},
	{
		"minecraft:brown_banner[rotation=7]",
		nil,
		NewMapping{
			477: 7560,
			573: 7560,
			718: 8096,
			393: 7053,
			404: 7054,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12848,
		},
	},
	{
		"minecraft:crimson_stairs[facing=south,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			718: 15345,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=none,south=tall,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 12512,
		},
	},
	{
		"minecraft:birch_stairs[facing=south,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			393: 4998,
			404: 4999,
			477: 5502,
			573: 5502,
			718: 5518,
		},
	},
	{
		"minecraft:blue_wall_banner[facing=west]",
		nil,
		NewMapping{
			477: 7663,
			573: 7663,
			718: 8199,
			393: 7156,
			404: 7157,
		},
	},
	{
		"minecraft:red_mushroom_block[down=true,east=false,north=false,south=false,up=true,west=true]",
		nil,
		NewMapping{
			393: 4079,
			404: 4080,
			477: 4583,
			573: 4583,
			718: 4597,
		},
	},
	{
		"minecraft:mossy_stone_brick_wall[east=tall,north=low,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12106,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=10,south=none,west=up]",
		nil,
		NewMapping{
			393: 3000,
			404: 3001,
			477: 3304,
			573: 3304,
			718: 3306,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=none,north=low,south=none,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16815,
		},
	},
	{
		"minecraft:note_block[instrument=didgeridoo,note=23,powered=true]",
		nil,
		NewMapping{
			718: 895,
			477: 894,
			573: 894,
		},
	},
	{
		"minecraft:prismarine_wall[east=tall,north=none,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 11433,
		},
	},
	{
		"minecraft:nether_brick_wall[east=tall,north=tall,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 13119,
		},
	},
	{
		"minecraft:repeater[delay=4,facing=north,locked=false,powered=true]",
		nil,
		NewMapping{
			393: 3563,
			404: 3564,
			477: 4067,
			573: 4067,
			718: 4081,
			4:   1518,
		},
	},
	{
		"minecraft:lever[face=ceiling,facing=west,powered=false]",
		nil,
		NewMapping{
			4:   1104,
			393: 3298,
			404: 3299,
			477: 3802,
			573: 3802,
			718: 3804,
		},
	},
	{
		"minecraft:quartz_stairs[facing=south,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			393: 5733,
			404: 5734,
			477: 6240,
			573: 6240,
			718: 6776,
		},
	},
	{
		"minecraft:red_nether_brick_wall[east=true,north=false,south=true,up=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10865,
			573: 10865,
		},
	},
	{
		"minecraft:purpur_stairs[facing=north,half=bottom,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			404: 8091,
			477: 8615,
			573: 8615,
			718: 9151,
			393: 8090,
		},
	},
	{
		"minecraft:brown_shulker_box[facing=east]",
		nil,
		NewMapping{
			477: 8815,
			573: 8815,
			718: 9351,
			4:   3701,
			393: 8290,
			404: 8291,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 14288,
		},
	},
	{
		"minecraft:dragon_head[rotation=13]",
		nil,
		NewMapping{
			477: 6067,
			573: 6067,
			718: 6603,
			393: 5564,
			404: 5565,
		},
	},
	{
		"minecraft:cyan_concrete",
		nil,
		NewMapping{
			718: 9447,
			4:   4025,
			393: 8386,
			404: 8387,
			477: 8911,
			573: 8911,
		},
	},
	{
		"minecraft:polished_andesite_stairs[facing=east,half=top,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			477: 10159,
			573: 10159,
			718: 10695,
		},
	},
	{
		"minecraft:warped_fence[east=false,north=true,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			718: 15116,
		},
	},
	{
		"minecraft:birch_trapdoor[facing=south,half=bottom,open=false,powered=true,waterlogged=true]",
		nil,
		NewMapping{
			393: 3749,
			404: 3750,
			477: 4253,
			573: 4253,
			718: 4267,
		},
	},
	{
		"minecraft:wither_skeleton_skull[rotation=3]",
		nil,
		NewMapping{
			393: 5474,
			404: 5475,
			477: 5977,
			573: 5977,
			718: 6513,
		},
	},
	{
		"minecraft:carrots[age=1]",
		nil,
		NewMapping{
			404: 5289,
			477: 5795,
			573: 5795,
			718: 6331,
			4:   2257,
			393: 5288,
		},
	},
	{
		"minecraft:skeleton_skull[rotation=12]",
		nil,
		NewMapping{
			393: 5463,
			404: 5464,
			477: 5966,
			573: 5966,
			718: 6502,
		},
	},
	{
		"minecraft:birch_sign[rotation=15,waterlogged=false]",
		nil,
		NewMapping{
			477: 3474,
			573: 3474,
			718: 3476,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=north,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9381,
			573: 9381,
			718: 9917,
		},
	},
	{
		"minecraft:granite_wall[east=low,north=none,south=none,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12277,
		},
	},
	{
		"minecraft:cut_red_sandstone_slab[type=double,waterlogged=true]",
		nil,
		NewMapping{
			477: 7870,
			573: 7870,
			718: 8406,
		},
	},
	{
		"minecraft:smooth_quartz_stairs[facing=north,half=top,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 9774,
			573: 9774,
			718: 10310,
		},
	},
	{
		"minecraft:crimson_trapdoor[facing=south,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			718: 15154,
		},
	},
	{
		"minecraft:smooth_sandstone_stairs[facing=south,half=top,shape=inner_left,waterlogged=false]",
		nil,
		NewMapping{
			477: 9716,
			573: 9716,
			718: 10252,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=low,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 16925,
		},
	},
	{
		"minecraft:polished_blackstone_wall[east=low,north=low,south=tall,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16947,
		},
	},
	{
		"minecraft:oak_stairs[facing=north,half=top,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			393: 1657,
			404: 1658,
			477: 1961,
			573: 1961,
			718: 1963,
		},
	},
	{
		"minecraft:note_block[instrument=bass,note=3,powered=false]",
		nil,
		NewMapping{
			404: 455,
			477: 455,
			573: 455,
			718: 456,
			393: 455,
		},
	},
	{
		"minecraft:iron_door[facing=south,half=upper,hinge=left,open=false,powered=false]",
		nil,
		NewMapping{
			393: 3322,
			404: 3323,
			477: 3826,
			573: 3826,
			718: 3828,
		},
	},
	{
		"minecraft:nether_brick_wall[east=none,north=tall,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12899,
		},
	},
	{
		"minecraft:warped_trapdoor[facing=south,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			718: 15218,
		},
	},
	{
		"minecraft:obsidian",
		nil,
		NewMapping{
			393: 1129,
			404: 1130,
			477: 1433,
			573: 1433,
			718: 1434,
			4:   784,
		},
	},
	{
		"minecraft:vine[east=true,north=true,south=false,up=true,west=true]",
		nil,
		NewMapping{
			573: 4776,
			718: 4792,
			393: 4272,
			404: 4273,
			477: 4776,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=none,south=tall,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12517,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=none,north=tall,south=low,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 16432,
		},
	},
	{
		"minecraft:cobblestone_wall[east=none,north=low,south=low,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 5709,
		},
	},
	{
		"minecraft:andesite_wall[east=low,north=none,south=low,up=false,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 13261,
		},
	},
	{
		"minecraft:trapped_chest[facing=east,type=right,waterlogged=false]",
		nil,
		NewMapping{
			718: 6645,
			393: 5602,
			404: 5603,
			477: 6109,
			573: 6109,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=side,power=11,south=none,west=up]",
		nil,
		NewMapping{
			718: 2739,
			393: 2433,
			404: 2434,
			477: 2737,
			573: 2737,
		},
	},
	{
		"minecraft:spruce_planks",
		nil,
		NewMapping{
			393: 16,
			404: 16,
			477: 16,
			573: 16,
			718: 16,
			4:   81,
		},
	},
	{
		"minecraft:diorite_wall[east=true,north=true,south=false,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 11045,
			573: 11045,
		},
	},
	{
		"minecraft:redstone_wire[east=up,north=up,power=15,south=none,west=side]",
		nil,
		NewMapping{
			404: 1895,
			477: 2198,
			573: 2198,
			718: 2200,
			393: 1894,
		},
	},
	{
		"minecraft:yellow_stained_glass_pane[east=true,north=true,south=false,waterlogged=false,west=true]",
		nil,
		NewMapping{
			393: 5954,
			404: 5955,
			477: 6461,
			573: 6461,
			718: 6997,
		},
	},
	{
		"minecraft:polished_blackstone_stairs[facing=north,half=bottom,shape=outer_right,waterlogged=false]",
		nil,
		NewMapping{
			718: 16684,
		},
	},
	{
		"minecraft:black_bed[facing=north,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 991,
			404: 991,
			477: 1291,
			573: 1291,
			718: 1292,
		},
	},
	{
		"minecraft:nether_brick_wall[east=true,north=true,south=true,up=true,waterlogged=false,west=true]",
		nil,
		NewMapping{
			477: 10717,
			573: 10717,
		},
	},
	{
		"minecraft:sandstone_wall[east=low,north=low,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 13956,
		},
	},
	{
		"minecraft:andesite_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 13430,
		},
	},
	{
		"minecraft:sandstone_wall[east=tall,north=tall,south=none,up=false,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 14079,
		},
	},
	{
		"minecraft:prismarine_wall[east=none,north=low,south=none,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 11234,
		},
	},
	{
		"minecraft:oak_trapdoor[facing=east,half=bottom,open=true,powered=false,waterlogged=false]",
		nil,
		NewMapping{
			718: 4170,
			4:   1543,
			393: 3652,
			404: 3653,
			477: 4156,
			573: 4156,
		},
	},
	{
		"minecraft:spruce_trapdoor[facing=south,half=top,open=false,powered=true,waterlogged=false]",
		nil,
		NewMapping{
			573: 4182,
			718: 4196,
			393: 3678,
			404: 3679,
			477: 4182,
		},
	},
	{
		"minecraft:powered_rail[powered=true,shape=ascending_west]",
		nil,
		NewMapping{
			4:   443,
			393: 1007,
			404: 1007,
			477: 1307,
			573: 1307,
			718: 1308,
		},
	},
	{
		"minecraft:redstone_wire[east=side,north=up,power=14,south=up,west=up]",
		nil,
		NewMapping{
			393: 2310,
			404: 2311,
			477: 2614,
			573: 2614,
			718: 2616,
		},
	},
	{
		"minecraft:warped_fence[east=false,north=false,south=false,waterlogged=true,west=false]",
		nil,
		NewMapping{
			718: 15124,
		},
	},
	{
		"minecraft:polished_blackstone_brick_wall[east=low,north=none,south=low,up=false,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 16467,
		},
	},
	{
		"minecraft:mushroom_stem[down=true,east=true,north=true,south=true,up=false,west=false]",
		nil,
		NewMapping{
			718: 4636,
			393: 4118,
			404: 4119,
			477: 4622,
			573: 4622,
		},
	},
	{
		"minecraft:end_stone_brick_wall[east=low,north=low,south=tall,up=true,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 14280,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=none,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 6212,
		},
	},
	{
		"minecraft:note_block[instrument=guitar,note=22,powered=true]",
		nil,
		NewMapping{
			393: 642,
			404: 642,
			477: 642,
			573: 642,
			718: 643,
		},
	},
	{
		"minecraft:end_portal_frame[eye=false,facing=north]",
		nil,
		NewMapping{
			393: 4630,
			404: 4631,
			477: 5134,
			573: 5134,
			718: 5150,
			4:   1922,
		},
	},
	{
		"minecraft:cobblestone_wall[east=tall,north=low,south=none,up=true,waterlogged=true,west=low]",
		nil,
		NewMapping{
			718: 5910,
		},
	},
	{
		"minecraft:stone_brick_wall[east=low,north=tall,south=none,up=true,waterlogged=false,west=low]",
		nil,
		NewMapping{
			718: 12671,
		},
	},
	{
		"minecraft:jungle_button[face=floor,facing=south,powered=false]",
		nil,
		NewMapping{
			393: 5378,
			404: 5379,
			477: 5885,
			573: 5885,
			718: 6421,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=east,half=bottom,shape=outer_left,waterlogged=true]",
		nil,
		NewMapping{
			393: 4488,
			404: 4489,
			477: 4992,
			573: 4992,
			718: 5008,
		},
	},
	{
		"minecraft:note_block[instrument=pling,note=23,powered=true]",
		nil,
		NewMapping{
			477: 1044,
			573: 1044,
			718: 1045,
		},
	},
	{
		"minecraft:crimson_door[facing=east,half=lower,hinge=right,open=true,powered=true]",
		nil,
		NewMapping{
			718: 15587,
		},
	},
	{
		"minecraft:fire[age=3,east=true,north=false,south=false,up=false,west=true]",
		nil,
		NewMapping{
			393: 1245,
			404: 1246,
			477: 1549,
			573: 1549,
			718: 1550,
		},
	},
	{
		"minecraft:green_bed[facing=east,occupied=false,part=foot]",
		nil,
		NewMapping{
			393: 971,
			404: 971,
			477: 1271,
			573: 1271,
			718: 1272,
		},
	},
	{
		"minecraft:polished_diorite_stairs[facing=south,half=top,shape=straight,waterlogged=true]",
		nil,
		NewMapping{
			477: 9393,
			573: 9393,
			718: 9929,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=low,north=low,south=none,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 11668,
		},
	},
	{
		"minecraft:granite_wall[east=tall,north=none,south=low,up=false,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 12400,
		},
	},
	{
		"minecraft:pink_wall_banner[facing=east]",
		nil,
		NewMapping{
			718: 8180,
			393: 7137,
			404: 7138,
			477: 7644,
			573: 7644,
		},
	},
	{
		"minecraft:lava[level=11]",
		nil,
		NewMapping{
			404: 61,
			477: 61,
			573: 61,
			718: 61,
			4:   171,
			393: 61,
		},
	},
	{
		"minecraft:redstone_wire[east=none,north=none,power=12,south=side,west=up]",
		nil,
		NewMapping{
			393: 3015,
			404: 3016,
			477: 3319,
			573: 3319,
			718: 3321,
		},
	},
	{
		"minecraft:end_stone_brick_stairs[facing=north,half=bottom,shape=inner_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 9547,
			573: 9547,
			718: 10083,
		},
	},
	{
		"minecraft:brick_wall[east=false,north=true,south=false,up=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			477: 10375,
			573: 10375,
		},
	},
	{
		"minecraft:blackstone_wall[east=low,north=low,south=tall,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 16091,
		},
	},
	{
		"minecraft:mossy_cobblestone_wall[east=tall,north=tall,south=low,up=true,waterlogged=true,west=tall]",
		nil,
		NewMapping{
			718: 6283,
		},
	},
	{
		"minecraft:glass_pane[east=true,north=true,south=false,waterlogged=true,west=true]",
		nil,
		NewMapping{
			573: 4719,
			718: 4735,
			393: 4215,
			404: 4216,
			477: 4719,
		},
	},
	{
		"minecraft:diorite_stairs[facing=east,half=bottom,shape=straight,waterlogged=false]",
		nil,
		NewMapping{
			477: 10244,
			573: 10244,
			718: 10780,
		},
	},
	{
		"minecraft:cobblestone_wall[east=low,north=tall,south=low,up=true,waterlogged=false,west=none]",
		nil,
		NewMapping{
			718: 5852,
		},
	},
	{
		"minecraft:red_sandstone_wall[east=tall,north=tall,south=low,up=false,waterlogged=false,west=tall]",
		nil,
		NewMapping{
			718: 11826,
		},
	},
	{
		"minecraft:fire[age=3,east=false,north=false,south=false,up=true,west=false]",
		nil,
		NewMapping{
			718: 1565,
			393: 1260,
			404: 1261,
			477: 1564,
			573: 1564,
		},
	},
	{
		"minecraft:dark_oak_slab[type=top,waterlogged=true]",
		nil,
		NewMapping{
			393: 7287,
			404: 7288,
			477: 7794,
			573: 7794,
			718: 8330,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=true,north=true,south=true,up=false,west=true]",
		nil,
		NewMapping{
			393: 4085,
			404: 4086,
			477: 4589,
			573: 4589,
			718: 4603,
		},
	},
	{
		"minecraft:red_mushroom_block[down=false,east=true,north=false,south=true,up=false,west=false]",
		nil,
		NewMapping{
			393: 4094,
			404: 4095,
			477: 4598,
			573: 4598,
			718: 4612,
		},
	},
	{
		"minecraft:target[power=1]",
		nil,
		NewMapping{
			718: 15761,
		},
	},
	{
		"minecraft:stone_brick_wall[east=none,north=tall,south=low,up=true,waterlogged=true,west=none]",
		nil,
		NewMapping{
			718: 12571,
		},
	},
	{
		"minecraft:jungle_fence_gate[facing=north,in_wall=false,open=true,powered=false]",
		nil,
		NewMapping{
			404: 7427,
			477: 7951,
			573: 7951,
			718: 8487,
			4:   2966,
			393: 7426,
		},
	},
	{
		"minecraft:light_gray_bed[facing=south,occupied=true,part=foot]",
		nil,
		NewMapping{
			477: 1181,
			573: 1181,
			718: 1182,
			393: 881,
			404: 881,
		},
	},
	{
		"minecraft:stone_brick_stairs[facing=south,half=top,shape=outer_right,waterlogged=true]",
		nil,
		NewMapping{
			477: 4944,
			573: 4944,
			718: 4960,
			393: 4440,
			404: 4441,
		},
	},
	{
		},
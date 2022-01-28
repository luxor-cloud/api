package main

var users = map[string]user{
	"589bfd56-c5bc-40d0-a502-19d8c354a9c4": {
		ID:   "589bfd56-c5bc-40d0-a502-19d8c354a9c4",
		Name: "xDRofflLOL",
	},
	"93cbc7c9-8e4f-4406-bd7e-66b33bee417b": {
		ID:   "93cbc7c9-8e4f-4406-bd7e-66b33bee417b",
		Name: "ihasdemskillz",
	},
	"e937678d-1ace-4952-b63b-a094524dd439": {
		ID:   "e937678d-1ace-4952-b63b-a094524dd439",
		Name: "speckMitWuerstchen",
	},
	"7333abbb-4f39-447f-b218-2cecca9a3108": {
		ID:   "7333abbb-4f39-447f-b218-2cecca9a3108",
		Name: "ichbineingeilertyp",
	},
	"2292e66b-50dc-4e1f-8fd2-f5cc1b541a0b": {
		ID:   "2292e66b-50dc-4e1f-8fd2-f5cc1b541a0b",
		Name: "dasisteinname",
	},
}

var serversByUser = map[string][]server{
	"589bfd56-c5bc-40d0-a502-19d8c354a9c4": []server{
		{
			ID:   "5fbb31ba-e1c0-45ca-a895-5179b0c342e4",
			Name: "mc02",
			IP:   "1.2.3.4",
		},
		{
			ID:   "52700f7b-14a4-4741-8b74-0d6a511cf681",
			Name: "HytaleServer",
			IP:   "113.140.73.224",
		},
		{
			ID:   "e562fb47-a211-4e24-a6b8-e9ab57511d5e",
			Name: "bierdose.net",
			IP:   "50.231.167.69",
		},
	},
	"93cbc7c9-8e4f-4406-bd7e-66b33bee417b": []server{
		{
			ID:   "875c3a4d-c6e3-4db7-bb4b-36d620338087",
			Name: "meinweinkeller",
			IP:   "4.109.172.18",
		},
		{
			ID:   "413e48b2-7661-4c2c-bfd6-4d6ed525e4ac",
			Name: "csgo-fsn1-dc15",
			IP:   "18.229.27.92",
		},
		{
			ID:   "6bdd4730-c454-4035-b4c4-fbf61554d71d",
			Name: "DASISTABERTOLL",
			IP:   "203.164.20.3",
		},
		{
			ID:   "67206e28-a0c8-40e5-8c27-ff12b4afbe15",
			Name: "bergwerkLABS",
			IP:   "2553:677d:3729:ffb5:98d1:d2b0:43c4:9437",
		},
		{
			ID:   "1a07cfac-7e44-446a-a6b1-bf9e6e674b31",
			Name: "bErgWerKLagS",
			IP:   "7665:8051:7d94:d459:a1dd:77d1:cc56:9cd1",
		},
		{
			ID:   "a48e979e-3da9-4e96-9034-99a6b8f29d0b",
			Name: "trüffelsindEKELHAFT",
			IP:   "a15d:c828:199d:e4ac:4a7d:1595:3ac0:74e6",
		},
		{
			ID:   "595361b3-d1d3-47b3-b65b-e3bfd0f76464",
			Name: "__soabernicht",
			IP:   "dc60:d549:7486:9be2:17a5:b356:f3e4:cb0a",
		},
	},
	"e937678d-1ace-4952-b63b-a094524dd439": []server{
		{
			ID:   "ad4d8989-b4e2-4bf7-8cc7-79ff06e8ab85",
			Name: "LegoUniverse",
			IP:   "47f1:5fa4:c189:7ef0:6a23:a0e2:6405:bf49",
		},
	},
	"7333abbb-4f39-447f-b218-2cecca9a3108": []server{
		{
			ID:   "b992e6a0-3d42-4d5c-a6f2-2d71e84cd690",
			Name: "L",
			IP:   "1764:e48c:fb8c:cb40:33d7:72f2:fe21:91f2",
		},
		{
			ID:   "823afc5f-3191-4cea-bc8c-41c38fccb7b5",
			Name: "mc1",
			IP:   "102.97.98.184",
		},
		{
			ID:   "d62684f2-8765-48ea-8198-d9fb73999a55",
			Name: "mc2",
			IP:   "199.18.136.27",
		},
		{
			ID:   "212443c4-671f-44b8-a9a8-b309c28f9767",
			Name: "mc3",
			IP:   "122.205.14.82",
		},
	},
	"2292e66b-50dc-4e1f-8fd2-f5cc1b541a0b": []server{
		{
			ID:   "4df75429-dc8e-4063-bd6b-bb0df5d6b2d6",
			Name: "battlefieldplay4free",
			IP:   "141.57.128.253",
		},
		{
			ID:   "cce7b731-98cd-4f2a-8843-5414096acb1f",
			Name: "123123123",
			IP:   "178.145.201.224",
		},
		{
			ID:   "1d41882a-3fad-42ff-85d8-037122106fe6",
			Name: "Dies",
			IP:   "5fb8:16c7:f507:5fd5:36f9:956c:9836:8eca",
		},
		{
			ID:   "d8f2807e-0693-4af4-9754-5f5156122eba",
			Name: "WorldOfWarcraft",
			IP:   "d117:d448:db69:f684:b5a5:7b0a:b5b1:733f",
		},
		{
			ID:   "c120b497-e098-4a8b-bf08-f1a232f28def",
			Name: "NEINDASKANNICHTSEIN",
			IP:   "5093:bdf7:a55a:3cc6:08ad:0c29:ff00:dcfa",
		},
	},
}

const state = {	
	walmartList: [
        {
            "id":0,
            "image": "https://www.walmart.com//i5.walmartimages.com/asr/bb905d9b-4193-443d-949c-3883338ebde3_1.7dbd3027f63b8f16f075a9ac4d860119.jpeg?odnWidth=282&odnHeight=282&odnBg=ffffff",
            "name": "HORMEL MARY KITCHEN Corned Beef Hash, 7.5 oz",
            "price": "$1.38",
            "rating": "4.8 Stars. 40 reviews.",
            "link": "https:/ip/HORMEL-MARY-KITCHEN-Corned-Beef-Hash-7-5-oz/37532131?wpa_bd=&wpa_pg_seller_id=F55CDC31AB754BB68FE0B39041159D63&wpa_ref_id=wpaqs:vGlEbZAHM6QiSTu7EBj-uaFXpDnruM0y6meBJU9RGeJe5FCq0mKjekxeIj-i--K5gQSJ6odAn4kKNkJmVifRIFHDpy1n9xEgpcE6wo0bCHa-pKN8In1FcnFw4eLrklgNFR1pIJzx8NKRXUdcALlb8Y4d-9wCueApA8ktGMXNz5NO7pU_DKpS6WPSgG8gvnqSlNjY_3WoYkVJIUxhnI3EXQ&wpa_tag=__tag__&wpa_aux_info=__aux_info__&wpa_pos=1&wpa_plmt=__plmt__&wpa_aduid=__aduid__"
        },
        {
            "id":1,
            "image": "https://www.walmart.com//i5.walmartimages.com/asr/b6ccc571-3db7-4c95-98c4-a18dc0f8ef85_2.890efc06d40e935239c92d473006aaea.jpeg?odnWidth=282&odnHeight=282&odnBg=ffffff",
            "name": "All Natural* 93% Lean/7% Fat Lean Ground Beef Tray, 1 lb",
            "price": "$6.38",
            "rating": "3.7 Stars. 30 reviews.",
            "link": "https:/ip/All-Natural-93-Lean-7-Fat-Lean-Ground-Beef-Tray-1-lb/824841960"
        },
        {
            "id":2,
            "image": "https://www.walmart.com//i5.walmartimages.com/asr/c488097d-cd02-491d-9ba1-06414593dad5_3.86f74d35f449351329d6723151047398.jpeg?odnWidth=282&odnHeight=282&odnBg=ffffff",
            "name": "All Natural* 80% Lean/20% Fat Ground Beef Chuck Tray, 2.25 lb",
            "price": "$9.88",
            "rating": "4.1 Stars. 38 reviews.",
            "link": "https:/ip/All-Natural-80-Lean-20-Fat-Ground-Beef-Chuck-Tray-2-25-lb/959024059"
        },
        {
            "id":3,
            "image": "https://www.walmart.com//i5.walmartimages.com/asr/681d5c16-b472-4b24-bc6e-a817e5a68c79.6cde6a88932cea8d318872342991bdc3.png?odnWidth=282&odnHeight=282&odnBg=ffffff",
            "name": "Beef Bottom Round Steak Thin, 0.34 - 2.0 lb",
            "price": "$7.45",
            "rating": "3.0 Stars. 2 reviews.",
            "link": "https:/ip/Beef-Bottom-Round-Steak-Thin-0-34-2-0-lb/150293351"
        },
        {
            "id":4,
            "image": "https://www.walmart.com//i5.walmartimages.com/asr/87f2f5e4-0459-4b80-b874-d149badd0af0_3.4df2d3808f0ea5d68b969a3186230aa6.jpeg?odnWidth=282&odnHeight=282&odnBg=ffffff",
            "name": "All Natural* 80% Lean/20% Fat Ground Beef Chuck Tray, 1 lb",
            "price": "$5.24",
            "rating": "4.1 Stars. 20 reviews.",
            "link": "https:/ip/All-Natural-80-Lean-20-Fat-Ground-Beef-Chuck-Tray-1-lb/479601462"
        },
        {
            "id":5,
            "image": "https://www.walmart.com//i5.walmartimages.com/asr/01614244-eea3-437d-b2d1-3fb0fd27fe36.f7607ba58afa85a0f8ad7d403c7e64ed.png?odnWidth=282&odnHeight=282&odnBg=ffffff",
            "name": "Beef Stew Meat Family Pack, 2.15 - 3.0 lb",
            "price": "$19.15",
            "rating": "4.1 Stars. 14 reviews.",
            "link": "https:/ip/Beef-Stew-Meat-Family-Pack-2-15-3-0-lb/39944484"
        },
        {
            "id":6,
            "image": "https://www.walmart.com//i5.walmartimages.com/asr/d2626571-5fde-42d8-be51-8cae90e6343e_2.fb3ee160748d66194ee7c70c2ac41812.jpeg?odnWidth=282&odnHeight=282&odnBg=ffffff",
            "name": "All Natural* 73% Lean/27% Fat Lean Ground Beef",
            "price": "$3.97",
            "rating": "4.3 Stars. 37 reviews.",
            "link": "https:/ip/All-Natural-73-Lean-27-Fat-Lean-Ground-Beef/15136790"
        },
        {
            "id":7,
            "image": "https://www.walmart.com//i5.walmartimages.com/asr/aecbd6bd-f14e-4fe8-a8af-856d7b6f131d_2.9d5a24660d75787acd2129937510b34e.jpeg?odnWidth=282&odnHeight=282&odnBg=ffffff",
            "name": "All Natural* 73% Lean/27% Fat Ground Beef Roll, 5 lb",
            "price": "$15.24",
            "rating": "3.2 Stars. 90 reviews.",
            "link": "https:/ip/All-Natural-73-Lean-27-Fat-Ground-Beef-Roll-5-lb/15136795"
        },
        {
            "id":8,
            "image": "https://www.walmart.com//i5.walmartimages.com/asr/9c9213e6-cda6-400a-9e6f-c76efe78dc80_3.861f62f13bc28245770602e5dff62d04.jpeg?odnWidth=282&odnHeight=282&odnBg=ffffff",
            "name": "All Natural* 93% Lean/7% Fat Lean Ground Beef Tray, 2.25 lb",
            "price": "$13.18",
            "rating": "4.4 Stars. 29 reviews.",
            "link": "https:/ip/All-Natural-93-Lean-7-Fat-Lean-Ground-Beef-Tray-2-25-lb/512398950"
        },
        {
            "id":9,
            "image": "https://www.walmart.com//i5.walmartimages.com/asr/4a551ab1-b0e3-4620-92f1-8cf6eed748c9.33d8653d9c5d00fd9c7a24f0031eaa3a.png?odnWidth=282&odnHeight=282&odnBg=ffffff",
            "name": "Beef Chuck Roast, 2.25 - 2.65 lb",
            "price": "$16.89",
            "rating": "3.5 Stars. 29 reviews.",
            "link": "https:/ip/Beef-Chuck-Roast-2-25-2-65-lb/55446514"
        },
        {
            "id":10,
            "image": "https://www.walmart.com//i5.walmartimages.com/asr/0cabd9d3-7428-46cf-9c37-97e3fc178954.9d2a1c6fdf55f7c944bf4d3cccf4ea1c.png?odnWidth=282&odnHeight=282&odnBg=ffffff",
            "name": "Beef Choice Angus Tri Tip Steak, 1.0 - 1.91 lb",
            "price": "$13.65",
            "rating": "3.5 Stars. 6 reviews.",
            "link": "https:/ip/Beef-Choice-Angus-Tri-Tip-Steak-1-0-1-91-lb/912016220"
        },
        
    ],
    targetList: [
        {
            "id":0,
            "image": "https://target.scene7.com/is/image/Target/GUEST_1f8f1293-0d68-4c23-83fa-76cd6b60a0de",
            "name": "All Natural 80/20 Ground Beef - 1lb - Good & Gather\u2122",
            "price": "https://target.scene7.com/is/image/Target/GUEST_1f8f1293-0d68-4c23-83fa-76cd6b60a0de",
            "rating": "3.5 out of 5 stars with 52 ratings",
            "link": "https://www.target.com/p/all-natural-80-20-ground-beef-1lb-good-38-gather-8482/-/A-13287606#lnk=sametab"
        },
        {
            "id":1,
            "image": "https://target.scene7.com/is/image/Target/GUEST_bf19be30-c059-4665-9a00-0ce662292e7c",
            "name": "USDA Choice Angus Beef Stew Meat - 1lb - Good & Gather\u2122",
            "price": "https://target.scene7.com/is/image/Target/GUEST_bf19be30-c059-4665-9a00-0ce662292e7c",
            "rating": "3.2 out of 5 stars with 20 ratings",
            "link": "https://www.target.com/p/usda-choice-angus-beef-stew-meat-1lb-good-38-gather-8482/-/A-14827741#lnk=sametab"
        },
        {
            "id":2,
            "image": "https://target.scene7.com/is/image/Target/GUEST_249496f1-3ea4-411e-9079-640072826289",
            "name": "Old Wisconsin Beef Sausage Sticks - 6oz",
            "price": "https://target.scene7.com/is/image/Target/GUEST_249496f1-3ea4-411e-9079-640072826289",
            "rating": "4.6 out of 5 stars with 26 ratings",
            "link": "https://www.target.com/p/old-wisconsin-beef-sausage-sticks-6oz/-/A-50428778#lnk=sametab"
        },
        {
            "id":3,
            "image": "https://target.scene7.com/is/image/Target/GUEST_f6fdad02-85d5-4335-aedd-591253f69d9b",
            "name": "USDA Choice Angus Beef Steak Strips - 14oz - Good & Gather\u2122",
            "price": "https://target.scene7.com/is/image/Target/GUEST_f6fdad02-85d5-4335-aedd-591253f69d9b",
            "rating": "2.9 out of 5 stars with 22 ratings",
            "link": "https://www.target.com/p/usda-choice-angus-beef-steak-strips-14oz-good-38-gather-8482/-/A-50030806#lnk=sametab"
        },
        {
            "id":4,
            "image": "https://target.scene7.com/is/image/Target/GUEST_5f6f3e32-c1e3-4258-aeeb-e4883e4ad99a",
            "name": "Organic 100% Grassfed 85/15 Ground Beef - 1lb - Good & Gather\u2122",
            "price": "https://target.scene7.com/is/image/Target/GUEST_5f6f3e32-c1e3-4258-aeeb-e4883e4ad99a",
            "rating": "4.2 out of 5 stars with 39 ratings",
            "link": "https://www.target.com/p/organic-100-grassfed-85-15-ground-beef-1lb-good-38-gather-8482/-/A-23974824#lnk=sametab"
        },
        {
            "id":5,
            "image": "https://target.scene7.com/is/image/Target/GUEST_9810f9e3-842c-4867-bec0-7c45a0c227a2",
            "name": "Impossible Burger Plant-Based Patties - 8oz",
            "price": "https://target.scene7.com/is/image/Target/GUEST_9810f9e3-842c-4867-bec0-7c45a0c227a2",
            "rating": "4.6 out of 5 stars with 5286 ratings",
            "link": "https://www.target.com/p/impossible-burger-plant-based-patties-8oz/-/A-79946726#lnk=sametab"
        },
        {
            "id":6,
            "image": "https://target.scene7.com/is/image/Target/GUEST_c96b704d-c22b-4f11-b7e8-62d56ca7ad37",
            "name": "Old Neighborhood Shaved Beef Steak - 14oz",
            "price": "https://target.scene7.com/is/image/Target/GUEST_c96b704d-c22b-4f11-b7e8-62d56ca7ad37",
            "rating": "4.3 out of 5 stars with 45 ratings",
            "link": "https://www.target.com/p/old-neighborhood-shaved-beef-steak-14oz/-/A-53318444#lnk=sametab"
        },
        {
            "id":7,
            "image": "https://target.scene7.com/is/image/Target/GUEST_a48beb63-385f-4cb4-9223-33040bc74dc7",
            "name": "All Natural 85/15 Ground Beef - 2lbs - Good & Gather\u2122",
            "price": "https://target.scene7.com/is/image/Target/GUEST_a48beb63-385f-4cb4-9223-33040bc74dc7",
            "rating": "4.2 out of 5 stars with 17 ratings",
            "link": "https://www.target.com/p/all-natural-85-15-ground-beef-2lbs-good-38-gather-8482/-/A-13287609#lnk=sametab"
        },
        {
            "id":8,
            "image": "https://target.scene7.com/is/image/Target/GUEST_c336d2b4-cff5-4004-a667-ad2960f64dc6",
            "name": "Impossible Burger Patties - Frozen - 1.5lbs/6ct",
            "price": "https://target.scene7.com/is/image/Target/GUEST_c336d2b4-cff5-4004-a667-ad2960f64dc6",
            "rating": "5 out of 5 stars with 7 ratings",
            "link": "https://www.target.com/p/impossible-burger-patties-frozen-1-5lbs-6ct/-/A-82539838#lnk=sametab"
        },
        {
            "id":9,
            "image": "https://target.scene7.com/is/image/Target/GUEST_ac0e305a-90f2-4448-94dd-0e597dc4a6e7",
            "name": "All Natural 85/15 Ground Beef - 1lb - Good & Gather\u2122",
            "price": "https://target.scene7.com/is/image/Target/GUEST_ac0e305a-90f2-4448-94dd-0e597dc4a6e7",
            "rating": "4.2 out of 5 stars with 26 ratings",
            "link": "https://www.target.com/p/all-natural-85-15-ground-beef-1lb-good-38-gather-8482/-/A-13288295#lnk=sametab"
        },
        {
            "id":10,
            "image": "https://target.scene7.com/is/image/Target/GUEST_6c7db3fc-c160-4451-a9be-0e86d3c360c3",
            "name": "Organic 100% Grassfed 90/10 Ground Beef - 1lb - Good & Gather\u2122",
            "price": "https://target.scene7.com/is/image/Target/GUEST_6c7db3fc-c160-4451-a9be-0e86d3c360c3",
            "rating": "4 out of 5 stars with 23 ratings",
            "link": "https://www.target.com/p/organic-100-grassfed-90-10-ground-beef-1lb-good-38-gather-8482/-/A-23974825#lnk=sametab"
        },
    ],
    walgreensList: [
        {
            "id":0,
            "image": "https://pics.drugstore.com/prodimg/624097/450.jpg",
            "name": "American FoodsGround Beef\u00a0\u00a0-\u00a016 oz  ",
            "price": "$4.29",
            "rating": "4.6 out of 5 total 63 reviews",
            "link": "https://www.walgreens.com/store/c/american-foods-ground-beef/ID=300402728-product"
        },
        {
            "id":1,
            "image": "https://pics.drugstore.com/prodimg/615311/450.jpg",
            "name": "PetShoppe PremiumBeef Tenders\u00a0\u00a0-\u00a012 oz  ",
            "price": "2/$18.00 or 1/$9.99",
            "rating": "4.6 out of 5 total 10 reviews",
            "link": "https://www.walgreens.com/store/c/petshoppe-premium-beef-tenders/ID=300394215-product"
        },
        {
            "id":2,
            "image": "https://pics.drugstore.com/prodimg/625469/450.jpg",
            "name": "PetShoppe PremiumBeef Jerky\u00a0\u00a0-\u00a012 oz  ",
            "price": "$12.99",
            "rating": "4.8 out of 5 total 15 reviews",
            "link": "https://www.walgreens.com/store/c/petshoppe-premium-beef-jerky/ID=300404129-product"
        },
        {
            "id":3,
            "image": "https://pics.drugstore.com/prodimg/509428/450.jpg",
            "name": "Libby'sCorned Beef\u00a0\u00a0-\u00a012 oz  ",
            "price": "$5.29",
            "rating": "4.6 out of 5 total 12 reviews",
            "link": "https://www.walgreens.com/store/c/libby's-corned-beef/ID=prod6082627-product"
        },
        {
            "id":4,
            "image": "https://pics.drugstore.com/prodimg/440738/450.jpg",
            "name": "Chef BoyardeeBeef Ravioli\u00a0\u00a0-\u00a015 oz  ",
            "price": "3/$3.00 or 1/$1.29",
            "rating": "2.9 out of 5 total 190 reviews",
            "link": "https://www.walgreens.com/store/c/chef-boyardee-beef-ravioli/ID=prod388-product"
        },
        {
            "id":5,
            "image": "https://pics.drugstore.com/prodimg/438230/450.jpg",
            "name": "Jack Link'sBeef Steak Teriyaki\u00a0\u00a0-\u00a02 oz  ",
            "price": "$2.99",
            "rating": "3.0 out of 5 total 2 reviews",
            "link": "https://www.walgreens.com/store/c/jack-link's-beef-steak-teriyaki/ID=prod6072549-product"
        },
        {
            "id":6,
            "image": "https://pics.drugstore.com/prodimg/597503/450.jpg",
            "name": "Nice!Beef Jerky Original\u00a0\u00a0-\u00a010 oz  ",
            "price": "$9.99",
            "rating": "4.0 out of 5 total 45 reviews",
            "link": "https://www.walgreens.com/store/c/nice!-beef-jerky-original/ID=prod6356032-product"
        },
        {
            "id":7,
            "image": "https://pics.drugstore.com/prodimg/623769/450.jpg",
            "name": "Jack Link'sBeef Jerky Teriyaki\u00a0\u00a0-\u00a010 oz  ",
            "price": "$12.99",
            "rating": "4.6 out of 5 total 18 reviews",
            "link": "https://www.walgreens.com/store/c/jack-link's-beef-jerky-teriyaki/ID=300402416-product"
        },
        {
            "id":8,
            "image": "https://pics.drugstore.com/prodimg/597560/450.jpg",
            "name": "Jack Link'sBeef Strips Teriyaki\u00a0\u00a0-\u00a02.85 oz  ",
            "price": "$5.99",
            "rating": "5.0 out of 5 total 1 reviews",
            "link": "https://www.walgreens.com/store/c/jack-link's-beef-strips-teriyaki/ID=prod6356057-product"
        },
        {
            "id":9,
            "image": "https://pics.drugstore.com/prodimg/567210/450.jpg",
            "name": "Jack Link'sBeef Jerky Teriyaki\u00a0\u00a0-\u00a02.85 oz  ",
            "price": "$5.99",
            "rating": "4.3 out of 5 total 11 reviews",
            "link": "https://www.walgreens.com/store/c/jack-link's-beef-jerky-teriyaki/ID=prod6286608-product"
        },
        {
            "id":10,
            "image": "https://pics.drugstore.com/prodimg/612500/450.jpg",
            "name": "Mary KitchenCorned Beef Hash\u00a0\u00a0-\u00a014 oz  ",
            "price": "3/$8.00 or 1/$2.99",
            "rating": "3.8 out of 5 total 36 reviews",
            "link": "https://www.walgreens.com/store/c/mary-kitchen-corned-beef-hash/ID=prod6377450-product"
        },
        
        {
            "id":11,
            "image": "https://pics.drugstore.com/prodimg/621080/450.jpg",
            "name": "BenefulOriginals with Beef Dog Food\u00a0\u00a0-\u00a012 lb  ",
            "price": "$15.99",
            "rating": "4.6 out of 5 total 1183 reviews",
            "link": "https://www.walgreens.com/store/c/beneful-originals-with-beef-dog-food/ID=300399645-product"
        },
        {
            "id":12,
            "image": "https://pics.drugstore.com/prodimg/625201/450.jpg",
            "name": "Jack Link'sBeef Steak Bites Original\u00a0\u00a0-\u00a01.5 oz  ",
            "price": "$2.49",
            "rating": "5.0 out of 5 total 1 reviews",
            "link": "https://www.walgreens.com/store/c/jack-link's-beef-steak-bites-original/ID=300403842-product"
        },
    ],
}

const mutations = {

	'ADD_STATUS' (state, id) {
		state.productList.forEach(function(item){
			if(item.testId==id){
				item.isAdd=true
			}
		})
	},
	'REMOVE_STATUS' (state, id) {
		state.productList.forEach(function(item){
			if(item.testId==id){
				item.isAdd=false
			}
		})
	}
}

const actions = {

}

const getters = {
	walmart: (state) => {
		return state.walmartList;
	},
	target: (state) => {
		return state.targetList;
	},
    walgreens: (state) => {
        return state.walgreensList;
    }
}

export default {
	state,
	mutations,
	actions,
	getters
}
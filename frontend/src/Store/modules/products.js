const state = {	
	walmartList: [
		{
			"id": 10,
			"image": "i5.walmartimages.com/asr/ed1192b2-48cc-4882-868b-e99a2f56911c_1.4d758f51f9f7dc625ef956cc04355bfa.png?odnWidth=282&odnHeight=282&odnBg=ffffff",
			"name": "Hormel Mary Kitchen Roast Beef Hash, 14 oz Can",
			"price": "None",
			"rating": "4.3 Stars. 391 reviews.",
			"link": "https://www.walmart.com//ip/Hormel-Mary-Kitchen-Roast-Beef-Hash-14-oz-Can/206684852"
		},
		{
			"id": 11,
			"image": "i5.walmartimages.com/asr/c00fce0e-6b68-4a78-82ab-e560b9efff6f.7c0d8dc07d68d722ba21dbb9472acb03.jpeg?odnWidth=282&odnHeight=282&odnBg=ffffff",
			"name": "Great Value Corned Beef, 12 oz Can",
			"price": "$3.34",
			"rating": "4.4 Stars. 439 reviews.",
			"link": "https://www.walmart.com//ip/Great-Value-Corned-Beef-12-oz-Can/45595282"
		},
		{
			"id": 12,
			"image": "i5.walmartimages.com/asr/ec69f481-74f2-43d9-b696-f76811e8b094.adda1ec7e4177f8a1d0bce23990c9b5a.jpeg?odnWidth=282&odnHeight=282&odnBg=ffffff",
			"name": "Beef Sirloin Tip Steak Thin, 0.85 - 1.61 lb",
			"price": "$9.93",
			"rating": "2.8 Stars. 17 reviews.",
			"link": "https://www.walmart.com//ip/Beef-Sirloin-Tip-Steak-Thin-0-85-1-61-lb/21553581"
		},
		{
			"id": 13,
			"image": "i5.walmartimages.com/asr/2e9e1449-ec2d-4835-9385-b7c060162126.eeacbcd092fb3d1f60e47c1b58095e31.jpeg?odnWidth=282&odnHeight=282&odnBg=ffffff",
			"name": "Great Value Corned Beef Hash, 14 oz Can",
			"price": "$1.74",
			"rating": "4.4 Stars. 25 reviews.",
			"link": "https://www.walmart.com//ip/Great-Value-Corned-Beef-Hash-14-oz-Can/848626120"
		},
		{
			"id": 14,
			"image": "i5.walmartimages.com/asr/4d3aaf7a-9ca4-4fe9-bbfe-5c2f498116c3.02297b1293285d894efb936cc6bcd5fb.jpeg?odnWidth=282&odnHeight=282&odnBg=ffffff",
			"name": "Beef Choice Angus Chuck Roast, 2.25 - 3.38 lb",
			"price": "$20.64",
			"rating": "3.1 Stars. 18 reviews.",
			"link": "https://www.walmart.com//ip/Beef-Choice-Angus-Chuck-Roast-2-25-3-38-lb/39944474"
		}
	],
    targetList: [
		{
			"id": 119,
			"image": "https://target.scene7.com/is/image/Target/GUEST_7c845c01-6cf1-40ce-b756-60d55faedcb8",
			"name": "USDA Choice Top Sirloin Steak - 1.18-2.24 lbs - price per lb - Good & Gather\u2122",
			"price": "$26.86",
			"rating": "3.4 out of 5 stars with 17 ratings",
			"link": "https://www.target.com/p/usda-choice-top-sirloin-steak-1-18-2-24-lbs-price-per-lb-good-38-gather-8482/-/A-54640789#lnk=sametab"
		},
		{
			"id": 120,
			"image": "https://target.scene7.com/is/image/Target/GUEST_f7992e1c-7127-46fb-a5b6-1de594076d4f",
			"name": "Laura's Lean Beef 96/4 Ground Beef - 1lb",
			"price": "$8.19",
			"rating": "3.8 out of 5 stars with 24 ratings",
			"link": "https://www.target.com/p/laura-s-lean-beef-96-4-ground-beef-1lb/-/A-13435602#lnk=sametab"
		},
		{
			"id": 121,
			"image": "https://target.scene7.com/is/image/Target/GUEST_59e18b4b-ffd1-4f10-bdbc-d92f7169728f",
			"name": "USDA Choice Angus Beef Chuck Tender Roast - 1.5-2.67 lbs - price per lb - Good & Gather\u2122",
			"price": "$20.00",
			"rating": "3.3 out of 5 stars with 16 ratings",
			"link": "https://www.target.com/p/usda-choice-angus-beef-chuck-tender-roast-1-5-2-67-lbs-price-per-lb-good-38-gather-8482/-/A-14915349#lnk=sametab"
		},
		{
			"id": 122,
			"image": "https://target.scene7.com/is/image/Target/GUEST_80c194a9-3e94-419b-bc71-873a04ac6e0d",
			"name": "Earth's Best Baked Mini Beef Meatballs - Frozen - 26oz",
			"price": "$10.99",
			"rating": "4.7 out of 5 stars with 29 ratings",
			"link": "https://www.target.com/p/earth-39-s-best-baked-mini-beef-meatballs-frozen-26oz/-/A-23974771#lnk=sametab"
		},
		{
			"id": 123,
			"image": "https://target.scene7.com/is/image/Target/GUEST_9989cfbb-4ac1-43ff-80e9-6edb2e444e81",
			"name": "USDA Choice Angus Petite Sirloin Steak - 0.68-1.13 lbs - price per lb - Good & Gather\u2122",
			"price": "$13.55",
			"rating": "4.1 out of 5 stars with 21 ratings",
			"link": "https://www.target.com/p/usda-choice-angus-petite-sirloin-steak-0-68-1-13-lbs-price-per-lb-good-38-gather-8482/-/A-54640782#lnk=sametab"
		},
		{
			"id": 124,
			"image": "https://target.scene7.com/is/image/Target/GUEST_c3e8c0f5-178f-4749-a561-88c469be8c33",
			"name": "All Natural 80/20 Ground Beef - 2lbs - Good & Gather\u2122",
			"price": "$10.79",
			"rating": "3.4 out of 5 stars with 7 ratings",
			"link": "https://www.target.com/p/all-natural-80-20-ground-beef-2lbs-good-38-gather-8482/-/A-14932043#lnk=sametab"
		},
		{
			"id": 125,
			"image": "https://target.scene7.com/is/image/Target/GUEST_d7a1e7a2-ab59-4ba6-824d-553d21f4757f",
			"name": "Italian Style Beef, Pork, & Chicken Meatballs - Frozen - 26oz - Good & Gather\u2122",
			"price": "$4.99",
			"rating": "3.9 out of 5 stars with 53 ratings",
			"link": "https://www.target.com/p/italian-style-beef-pork-38-chicken-meatballs-frozen-26oz-good-38-gather-8482/-/A-13379453#lnk=sametab"
		},
		{
			"id": 126,
			"image": "https://target.scene7.com/is/image/Target/GUEST_ff1e8eed-9cf4-4156-85a2-94cfe7bd1a1c",
			"name": "USDA Choice Angus Beef Chuck Steak - 0.93-1.55 lbs - price per lb - Good & Gather\u2122",
			"price": "$12.38",
			"rating": "3.8 out of 5 stars with 4 ratings",
			"link": "https://www.target.com/p/usda-choice-angus-beef-chuck-steak-0-93-1-55-lbs-price-per-lb-good-38-gather-8482/-/A-14933842#lnk=sametab"
		},
		{
			"id": 127,
			"image": "https://target.scene7.com/is/image/Target/GUEST_4cdf9af4-ef79-4da2-92e4-c51a31ca3799",
			"name": "USDA Choice Angus Beef Skirt Steak - 0.75-1.6 lbs - price per lb - Good & Gather\u2122",
			"price": "$15.98",
			"rating": "3.7 out of 5 stars with 7 ratings",
			"link": "https://www.target.com/p/usda-choice-angus-beef-skirt-steak-0-75-1-6-lbs-price-per-lb-good-38-gather-8482/-/A-14911555#lnk=sametab"
		}
	],
    walgreensList: [
		{
			"id": 244,
			"image": "https://pics.drugstore.com/prodimg/597562/450.jpg",
			"name": "Jack Link'sSnack Stick Original\u00a0\u00a0-\u00a01 oz  ",
			"price": "$1.29",
			"rating": "5.0 out of 5 total 2 reviews",
			"link": "https://www.walgreens.com/store/c/jack-link's-snack-stick-original/ID=prod6356056-product"
		},
		{
			"id": 245,
			"image": "https://pics.drugstore.com/prodimg/482636/450.jpg",
			"name": "NongshimNoodle Bowl Spicy Chicken\u00a0\u00a0-\u00a03.03 oz  ",
			"price": "$1.29",
			"rating": "No Rating",
			"link": "https://www.walgreens.com/store/c/nongshim-noodle-bowl-spicy-chicken/ID=prod6106988-product"
		},
		{
			"id": 246,
			"image": "https://pics.drugstore.com/prodimg/468875/450.jpg",
			"name": "Nice!Microwavable Cheeseburger\u00a0\u00a0-\u00a04.9 oz  ",
			"price": "2/$2.00 or 1/$1.29",
			"rating": "4.4 out of 5 total 41 reviews",
			"link": "https://www.walgreens.com/store/c/nice!-microwavable-cheeseburger/ID=prod6163624-product"
		},
		{
			"id": 247,
			"image": "https://pics.drugstore.com/prodimg/607898/450.jpg",
			"name": "NutrishWet Dog Food\u00a0\u00a0-\u00a08 OZ  ",
			"price": "2/$6.00 or 1/$3.19",
			"rating": "4.8 out of 5 total 223 reviews",
			"link": "https://www.walgreens.com/store/c/nutrish-wet-dog-food/ID=prod6383969-product"
		}
	],
}

const mutations = {

	'SET_PRODUCTS' (state, products) {
		state.walmartList = products[0].products;
        state.targetList = products[1].products;
        state.walgreensList = products[2].products;
	},
}

const actions = {
    setProduct({commit},products){
        commit('SET_PRODUCTS',products)
    }
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
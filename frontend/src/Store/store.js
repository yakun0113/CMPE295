import { createStore } from 'vuex';
import products from './modules/products';
import user from './modules/user';
const store = createStore({
		modules: {	
			products,
			user,
	}
  })

export default store;


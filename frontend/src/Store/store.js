import { createStore } from 'vuex';
import products from './modules/products';
import * as actions from './actions';

const store = createStore({
	actions,
		modules: {	
			products,
	}
  })

export default store;


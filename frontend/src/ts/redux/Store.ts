import {createStore, Store} from 'redux'
import {rootReducer} from '.'

import {StateTypes} from '..'
//sync store across tabs
//redux-state-sync - https://www.npmjs.com/package/redux-state-sync
export const store: Store<StateTypes> = createStore(rootReducer);



import {ComponentReducer,CustomerReducer} from '.'
import { combineReducers } from 'redux'

export const rootReducer = combineReducers({
    Components : ComponentReducer,
    Customers  : CustomerReducer
})







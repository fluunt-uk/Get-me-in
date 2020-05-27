import { AnyAction } from "redux";
import {SET_STATUS } from ".";
// import { GlobalComponentValues } from "../helpers";

const initialState : object = {LoginPageState:false}

export default function ComponentReducer(state = initialState, action:AnyAction):object{
    const payload = action.payload
    switch (action.type){
        case SET_STATUS:
            return Object.assign({}, state, payload);
        default: 
            return state
    }   
}



// import {DefaultStoreResolver,NestedValueResolver,NumCheck} from '..'
import { AnyAction } from "redux";
import 
{
    CustomerType,
    SET_DEFAULT,
    // ADD_RECEIVER,
    // REMOVE_RECEIVER,
    // MODIFY_RECEIVER,
    // MODIFY_RECEIVER_ELEMENT
} from "..";

const initialState:Array<CustomerType> =  []

export default function ReceiverReducer(state = initialState, action:AnyAction):Array<CustomerType>{
    let result:Array<object> =[]
    const payload = action.payload

    switch (action.type){
        case SET_DEFAULT:
        // case ADD_RECEIVER:
        // case REMOVE_RECEIVER:
        // case MODIFY_RECEIVER:
        // case MODIFY_RECEIVER_ELEMENT:
          

        default:
            return state
            
    }
}


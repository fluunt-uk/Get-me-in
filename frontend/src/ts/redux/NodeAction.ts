
import 
{
    GlobalComponentValues, 
    StateTypes,
    ActionTypes
} from '..'

//change below

export const SET_DEFAULT = "SET_DEFAULT"; 
export const SET_STATUS = 'SET_STATUS'

export function SetDefault(payload:StateTypes): ActionTypes {
    return { type: SET_DEFAULT, payload }
};
export function SetStatus(payload:GlobalComponentValues): ActionTypes {
    return { type: SET_STATUS, payload }
};





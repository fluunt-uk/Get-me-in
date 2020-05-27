//Defines type for actions that change the state
export interface ActionTypes
{
    type: string;
    payload: StateTypes |CustomerType | GlobalComponentValues | string;
}

//Contains all props including Arrays of State Values and Actions to change state
export interface PropItems
{
    Customers: Array<CustomerType>
    Components: GlobalComponentValues
    Default: Function
    Status: Function
    dispatch: Function
    history: any
}

//Type for values that will be in the state
export interface StateTypes
{
    Components: GlobalComponentValues
    Customers: Array<CustomerType>
}

//Global Values that effect components
export interface GlobalComponentValues
{
    LoginPageState:Boolean
} 

//Customer Information
export interface CustomerType{}

//Used for writing information to immutables
export interface InternalObjects
{
    [key : string] : string | number;
}


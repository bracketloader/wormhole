/**
* This file was automatically generated by @cosmwasm/ts-codegen@0.24.0.
* DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema file,
* and run the @cosmwasm/ts-codegen generate command to regenerate this file.
*/

export type Binary = string;
export type Uint128 = string;
export interface InstantiateMsg {
  asset_address: Binary;
  asset_chain: number;
  decimals: number;
  init_hook?: InitHook | null;
  mint?: InitMint | null;
  name: string;
  symbol: string;
  [k: string]: unknown;
}
export interface InitHook {
  contract_addr: string;
  msg: Binary;
  [k: string]: unknown;
}
export interface InitMint {
  amount: Uint128;
  recipient: string;
  [k: string]: unknown;
}
export type ExecuteMsg = {
  transfer: {
    amount: Uint128;
    recipient: string;
    [k: string]: unknown;
  };
} | {
  burn: {
    account: string;
    amount: Uint128;
    [k: string]: unknown;
  };
} | {
  send: {
    amount: Uint128;
    contract: string;
    msg: Binary;
    [k: string]: unknown;
  };
} | {
  mint: {
    amount: Uint128;
    recipient: string;
    [k: string]: unknown;
  };
} | {
  increase_allowance: {
    amount: Uint128;
    expires?: Expiration | null;
    spender: string;
    [k: string]: unknown;
  };
} | {
  decrease_allowance: {
    amount: Uint128;
    expires?: Expiration | null;
    spender: string;
    [k: string]: unknown;
  };
} | {
  transfer_from: {
    amount: Uint128;
    owner: string;
    recipient: string;
    [k: string]: unknown;
  };
} | {
  send_from: {
    amount: Uint128;
    contract: string;
    msg: Binary;
    owner: string;
    [k: string]: unknown;
  };
} | {
  burn_from: {
    amount: Uint128;
    owner: string;
    [k: string]: unknown;
  };
} | {
  update_metadata: {
    name: string;
    symbol: string;
    [k: string]: unknown;
  };
};
export type Expiration = {
  at_height: number;
} | {
  at_time: Timestamp;
} | {
  never: {
    [k: string]: unknown;
  };
};
export type Timestamp = Uint64;
export type Uint64 = string;
export type QueryMsg = {
  wrapped_asset_info: {
    [k: string]: unknown;
  };
} | {
  balance: {
    address: string;
    [k: string]: unknown;
  };
} | {
  token_info: {
    [k: string]: unknown;
  };
} | {
  allowance: {
    owner: string;
    spender: string;
    [k: string]: unknown;
  };
};
export interface AllowanceResponse {
  allowance: Uint128;
  expires: Expiration;
  [k: string]: unknown;
}
export interface BalanceResponse {
  balance: Uint128;
  [k: string]: unknown;
}
export interface TokenInfoResponse {
  decimals: number;
  name: string;
  symbol: string;
  total_supply: Uint128;
  [k: string]: unknown;
}
export type Addr = string;
export interface WrappedAssetInfoResponse {
  asset_address: Binary;
  asset_chain: number;
  bridge: Addr;
  [k: string]: unknown;
}
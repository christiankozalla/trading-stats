import PocketBase, { type RecordService, type RecordModel } from 'pocketbase';

interface TypedPocketBase extends PocketBase {
  collection(idOrName: string): RecordService;
  collection(idOrName: 'trades'): RecordService<Trade>;
  collection(idOrName: 'trading_accounts'): RecordService<TradingAccount>;
  collection(idOrName: 'profit_loss'): RecordService<ProfitLoss>;
}

export const pb = new PocketBase(
  import.meta.env.POCKETBASE_BASE_URL || 'http://localhost:8090'
) as TypedPocketBase;

export type ProfitLoss = {
  DateTime_close: string;
  ProfitLoss_dollar: number;
} & RecordModel;

export type TradingAccount = {
  name: string;
} & RecordModel;

export type User = {
  avatar: string;
  collectionId: string;
  collectionName: string;
  created: string;
  emailVisibility: boolean;
  email?: string;
  id: string;
  name: string;
  updated: string;
  username: string;
  verified: boolean;
} & RecordModel;

export type Trade = {
  BuySell_close: string;
  BuySell_open: string;
  FillPrice_close: number;
  FillPrice_open: number;
  ProfitLoss: number;
  Multiplier: number;
  Quantity_close: number;
  Quantity_open: number;
  Symbol: string;
  ShortSymbol: string;
  account: string;
  collectionId: string;
  collectionName: string;
  id: string;
  t1_DateTime: string;
  user: string;
} & RecordModel;

export type Collections = 'trades' | 'users' | 'raw_trades' | 'trade_log_files' | 'profit_loss';

import PocketBase, { type RecordService, type RecordModel } from 'pocketbase';
import { supportedOrFallbackLocale } from './router/helpers';
import { ref } from 'vue';

interface TypedPocketBase extends PocketBase {
  collection(idOrName: string): RecordService;
  collection(idOrName: 'trades'): RecordService<Trade>;
  collection(idOrName: 'trading_accounts'): RecordService<TradingAccount>;
  collection(idOrName: 'profit_loss'): RecordService<ProfitLoss>;
}

// import.meta.env.BASE_URL is provided by Vite (default "/") which works, because Vue and Pocketbase run on the same domain
export const pb = new PocketBase(import.meta.env.BASE_URL) as TypedPocketBase;

// Alternative implementation idea:
// since it is problematic that pb and all its methods is not fully reactive and not integrated in Vue's reactivity system
// and since pb.authStore.onChange handler must fire immediately, but then always for the root route ("/"), useRoute/useRouter composables are really usable inside
// -> make a Pinia store, that initializes pb and then returns the instance, and a couple of Vue's reactive primitive only to keep the UI in sync with the auth state from pb
export const isAuthenticated = ref<boolean>(false);
pb.authStore.onChange((token, _model) => {
  if (!token) {
    const currentLocale = supportedOrFallbackLocale(window.location.pathname.slice(1, 3));
    isAuthenticated.value = false;
    window.location.pathname = `/${currentLocale}/login-signup`;
  } else {
    if (isAuthenticated.value !== true) isAuthenticated.value = true;
  }
});

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

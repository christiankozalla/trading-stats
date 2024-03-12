import PocketBase, { type RecordModel } from 'pocketbase';

const pb = new PocketBase(import.meta.env.POCKETBASE_BASE_URL || 'http://localhost:8090');

export type TradingAccount = {
  name: string
} & RecordModel


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
}

export { pb };

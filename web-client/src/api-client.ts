import PocketBase from 'pocketbase';

const pb = new PocketBase(import.meta.env.POCKETBASE_BASE_URL || 'http://localhost:8090');

console.log('api client isValid', pb.authStore.isValid);
console.log('api client isAuthRecord', pb.authStore.isAuthRecord);
console.log('api client record/record', pb.authStore.model);

interface CreateUserData {
  email: string;
  password: string;
  passwordConfirm: string;
  name?: string;
  username?: string;
}

export interface User {
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

export { pb,  };

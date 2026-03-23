export interface AssetsDb {
  CreatedAt: string;
  DeletedAt: string | null;
  ID: number;
  UpdatedAt: string;
  assetType: string;
  category: string;
  currency: string;
  description: string;
  liquidityScore: number;
  name: string;
  ownerType: string;
  status: string;
  total: string;
}

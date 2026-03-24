import type { RadioGroupItem } from "@nuxt/ui";
import type { InjectionKey } from "vue";
import type { AssetsDb } from "./asset";
import type { AsyncDataRequestStatus } from "#app";
import type { CategoryDb } from "./category";
import type { TransactionDb } from "./transaction";

export type BaseModalStatus = "add" | "edit" | "delete" | "detail" | null;

// >>>>>> ASSET SECTION <<<<<<

export interface AssetProvideInjectKey {
  modalOpen: Ref<BaseModalStatus>;
  assetId: Ref<number | null>;
  updateModalOpen: (status: BaseModalStatus, id?: number) => void;

  constant: {
    liquidityScoreMapping: Record<number, string>;
    statusItem: Ref<RadioGroupItem[]>;
    currencyItem: Ref<RadioGroupItem[]>;
  };

  fetcher: {
    data: Ref<{ message: string; data: AssetsDb[] } | undefined>;
    status: Ref<AsyncDataRequestStatus>;
  };
}

export const assetProvideKey: InjectionKey<AssetProvideInjectKey> = Symbol(
  "AssetProvideInjectKey",
);

// >>>>>> TRANSACTION SECTION <<<<<<
export interface TransactionProviderInjectKey {
  modal: {
    modalOpen: Ref<BaseModalStatus>;
    transactionId: Ref<string | null>;
    updateModalOpen: (status: BaseModalStatus, id?: string) => void;
  };

  resources: {
    categories: Ref<CategoryDb[]>;
    assets: Ref<AssetsDb[]>;
  };

  fetcher: {
    data: Ref<{ message: string; data: TransactionDb[] } | undefined>;
    status: Ref<AsyncDataRequestStatus>;
  };
}

export const transactionProviderKey: InjectionKey<TransactionProviderInjectKey> =
  Symbol("TransactionProviderInjectKey");

import type { RadioGroupItem } from "@nuxt/ui";
import type { InjectionKey } from "vue";
import type { AssetsDb } from "./asset";
import type { AsyncDataRequestStatus } from "#app";

export type BaseModalStatus = "add" | "edit" | "delete" | "detail" | null;

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

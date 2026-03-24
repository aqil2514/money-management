import { serverUrl } from "~/constants/server-url";
import type { AssetsDb } from "~/types/asset";
import type { CategoryDb } from "~/types/category";
import type { BaseModalStatus } from "~/types/injectKey";
import type { TransactionDb } from "~/types/transaction";

const getModal = () => {
  const modalOpen = ref<BaseModalStatus>(null);
  const transactionId = ref<string | null>(null);
  const updateModalOpen = (
    status: BaseModalStatus,
    id: string | null = null,
  ) => {
    modalOpen.value = status;
    transactionId.value = id;
  };
  return {
    modalOpen,
    transactionId,
    updateModalOpen,
  };
};

export const useTransactionPage = async () => {
  const modal = getModal();

  const categoriesPromise = useAsyncData("categories", () =>
    $fetch<{ message: string; data: CategoryDb[] }>(`${serverUrl}/category`),
  );

  const assetsPromise = useAsyncData("assets", () =>
    $fetch<{ message: string; data: AssetsDb[] }>(`${serverUrl}/asset`),
  );

  const fetcher = useLazyAsyncData<{
    message: string;
    data: TransactionDb[];
  }>("transaction", () => $fetch(`${serverUrl}/transactions`));

  const [{ data: rawCategories }, { data: rawAssets }] = await Promise.all([
    categoriesPromise,
    assetsPromise,
  ]);

  const resources = {
    categories: computed(() => rawCategories.value?.data ?? []),
    assets: computed(() => rawAssets.value?.data ?? []),
  };

  return {
    modal,
    resources,
    fetcher: {
      data: fetcher.data,
      status: fetcher.status,
    },
  };
};

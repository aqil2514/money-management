import type { RadioGroupItem } from "@nuxt/ui";
import { serverUrl } from "~/constants/server-url";
import type { AssetsDb } from "~/types/asset";
import type { BaseModalStatus } from "~/types/injectKey";

export const useUseAssetPage = () => {
  const modalOpen = ref<BaseModalStatus>(null);
  const assetId = ref<number | null>(null);
  const updateModalOpen = (status: BaseModalStatus, id?: number | null) => {
    modalOpen.value = status;
    assetId.value = id ? id : null;
  };

  const liquidityScoreMapping: Record<number, string> = {
    1: "Sangat Tidak Likuid (Butuh waktu > 1 tahun, misal: Properti/Tanah)",
    2: "Tidak Likuid (Butuh waktu bulanan, misal: Kendaraan/Koleksi)",
    3: "Moderat (Butuh beberapa hari/minggu, misal: Emas fisik/Reksadana Terproteksi)",
    4: "Likuid (Bisa cair dalam < 2 hari kerja, misal: Saham/Reksadana Pasar Uang)",
    5: "Sangat Likuid (Instan/Kas, misal: Saldo Bank/Dompet Digital)",
  };

  const statusItem = shallowRef<RadioGroupItem[]>([
    {
      value: "active",
      label: "Aktif",
    },
    {
      value: "nonactive",
      label: "Non Aktif",
    },
  ]);

  const currencyItem = shallowRef<RadioGroupItem[]>([
    {
      value: "idr",
      label: "Rp - Rupiah",
    },
    {
      value: "usd",
      label: "$ - Dolar",
    },
  ]);

  const { data, status } = useFetch<{ message: string; data: AssetsDb[] }>(
    `${serverUrl}/asset`,
    {
      key: "fetch-asset",
    },
  );

  return {
    modalOpen,
    assetId,
    updateModalOpen,
    liquidityScoreMapping,
    statusItem,
    currencyItem,

    fetcher: {
      data,
      status,
    },
  };
};

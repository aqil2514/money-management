import { CalendarDate, Time } from "@internationalized/date";
import type { FormSubmitEvent } from "@nuxt/ui";
import {
  defaultTransaction,
  type TransactionSchemaType,
} from "~/schemas/transaction-schema";
import { serverUrl } from "~/constants/server-url";

export const useTransactionForm = () => {
  const toast = useToast();
  const isLoading = ref(false);

  const now = new Date();

  const state = reactive({
    ...defaultTransaction,
    time: new Time(now.getHours(), now.getMinutes()),
  }) as TransactionSchemaType;

  async function onSubmit(event: FormSubmitEvent<TransactionSchemaType>) {
    const rawDate: CalendarDate = event.data.date;
    const rawTime: Time = event.data.time;

    const createdDate = new Date(
      rawDate.year,
      rawDate.month,
      rawDate.day,
      rawTime.hour,
      rawTime.minute,
    );

    const payload = {
      ...event.data,
      date: createdDate.toISOString(),
    };

    try {
      isLoading.value = true;
      await $fetch(`${serverUrl}/transactions`, {
        method: "POST",
        body: payload,
      });

      await refreshNuxtData("list-transactions");

      toast.add({
        title: "Berhasil!",
        description: "Data transaksi telah tersimpan di server Go.",
        color: "success",
      });
    } catch (error: any) {
      console.error(error);
      toast.add({
        title: "Gagal!",
        description:
          error.data?.message || "Tidak dapat terhubung ke server Go.",
        color: "error",
      });
    } finally {
      isLoading.value = false;
    }
  }

  return {
    state,
    onSubmit,
    isLoading,
  };
};

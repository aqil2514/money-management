// composables/useTransactionForm.ts
import * as z from "zod";
import { CalendarDateTime } from "@internationalized/date";
import type { FormSubmitEvent } from "@nuxt/ui";

export const useTransactionForm = () => {
  const toast = useToast();

  // 1. Schema
  const schema = z.object({
    nominal: z.number().min(1, "Nominal tidak boleh 0"),
    category: z.string(),
    assetName: z.string(),
    note: z.string().optional(),
  });

  type Schema = z.output<typeof schema>;

  // 2. State Tanggal (Shallow)
  const now = new Date();
  const transactionDate = shallowRef(
    new CalendarDateTime(
      now.getFullYear(),
      now.getMonth() + 1,
      now.getDate(),
      now.getHours(),
      now.getMinutes(),
      now.getSeconds(),
    ),
  );

  // 3. State Form
  const state = reactive({
    nominal: 0,
    category: "",
    assetName: "",
    note: "",
  });

  // 4. Submit Handler
  async function onSubmit(event: FormSubmitEvent<Schema>) {
    try {
      const payload = {
        ...event.data,
        date: transactionDate.value.toDate("UTC"),
      };

      const response = await $fetch("http://localhost:8000/api/transactions", {
        method: "POST",
        body: payload,
      });

      toast.add({
        title: "Berhasil!",
        description: "Data transaksi telah tersimpan di server Go.",
        color: "success",
      });

      console.log("Respon dari Go:", response);
    } catch (error: any) {
      console.error(error);
      toast.add({
        title: "Gagal!",
        description:
          error.data?.message || "Tidak dapat terhubung ke server Go.",
        color: "error",
      });
    }
  }

  return {
    schema,
    state,
    transactionDate,
    onSubmit,
  };
};

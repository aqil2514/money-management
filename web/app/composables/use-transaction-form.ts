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
    const payload = {
      ...event.data,
      date: transactionDate.value.toDate("UTC"),
    };
    toast.add({ title: "Tambah data sukses" });
    console.log("Payload:", payload);

    // Anda bisa tambahkan logic reset di sini
    // Object.assign(state, { nominal: 0, note: '' })
  }

  return {
    schema,
    state,
    transactionDate,
    onSubmit,
  };
};

<script setup lang="ts">
import { z } from "zod";
import { UserRole } from "@src/utils/enums";
import { toTypedSchema } from "@vee-validate/zod";
import { useForm } from "vee-validate";
import { useUserData } from "@src/stores/user";
import { useNotyf } from "@src/composable/useNotyf";

const userStore = useUserData();
const notyf = useNotyf();

const opened = ref(false);
const userDraft = computed<Partial<UserData> & { password: string; repassword: string }>(
  () => ({
    name: "",
    role: UserRole.Normal,
    password: "",
    repassword: "",
  })
);

const zodSchema = z
  .object({
    name: z.string({
      required_error: "用户名称不能为空",
    }),
    password: z
      .string({
        required_error: "密码不能为空",
      })
      .min(3, {
        message: "密码不能少于3个字符",
      }),
    repassword: z.string().optional(),
    role: z.number(),
  })
  .refine((data) => data.password === data.repassword, {
    message: "两次输入的密码不一致",
    path: ["repassword"],
  });
const validationSchema = toTypedSchema(zodSchema);
const { handleSubmit } = useForm({
  validationSchema,
  initialValues: userDraft,
});

useListener(Signal.OpenNewUser, () => {
  opened.value = true;
});

const createUser = handleSubmit(async (values) => {
  const res = await userStore.$add({
    name: values.name,
    password: values.password,
    role: values.role,
  });
  if (res.error) {
    notyf.error(res.message);
  } else {
    notyf.success("创建成功");
    opened.value = false;
  }
});
</script>
<template>
  <VModal
    is="form"
    method="post"
    novalidate
    autocomplete="off"
    size="medium"
    :open="opened"
    title="添加用户"
    actions="right"
    cancel-label="取消"
    @submit.prevent="createUser"
    @close="opened = false"
  >
    <template #content>
      <div class="modal-form">
        <VField id="name" v-slot="{ field }" horizontal label="用户名称 *">
          <VControl fullwidth>
            <VInput type="text" placeholder="用户名称" />
            <p v-if="field?.errorMessage" class="help is-danger">
              {{ field.errorMessage }}
            </p>
          </VControl>
        </VField>
        <VField id="password" v-slot="{ field }" horizontal label="密码 *">
          <VControl fullwidth>
            <VInput type="password" placeholder="密码" />
            <p v-if="field?.errorMessage" class="help is-danger">
              {{ field.errorMessage }}
            </p>
          </VControl>
        </VField>
        <VField id="repassword" v-slot="{ field }" horizontal label="再次输入密码">
          <VControl fullwidth>
            <VInput
              v-model="userDraft.repassword"
              type="password"
              placeholder="再次输入密码"
            />
            <p v-if="field?.errorMessage" class="help is-danger">
              {{ field.errorMessage }}
            </p>
          </VControl>
        </VField>
        <VField id="role" v-slot="{ field }" horizontal label="类型">
          <VControl fullwidth>
            <VSelect class="is-rounded">
              <VOption :value="UserRole.Normal">普通用户</VOption>
              <VOption :value="UserRole.Admin">管理员</VOption>
            </VSelect>
          </VControl>
          <p v-if="field?.errorMessage" class="help is-danger">
            {{ field.errorMessage }}
          </p>
        </VField>
      </div>
    </template>
    <template #action>
      <VButton type="submit" color="primary" raised> 创建 </VButton>
    </template>
  </VModal>
</template>

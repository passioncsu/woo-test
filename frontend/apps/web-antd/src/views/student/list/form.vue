<script lang="ts" setup>
import { z } from '@vben/common-ui';
import { message } from 'ant-design-vue';
import { ref } from 'vue';

import { useVbenForm } from '#/adapter/form';
import { createStudentApi, updateStudentApi } from '#/api/modules/student';

const emit = defineEmits<{
  success: [];
}>();

const isEdit = ref(false);
const editId = ref<number>();
const submitting = ref(false);

const phoneRule = z.string().regex(/^$|^1[3-9]\d{9}$/, '请输入正确的手机号').optional();
const emailRule = z.string().regex(/^$|^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$/, '请输入正确的邮箱').optional();

const [Form, formApi] = useVbenForm({
  schema: [
    {
      component: 'Input',
      componentProps: { placeholder: '请输入学号', maxlength: 50 },
      fieldName: 'student_no',
      label: '学号',
      rules: 'required',
    },
    {
      component: 'Input',
      componentProps: { placeholder: '请输入姓名', maxlength: 100 },
      fieldName: 'name',
      label: '姓名',
      rules: 'required',
    },
    {
      component: 'Select',
      componentProps: {
        placeholder: '请选择性别',
        allowClear: true,
        options: [
          { label: '男', value: '男' },
          { label: '女', value: '女' },
        ],
      },
      fieldName: 'gender',
      label: '性别',
    },
    {
      component: 'Input',
      componentProps: { placeholder: '请输入专业', maxlength: 100 },
      fieldName: 'major',
      label: '专业',
    },
    {
      component: 'Input',
      componentProps: { placeholder: '请输入班级', maxlength: 100 },
      fieldName: 'className',
      label: '班级',
    },
    {
      component: 'DatePicker',
      componentProps: { placeholder: '请选择出生日期', class: 'w-full', valueFormat: 'YYYY-MM-DD' },
      fieldName: 'birthday',
      label: '出生日期',
    },
    {
      component: 'Select',
      componentProps: {
        placeholder: '请选择状态',
        options: [
          { label: '在读', value: 1 },
          { label: '休学', value: 2 },
          { label: '毕业', value: 3 },
        ],
      },
      fieldName: 'status',
      label: '状态',
      defaultValue: 1,
    },
    {
      component: 'Input',
      componentProps: { placeholder: '请输入手机号', maxlength: 11 },
      fieldName: 'phone',
      label: '手机号',
      rules: phoneRule,
    },
    {
      component: 'Input',
      componentProps: { placeholder: '请输入邮箱', maxlength: 100 },
      fieldName: 'email',
      label: '邮箱',
      rules: emailRule,
    },
    {
      component: 'Input',
      componentProps: { placeholder: '请输入地址', maxlength: 255 },
      fieldName: 'address',
      label: '地址',
    },
  ],
  commonConfig: {
    componentProps: {
      class: 'w-full',
    },
  },
  layout: 'horizontal',
  wrapperClass: 'grid-cols-1 md:grid-cols-2',
});

async function handleSubmit() {
  const result = await formApi.validate();
  if (Object.keys(result?.errors ?? {}).length > 0) {
    return;
  }

  const rawValues = formApi.getValues();

  const values: Record<string, any> = {};

  if (isEdit.value) {
    // 编辑模式：发送所有字段（含空串），支持清空
    for (const [key, value] of Object.entries(rawValues)) {
      if (value !== undefined) {
        values[key] = value === null ? '' : value;
      }
    }
  } else {
    // 创建模式：过滤空值
    for (const [key, value] of Object.entries(rawValues)) {
      if (value !== undefined && value !== null && value !== '') {
        values[key] = value;
      }
    }
  }

  // 字段映射: className -> class（避免JS保留字冲突）
  if (values.className !== undefined) {
    values['class'] = values.className;
    delete values.className;
  }

  submitting.value = true;
  try {
    if (isEdit.value && editId.value) {
      await updateStudentApi(editId.value, values);
      message.success('更新成功');
    } else {
      await createStudentApi(values as any);
      message.success('创建成功');
    }
    emit('success');
  } catch {
    // 错误已由 request interceptor 处理
  } finally {
    submitting.value = false;
  }
}

function setFormData(data: Record<string, any>) {
  isEdit.value = !!data.id;
  editId.value = data.id;
  // 字段映射: class -> className（避免JS保留字冲突）
  const mapped = { ...data };
  if (mapped['class'] !== undefined) {
    mapped.className = mapped['class'];
    delete mapped['class'];
  }
  formApi.setValues(mapped);
}

function resetFields() {
  isEdit.value = false;
  editId.value = undefined;
  formApi.resetForm();
}

defineExpose({ setFormData, resetFields });
</script>

<template>
  <div class="p-4">
    <Form />
    <div class="mt-4 flex justify-end gap-3">
      <a-button @click="resetFields">重置</a-button>
      <a-button type="primary" :loading="submitting" @click="handleSubmit">提交</a-button>
    </div>
  </div>
</template>

<template>
  <div class="input__wrapper">
    <label class="input-field-label" :for="id">
      <span>{{ fieldName }}</span>
      <input class="input-field" :id="id" :type="fieldType" :name="fieldName" autocomplete="off"
             :placeholder="fieldName" v-model="value">
    </label>
    <span v-if="errorsExist" class="invalid-field">{{ errorText }}</span>
  </div>
</template>

<script>
export default {
  name: 'InputField',
  props: {
    id: { type: String, required: true },
    fieldType: { type: String, default: "text" },
    fieldName: { type: String, default: "" },
    errors: { type: Object, default: () => ({}) }
  },
  data() {
    return {
      value: "",
    }
  },
  computed: {
    errorsExist() {
      return Object.keys(this.$props.errors).length !== 0;
    },
    errorText() {
      return this.$props.errors[this.$props.id].message;
    }
  },
  watch: {
    value: function() {
      this.$emit('getValue', this.$props.id, this.value);
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.input-field-label {
  display: flex;
  flex-direction: column;
}
.input-field-label span {
  align-self: flex-start;
  font-weight: 500;
}
.input__wrapper {
  margin-top: 15px;
}
.input__wrapper:last-child {
  margin-bottom: 20px;
}
.input-field {
  font-size: 16px;
  font-family: inherit;
  padding: 0.25em 0.5em;
  background-color: #ffffff;
  border: 1px solid #6F8BE9;
  border-radius: 4px;
  transition: 180ms box-shadow ease-in-out;
}
.input-field:focus {
  //border-color: hsl(245, 53%, 70%);
  box-shadow: 0 0 0 1px #6F8BE9;
  outline: 2px solid transparent;
}
</style>

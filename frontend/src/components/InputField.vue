<template>
  <div class="input__wrapper">
    <label class="input-field-label" :for="data.id">
      <span>{{ data.data.name }}</span>
      <input class="input-field" :id="data.id" type="text" :name="data.data.name" autocomplete="off"
             :placeholder="data.data.name" v-model="value">
    </label>
    <span v-if="errorsExist" class="invalid-field">{{ errorText }}</span>
  </div>
</template>

<script>
export default {
  name: 'InputField',
  props: {
    data: { type: Object },
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
      this.$emit('getValue', { id: this.$props.data.id, label: this.value });
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
//border-color: hsl(245, 53%, 70%); box-shadow: 0 0 0 1px #6F8BE9;
  outline: 2px solid transparent;
}
</style>

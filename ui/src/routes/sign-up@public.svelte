<script lang="ts">
  import { goto } from "$app/navigation";
  import type { SignUpFormErrors } from "$lib/domain/sign-up";
  import { validate as validateSignInForm, signUp } from "$lib/domain/sign-up";
  import { userStore } from "$lib/domain/user";
  import { login } from "$lib/domain/sign-in";
  import { createPmc } from "$lib/domain/pmc";

  let email = "";
  let username = "";
  let password = "";
  let pmcName = "";
  let passwordConfirmation = "";
  let errors: SignUpFormErrors;
  let apiError: string | null;

  const handleSubmit = () => {
    const data = {
      email,
      username,
      password,
      passwordConfirmation,
      pmcName,
    };

    errors = validateSignInForm(data);

    if (errors) return;

    // @TODO: handle edge case where user is created but pmc creation fails.
    signUp(data).then((res) => {
      if (res.error) {
        apiError = res.error;
        return;
      }

      apiError = null;
      userStore.update({
        id: res.id,
        email: res.email,
        username: res.username,
      });

      login({ email, password })
        .then((res) => {
          if (res.error) {
            apiError = res.error;
            return;
          }

          apiError = null;
          sessionStorage.setItem("token", res.token);

          createPmc({ name: pmcName })
            .then((res) => {
              if (res.error) {
                apiError = res.error;
                return;
              }

              apiError = null;
              goto("/pmc");
            })
            .catch((err) => {
              apiError = err;
            });
        })
        .catch((err) => {
          apiError = err;
        });
    });
  };
</script>

<div class="flex justify-center h-screen items-center">
  <div class="card w-96 bg-slate-400 shadow-xl text-black nav-offset card-min-height">
    <div class="card-body">
      <div class="justify-between">
        <h1 class="card-title pb-5">Sign Up</h1>
        <form id="sign-up" on:submit|preventDefault={handleSubmit}>
          <div class="form-control w-full max-w-xs">
            <label for="email" class="label">
              <span class="label-text-alt text-black">Email</span>
            </label>
            <input
              bind:value={email}
              name="email"
              type="email"
              placeholder="Email"
              class="input input-bordered w-full max-w-xs text-slate-400"
            />
            <span class="label-text-alt text-rose-800 p-1 input-error">
              {#if errors && errors.email}
                {errors.email}
              {/if}
            </span>
          </div>
          <div class="form-control w-full max-w-xs">
            <label for="Username" class="label">
              <span class="label-text-alt text-black">Username</span>
            </label>
            <input
              bind:value={username}
              name="username"
              type="text"
              placeholder="Username"
              class="input input-bordered w-full max-w-xs text-slate-400"
            />
            <span class="label-text-alt text-rose-800 p-1 input-error">
              {#if errors && errors.username}
                {errors.username}
              {/if}
            </span>
          </div>
          <div class="form-control w-full max-w-xs">
            <label for="pmcName" class="label">
              <span class="label-text-alt text-black">PMC Name</span>
            </label>
            <input
              bind:value={pmcName}
              name="pmcName"
              type="text"
              placeholder="Pmc name"
              class="input input-bordered w-full max-w-xs text-slate-400"
            />
            <span class="label-text-alt text-rose-800 p-1 input-error">
              {#if errors && errors.pmcName}
                {errors.pmcName}
              {/if}
            </span>
          </div>
          <div class="form-control w-full max-w-xs">
            <label for="password" class="label">
              <span class="label-text-alt text-black">Password</span>
            </label>
            <input
              bind:value={password}
              name="password"
              type="password"
              placeholder="Password"
              class="input input-bordered w-full max-w-xs text-slate-400"
            />
            <span class="label-text-alt text-rose-800 p-1 input-error">
              {#if errors && errors.password}
                {errors.password}
              {/if}
            </span>
          </div>
          <div class="form-control w-full max-w-xs pb-8">
            <label for="password-confirmation" class="label">
              <span class="label-text-alt text-black">Password Confirmation</span>
            </label>
            <input
              bind:value={passwordConfirmation}
              name="password-confirmation"
              type="password"
              placeholder="Password Confirmation"
              class="input input-bordered w-full max-w-xs text-slate-400"
            />
            <span class="label-text-alt text-rose-800 p-1 input-error">
              {#if errors && errors.passwordConfirmation}
                {errors.passwordConfirmation}
              {/if}
            </span>
          </div>
          <div class="card-actions">
            <span class="label-text-alt text-rose-800 p-1 input-error">
              {#if apiError}
                {apiError}
              {/if}
            </span>
            <button for="sign-up" class="btn btn-primary btn-block" type="submit">Sign Up</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</div>

<style>
  .nav-offset {
    margin-bottom: 30vh;
  }

  .input-error {
    min-height: 16px;
  }

  .card-min-height {
    min-height: 400px;
  }
</style>

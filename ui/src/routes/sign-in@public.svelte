<script lang="ts">
  // @TODO: Re-style / unfuck sign-in and sign up pages
  // @TODO: Fix broken layout issues
  // @TODO: Add stores
  // @TODO: find a way to include auth header on all requests without explicitly handling it
  import { goto } from "$app/navigation";
  import { validate, login } from "$lib/domain/sign-in";
  import type { LoginFormErrors } from "$lib/domain/sign-in";

  let email = "";
  let password = "";
  let errors: LoginFormErrors;
  let apiError: string | null;

  const handleSubmit = () => {
    errors = validate({ email, password });
    if (errors) return;
    login({ email, password })
      .then((res) => {
        if (res.error) {
          apiError = res.error;
          return;
        }

        apiError = null;
        sessionStorage.setItem("token", res.token);
        goto("/pmc");
      })
      .catch((err) => {
        apiError = err;
      });
  };
</script>

<div class="flex justify-center h-screen items-center">
  <div class="card w-96 bg-slate-400 shadow-xl h-1/3 text-black nav-offset card-min-height">
    <div class="card-body">
      <div class="justify-between">
        <h1 class="card-title pb-5">Sign In</h1>
        <form id="sign-in" on:submit|preventDefault={handleSubmit}>
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
          <div class="form-control w-full max-w-xs pb-14">
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
          <div class="card-actions">
            <span class="label-text-alt text-rose-800 p-1 input-error">
              {#if apiError}
                {apiError}
              {/if}
            </span>
            <button for="sign-in" class="btn btn-primary btn-block" type="submit">Sign In</button>
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

def _gpg_sign_impl(ctx):
  output_file = ctx.actions.declare_file(ctx.file.base.basename + ctx.attr.suffix)
  command = "echo ${GPG_PASS} | gpg --pinentry-mode loopback --digest-algo SHA512 --passphrase-fd 0 --output %s --detach-sig %s" % (output_file.path, ctx.file.base.path)
  ctx.actions.run_shell(
    command = command,
    use_default_shell_env = True,
    inputs = [ctx.file.base],
    outputs = [output_file],
    progress_message = "Signing binary",
    mnemonic = "gpg",
  )
  runfiles = ctx.runfiles(files = [output_file])
  return [DefaultInfo(executable = output_file, runfiles = runfiles)]

gpg_sign = rule(
  implementation = _gpg_sign_impl,
  attrs = {
    "base": attr.label(allow_single_file=True),
    "suffix": attr.string(default=".sig"),
  },
)

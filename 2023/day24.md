# Signing

The process of signing involves... well, signing an artifact with a key, and later verifying that this artifact has not been tampered with.

An "artifact" in this scenario can be anything

- [code](https://venafi.com/machine-identity-basics/what-is-code-signing/#item-1)
- [git commit](https://docs.github.com/en/authentication/managing-commit-signature-verification/signing-commits)
- [container images](https://docs.sigstore.dev/cosign/overview/)

Signing and verifying the signature ensures that the artifact(container) we pulled from the registry is the same one that we pushed.
This secures us from supply chain and man-in-the-middle attack where we download something different that we wanted.

The CI workflow would look like this:

0. Developer pushes code to Git
1. CI builds the code into a container
2. **CI signs the container with our private key**
3. CI pushes the signed container to our registry

And then when we want to deploy this image:

1. Pull the image
2. **Verify the signature with our public key**
   1. If signature does not match, fail the deploy - image is probably compromised
3. If signature does match, proceed with the deploy

This workflow is based on public-private key cryptography.
When you sign something with your private key, everyone that has access to your public key can verify that this was signed by you.

And since the public key is... well, public, that means everyone.

## The danger of NOT signing your images

If you are not signing your container images, there is the danger that someone will replace an image in your repository with another image that is malicious.

For example, you can push the `my-repo/my-image:1.0.0` image to your repository, but image tags, even versioned ones (like `1.0.0`) are mutable.
So an attacker that has access to your repo can push another image, tag it the same way, and this way it will override your image.
Then, when you go an deploy this image, the image that will get deployed is the one that attacked forged.
This will probably be a maliciuos one.
For example, on that has malware, is stealing data, or using your infrastructure for mining crypto currencies.

This problem can be solved by signing your images, because when you sign an images, then you can later verify that what you pull is what you uploaded in the first place.

So let's take a look at how we can do this via a tool called [cosign](https://docs.sigstore.dev/cosign/overview/).

## Signing container images

First, download the tool, following the instructions for your OS [here](https://docs.sigstore.dev/cosign/installation/).

Generate a key-pair if you don't have one:

```console
cosign generate-key-pair
```

This will output two files in the current folder:

- `cosign.key` - your private key.
DO NOT SHARE WITH ANYONE.
- `cosign.pub` - your public key.
Share with whoever needs it.

We can use the private key to sign an image:

```console
$ cosign sign --key cosign.key asankov/signed
Enter password for private key:

Pushing signature to: index.docker.io/asankov/signed
```

This command signed the `asankov/signed` contaner image and pushed the signature to the container repo.

## Verifying signatures

Now that we have signed the image, let's verify the signature.

For that, we need our public key:

```console
$ cosign verify --key=cosign.pub asankov/signed | jq

Verification for index.docker.io/asankov/signed:latest --
The following checks were performed on each of these signatures:
  - The cosign claims were validated
  - The signatures were verified against the specified public key
[
  {
    "critical": {
      "identity": {
        "docker-reference": "index.docker.io/asankov/signed"
      },
      "image": {
        "docker-manifest-digest": "sha256:93d62c92b70efc512379cf89317eaf41b8ce6cba84a5e69507a95a7f15708506"
      },
      "type": "cosign container image signature"
    },
    "optional": null
  }
]
```

The output of this command showed us that the image is signed by the key we expected.
Since we are the only ones that have access to our private key, this means that no one except us could have pushed this image and signature to the container repo.
Hence, the contents of this image have not been tampered with since we pushed it.

Let's try to verify an image that we have NOT signed.

```console
$ cosign verify --key=cosign.pub asankov/not-signed
Error: no matching signatures:

main.go:62: error during command execution: no matching signatures:
```

Just as expected, `cosign` could not verify the signature of this image (because there was not one).

In this example, this image (`asankov/not-signed`) is not signed at all, but we would have gotten the same error if someone had signed this image with different key than the one we are using to verify it.

### Verifying signatures in Kubernetes

In the previous example, we were verifying the signatures by hand.
However, that is good only for demo purposes or for playing around with the tool.

In a real-world scenario, you would want this verification to be done automatically at the time of deploy.

Fortunately, there are many `cosign` integrations for doing that.

For example, if we are using Kubernetes, we can deploy a validating webhook that will audit all new deployments and verify that the container images used by them are signed.

For Kubernetes you can choose from 3 existing integrations - [Gatekeeper](https://github.com/sigstore/cosign-gatekeeper-provider), [Kyverno](https://kyverno.io/docs/writing-policies/verify-images/) or [Conaisseur](https://github.com/sse-secure-systems/connaisseur#what-is-connaisseur).
You can choose one of the three depending on your preference, or if you are already using them for something else.

## Dangers to be aware of

As with everything else, signing images is not a silver bullet and will not solve all your security problems.

There is still the problem that your private keys might leak, in which case everyone can sign everything and it will still pass your signature check.

However, integrating signing into your workflow adds yet another layer of defence and one more hoop for attackers to jump over.

## Summary

Signing artifacts prevents supply-chain and man-in-the-middle attacks, by allowing you to verify the integrity of your artifacts.

[Sigstore](https://sigstore.dev/) and [cosign](https://docs.sigstore.dev/cosign/overview/) are useful tools to sign your artifacts and they come with many integrations to choose from.
See you on [Day 25](day25.md).

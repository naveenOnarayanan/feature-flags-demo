import 'package:launchdarkly_flutter_client_sdk/launchdarkly_flutter_client_sdk.dart';

class SampleDart {
  final LDClient ldClient;

  SampleDart(String sdkKey) : ldClient = LDClient(sdkKey);

  void sampleMethod1() async {
    LDUser user = LDUser(
      key: "context-key-123abc",
      name: "Sandy",
    );

    await ldClient.start(user);

    bool flagValue = await ldClient.boolVariation("is-dark-mode-enabled", false);

    if (flagValue) {
      // The code to run if the feature is on
      System.out.println("Dark mode!");
    } else {
      // The code to run if the feature is off
      bool otherFlag = await ldClient.boolVariation("is-ameyas-weird-blue-theme", false);
      if (otherFlag) {
        System.out.println("Why? Just... why?");
      } else {
        System.out.println("Light mode it is, then!");
      }
    }

    await ldClient.close();
  }
}

void main() {
  SampleDart sample = SampleDart("sdk-key-123abc");
  sample.sampleMethod1();
}

#include <stdio.h>
#include <string.h>

int main() {
    char* obs_flag =
            "\x3c\x36\x3b\x3d\x21\x39\x6a\x2f"
            "\x34\x2e\x77\x39\x32\x6a\x39\x2f"
            "\x36\x3b\x77\x39\x3f\x28\x3f\x3b"
            "\x36\x77\x1c\x0e\x0d\x27\x00";

    char password[512]; // input buffer
    char true_flag[31];

    char* halloweenie_message = "What is the best and sp00kiest breakfast cereal?";
    char* password_prompt = "Please enter the passphrase: ";
    char* not_flag = "notflag{you-guessed-it---this-is-not-the-flag}";
    char* passphrase = "c0unt-ch0cula";
    int res = 0;        // result of functions

    for (int i = 0; i < 31; i++) {
        if (obs_flag[i] != 0) {
            true_flag[i] = obs_flag[i] ^ 'Z';
        }
    }
    true_flag[30] = '\x0';

    printf("%s\n", halloweenie_message);
    printf("%s", password_prompt);
    scanf("%s", password);

    res = strcmp(password, passphrase);
    if (res == 0) {
        printf("%s\n", true_flag);
    } else {
        printf("%s\n", not_flag);
    }

    return 0;
}

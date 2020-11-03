import 'package:divoc/base/routes.dart';
import 'package:divoc/base/theme.dart';
import 'package:divoc/data_source/network.dart';
import 'package:divoc/generated/l10n.dart';
import 'package:divoc/home/flow_screen.dart';
import 'package:divoc/home/home_page.dart';
import 'package:divoc/home/home_repository.dart';
import 'package:divoc/login/auth_repository.dart';
import 'package:divoc/login/login_page.dart';
import 'package:flutter/material.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:flutter_portal/flutter_portal.dart';
import 'package:provider/provider.dart';

class ProviderApp extends StatelessWidget {
  final AuthRepository authRepository;
  final HomeRepository homeRepository;
  final ApiClient apiClient;

  ProviderApp({
    @required this.authRepository,
    @required this.homeRepository,
    @required this.apiClient,
  });

  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        Provider<AuthRepository>(create: (_) => authRepository),
        Provider<ApiClient>(create: (_) => apiClient),
        Provider<HomeRepository>(create: (_) => homeRepository),
      ],
      child: Portal(
        child: MaterialApp(
          theme: DivocTheme.appTheme,
          localizationsDelegates: [
            DivocLocalizations.delegate,
            GlobalMaterialLocalizations.delegate,
            GlobalCupertinoLocalizations.delegate,
            GlobalWidgetsLocalizations.delegate
          ],
          supportedLocales: [Locale('en', ''), Locale('hi', '')],
          home: authRepository.isLoggedIn ? HomePage() : LoginPage(),
          routes: {
            DivocRoutes.home: (context) => HomePage(),
            DivocRoutes.login: (context) => LoginPage(),
          },
        ),
      ),
    );
  }
}

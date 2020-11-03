import 'package:divoc/base/common_widget.dart';
import 'package:divoc/generated/l10n.dart';
import 'package:flutter/material.dart';
import 'package:intl/intl.dart';

import 'navigation_flow.dart';

class NewUserEnrollForm extends StatefulWidget {
  final RouteInfo routeInfo;

  NewUserEnrollForm(this.routeInfo);

  @override
  _NewUserEnrollFormState createState() => _NewUserEnrollFormState();
}

class _NewUserEnrollFormState extends State<NewUserEnrollForm> {
  final textController = new TextEditingController();

  @override
  void dispose() {
    textController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return DivocForm(
      title: DivocLocalizations.of(context).titleDetailsRecipient,
      child: SingleChildScrollView(
        child: Column(
          children: [
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: TextFormField(
                keyboardType: TextInputType.name,
                onSaved: (value) {},
                validator: (value) {
                  var msg = value.isEmpty ? "Cannot be Empty" : null;
                  return msg;
                },
                decoration: InputDecoration(
                  labelText: "Name",
                  border: OutlineInputBorder(),
                ),
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: TextFormField(
                keyboardType: TextInputType.emailAddress,
                onSaved: (value) {},
                validator: (value) {
                  var msg = value.isEmpty ? "Cannot be Empty" : null;
                  return msg;
                },
                decoration: InputDecoration(
                  labelText: "Email",
                  border: OutlineInputBorder(),
                ),
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: TextFormField(
                keyboardType: TextInputType.phone,
                onSaved: (value) {},
                validator: (value) {
                  var msg = value.isEmpty ? "Cannot be Empty" : null;
                  return msg;
                },
                decoration: InputDecoration(
                  labelText: "Mobile",
                  prefixText: "+91 ",
                  border: OutlineInputBorder(),
                ),
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: TextFormField(
                readOnly: true,
                controller: textController,
                showCursor: true,
                keyboardType: TextInputType.datetime,
                onSaved: (value) {},
                validator: (value) {
                  var msg = value.isEmpty ? "Cannot be Empty" : null;
                  return msg;
                },
                decoration: InputDecoration(
                  labelText: "DOB",
                  prefixIcon: IconButton(
                    icon: Icon(Icons.date_range),
                    onPressed: () async {
                      final selectTime = await _selectTime(context);
                      if (selectTime != null) {
                        String formattedDate =
                            DateFormat('dd-MM-yyyy').format(selectTime);
                        textController.text = formattedDate;
                      }
                    },
                  ),
                  border: OutlineInputBorder(),
                ),
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: TextFormField(
                keyboardType: TextInputType.datetime,
                onSaved: (value) {},
                validator: (value) {
                  var msg = value.isEmpty ? "Cannot be Empty" : null;
                  return msg;
                },
                decoration: InputDecoration(
                  labelText: "Nationality",
                  border: OutlineInputBorder(),
                ),
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: FormButton(
                text: "DONE",
                onPressed: () {
                  NavigationFormFlow.push(context,
                      widget.routeInfo.nextRoutesMeta[0].fullNextRoutePath);
                },
              ),
            )
          ],
        ),
      ),
    );
  }

  Future<DateTime> _selectTime(BuildContext context) async {
    /*TimeOfDay selectedTime = TimeOfDay.now();
    final TimeOfDay datePicked = await showTimePicker(
        context: context,
        initialTime: selectedTime,
        builder: (BuildContext context, Widget child) {
          return MediaQuery(
            data: MediaQuery.of(context).copyWith(alwaysUse24HourFormat: false),
            child: child,
          );
        });

    return datePicked;*/

    final datePicker = await showDatePicker(
        context: context,
        initialDate: DateTime.now(),
        //which date will display when user open the picker
        firstDate: DateTime(1950),
        //what will be the previous supported year in picker
        lastDate: DateTime.now());
    return datePicker;

    /* if (datePicked != null && datePicked != selectedTime )
      setState(() {
        selectedTime = datePicked;
      });*/
  }
}

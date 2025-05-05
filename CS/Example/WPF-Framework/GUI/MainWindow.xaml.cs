using System.Text;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;

namespace WPF_Freamwork_GUI;

/// <summary>
/// Interaction logic for MainWindow.xaml
/// </summary>
public partial class MainWindow : Window {
  public MainWindow() {
    InitializeComponent();
  }

  private void ElementAdd(object sender, RoutedEventArgs e) {
    // IsNullOrWhiteSpace is ListBox in the input field checking if the input is empty
    // Contains the same name in the ListBox
    if (!string.IsNullOrWhiteSpace(InputElement.Text) && !ListNames.Items.Contains(InputElement.Text)) {
      ListNames.Items.Add(InputElement.Text);
      InputElement.Clear();
    }
  }

  private void ElementClear(object sender, RoutedEventArgs e) {
    ListNames.Items.Clear();
  }
}